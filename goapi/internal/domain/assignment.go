package domain

import (
	"time"

	"gorm.io/gorm"
)

// PhoneAssignment 号码记录表
type PhoneAssignment struct {
	ID                     int64          `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	CreatedAt              time.Time      `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"comment:更新时间" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at"`
	BusinessName           string         `gorm:"type:varchar(255);not null;comment:业务名称" json:"business_name"`
	BusinessCode           string         `gorm:"type:varchar(50);not null;comment:业务编码;index:idx_business_code" json:"business_code"`
	MerchantNo             string         `gorm:"type:varchar(50);not null;comment:商户号;index:idx_merchant_no" json:"merchant_no"`
	MerchantName           string         `gorm:"type:varchar(255);not null;comment:商户名称" json:"merchant_name"`
	PhoneNumber            *string        `gorm:"type:varchar(50);comment:获取到的手机号;index:idx_phone_number" json:"phone_number"`
	VerificationCode       *string        `gorm:"type:varchar(50);comment:获取到的验证码" json:"verification_code"`
	FetchCount             *int           `gorm:"default:0;comment:获取验证码次数" json:"fetch_count"`
	Status                 *string        `gorm:"type:varchar(20);default:pending;comment:状态 (pending:待取码, completed:已完成, expired:已过期, failed:失败);index:idx_status" json:"status"`
	ProviderCost           *float64       `gorm:"type:decimal(10,4);comment:渠道成本" json:"provider_cost"`
	MerchantFee            *float64       `gorm:"type:decimal(10,4);comment:商户费用" json:"merchant_fee"`
	Profit                 *float64       `gorm:"type:decimal(10,4);comment:利润" json:"profit"`
	Remark                 *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`
	CustomerID             *int64         `gorm:"comment:客户ID, 关联到sms_customers.id;index:idx_customer_id" json:"customer_id"`
	ProviderID             *int64         `gorm:"comment:服务商ID, 关联到sms_providers.id;index:idx_provider_id" json:"provider_id"`
	PlatformBusinessTypeID *int64         `gorm:"comment:平台业务类型ID" json:"platform_business_type_id"`

	// 关联关系
	Customer             *Customer             `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Provider             *Provider             `gorm:"foreignKey:ProviderID" json:"provider,omitempty"`
	PlatformBusinessType *PlatformBusinessType `gorm:"foreignKey:PlatformBusinessTypeID" json:"platform_business_type,omitempty"`
}

// TableName 指定表名
func (PhoneAssignment) TableName() string {
	return "sms_phone_assignments"
}
