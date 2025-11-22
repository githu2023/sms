package providers

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sms-platform/goapi/internal/service"
	"sync"
	"time"
)

// MockProvider implements the service.Provider interface for testing purposes.
type MockProvider struct {
	info                 *service.ProviderInfo
	healthy              bool
	phoneNumbersInUse    map[string]struct{} // Tracks phone numbers currently "assigned"
	codesReceived        map[string]string   // Maps phone number to received code
	mu                   sync.RWMutex
	successRate          int // Percentage (0-100)
	minDelayMs           int
	maxDelayMs           int
	mockBusinessTypes    map[string]struct{}
	mockCardTypes        map[string]struct{}
}

// NewMockProvider creates a new instance of MockProvider.
func NewMockProvider(id string, name string, priority int, successRate int, minDelayMs int, maxDelayMs int) *MockProvider {
	if successRate < 0 || successRate > 100 {
		successRate = 100
	}
	if minDelayMs < 0 {
		minDelayMs = 0
	}
	if maxDelayMs < minDelayMs {
		maxDelayMs = minDelayMs
	}

	return &MockProvider{
		info: &service.ProviderInfo{
			ID:                 id,
			Name:               name,
			Type:               "mock",
			Priority:           priority,
			CostPerSMS:         0.01,
			SupportedCountries: []string{"US", "CN"},
			RateLimit:          1000,
			Timeout:            5 * time.Second,
		},
		healthy:           true,
		phoneNumbersInUse: make(map[string]struct{}),
		codesReceived:     make(map[string]string),
		successRate:       successRate,
		minDelayMs:        minDelayMs,
		maxDelayMs:        maxDelayMs,
		mockBusinessTypes: map[string]struct{}{
			"qq":     {},
			"wechat": {},
			"test":   {},
		},
		mockCardTypes: map[string]struct{}{
			"physical": {},
			"virtual":  {},
		},
	}
}

// GetProviderInfo returns metadata about the mock provider.
func (mp *MockProvider) GetProviderInfo() *service.ProviderInfo {
	return mp.info
}

// IsHealthy returns the health status of the mock provider.
func (mp *MockProvider) IsHealthy(ctx context.Context) bool {
	mp.mu.RLock()
	defer mp.mu.RUnlock()
	return mp.healthy
}

// SetHealthy allows setting the health status for testing.
func (mp *MockProvider) SetHealthy(healthy bool) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.healthy = healthy
}

// simulateLatency simulates network latency.
func (mp *MockProvider) simulateLatency() {
	if mp.minDelayMs > 0 || mp.maxDelayMs > 0 {
		delay := rand.Intn(mp.maxDelayMs-mp.minDelayMs+1) + mp.minDelayMs
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

// simulateFailure determines if a mock operation should fail based on successRate.
func (mp *MockProvider) simulateFailure() bool {
	return rand.Intn(100) >= mp.successRate
}

// GetPhone simulates requesting a phone number.
func (mp *MockProvider) GetPhone(ctx context.Context, customerID int64, businessType, cardType string) (*service.PhoneResponse, error) {
	mp.simulateLatency()

	if mp.simulateFailure() {
		return nil, errors.New("mock provider failed to get phone (simulated error)")
	}

	mp.mu.Lock()
	defer mp.mu.Unlock()

	// Validate business type and card type
	if _, ok := mp.mockBusinessTypes[businessType]; !ok {
		return nil, fmt.Errorf("unsupported business type: %s", businessType)
	}
	if _, ok := mp.mockCardTypes[cardType]; !ok {
		return nil, fmt.Errorf("unsupported card type: %s", cardType)
	}

	// Generate a mock phone number
	var phoneNumber string
	for {
		phoneNumber = fmt.Sprintf("1%010d", rand.Int63n(1e10)) // US-like 11-digit number
		if _, inUse := mp.phoneNumbersInUse[phoneNumber]; !inUse {
			break
		}
	}
	mp.phoneNumbersInUse[phoneNumber] = struct{}{}

	return &service.PhoneResponse{
		PhoneNumber: phoneNumber,
		CountryCode: "US",
		Cost:        mp.info.CostPerSMS,
		ValidUntil:  time.Now().Add(5 * time.Minute), // Mock valid for 5 minutes
		ProviderID:  mp.info.ID,
	}, nil
}

// GetCode simulates retrieving an SMS verification code.
func (mp *MockProvider) GetCode(ctx context.Context, customerID int64, phoneNumber string, timeout time.Duration) (*service.CodeResponse, error) {
	mp.simulateLatency()

	if mp.simulateFailure() {
		return nil, errors.New("mock provider failed to get code (simulated error)")
	}

	mp.mu.RLock()
	// Check if the phone number was "assigned" by this mock provider
	if _, inUse := mp.phoneNumbersInUse[phoneNumber]; !inUse {
		mp.mu.RUnlock()
		return nil, fmt.Errorf("phone number %s not assigned by this mock provider", phoneNumber)
	}
	mp.mu.RUnlock()

	// Simulate waiting for code (long polling)
	codeChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		// Simulate random delay before code arrives (or doesn't)
		time.Sleep(time.Duration(rand.Intn(int(timeout.Seconds()/2))) * time.Second) // Code arrives within half the timeout

		mp.mu.Lock()
		defer mp.mu.Unlock()

		if true { // Always succeed for testing purposes
			code := fmt.Sprintf("%06d", rand.Intn(1000000))
			mp.codesReceived[phoneNumber] = code
			codeChan <- code
		} else {
			errChan <- errors.New("simulated: code not received within expected time")
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case code := <-codeChan:
		return &service.CodeResponse{
			Code:       code,
			Message:    "Mock code received",
			ReceivedAt: time.Now(),
			ProviderID: mp.info.ID,
		}, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(timeout):
		return nil, errors.New("mock provider: get code timed out")
	}
}
