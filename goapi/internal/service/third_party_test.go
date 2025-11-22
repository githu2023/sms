package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProvider for testing
type MockProvider struct {
	mock.Mock
	id       string
	name     string
	healthy  bool
	priority int
	cost     float64
}

func NewMockProvider(id, name string, priority int, cost float64) *MockProvider {
	return &MockProvider{
		id:       id,
		name:     name,
		healthy:  true,
		priority: priority,
		cost:     cost,
	}
}

func (m *MockProvider) GetProviderInfo() *ProviderInfo {
	return &ProviderInfo{
		ID:                 m.id,
		Name:               m.name,
		Type:               "mock",
		Priority:           m.priority,
		CostPerSMS:         m.cost,
		SupportedCountries: []string{"US", "CN"},
		RateLimit:          60,
		Timeout:            30 * time.Second,
		Metadata:           map[string]string{"version": "1.0.0"},
	}
}

func (m *MockProvider) IsHealthy(ctx context.Context) bool {
	return m.healthy
}

func (m *MockProvider) GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error) {
	args := m.Called(ctx, businessType, cardType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*PhoneResponse), args.Error(1)
}

func (m *MockProvider) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*CodeResponse, error) {
	args := m.Called(ctx, phoneNumber, timeout)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CodeResponse), args.Error(1)
}

func (m *MockProvider) SetHealthy(healthy bool) {
	m.healthy = healthy
}

func TestThirdPartyService_RegisterProvider(t *testing.T) {
	service := NewThirdPartyService()

	// Test successful registration
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	err := service.RegisterProvider(mockProvider)
	assert.NoError(t, err)

	// Test duplicate registration
	duplicate := NewMockProvider("test-provider", "Duplicate Provider", 2, 0.15)
	err = service.RegisterProvider(duplicate)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")

	// Test nil provider
	err = service.RegisterProvider(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be nil")
}

func TestThirdPartyService_GetProviders(t *testing.T) {
	service := NewThirdPartyService()

	// Initially empty
	providers := service.GetProviders()
	assert.Empty(t, providers)

	// Register a provider
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	service.RegisterProvider(mockProvider)

	// Should have one provider
	providers = service.GetProviders()
	assert.Len(t, providers, 1)
	assert.Equal(t, "test-provider", providers[0].GetProviderInfo().ID)
}

func TestThirdPartyService_GetProviderByID(t *testing.T) {
	service := NewThirdPartyService()
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	service.RegisterProvider(mockProvider)

	// Test existing provider
	foundProvider, err := service.GetProviderByID("test-provider")
	assert.NoError(t, err)
	assert.NotNil(t, foundProvider)
	assert.Equal(t, "test-provider", foundProvider.GetProviderInfo().ID)

	// Test non-existing provider
	_, err = service.GetProviderByID("non-existent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestThirdPartyService_HealthCheck(t *testing.T) {
	service := NewThirdPartyService()

	// Register healthy provider
	healthyProvider := NewMockProvider("healthy", "Healthy Provider", 1, 0.10)
	healthyProvider.SetHealthy(true)
	service.RegisterProvider(healthyProvider)

	// Register unhealthy provider
	unhealthyProvider := NewMockProvider("unhealthy", "Unhealthy Provider", 2, 0.15)
	unhealthyProvider.SetHealthy(false)
	service.RegisterProvider(unhealthyProvider)

	// Perform health check
	ctx := context.Background()
	results := service.HealthCheck(ctx)

	assert.Len(t, results, 2)
	assert.True(t, results["healthy"])
	assert.False(t, results["unhealthy"])
}

func TestThirdPartyService_GetPhone_Success(t *testing.T) {
	service := NewThirdPartyService()
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	mockProvider.SetHealthy(true)

	// Mock successful phone response
	expectedResponse := &PhoneResponse{
		PhoneNumber: "+15551234567",
		CountryCode: "US",
		Cost:        0.10,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "test-provider",
	}
	mockProvider.On("GetPhone", mock.Anything, "verification", "any").Return(expectedResponse, nil)

	service.RegisterProvider(mockProvider)

	ctx := context.Background()
	response, err := service.GetPhone(ctx, "verification", "any")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedResponse.PhoneNumber, response.PhoneNumber)
	assert.Equal(t, expectedResponse.ProviderID, response.ProviderID)
	mockProvider.AssertExpectations(t)
}

func TestThirdPartyService_GetPhone_NoProviders(t *testing.T) {
	service := NewThirdPartyService()

	ctx := context.Background()
	_, err := service.GetPhone(ctx, "verification", "any")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no healthy providers")
}

func TestThirdPartyService_GetPhone_AllProvidersUnhealthy(t *testing.T) {
	service := NewThirdPartyService()
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	mockProvider.SetHealthy(false) // Unhealthy
	service.RegisterProvider(mockProvider)

	ctx := context.Background()
	_, err := service.GetPhone(ctx, "verification", "any")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no healthy providers")
}

func TestThirdPartyService_GetCode_Success(t *testing.T) {
	service := NewThirdPartyService()
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	mockProvider.SetHealthy(true)

	phoneNumber := "+15551234567"
	expectedCodeResponse := &CodeResponse{
		Code:       "123456",
		Message:    "Your verification code is 123456",
		ReceivedAt: time.Now(),
		ProviderID: "test-provider",
	}

	mockProvider.On("GetCode", mock.Anything, phoneNumber, 5*time.Second).Return(expectedCodeResponse, nil)
	service.RegisterProvider(mockProvider)

	ctx := context.Background()
	response, err := service.GetCode(ctx, phoneNumber, 5*time.Second)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedCodeResponse.Code, response.Code)
	assert.Equal(t, expectedCodeResponse.ProviderID, response.ProviderID)
	mockProvider.AssertExpectations(t)
}

func TestThirdPartyService_GetCode_PhoneNotFound(t *testing.T) {
	service := NewThirdPartyService()
	mockProvider := NewMockProvider("test-provider", "Test Provider", 1, 0.10)
	mockProvider.SetHealthy(true)

	phoneNumber := "+15551234567"
	mockProvider.On("GetCode", mock.Anything, phoneNumber, 5*time.Second).Return(nil, assert.AnError)
	service.RegisterProvider(mockProvider)

	ctx := context.Background()
	_, err := service.GetCode(ctx, phoneNumber, 5*time.Second)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no provider could retrieve code")
	mockProvider.AssertExpectations(t)
}

func TestThirdPartyService_MultipleProviders_LoadBalancing(t *testing.T) {
	service := NewThirdPartyService()

	// Register multiple healthy providers
	provider1 := NewMockProvider("provider-1", "Provider 1", 1, 0.10)
	provider1.SetHealthy(true)
	expectedResponse1 := &PhoneResponse{
		PhoneNumber: "+15551111111",
		CountryCode: "US",
		Cost:        0.10,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "provider-1",
	}
	provider1.On("GetPhone", mock.Anything, "verification", "any").Return(expectedResponse1, nil).Maybe()
	service.RegisterProvider(provider1)

	provider2 := NewMockProvider("provider-2", "Provider 2", 2, 0.15)
	provider2.SetHealthy(true)
	expectedResponse2 := &PhoneResponse{
		PhoneNumber: "+15552222222",
		CountryCode: "US",
		Cost:        0.15,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "provider-2",
	}
	provider2.On("GetPhone", mock.Anything, "verification", "any").Return(expectedResponse2, nil).Maybe()
	service.RegisterProvider(provider2)

	ctx := context.Background()

	// Make a request to ensure it works with multiple providers
	response, err := service.GetPhone(ctx, "verification", "any")
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Contains(t, []string{"provider-1", "provider-2"}, response.ProviderID)
}

func TestThirdPartyService_ProviderFailover(t *testing.T) {
	service := NewThirdPartyService()

	// First provider that always fails
	failingProvider := NewMockProvider("failing-provider", "Failing Provider", 1, 0.10)
	failingProvider.SetHealthy(true)
	failingProvider.On("GetPhone", mock.Anything, "verification", "any").Return(nil, assert.AnError).Maybe()
	service.RegisterProvider(failingProvider)

	// Second provider that works
	workingProvider := NewMockProvider("working-provider", "Working Provider", 2, 0.15)
	workingProvider.SetHealthy(true)
	expectedResponse := &PhoneResponse{
		PhoneNumber: "+15551234567",
		CountryCode: "US",
		Cost:        0.15,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "working-provider",
	}
	workingProvider.On("GetPhone", mock.Anything, "verification", "any").Return(expectedResponse, nil).Maybe()
	service.RegisterProvider(workingProvider)

	ctx := context.Background()
	response, err := service.GetPhone(ctx, "verification", "any")

	// Should succeed with one of the providers
	// Due to random shuffling, we can't predict which will be called first
	// But the request should succeed
	assert.NoError(t, err)
	assert.NotNil(t, response)
}
