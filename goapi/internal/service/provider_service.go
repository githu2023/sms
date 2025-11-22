package service

import (
	"context"
	"sms-platform/goapi/internal/domain"
)

// ProviderService defines the interface for provider related business logic.
type ProviderService interface {
	CreateProvider(ctx context.Context, name, apiConfig string, isEnabled bool) (*domain.Provider, error)
	GetProviderByID(ctx context.Context, id int) (*domain.Provider, error)
	ListProviders(ctx context.Context) ([]*domain.Provider, error)
	UpdateProvider(ctx context.Context, provider *domain.Provider) error
	DeleteProvider(ctx context.Context, id int) error
}

type providerService struct {
	repo domain.ProviderRepository
}

// NewProviderService creates a new provider service.
func NewProviderService(repo domain.ProviderRepository) ProviderService {
	return &providerService{repo: repo}
}

func (s *providerService) CreateProvider(ctx context.Context, name, apiConfig string, isEnabled bool) (*domain.Provider, error) {
	provider := &domain.Provider{
		Name:      name,
		APIConfig: apiConfig,
		IsEnabled: isEnabled,
	}
	err := s.repo.Create(ctx, provider)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (s *providerService) GetProviderByID(ctx context.Context, id int) (*domain.Provider, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *providerService) ListProviders(ctx context.Context) ([]*domain.Provider, error) {
	return s.repo.FindAll(ctx)
}

func (s *providerService) UpdateProvider(ctx context.Context, provider *domain.Provider) error {
	return s.repo.Update(ctx, provider)
}

func (s *providerService) DeleteProvider(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
