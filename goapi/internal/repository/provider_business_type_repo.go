package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

type providerBusinessTypeRepository struct {
	db *gorm.DB
}

// NewProviderBusinessTypeRepository 创建运营商业务类型Repository
func NewProviderBusinessTypeRepository(db *gorm.DB) domain.ProviderBusinessTypeRepository {
	return &providerBusinessTypeRepository{db: db}
}

func (r *providerBusinessTypeRepository) FindByID(ctx context.Context, id int64) (*domain.ProviderBusinessType, error) {
	var businessType domain.ProviderBusinessType
	err := r.db.WithContext(ctx).First(&businessType, id).Error
	if err != nil {
		return nil, err
	}
	return &businessType, nil
}

func (r *providerBusinessTypeRepository) FindByProviderCodeAndBusinessCode(ctx context.Context, providerCode, businessCode string) (*domain.ProviderBusinessType, error) {
	var businessType domain.ProviderBusinessType
	err := r.db.WithContext(ctx).
		Where("provider_code = ? AND business_code = ? AND (status IS NULL OR status = ?)", providerCode, businessCode, true).
		First(&businessType).Error
	if err != nil {
		return nil, err
	}
	return &businessType, nil
}

func (r *providerBusinessTypeRepository) FindByProviderCode(ctx context.Context, providerCode string) ([]*domain.ProviderBusinessType, error) {
	var businessTypes []*domain.ProviderBusinessType
	err := r.db.WithContext(ctx).
		Where("provider_code = ? AND (status IS NULL OR status = ?)", providerCode, true).
		Find(&businessTypes).Error
	return businessTypes, err
}

