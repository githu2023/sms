package domain

import "context"

// Provider 第三方服务商表
type Provider struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"not null" json:"name"`                    // 服务商名称
	APIConfig string `gorm:"type:json" json:"-"`                      // 服务商的API配置 (如URL, key等)
	IsEnabled bool   `gorm:"not null;default:true" json:"is_enabled"` // 是否启用该服务商
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
