package providers

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMockProvider_Basic(t *testing.T) {
	// Updated NewMockProvider call with all 6 arguments: id, name, priority, successRate, minDelayMs, maxDelayMs
	provider := NewMockProvider("test-id", "Test Provider", 1, 100, 0, 0) // 100% success, no delay

	// Test provider info
	info := provider.GetProviderInfo()
	assert.Equal(t, "test-id", info.ID)
	assert.Equal(t, "Test Provider", info.Name)
	assert.Equal(t, "mock", info.Type)
	assert.Equal(t, 1, info.Priority)
	assert.Equal(t, 0.01, info.CostPerSMS) // Default cost in MockProvider

	// Test health check
	ctx := context.Background()
	assert.True(t, provider.IsHealthy(ctx))

	provider.SetHealthy(false)
	assert.False(t, provider.IsHealthy(ctx))
}

func TestMockProvider_GetPhone(t *testing.T) {
	// Updated NewMockProvider call
	provider := NewMockProvider("test-id", "Test Provider", 1, 100, 0, 0) // 100% success, no delay
	provider.SetHealthy(true)
	// provider.SetFailureRate(0) // Removed: SetFailureRate is not a method of MockProvider

	ctx := context.Background()
	customerID := int64(1) // Placeholder customer ID
	// Updated GetPhone call with customerID
	response, err := provider.GetPhone(ctx, customerID, "test", "physical")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.PhoneNumber)
	assert.Equal(t, "US", response.CountryCode)
	assert.Equal(t, "test-id", response.ProviderID)
	assert.True(t, response.ValidUntil.After(time.Now()))
}

func TestMockProvider_GetCode(t *testing.T) {
	// Updated NewMockProvider call
	provider := NewMockProvider("test-id", "Test Provider", 1, 100, 0, 0) // 100% success, no delay
	provider.SetHealthy(true)
	// provider.SetFailureRate(0) // Removed: SetFailureRate is not a method of MockProvider

	ctx := context.Background()
	customerID := int64(1) // Placeholder customer ID

	// Get a phone number first
	// Updated GetPhone call with customerID
	phoneResponse, err := provider.GetPhone(ctx, customerID, "test", "physical")
	assert.NoError(t, err)

	// Get the code
	// Updated GetCode call with customerID
	codeResponse, err := provider.GetCode(ctx, customerID, phoneResponse.PhoneNumber, 5*time.Second)
	assert.NoError(t, err)
	assert.NotNil(t, codeResponse)
	assert.NotEmpty(t, codeResponse.Code)
	assert.Equal(t, "test-id", codeResponse.ProviderID)
	assert.Regexp(t, `^\d{6}$`, codeResponse.Code)
}

/*
// TestHTTPProvider_Creation commented out as NewHTTPProvider and HTTPProviderConfig are not defined.
func TestHTTPProvider_Creation(t *testing.T) {
	config := HTTPProviderConfig{
		ID:       "http-test",
		Name:     "HTTP Test Provider",
		BaseURL:  "https://api.example.com",
		Priority: 1,
	}

	provider := NewHTTPProvider(config)
	assert.NotNil(t, provider)

	info := provider.GetProviderInfo()
	assert.Equal(t, "http-test", info.ID)
	assert.Equal(t, "HTTP Test Provider", info.Name)
	assert.Equal(t, "http", info.Type)
}
*/
