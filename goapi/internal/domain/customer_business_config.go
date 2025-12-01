package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// CustomerBusinessConfig 商户业务配置表
type CustomerBusinessConfig struct {
	ID                     int64          `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CustomerID             int64          `gorm:"not null;comment:商户ID" json:"customer_id"`
	PlatformBusinessTypeID int64          `gorm:"not null;comment:平台业务类型ID" json:"platform_business_type_id"`
	BusinessCode           string         `gorm:"type:varchar(50);not null;comment:业务编码;index" json:"business_code"`
	BusinessName           string         `gorm:"type:varchar(255);not null;comment:业务名称" json:"business_name"`
	Cost                   float64        `gorm:"type:decimal(10,4);not null;default:0.0000;comment:业务成本（单价）" json:"cost"`
	Weight                 int            `gorm:"not null;default:1;comment:权重（用于随机选择，权重越高被选中概率越大）" json:"weight"`
	Status                 *bool          `gorm:"default:true;comment:是否启用" json:"status"`
}

// TableName specifies the table name for CustomerBusinessConfig
func (CustomerBusinessConfig) TableName() string {
	return "sms_customer_business_config"
}

// CustomerBusinessConfigRepository defines the interface for customer business config data operations
type CustomerBusinessConfigRepository interface {
	Create(ctx context.Context, config *CustomerBusinessConfig) error
	FindByCustomerID(ctx context.Context, customerID int64) ([]*CustomerBusinessConfig, error)
	FindByCustomerIDAndBusinessCode(ctx context.Context, customerID int64, businessCode string) (*CustomerBusinessConfig, error)
	Update(ctx context.Context, config *CustomerBusinessConfig) error
	Delete(ctx context.Context, id int64) error
	FindByCustomerIDAndEnabled(ctx context.Context, customerID int64) ([]*CustomerBusinessConfig, error)
}
