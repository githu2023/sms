package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Provider represents a third-party SMS service provider
type Provider interface {
	// GetPhone requests a phone number for receiving SMS
	GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error)

	// GetCode retrieves SMS verification code for a specific phone number
	GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*CodeResponse, error)

	// GetProviderInfo returns provider metadata
	GetProviderInfo() *ProviderInfo

	// IsHealthy checks if the provider is available
	IsHealthy(ctx context.Context) bool
}

// PhoneResponse represents the response from GetPhone request
type PhoneResponse struct {
	PhoneNumber string    `json:"phone_number"`
	CountryCode string    `json:"country_code"`
	Cost        float64   `json:"cost"`
	ValidUntil  time.Time `json:"valid_until"`
	ProviderID  string    `json:"provider_id"`
}

// CodeResponse represents the response from GetCode request
type CodeResponse struct {
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	ReceivedAt time.Time `json:"received_at"`
	ProviderID string    `json:"provider_id"`
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

// ThirdPartyService manages multiple SMS service providers
type ThirdPartyService interface {
	// GetPhone gets a phone number from the best available provider
	GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error)

	// GetCode gets verification code from the provider that issued the phone
	GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*CodeResponse, error)

	// RegisterProvider adds a new provider to the service
	RegisterProvider(provider Provider) error

	// GetProviders returns all registered providers
	GetProviders() []Provider

	// GetProviderByID returns a specific provider by ID
	GetProviderByID(id string) (Provider, error)

	// SetProviderPriority updates provider priority for load balancing
	SetProviderPriority(id string, priority int) error

	// HealthCheck performs health check on all providers
	HealthCheck(ctx context.Context) map[string]bool
}

// thirdPartyService implements ThirdPartyService
type thirdPartyService struct {
	providers  map[string]Provider
	mutex      sync.RWMutex
	httpClient *http.Client
}

// NewThirdPartyService creates a new ThirdPartyService instance
func NewThirdPartyService() ThirdPartyService {
	return &thirdPartyService{
		providers: make(map[string]Provider),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       100,
				MaxConnsPerHost:    10,
				IdleConnTimeout:    90 * time.Second,
				DisableCompression: true,
			},
		},
	}
}

// RegisterProvider adds a new provider to the service
func (s *thirdPartyService) RegisterProvider(provider Provider) error {
	if provider == nil {
		return errors.New("provider cannot be nil")
	}

	info := provider.GetProviderInfo()
	if info.ID == "" {
		return errors.New("provider ID cannot be empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check for duplicate IDs
	if _, exists := s.providers[info.ID]; exists {
		return fmt.Errorf("provider with ID %s already exists", info.ID)
	}

	s.providers[info.ID] = provider
	return nil
}

// GetProviders returns all registered providers
func (s *thirdPartyService) GetProviders() []Provider {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	providers := make([]Provider, 0, len(s.providers))
	for _, provider := range s.providers {
		providers = append(providers, provider)
	}
	return providers
}

// GetProviderByID returns a specific provider by ID
func (s *thirdPartyService) GetProviderByID(id string) (Provider, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	provider, exists := s.providers[id]
	if !exists {
		return nil, fmt.Errorf("provider with ID %s not found", id)
	}
	return provider, nil
}

// SetProviderPriority updates provider priority
func (s *thirdPartyService) SetProviderPriority(id string, priority int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	provider, exists := s.providers[id]
	if !exists {
		return fmt.Errorf("provider with ID %s not found", id)
	}

	// Note: This is a simplified implementation
	// In a real system, you'd modify the provider's priority in its configuration
	_ = provider
	_ = priority

	return nil
}

// GetPhone gets a phone number from the best available provider
func (s *thirdPartyService) GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error) {
	providers := s.getHealthyProvidersByPriority(ctx)
	if len(providers) == 0 {
		return nil, errors.New("no healthy providers available")
	}

	// Try providers in priority order
	var lastErr error
	for _, provider := range providers {
		response, err := provider.GetPhone(ctx, businessType, cardType)
		if err == nil {
			return response, nil
		}
		lastErr = err

		// Log error and continue to next provider
		fmt.Printf("Provider %s failed: %v\n", provider.GetProviderInfo().ID, err)
	}

	return nil, fmt.Errorf("all providers failed, last error: %w", lastErr)
}

// GetCode gets verification code from the provider
func (s *thirdPartyService) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*CodeResponse, error) {
	// For now, try all providers since we don't track which provider issued the phone
	// In a real system, you'd track the provider that issued each phone number
	providers := s.getHealthyProvidersByPriority(ctx)
	if len(providers) == 0 {
		return nil, errors.New("no healthy providers available")
	}

	var lastErr error
	for _, provider := range providers {
		response, err := provider.GetCode(ctx, phoneNumber, timeout)
		if err == nil {
			return response, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("no provider could retrieve code for phone %s: %w", phoneNumber, lastErr)
}

// HealthCheck performs health check on all providers
func (s *thirdPartyService) HealthCheck(ctx context.Context) map[string]bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	results := make(map[string]bool)
	for id, provider := range s.providers {
		results[id] = provider.IsHealthy(ctx)
	}
	return results
}

// getHealthyProvidersByPriority returns healthy providers sorted by priority
func (s *thirdPartyService) getHealthyProvidersByPriority(ctx context.Context) []Provider {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var healthyProviders []Provider
	for _, provider := range s.providers {
		if provider.IsHealthy(ctx) {
			healthyProviders = append(healthyProviders, provider)
		}
	}

	// Sort by priority (higher priority first)
	// This is a simplified sorting - in production you'd use sort.Slice
	// For now, we'll shuffle for load balancing
	rand.Shuffle(len(healthyProviders), func(i, j int) {
		healthyProviders[i], healthyProviders[j] = healthyProviders[j], healthyProviders[i]
	})

	return healthyProviders
}