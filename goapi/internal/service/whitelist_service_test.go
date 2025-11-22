package service

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockWhitelistRepository is a mock implementation of the WhitelistRepository interface
type MockWhitelistRepository struct {
	mock.Mock
}

func (m *MockWhitelistRepository) Create(ctx context.Context, whitelist *domain.IPWhitelist) error {
	args := m.Called(ctx, whitelist)
	return args.Error(0)
}

func (m *MockWhitelistRepository) FindByID(ctx context.Context, id int64) (*domain.IPWhitelist, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.IPWhitelist), args.Error(1)
}

func (m *MockWhitelistRepository) FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.IPWhitelist, int64, error) {
	args := m.Called(ctx, customerID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.IPWhitelist), args.Get(1).(int64), args.Error(2)
}

func (m *MockWhitelistRepository) FindByCustomerIDAndIP(ctx context.Context, customerID int64, ipAddress string) (*domain.IPWhitelist, error) {
	args := m.Called(ctx, customerID, ipAddress)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.IPWhitelist), args.Error(1)
}

func (m *MockWhitelistRepository) Update(ctx context.Context, whitelist *domain.IPWhitelist) error {
	args := m.Called(ctx, whitelist)
	return args.Error(0)
}

func (m *MockWhitelistRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockWhitelistRepository) DeleteByCustomerIDAndIP(ctx context.Context, customerID int64, ipAddress string) error {
	args := m.Called(ctx, customerID, ipAddress)
	return args.Error(0)
}

// TestAddWhitelist tests adding a new whitelist entry with valid IP
func TestAddWhitelist(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.1"
	notes := "Office IP"

	mockRepo.On("Create", ctx, mock.MatchedBy(func(w *domain.IPWhitelist) bool {
		return w.CustomerID == customerID && w.IPAddress == ipAddress && w.Notes == notes
	})).Return(nil)

	err := service.AddWhitelist(ctx, customerID, ipAddress, notes)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestAddWhitelist_WithCIDR tests adding a whitelist entry with valid CIDR
func TestAddWhitelist_WithCIDR(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.0/24"
	notes := "Office Network"

	mockRepo.On("Create", ctx, mock.MatchedBy(func(w *domain.IPWhitelist) bool {
		return w.CustomerID == customerID && w.IPAddress == ipAddress
	})).Return(nil)

	err := service.AddWhitelist(ctx, customerID, ipAddress, notes)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestAddWhitelist_InvalidIP tests adding a whitelist with invalid IP
func TestAddWhitelist_InvalidIP(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "invalid-ip-address"
	notes := "Invalid IP"

	err := service.AddWhitelist(ctx, customerID, ipAddress, notes)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidIPFormat, err)
	mockRepo.AssertNotCalled(t, "Create")
}

// TestAddWhitelist_InvalidCIDR tests adding a whitelist with invalid CIDR
func TestAddWhitelist_InvalidCIDR(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.0/33" // Invalid CIDR (mask > 32)
	notes := "Invalid CIDR"

	err := service.AddWhitelist(ctx, customerID, ipAddress, notes)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidIPFormat, err)
	mockRepo.AssertNotCalled(t, "Create")
}

// TestDeleteWhitelist tests deleting a whitelist entry
func TestDeleteWhitelist(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.1"

	mockRepo.On("DeleteByCustomerIDAndIP", ctx, customerID, ipAddress).Return(nil)

	err := service.DeleteWhitelist(ctx, customerID, ipAddress)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestDeleteWhitelist_NotFound tests deleting a non-existent whitelist entry
func TestDeleteWhitelist_NotFound(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.1"

	mockRepo.On("DeleteByCustomerIDAndIP", ctx, customerID, ipAddress).Return(gorm.ErrRecordNotFound)

	err := service.DeleteWhitelist(ctx, customerID, ipAddress)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// TestListWhitelists tests listing whitelists with pagination
func TestListWhitelists(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	page := 1
	limit := 20

	whitelists := []*domain.IPWhitelist{
		{
			ID:         1,
			CustomerID: customerID,
			IPAddress:  "192.168.1.1",
			Notes:      "Office",
			CreatedAt:  time.Now(),
		},
		{
			ID:         2,
			CustomerID: customerID,
			IPAddress:  "192.168.1.0/24",
			Notes:      "Network",
			CreatedAt:  time.Now(),
		},
	}

	mockRepo.On("FindByCustomerID", ctx, customerID, limit, (page-1)*limit).Return(whitelists, int64(2), nil)

	list, total, err := service.ListWhitelists(ctx, customerID, page, limit)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, list, 2)
	assert.Equal(t, "192.168.1.1", list[0].IPAddress)
	assert.Equal(t, "192.168.1.0/24", list[1].IPAddress)
	mockRepo.AssertExpectations(t)
}

// TestListWhitelists_Empty tests listing whitelists when none exist
func TestListWhitelists_Empty(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	page := 1
	limit := 20

	mockRepo.On("FindByCustomerID", ctx, customerID, limit, (page-1)*limit).Return([]*domain.IPWhitelist{}, int64(0), nil)

	list, total, err := service.ListWhitelists(ctx, customerID, page, limit)

	assert.NoError(t, err)
	assert.Equal(t, int64(0), total)
	assert.Len(t, list, 0)
	mockRepo.AssertExpectations(t)
}

// TestListWhitelists_DatabaseError tests listing whitelists when database error occurs
func TestListWhitelists_DatabaseError(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	page := 1
	limit := 20

	dbErr := errors.New("database connection error")
	mockRepo.On("FindByCustomerID", ctx, customerID, limit, (page-1)*limit).Return(nil, int64(0), dbErr)

	list, total, err := service.ListWhitelists(ctx, customerID, page, limit)

	assert.Error(t, err)
	assert.Equal(t, dbErr, err)
	assert.Nil(t, list)
	assert.Equal(t, int64(0), total)
	mockRepo.AssertExpectations(t)
}

// TestGetWhitelist tests getting a specific whitelist entry
func TestGetWhitelist(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.1"

	whitelist := &domain.IPWhitelist{
		ID:         1,
		CustomerID: customerID,
		IPAddress:  ipAddress,
		Notes:      "Office",
		CreatedAt:  time.Now(),
	}

	mockRepo.On("FindByCustomerIDAndIP", ctx, customerID, ipAddress).Return(whitelist, nil)

	result, err := service.GetWhitelist(ctx, customerID, ipAddress)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, ipAddress, result.IPAddress)
	assert.Equal(t, "Office", result.Notes)
	mockRepo.AssertExpectations(t)
}

// TestGetWhitelist_NotFound tests getting a non-existent whitelist entry
func TestGetWhitelist_NotFound(t *testing.T) {
	mockRepo := new(MockWhitelistRepository)
	service := NewWhitelistService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	ipAddress := "192.168.1.1"

	mockRepo.On("FindByCustomerIDAndIP", ctx, customerID, ipAddress).Return(nil, gorm.ErrRecordNotFound)

	result, err := service.GetWhitelist(ctx, customerID, ipAddress)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// TestIsValidIPOrCIDR tests IP and CIDR validation
func TestIsValidIPOrCIDR(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid IPv4", "192.168.1.1", true},
		{"Valid IPv4 CIDR", "192.168.1.0/24", true},
		{"Valid IPv6", "2001:db8::1", true},
		{"Valid IPv6 CIDR", "2001:db8::/32", true},
		{"Invalid IP", "invalid", false},
		{"Invalid CIDR", "192.168.1.0/33", false},
		{"Empty string", "", false},
		{"Partial IP", "192.168.1", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidIPOrCIDR(tc.input)
			assert.Equal(t, tc.expected, result, "IP validation failed for %s", tc.input)
		})
	}
}
