package service

import (
	"context"
	"sms-platform/goapi/internal/domain"
)

// BusinessService defines the interface for business type related business logic.
type BusinessService interface {
	CreateBusinessType(ctx context.Context, name, code string, isEnabled bool) (*domain.BusinessType, error)
	GetBusinessTypeByID(ctx context.Context, id int) (*domain.BusinessType, error)
	GetBusinessTypeByCode(ctx context.Context, code string) (*domain.BusinessType, error)
	ListBusinessTypes(ctx context.Context) ([]*domain.BusinessType, error)
	UpdateBusinessType(ctx context.Context, businessType *domain.BusinessType) error
	DeleteBusinessType(ctx context.Context, id int) error
}

type businessService struct {
	repo domain.BusinessTypeRepository
}

// NewBusinessService creates a new business service.
func NewBusinessService(repo domain.BusinessTypeRepository) BusinessService {
	return &businessService{repo: repo}
}

func (s *businessService) CreateBusinessType(ctx context.Context, name, code string, isEnabled bool) (*domain.BusinessType, error) {
	businessType := &domain.BusinessType{
		Name:      name,
		Code:      code,
		IsEnabled: isEnabled,
	}
	err := s.repo.Create(ctx, businessType)
	if err != nil {
		return nil, err
	}
	return businessType, nil
}

func (s *businessService) GetBusinessTypeByID(ctx context.Context, id int) (*domain.BusinessType, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *businessService) GetBusinessTypeByCode(ctx context.Context, code string) (*domain.BusinessType, error) {
	return s.repo.FindByCode(ctx, code)
}

func (s *businessService) ListBusinessTypes(ctx context.Context) ([]*domain.BusinessType, error) {
	return s.repo.FindAll(ctx)
}

func (s *businessService) UpdateBusinessType(ctx context.Context, businessType *domain.BusinessType) error {
	return s.repo.Update(ctx, businessType)
}

func (s *businessService) DeleteBusinessType(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
