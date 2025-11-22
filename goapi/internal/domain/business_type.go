package domain

import "context"

// BusinessType 业务类型表
type BusinessType struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"not null" json:"name"`                    // 业务名称, 例如 "腾讯QQ"
	Code      string `gorm:"unique;not null" json:"code"`             // 业务代码, 例如 "qq"
	IsEnabled bool   `gorm:"not null;default:true" json:"is_enabled"` // 是否开放该业务
}

// BusinessTypeRepository defines the interface for business type data operations.
type BusinessTypeRepository interface {
	Create(ctx context.Context, businessType *BusinessType) error
	FindByCode(ctx context.Context, code string) (*BusinessType, error)
	FindByID(ctx context.Context, id int) (*BusinessType, error)
	FindAll(ctx context.Context) ([]*BusinessType, error)
	Update(ctx context.Context, businessType *BusinessType) error
	Delete(ctx context.Context, id int) error
}

// TableName 指定表名
func (BusinessType) TableName() string {
	return "sms_business_types"
}
