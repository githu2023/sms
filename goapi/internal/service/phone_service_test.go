package service

import (
	"context"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockThirdPartyServiceInterface for testing
type MockThirdPartyServiceInterface struct {
	mock.Mock
}

func (m *MockThirdPartyServiceInterface) GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error) {
	args := m.Called(ctx, businessType, cardType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*PhoneResponse), args.Error(1)
}

func (m *MockThirdPartyServiceInterface) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*CodeResponse, error) {
	args := m.Called(ctx, phoneNumber, timeout)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CodeResponse), args.Error(1)
}

func (m *MockThirdPartyServiceInterface) HealthCheck(ctx context.Context) map[string]bool {
	args := m.Called(ctx)
	return args.Get(0).(map[string]bool)
}

// MockLogRepository for testing
type MockLogRepository struct {
	mock.Mock
}

func (m *MockLogRepository) Create(ctx context.Context, log *domain.APILog) error {
	args := m.Called(ctx, log)
	return args.Error(0)
}

func (m *MockLogRepository) FindByID(ctx context.Context, id int64) (*domain.APILog, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.APILog), args.Error(1)
}

func (m *MockLogRepository) FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.APILog, int64, error) {
	args := m.Called(ctx, customerID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.APILog), args.Get(1).(int64), args.Error(2)
}

func (m *MockLogRepository) FindByPath(ctx context.Context, path string, limit, offset int) ([]*domain.APILog, int64, error) {
	args := m.Called(ctx, path, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.APILog), args.Get(1).(int64), args.Error(2)
}

func (m *MockLogRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]*domain.APILog, int64, error) {
	args := m.Called(ctx, startDate, endDate, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.APILog), args.Get(1).(int64), args.Error(2)
}

func (m *MockLogRepository) FindByCustomerIDAndPath(ctx context.Context, customerID int64, path string, limit, offset int) ([]*domain.APILog, int64, error) {
	args := m.Called(ctx, customerID, path, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.APILog), args.Get(1).(int64), args.Error(2)
}

func (m *MockLogRepository) DeleteOldLogs(ctx context.Context, beforeDate time.Time) (int64, error) {
	args := m.Called(ctx, beforeDate)
	return args.Get(0).(int64), args.Error(1)
}

func TestPhoneService_GetPhone_Success(t *testing.T) {
	// Setup mocks
	mockThirdPartyService := &MockThirdPartyServiceInterface{}
	mockTransactionRepo := &MockTransactionRepository{}
	mockLogRepo := &MockLogRepository{}

	// Expected phone response
	expectedPhoneResponse := &PhoneResponse{
		PhoneNumber: "+15551234567",
		CountryCode: "US",
		Cost:        0.10,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "test-provider",
	}

	// Setup mock expectations
	mockThirdPartyService.On("GetPhone", mock.Anything, "verification", "any").Return(expectedPhoneResponse, nil)
	mockTransactionRepo.On("GetBalance", mock.Anything, int64(1)).Return(1.00, nil)
	mockTransactionRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Transaction")).Return(nil)
	mockLogRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.APILog")).Return(nil).Maybe() // Make it optional since it runs in goroutine

	// Create service
	service := NewPhoneService(mockThirdPartyService, mockTransactionRepo, mockLogRepo, nil)

	// Test
	ctx := context.Background()
	result, err := service.GetPhone(ctx, 1, "verification", "any")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedPhoneResponse.PhoneNumber, result.PhoneNumber)
	assert.Equal(t, expectedPhoneResponse.CountryCode, result.CountryCode)
	assert.Equal(t, expectedPhoneResponse.Cost, result.Cost)
	assert.Equal(t, expectedPhoneResponse.ProviderID, result.ProviderID)
	assert.True(t, result.ValidUntil.After(time.Now()))

	mockThirdPartyService.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestPhoneService_GetPhone_InsufficientBalance(t *testing.T) {
	// Setup mocks
	mockThirdPartyService := &MockThirdPartyServiceInterface{}
	mockTransactionRepo := &MockTransactionRepository{}
	mockLogRepo := &MockLogRepository{}

	expectedPhoneResponse := &PhoneResponse{
		PhoneNumber: "+15551234567",
		CountryCode: "US",
		Cost:        0.50, // Higher cost than available balance
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "test-provider",
	}

	// Setup mock expectations
	mockThirdPartyService.On("GetPhone", mock.Anything, "verification", "any").Return(expectedPhoneResponse, nil)
	mockTransactionRepo.On("GetBalance", mock.Anything, int64(1)).Return(0.10, nil)                    // Low balance
	mockLogRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.APILog")).Return(nil).Maybe() // Make it optional since it runs in goroutine

	// Create service
	service := NewPhoneService(mockThirdPartyService, mockTransactionRepo, mockLogRepo, nil)

	// Test
	ctx := context.Background()
	_, err := service.GetPhone(ctx, 1, "verification", "any")

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, common.ErrInsufficientBalance, err)

	mockThirdPartyService.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestPhoneService_GetCode_Success(t *testing.T) {
	// Setup mocks
	mockThirdPartyService := &MockThirdPartyServiceInterface{}
	mockTransactionRepo := &MockTransactionRepository{}
	mockLogRepo := &MockLogRepository{}

	phoneNumber := "+15551234567"
	expectedCodeResponse := &CodeResponse{
		Code:       "123456",
		Message:    "Your verification code is 123456",
		ReceivedAt: time.Now(),
		ProviderID: "test-provider",
	}

	// Setup mock expectations
	mockThirdPartyService.On("GetCode", mock.Anything, phoneNumber, 30*time.Second).Return(expectedCodeResponse, nil)
	mockLogRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.APILog")).Return(nil).Maybe() // Make it optional since it runs in goroutine

	// Create service
	service := NewPhoneService(mockThirdPartyService, mockTransactionRepo, mockLogRepo, nil)

	// Test
	ctx := context.Background()
	result, err := service.GetCode(ctx, 1, phoneNumber, 30*time.Second)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedCodeResponse.Code, result.Code)
	assert.Equal(t, expectedCodeResponse.Message, result.Message)
	assert.Equal(t, expectedCodeResponse.ProviderID, result.ProviderID)

	mockThirdPartyService.AssertExpectations(t)
	mockLogRepo.AssertExpectations(t)
}
