package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// PlatformProviderBusinessMapping 平台运营商业务映射表
type PlatformProviderBusinessMapping struct {
	ID                     uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PlatformBusinessTypeID *int64         `gorm:"comment:平台业务ID（关联sms_platform_business_types表的ID）" json:"platform_business_type_id"`
	PlatformBusinessCode   *string        `gorm:"type:varchar(50);comment:平台业务编码" json:"platform_business_code"`
	ProviderBusinessTypeID *int64         `gorm:"comment:三方业务ID（关联sms_providers_business_types表的ID）" json:"provider_business_type_id"`
	ProviderCode           *string        `gorm:"type:varchar(50);comment:三方编码" json:"provider_code"`
	BusinessCode           *string        `gorm:"type:varchar(50);comment:三方业务编码" json:"business_code"`
	Weight                 *int           `gorm:"comment:权重（用于随机选择，权重越高被选中概率越大）" json:"weight"`
	Status                 *bool          `gorm:"comment:是否启用该映射" json:"status"`
	Remark                 *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`
}

func (PlatformProviderBusinessMapping) TableName() string {
	return "sms_platform_provider_business_mapping"
}

// PlatformProviderBusinessMappingRepository 平台运营商业务映射Repository接口
type PlatformProviderBusinessMappingRepository interface {
	FindByPlatformBusinessTypeID(ctx context.Context, platformBusinessTypeID int64) ([]*PlatformProviderBusinessMapping, error)
	FindByPlatformBusinessCode(ctx context.Context, platformBusinessCode string) ([]*PlatformProviderBusinessMapping, error)
}
