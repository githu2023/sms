package service

import (
	"context"
	"sms-platform/goapi/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBusinessTypeRepository is a mock implementation of the BusinessTypeRepository interface.
type MockBusinessTypeRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeRepository) Create(ctx context.Context, businessType *domain.BusinessType) error {
	args := m.Called(ctx, businessType)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) FindByCode(ctx context.Context, code string) (*domain.BusinessType, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindByID(ctx context.Context, id int) (*domain.BusinessType, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindAll(ctx context.Context) ([]*domain.BusinessType, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) Update(ctx context.Context, businessType *domain.BusinessType) error {
	args := m.Called(ctx, businessType)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestBusinessService_CreateBusinessType(t *testing.T) {
	mockRepo := new(MockBusinessTypeRepository)
	businessService := NewBusinessService(mockRepo)
	ctx := context.Background()

	businessType := &domain.BusinessType{
		Name:      "TestBT",
		Code:      "test_bt",
		IsEnabled: true,
	}

	mockRepo.On("Create", ctx, businessType).Return(nil)

	createdBT, err := businessService.CreateBusinessType(ctx, businessType.Name, businessType.Code, businessType.IsEnabled)
	assert.NoError(t, err)
	assert.Equal(t, businessType.Name, createdBT.Name)
	mockRepo.AssertExpectations(t)
}

func TestBusinessService_GetBusinessTypeByID(t *testing.T) {
	mockRepo := new(MockBusinessTypeRepository)
	businessService := NewBusinessService(mockRepo)
	ctx := context.Background()

	businessType := &domain.BusinessType{ID: 1, Name: "TestBT"}

	mockRepo.On("FindByID", ctx, 1).Return(businessType, nil)

	foundBT, err := businessService.GetBusinessTypeByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, businessType.ID, foundBT.ID)
	mockRepo.AssertExpectations(t)
}

func TestBusinessService_GetBusinessTypeByCode(t *testing.T) {
	mockRepo := new(MockBusinessTypeRepository)
	businessService := NewBusinessService(mockRepo)
	ctx := context.Background()

	businessType := &domain.BusinessType{ID: 1, Code: "test_code"}

	mockRepo.On("FindByCode", ctx, "test_code").Return(businessType, nil)

	foundBT, err := businessService.GetBusinessTypeByCode(ctx, "test_code")
	assert.NoError(t, err)
	assert.Equal(t, businessType.ID, foundBT.ID)
	mockRepo.AssertExpectations(t)
}

func TestBusinessService_ListBusinessTypes(t *testing.T) {
	mockRepo := new(MockBusinessTypeRepository)
	businessService := NewBusinessService(mockRepo)
	ctx := context.Background()

	businessTypes := []*domain.BusinessType{
		{ID: 1, Name: "BT1"},
		{ID: 2, Name: "BT2"},
	}

	mockRepo.On("FindAll", ctx).Return(businessTypes, nil)

	listedBTs, err := businessService.ListBusinessTypes(ctx)
	assert.NoError(t, err)
	assert.Len(t, listedBTs, 2)
	mockRepo.AssertExpectations(t)
}

func TestBusinessService_UpdateBusinessType(t *testing.T) {
	mockRepo := new(MockBusinessTypeRepository)
	businessService := NewBusinessService(mockRepo)
	ctx := context.Background()

	businessType := &domain.BusinessType{ID: 1, Name: "UpdatedBT"}

	mockRepo.On("Update", ctx, businessType).Return(nil)

	err := businessService.UpdateBusinessType(ctx, businessType)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBusinessService_DeleteBusinessType(t *testing.T) {
	mockRepo := new(MockBusinessTypeRepository)
	businessService := NewBusinessService(mockRepo)
	ctx := context.Background()

	mockRepo.On("Delete", ctx, 1).Return(nil)

	err := businessService.DeleteBusinessType(ctx, 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
