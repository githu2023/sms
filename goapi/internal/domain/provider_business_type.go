package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// ProviderBusinessType 三方渠道与业务关系管理表
type ProviderBusinessType struct {
	ID           int64          `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ProviderID   int            `gorm:"not null;comment:三方渠道ID（关联sms_providers表的ID）;uniqueIndex:uk_provider_business" json:"provider_id"`
	ProviderCode string         `gorm:"type:varchar(50);not null;comment:三方编码;index:idx_provider_code" json:"provider_code"`
	BusinessName *string        `gorm:"type:varchar(255);comment:业务名称" json:"business_name"`
	BusinessCode *string        `gorm:"type:varchar(50);comment:业务编码;index:idx_business_code" json:"business_code"`
	Price        *float64       `gorm:"type:decimal(10,4);default:0.0000;comment:该渠道该业务的价格" json:"price"`
	Status       *bool          `gorm:"default:1;comment:该渠道是否支持该业务" json:"status"`
	Remark       *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`

	// 关联关系
	Provider *Provider `gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE" json:"provider,omitempty"`
}

func (ProviderBusinessType) TableName() string {
	return "sms_providers_business_types"
}

// ProviderBusinessTypeRepository 运营商业务类型Repository接口
type ProviderBusinessTypeRepository interface {
	FindByID(ctx context.Context, id int64) (*ProviderBusinessType, error)
	FindByProviderCodeAndBusinessCode(ctx context.Context, providerCode, businessCode string) (*ProviderBusinessType, error)
	FindByProviderCode(ctx context.Context, providerCode string) ([]*ProviderBusinessType, error)
}
