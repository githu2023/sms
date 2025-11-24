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
	// 新增：获取用户分配的业务类型
	GetBusinessTypesForCustomer(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error)
}

type businessService struct {
	businessTypeRepo           domain.BusinessTypeRepository
	customerBusinessConfigRepo domain.CustomerBusinessConfigRepository
}

// NewBusinessService creates a new business service.
func NewBusinessService(businessTypeRepo domain.BusinessTypeRepository, customerBusinessConfigRepo domain.CustomerBusinessConfigRepository) BusinessService {
	return &businessService{
		businessTypeRepo:           businessTypeRepo,
		customerBusinessConfigRepo: customerBusinessConfigRepo,
	}
}

func (s *businessService) CreateBusinessType(ctx context.Context, name, code string, isEnabled bool) (*domain.BusinessType, error) {
	businessType := &domain.BusinessType{
		Name:      name,
		Code:      code,
		IsEnabled: isEnabled,
	}
	err := s.businessTypeRepo.Create(ctx, businessType)
	if err != nil {
		return nil, err
	}
	return businessType, nil
}

func (s *businessService) GetBusinessTypeByID(ctx context.Context, id int) (*domain.BusinessType, error) {
	return s.businessTypeRepo.FindByID(ctx, id)
}

func (s *businessService) GetBusinessTypeByCode(ctx context.Context, code string) (*domain.BusinessType, error) {
	return s.businessTypeRepo.FindByCode(ctx, code)
}

func (s *businessService) ListBusinessTypes(ctx context.Context) ([]*domain.BusinessType, error) {
	return s.businessTypeRepo.FindAll(ctx)
}

func (s *businessService) UpdateBusinessType(ctx context.Context, businessType *domain.BusinessType) error {
	return s.businessTypeRepo.Update(ctx, businessType)
}

func (s *businessService) DeleteBusinessType(ctx context.Context, id int) error {
	return s.businessTypeRepo.Delete(ctx, id)
}

// GetBusinessTypesForCustomer 获取用户分配的业务类型
func (s *businessService) GetBusinessTypesForCustomer(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error) {
	return s.customerBusinessConfigRepo.FindByCustomerIDAndEnabled(ctx, customerID)
}
