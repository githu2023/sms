package service

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/domain"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPhoneAssignmentRepository for testing
type MockPhoneAssignmentRepository struct {
	mock.Mock
}

func (m *MockPhoneAssignmentRepository) Create(ctx context.Context, tx *gorm.DB, assignment *domain.PhoneAssignment) error {
	args := m.Called(ctx, tx, assignment)
	return args.Error(0)
}

func (m *MockPhoneAssignmentRepository) FindByPhone(ctx context.Context, tx *gorm.DB, phone string) (*domain.PhoneAssignment, error) {
	args := m.Called(ctx, tx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PhoneAssignment), args.Error(1)
}

func (m *MockPhoneAssignmentRepository) FindActiveByCustomerIDAndPhone(ctx context.Context, tx *gorm.DB, customerID int64, phone string) (*domain.PhoneAssignment, error) {
	args := m.Called(ctx, tx, customerID, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PhoneAssignment), args.Error(1)
}

func (m *MockPhoneAssignmentRepository) FindRecentByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.PhoneAssignment, int64, error) {
	args := m.Called(ctx, customerID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.PhoneAssignment), args.Get(1).(int64), args.Error(2)
}

func (m *MockPhoneAssignmentRepository) Update(ctx context.Context, tx *gorm.DB, assignment *domain.PhoneAssignment) error {
	args := m.Called(ctx, tx, assignment)
	return args.Error(0)
}

func (m *MockPhoneAssignmentRepository) UpdateVerificationCode(ctx context.Context, tx *gorm.DB, id int64, code string, receivedAt time.Time) error {
	args := m.Called(ctx, tx, id, code, receivedAt)
	return args.Error(0)
}

func (m *MockPhoneAssignmentRepository) UpdateStatus(ctx context.Context, tx *gorm.DB, id int64, status int) error {
	args := m.Called(ctx, tx, id, status)
	return args.Error(0)
}

func (m *MockPhoneAssignmentRepository) FindExpiredAssignments(ctx context.Context, limit int) ([]*domain.PhoneAssignment, error) {
	args := m.Called(ctx, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.PhoneAssignment), args.Error(1)
}

func (m *MockPhoneAssignmentRepository) CountByCustomerID(ctx context.Context, customerID int64) (int64, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).(int64), args.Error(1)
}

// Tests

func TestAssignmentService_GetAssignments_Success(t *testing.T) {
	// Setup
	mockAssignmentRepo := &MockPhoneAssignmentRepository{}

	service := NewAssignmentService(mockAssignmentRepo, nil, nil)

	// Mock data
	now := time.Now()
	expireTime := now.Add(30 * time.Minute)
	assignments := []*domain.PhoneAssignment{
		{
			ID:             1,
			CustomerID:     1,
			ProviderID:     "1",
			BusinessTypeID: 1,
			CardType:       "virtual",
			PhoneNumber:    "+15551234567",
			Cost:           0.10,
			Status:         2,
			CreatedAt:      now,
			ExpiresAt:      &expireTime,
		},
	}

	// Setup expectations
	mockAssignmentRepo.On("FindRecentByCustomerID", mock.Anything, int64(1), 20, 0).Return(assignments, int64(1), nil)

	// Test
	result, total, err := service.GetAssignments(context.Background(), 1, 1, 20, 0, "", nil, nil)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, int64(1), total)

	// Check assignment
	assignment := result[0]
	assert.Equal(t, int64(1), assignment.ID)
	assert.Equal(t, "+15551234567", assignment.PhoneNumber)

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_GetAssignments_RepositoryError(t *testing.T) {
	// Setup
	mockAssignmentRepo := &MockPhoneAssignmentRepository{}

	service := NewAssignmentService(mockAssignmentRepo, nil, nil)

	// Setup expectations
	mockAssignmentRepo.On("FindRecentByCustomerID", mock.Anything, int64(1), 20, 0).Return(nil, int64(0), errors.New("database error"))

	// Test
	_, _, err := service.GetAssignments(context.Background(), 1, 1, 20, 0, "", nil, nil)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_GetCostStatistics_Success(t *testing.T) {
	// Setup
	mockAssignmentRepo := &MockPhoneAssignmentRepository{}

	service := NewAssignmentService(mockAssignmentRepo, nil, nil)

	// Mock data
	now := time.Now()
	expireTime := now.Add(30 * time.Minute)
	assignments := []*domain.PhoneAssignment{
		{
			ID:             1,
			CustomerID:     1,
			ProviderID:     "1",
			BusinessTypeID: 1,
			CardType:       "virtual",
			PhoneNumber:    "+15551234567",
			Cost:           0.10,
			Status:         2,
			CreatedAt:      now,
			ExpiresAt:      &expireTime,
		},
		{
			ID:             2,
			CustomerID:     1,
			ProviderID:     "1",
			BusinessTypeID: 1,
			CardType:       "virtual",
			PhoneNumber:    "+15551234568",
			Cost:           0.15,
			Status:         2,
			CreatedAt:      now,
			ExpiresAt:      &expireTime,
		},
	}

	// Setup expectations
	startDate := now.AddDate(0, 0, -1)
	endDate := now.AddDate(0, 0, 1)
	mockAssignmentRepo.On("FindRecentByCustomerID", mock.Anything, int64(1), 9999999, 0).Return(assignments, int64(2), nil)

	// Test
	stats, err := service.GetCostStatistics(context.Background(), 1, &startDate, &endDate)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 0.25, stats.TotalCost)
	assert.Equal(t, int64(2), stats.TotalCount)

	mockAssignmentRepo.AssertExpectations(t)
}

func TestAssignmentService_GetCostStatistics_InvalidTimeRange(t *testing.T) {
	// Setup
	mockAssignmentRepo := &MockPhoneAssignmentRepository{}

	service := NewAssignmentService(mockAssignmentRepo, nil, nil)

	// Test with end date before start date
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, -1) // 1 day before start

	stats, err := service.GetCostStatistics(context.Background(), 1, &startDate, &endDate)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.Equal(t, ErrInvalidTimeRange, err)
}
