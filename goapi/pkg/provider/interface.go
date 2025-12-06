package provider

import (
	"context"
	"fmt"
	"time"
)

// SMSProvider represents a third-party SMS service provider interface
// This interface defines the contract that all SMS providers must implement
type SMSProvider interface {
	// GetPhone requests a phone number for receiving SMS
	// businessType: the type of business service (e.g., "qq", "wechat", "test")
	// cardType: the type of card ("physical" or "virtual")
	GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error)

	// GetCode retrieves SMS verification code for a specific phone number
	// phoneNumber: the phone number to receive SMS
	// timeout: maximum time to wait for the code
	// extId: optional external ID from provider (e.g., extId for BigBus666), if provided, will be used instead of phoneNumber lookup
	GetCode(ctx context.Context, phoneNumber string, timeout time.Duration, extId ...string) (*CodeResponse, error)

	// GetProviderInfo returns provider metadata and configuration
	GetProviderInfo() *ProviderInfo

	// IsHealthy checks if the provider is currently available and functional
	IsHealthy(ctx context.Context) bool

	// SetHealthy allows manual control of provider health status (primarily for testing)
	SetHealthy(healthy bool)

	// release phone number
	// phoneNumber: the phone number to release
	// extId: optional external ID from provider, if provided, will be used instead of phoneNumber lookup
	ReleasePhone(ctx context.Context, phoneNumber string, extId ...string) error
}

// PhoneResponse represents the response from GetPhone request
type PhoneResponse struct {
	PhoneNumber string    `json:"phone_number"` // The assigned phone number
	CountryCode string    `json:"country_code"` // Country code (e.g., "US", "CN")
	Cost        float64   `json:"cost"`         // Cost for this phone number
	ValidUntil  time.Time `json:"valid_until"`  // When this assignment expires
	ProviderID  string    `json:"provider_id"`  // ID of the provider that supplied this number
	ExtId       string    `json:"ext_id"`       // External ID from provider (e.g., extId for BigBus666), used for getting code and releasing phone
}

// CodeResponse represents the response from GetCode request
type CodeResponse struct {
	Code       string    `json:"code"`        // The verification code received
	Message    string    `json:"message"`     // Additional message or status
	ReceivedAt time.Time `json:"received_at"` // When the code was received
	ProviderID string    `json:"provider_id"` // ID of the provider
}

// ProviderInfo contains metadata about a provider
type ProviderInfo struct {
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	Type               string            `json:"type"` // "mock", "http", "api"
	Priority           int               `json:"priority"`
	CostPerSMS         float64           `json:"cost_per_sms"`
	SupportedCountries []string          `json:"supported_countries"`
	RateLimit          int               `json:"rate_limit"` // requests per minute
	Timeout            time.Duration     `json:"timeout"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}

// ProviderError represents a structured error from SMS providers
type ProviderError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *ProviderError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewProviderError creates a new provider error
func NewProviderError(code, message string) *ProviderError {
	return &ProviderError{
		Code:    code,
		Message: message,
	}
}

// IsProviderError checks if an error is a ProviderError
func IsProviderError(err error) bool {
	_, ok := err.(*ProviderError)
	return ok
}

// Common error types for SMS providers
var (
	ErrProviderUnavailable  = NewProviderError("PROVIDER_UNAVAILABLE", "Provider is currently unavailable")
	ErrUnsupportedService   = NewProviderError("UNSUPPORTED_SERVICE", "Business type not supported by provider")
	ErrNoPhoneAvailable     = NewProviderError("NO_PHONE_AVAILABLE", "No phone numbers available")
	ErrCodeTimeout          = NewProviderError("CODE_TIMEOUT", "Timeout waiting for verification code")
	ErrCodeNotReceived      = NewProviderError("CODE_NOT_RECEIVED", "Verification code not received yet")
	ErrPhoneAlreadyReleased = NewProviderError("PHONE_ALREADY_RELEASED", "Phone number has already been released by provider")
	ErrInvalidPhoneNumber   = NewProviderError("INVALID_PHONE", "Invalid phone number")
	ErrProviderError        = NewProviderError("PROVIDER_ERROR", "Provider internal error")
)
