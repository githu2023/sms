package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

type platformBusinessTypeRepository struct {
	db *gorm.DB
}

// NewPlatformBusinessTypeRepository 创建平台业务类型Repository
func NewPlatformBusinessTypeRepository(db *gorm.DB) domain.PlatformBusinessTypeRepository {
	return &platformBusinessTypeRepository{db: db}
}

func (r *platformBusinessTypeRepository) FindByID(ctx context.Context, id int64) (*domain.PlatformBusinessType, error) {
	var businessType domain.PlatformBusinessType
	err := r.db.WithContext(ctx).First(&businessType, id).Error
	if err != nil {
		return nil, err
	}
	return &businessType, nil
}

func (r *platformBusinessTypeRepository) FindByCode(ctx context.Context, code string) (*domain.PlatformBusinessType, error) {
	var businessType domain.PlatformBusinessType
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&businessType).Error
	if err != nil {
		return nil, err
	}
	return &businessType, nil
}

func (r *platformBusinessTypeRepository) FindAll(ctx context.Context) ([]*domain.PlatformBusinessType, error) {
	var businessTypes []*domain.PlatformBusinessType
	err := r.db.WithContext(ctx).Find(&businessTypes).Error
	return businessTypes, err
}

