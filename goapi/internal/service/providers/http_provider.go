package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"sms-platform/goapi/internal/service"
)

// HTTPProvider implements Provider interface for HTTP-based SMS providers
type HTTPProvider struct {
	info       *service.ProviderInfo
	httpClient *http.Client
	baseURL    string
	apiKey     string
	headers    map[string]string
}

// HTTPProviderConfig holds configuration for HTTP provider
type HTTPProviderConfig struct {
	ID         string
	Name       string
	BaseURL    string
	APIKey     string
	Priority   int
	CostPerSMS float64
	Timeout    time.Duration
	Headers    map[string]string
}

// NewHTTPProvider creates a new HTTP-based provider
func NewHTTPProvider(config HTTPProviderConfig) *HTTPProvider {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	if config.Headers == nil {
		config.Headers = make(map[string]string)
	}

	// Set default headers
	config.Headers["Content-Type"] = "application/json"
	config.Headers["User-Agent"] = "SMS-Platform/1.0"

	if config.APIKey != "" {
		config.Headers["Authorization"] = "Bearer " + config.APIKey
	}

	return &HTTPProvider{
		info: &service.ProviderInfo{
			ID:                 config.ID,
			Name:               config.Name,
			Type:               "http",
			Priority:           config.Priority,
			CostPerSMS:         config.CostPerSMS,
			SupportedCountries: []string{"US", "CN", "GB", "FR", "DE", "JP"},
			RateLimit:          60, // 60 requests per minute
			Timeout:            config.Timeout,
			Metadata: map[string]string{
				"base_url": config.BaseURL,
				"version":  "1.0.0",
			},
		},
		httpClient: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				MaxIdleConns:       100,
				MaxConnsPerHost:    10,
				IdleConnTimeout:    90 * time.Second,
				DisableCompression: true,
			},
		},
		baseURL: config.BaseURL,
		apiKey:  config.APIKey,
		headers: config.Headers,
	}
}

// GetProviderInfo returns provider metadata
func (p *HTTPProvider) GetProviderInfo() *service.ProviderInfo {
	return p.info
}

// IsHealthy checks if the provider is available
func (p *HTTPProvider) IsHealthy(ctx context.Context) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", p.baseURL+"/health", nil)
	if err != nil {
		return false
	}

	// Add headers
	for key, value := range p.headers {
		req.Header.Set(key, value)
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// GetPhone requests a phone number for receiving SMS
func (p *HTTPProvider) GetPhone(ctx context.Context, businessType, cardType string) (*service.PhoneResponse, error) {
	reqBody := map[string]interface{}{
		"business_type": businessType,
		"card_type":     cardType,
		"country":       "US",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/get_phone", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	for key, value := range p.headers {
		req.Header.Set(key, value)
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var response struct {
		PhoneNumber  string  `json:"phone_number"`
		CountryCode  string  `json:"country_code"`
		Cost         float64 `json:"cost"`
		ValidMinutes int     `json:"valid_minutes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &service.PhoneResponse{
		PhoneNumber: response.PhoneNumber,
		CountryCode: response.CountryCode,
		Cost:        response.Cost,
		ValidUntil:  time.Now().Add(time.Duration(response.ValidMinutes) * time.Minute),
		ProviderID:  p.info.ID,
	}, nil
}

// GetCode retrieves SMS verification code for a specific phone number
func (p *HTTPProvider) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration) (*service.CodeResponse, error) {
	reqBody := map[string]interface{}{
		"phone_number":    phoneNumber,
		"timeout_seconds": int(timeout.Seconds()),
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/get_code", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	for key, value := range p.headers {
		req.Header.Set(key, value)
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var response struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Timestamp int64  `json:"timestamp"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &service.CodeResponse{
		Code:       response.Code,
		Message:    response.Message,
		ReceivedAt: time.Unix(response.Timestamp, 0),
		ProviderID: p.info.ID,
	}, nil
}
