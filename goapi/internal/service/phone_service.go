package service

import (
	"context"
	"fmt"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/repository"
	"time"

	"gorm.io/gorm"
)

// PhoneService handles phone number and verification code operations
type PhoneService struct {
	thirdPartyService ThirdPartyServiceInterface
	transactionRepo   repository.TransactionRepository
	logRepo           repository.LogRepository
	db                *gorm.DB
}

// ThirdPartyServiceInterface defines methods for third-party service
type ThirdPartyServiceInterface interface {
	GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error)
	GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*CodeResponse, error)
	HealthCheck(ctx context.Context) map[string]bool
}

// PhoneServiceInterface defines the phone service methods
type PhoneServiceInterface interface {
	GetPhone(ctx context.Context, customerID uint, businessType, cardType string) (*GetPhoneResult, error)
	GetCode(ctx context.Context, customerID uint, phoneNumber string, timeout time.Duration) (*GetCodeResult, error)
}

// GetPhoneResult represents the result of getting a phone number
type GetPhoneResult struct {
	PhoneNumber string    `json:"phone_number"`
	CountryCode string    `json:"country_code"`
	Cost        float64   `json:"cost"`
	ValidUntil  time.Time `json:"valid_until"`
	ProviderID  string    `json:"provider_id"`
	Balance     float64   `json:"remaining_balance"`
}

// GetCodeResult represents the result of getting a verification code
type GetCodeResult struct {
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	ReceivedAt time.Time `json:"received_at"`
	ProviderID string    `json:"provider_id"`
}

// NewPhoneService creates a new phone service instance
func NewPhoneService(
	thirdPartyService ThirdPartyServiceInterface,
	transactionRepo repository.TransactionRepository,
	logRepo repository.LogRepository,
	db *gorm.DB,
) *PhoneService {
	return &PhoneService{
		thirdPartyService: thirdPartyService,
		transactionRepo:   transactionRepo,
		logRepo:           logRepo,
		db:                db,
	}
}

// GetPhone requests a phone number for receiving SMS
func (s *PhoneService) GetPhone(ctx context.Context, customerID uint, businessType, cardType string) (*GetPhoneResult, error) {
	// Get phone number from third-party service
	phoneResponse, err := s.thirdPartyService.GetPhone(ctx, businessType, cardType)
	if err != nil {
		s.logAPICall(ctx, customerID, "get_phone", "failed", fmt.Sprintf("Provider error: %v", err))
		return nil, fmt.Errorf("failed to get phone number: %w", err)
	}

	// Check current balance
	balance, err := s.transactionRepo.GetBalance(ctx, int64(customerID))
	if err != nil {
		s.logAPICall(ctx, customerID, "get_phone", "failed", fmt.Sprintf("Balance check error: %v", err))
		return nil, fmt.Errorf("failed to check balance: %w", err)
	}

	// Check if customer has sufficient balance
	if balance < phoneResponse.Cost {
		s.logAPICall(ctx, customerID, "get_phone", "failed", fmt.Sprintf("Insufficient balance: %.2f < %.2f", balance, phoneResponse.Cost))
		return nil, common.ErrInsufficientBalance
	}

	// Calculate new balance
	newBalanceAfter := balance - phoneResponse.Cost

	// Deduct the cost from customer's balance
	deductTx := &domain.Transaction{
		CustomerID:    int64(customerID),
		Amount:        -phoneResponse.Cost, // Negative for deduction
		BalanceBefore: balance,
		BalanceAfter:  newBalanceAfter,
		Type:          "2", // API consumption
		Notes:         fmt.Sprintf("Phone number rental: %s (Provider: %s)", phoneResponse.PhoneNumber, phoneResponse.ProviderID),
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.Create(ctx, deductTx); err != nil {
		s.logAPICall(ctx, customerID, "get_phone", "failed", fmt.Sprintf("Transaction creation error: %v", err))
		return nil, fmt.Errorf("failed to create deduction transaction: %w", err)
	}

	// Log successful API call
	s.logAPICall(ctx, customerID, "get_phone", "success",
		fmt.Sprintf("Phone: %s, Cost: %.2f, Provider: %s",
			phoneResponse.PhoneNumber, phoneResponse.Cost, phoneResponse.ProviderID))

	return &GetPhoneResult{
		PhoneNumber: phoneResponse.PhoneNumber,
		CountryCode: phoneResponse.CountryCode,
		Cost:        phoneResponse.Cost,
		ValidUntil:  phoneResponse.ValidUntil,
		ProviderID:  phoneResponse.ProviderID,
		Balance:     newBalanceAfter,
	}, nil
} // GetCode retrieves SMS verification code for a specific phone number
func (s *PhoneService) GetCode(ctx context.Context, customerID uint, phoneNumber string, timeout time.Duration) (*GetCodeResult, error) {
	// Validate timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	if timeout > 5*time.Minute {
		timeout = 5 * time.Minute
	}

	// Create a timeout context
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Get verification code from third-party service
	codeResponse, err := s.thirdPartyService.GetCode(ctxWithTimeout, phoneNumber, timeout)
	if err != nil {
		s.logAPICall(ctx, customerID, "get_code", "failed",
			fmt.Sprintf("Phone: %s, Error: %v", phoneNumber, err))

		// Check if it's a timeout error
		if ctxWithTimeout.Err() == context.DeadlineExceeded {
			return nil, common.ErrCodeTimeout
		}
		return nil, fmt.Errorf("failed to get verification code: %w", err)
	}

	// Log successful API call
	s.logAPICall(ctx, customerID, "get_code", "success",
		fmt.Sprintf("Phone: %s, Code received, Provider: %s",
			phoneNumber, codeResponse.ProviderID))

	return &GetCodeResult{
		Code:       codeResponse.Code,
		Message:    codeResponse.Message,
		ReceivedAt: codeResponse.ReceivedAt,
		ProviderID: codeResponse.ProviderID,
	}, nil
}

// logAPICall logs API calls for audit purposes
func (s *PhoneService) logAPICall(ctx context.Context, customerID uint, operation, status, details string) {
	// Create log entry in a separate goroutine to avoid blocking
	go func() {
		logEntry := &domain.APILog{
			CustomerID:   int64(customerID),
			RequestIP:    s.getIPFromContext(ctx),
			RequestPath:  fmt.Sprintf("/api/phone/%s", operation),
			RequestBody:  fmt.Sprintf("Status: %s, Details: %s", status, details),
			ResponseCode: 200, // Default success code
			CreatedAt:    time.Now(),
		}

		if status == "failed" {
			logEntry.ResponseCode = 400
		}

		// Best effort logging - ignore errors
		s.logRepo.Create(context.Background(), logEntry)
	}()
}

// getIPFromContext extracts IP address from context
func (s *PhoneService) getIPFromContext(ctx context.Context) string {
	if ip, ok := ctx.Value("ip_address").(string); ok {
		return ip
	}
	return "unknown"
}

// getUserAgentFromContext extracts user agent from context
func (s *PhoneService) getUserAgentFromContext(ctx context.Context) string {
	if ua, ok := ctx.Value("user_agent").(string); ok {
		return ua
	}
	return "unknown"
}
