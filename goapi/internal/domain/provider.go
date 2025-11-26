package domain

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// ProviderExtraConfig 运营商额外配置（JSON格式）
// 用于存储不同运营商的特殊配置，如 projectName、customParams 等
type ProviderExtraConfig map[string]interface{}

// Value 实现 driver.Valuer 接口，用于数据库存储
func (c ProviderExtraConfig) Value() (driver.Value, error) {
	if c == nil || len(c) == 0 {
		return nil, nil
	}
	return json.Marshal(c)
}

// Scan 实现 sql.Scanner 接口，用于数据库读取
func (c *ProviderExtraConfig) Scan(value interface{}) error {
	if value == nil {
		*c = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	if len(bytes) == 0 {
		*c = nil
		return nil
	}

	return json.Unmarshal(bytes, c)
}

// Provider 第三方运营商表
type Provider struct {
	ID          int                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        *string              `gorm:"type:varchar(255);comment:三方名称" json:"name"`
	Code        *string              `gorm:"type:varchar(50);comment:三方编码" json:"code"`
	APIGateway  *string              `gorm:"type:text;comment:三方API网关地址" json:"api_gateway"`
	MerchantID  *string              `gorm:"type:varchar(255);comment:三方商户号" json:"merchant_id"`
	MerchantKey *string              `gorm:"type:text;comment:三方商户key" json:"merchant_key"`
	Status      *bool                `gorm:"comment:启用状态" json:"status"`
	Remark      *string              `gorm:"type:varchar(500);comment:备注" json:"remark"`
	APIConfig   *string              `gorm:"type:text;comment:服务商的API配置 (已废弃，使用ExtraConfig)" json:"api_config"`
	ExtraConfig *ProviderExtraConfig `gorm:"type:json;comment:运营商额外配置(JSON格式，用于存储特殊配置如projectName等)" json:"extra_config"`
	IsEnabled   *bool                `gorm:"comment:是否启用该服务商" json:"is_enabled"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	DeletedAt   gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
}

// ProviderRepository defines the interface for provider data operations.
type ProviderRepository interface {
	Create(ctx context.Context, provider *Provider) error
	FindByID(ctx context.Context, id int) (*Provider, error)
	FindByCode(ctx context.Context, code string) (*Provider, error)
	FindAll(ctx context.Context) ([]*Provider, error)
	Update(ctx context.Context, provider *Provider) error
	Delete(ctx context.Context, id int) error
}

// TableName 指定表名
func (Provider) TableName() string {
	return "sms_providers"
}
