package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

type platformProviderBusinessMappingRepository struct {
	db *gorm.DB
}

// NewPlatformProviderBusinessMappingRepository 创建平台运营商业务映射Repository
func NewPlatformProviderBusinessMappingRepository(db *gorm.DB) domain.PlatformProviderBusinessMappingRepository {
	return &platformProviderBusinessMappingRepository{db: db}
}

func (r *platformProviderBusinessMappingRepository) FindByPlatformBusinessTypeID(ctx context.Context, platformBusinessTypeID int64) ([]*domain.PlatformProviderBusinessMapping, error) {
	var mappings []*domain.PlatformProviderBusinessMapping
	err := r.db.WithContext(ctx).
		Where("platform_business_type_id = ? AND status = ?", platformBusinessTypeID, true).
		Find(&mappings).Error
	return mappings, err
}

func (r *platformProviderBusinessMappingRepository) FindByPlatformBusinessCode(ctx context.Context, platformBusinessCode string) ([]*domain.PlatformProviderBusinessMapping, error) {
	var mappings []*domain.PlatformProviderBusinessMapping
	err := r.db.WithContext(ctx).
		Where("platform_business_code = ? AND status = ?", platformBusinessCode, true).
		Find(&mappings).Error
	return mappings, err
}

