package service

import (
	"context"
	"sms-platform/goapi/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProviderRepository is a mock implementation of the ProviderRepository interface.
type MockProviderRepository struct {
	mock.Mock
}

func (m *MockProviderRepository) Create(ctx context.Context, provider *domain.Provider) error {
	args := m.Called(ctx, provider)
	return args.Error(0)
}

func (m *MockProviderRepository) FindByID(ctx context.Context, id int) (*domain.Provider, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Provider), args.Error(1)
}

func (m *MockProviderRepository) FindByCode(ctx context.Context, code string) (*domain.Provider, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Provider), args.Error(1)
}

func (m *MockProviderRepository) FindAll(ctx context.Context) ([]*domain.Provider, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Provider), args.Error(1)
}

func (m *MockProviderRepository) Update(ctx context.Context, provider *domain.Provider) error {
	args := m.Called(ctx, provider)
	return args.Error(0)
}

func (m *MockProviderRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProviderService_CreateProvider(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	ctx := context.Background()

	name := "TestProvider"
	apiConfig := `{"url":"http://test.com"}`
	enabled := true
	provider := &domain.Provider{
		Name:      &name,
		APIConfig: &apiConfig,
		IsEnabled: &enabled,
	}

	mockRepo.On("Create", ctx, provider).Return(nil)

	createdProvider, err := providerService.CreateProvider(ctx, *provider.Name, *provider.APIConfig, *provider.IsEnabled)
	assert.NoError(t, err)
	assert.Equal(t, provider.Name, createdProvider.Name)
	mockRepo.AssertExpectations(t)
}

func TestProviderService_GetProviderByID(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	ctx := context.Background()

	name := "TestProvider"
	provider := &domain.Provider{ID: 1, Name: &name}

	mockRepo.On("FindByID", ctx, 1).Return(provider, nil)

	foundProvider, err := providerService.GetProviderByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, provider.ID, foundProvider.ID)
	mockRepo.AssertExpectations(t)
}

func TestProviderService_ListProviders(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	ctx := context.Background()

	p1Name := "P1"
	p2Name := "P2"
	providers := []*domain.Provider{
		{ID: 1, Name: &p1Name},
		{ID: 2, Name: &p2Name},
	}

	mockRepo.On("FindAll", ctx).Return(providers, nil)

	listedProviders, err := providerService.ListProviders(ctx)
	assert.NoError(t, err)
	assert.Len(t, listedProviders, 2)
	mockRepo.AssertExpectations(t)
}

func TestProviderService_UpdateProvider(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	ctx := context.Background()

	updatedName := "UpdatedProvider"
	provider := &domain.Provider{ID: 1, Name: &updatedName}

	mockRepo.On("Update", ctx, provider).Return(nil)

	err := providerService.UpdateProvider(ctx, provider)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProviderService_DeleteProvider(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	ctx := context.Background()

	mockRepo.On("Delete", ctx, 1).Return(nil)

	err := providerService.DeleteProvider(ctx, 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
