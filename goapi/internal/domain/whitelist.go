package domain

import (
	"time"

	"gorm.io/gorm"
)

// IPWhitelist IP白名单表
type IPWhitelist struct {
	ID         int64          `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	CreatedAt  time.Time      `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"comment:更新时间" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at"`
	CustomerID int64          `gorm:"not null;comment:客户ID;index:idx_customer_id" json:"customer_id"`
	IPAddress  *string        `gorm:"type:varchar(45);comment:IP地址" json:"ip_address"`
	Status     *bool          `gorm:"comment:启用状态" json:"status"`
	Remark     *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`

	// 关联关系
	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

// TableName 指定表名
func (IPWhitelist) TableName() string {
	return "sms_ip_whitelist"
}
