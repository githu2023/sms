package domain

import (
	"time"
)

// PhoneAssignment 手机号分配记录表
type PhoneAssignment struct {
	ID               int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID       int64      `gorm:"not null;index" json:"customer_id"`          // 客户ID
	ProviderID       string     `gorm:"not null" json:"provider_id"`                // 服务商ID (Changed to string to match ProviderInfo.ID)
	BusinessTypeID   int        `gorm:"not null" json:"business_type_id"`           // 业务类型ID
	CardType         string     `gorm:"size:50;not null" json:"card_type"`          // 卡类型 (例如: physical, virtual)
	PhoneNumber      string     `gorm:"size:50;not null;index" json:"phone_number"` // 获取到的手机号
	VerificationCode string     `gorm:"size:50" json:"verification_code"`           // 获取到的验证码
	Cost             float64    `gorm:"type:decimal(10,4);not null" json:"cost"`    // 本次操作的费用
	Status           int        `gorm:"not null" json:"status"`                     // 状态 (1:待取码, 2:已完成, 3:已过期, 4:失败)
	ExpiresAt        *time.Time `json:"expires_at"`                                 // 手机号锁定的过期时间
	CreatedAt        time.Time  `json:"created_at"`                                 // 创建时间
	UpdatedAt        time.Time  `json:"updated_at"`                                 // 更新时间
}

// TableName 指定表名
func (PhoneAssignment) TableName() string {
	return "sms_phone_assignments"
}
