package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

type businessTypeRepository struct {
	db *gorm.DB
}

// NewBusinessTypeRepository creates a new business type repository.
func NewBusinessTypeRepository(db *gorm.DB) domain.BusinessTypeRepository {
	return &businessTypeRepository{db: db}
}

func (r *businessTypeRepository) Create(ctx context.Context, businessType *domain.BusinessType) error {
	return r.db.WithContext(ctx).Create(businessType).Error
}

func (r *businessTypeRepository) FindByCode(ctx context.Context, code string) (*domain.BusinessType, error) {
	var businessType domain.BusinessType
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&businessType).Error; err != nil {
		return nil, err
	}
	return &businessType, nil
}

func (r *businessTypeRepository) FindByID(ctx context.Context, id int) (*domain.BusinessType, error) {
	var businessType domain.BusinessType
	if err := r.db.WithContext(ctx).First(&businessType, id).Error; err != nil {
		return nil, err
	}
	return &businessType, nil
}

func (r *businessTypeRepository) FindAll(ctx context.Context) ([]*domain.BusinessType, error) {
	var businessTypes []*domain.BusinessType
	if err := r.db.WithContext(ctx).Find(&businessTypes).Error; err != nil {
		return nil, err
	}
	return businessTypes, nil
}

func (r *businessTypeRepository) Update(ctx context.Context, businessType *domain.BusinessType) error {
	return r.db.WithContext(ctx).Save(businessType).Error
}

func (r *businessTypeRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.BusinessType{}, id).Error
}
