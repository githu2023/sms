package domain

import (
	"time"

	"gorm.io/gorm"
)

// Transaction 客户余额交易记录表
type Transaction struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID    int64          `gorm:"not null;comment:客户ID;index:idx_customer_id" json:"customer_id"`
	Amount        *float32       `gorm:"comment:变动金额 (正数为增加, 负数为减少)" json:"amount"`
	BalanceBefore *float32       `gorm:"comment:变动前余额" json:"balance_before"`
	BalanceAfter  *float32       `gorm:"comment:变动后余额" json:"balance_after"`
	FrozenBefore  *float32       `gorm:"comment:变动前冻结金额" json:"frozen_before"`
	FrozenAfter   *float32       `gorm:"comment:变动后冻结金额" json:"frozen_after"`
	Type          *string        `gorm:"type:varchar(10);comment:交易类型 (1:充值, 2:拉号消费, 3:拉号回退, 4:上分, 5:下分, 6:预冻结, 7:解冻, 8:冻结转实扣)" json:"type"`
	ReferenceID   *int64         `gorm:"comment:关联的业务ID, 例如sms_phone_assignments.id" json:"reference_id"`
	Notes         *string        `gorm:"type:varchar(255);comment:备注" json:"notes"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// 关联关系
	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

// TableName 指定表名
func (Transaction) TableName() string {
	return "sms_transactions"
}
