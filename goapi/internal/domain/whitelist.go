package domain

import "time"

// IPWhitelist IP白名单表
type IPWhitelist struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID int64     `gorm:"uniqueIndex:uniq_customer_ip;not null" json:"customer_id"`        // 客户ID
	IPAddress  string    `gorm:"uniqueIndex:uniq_customer_ip;size:45;not null" json:"ip_address"` // 白名单IP或IP段
	Notes      string    `json:"notes"`                                                           // 备注, 例如 "办公室IP"
	CreatedAt  time.Time `json:"created_at"`                                                      // 创建时间
}

// TableName 指定表名
func (IPWhitelist) TableName() string {
	return "sms_ip_whitelists"
}
