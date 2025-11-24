package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Provider 第三方运营商表
type Provider struct {
	ID          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        *string        `gorm:"type:varchar(255);comment:三方名称" json:"name"`
	Code        *string        `gorm:"type:varchar(50);comment:三方编码" json:"code"`
	APIGateway  *string        `gorm:"type:text;comment:三方API网关地址" json:"api_gateway"`
	MerchantID  *string        `gorm:"type:varchar(255);comment:三方商户号" json:"merchant_id"`
	MerchantKey *string        `gorm:"type:text;comment:三方商户key" json:"merchant_key"`
	Status      *bool          `gorm:"comment:启用状态" json:"status"`
	Remark      *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`
	APIConfig   *string        `gorm:"type:text;comment:服务商的API配置 (如URL, key等)" json:"api_config"`
	IsEnabled   *bool          `gorm:"comment:是否启用该服务商" json:"is_enabled"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ProviderRepository defines the interface for provider data operations.
type ProviderRepository interface {
	Create(ctx context.Context, provider *Provider) error
	FindByID(ctx context.Context, id int) (*Provider, error)
	FindAll(ctx context.Context) ([]*Provider, error)
	Update(ctx context.Context, provider *Provider) error
	Delete(ctx context.Context, id int) error
}

// TableName 指定表名
func (Provider) TableName() string {
	return "sms_providers"
}
