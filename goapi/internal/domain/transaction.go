package domain

import "time"

// Transaction 交易记录表
type Transaction struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID    int64     `gorm:"not null;index" json:"customer_id"`                 // 客户ID
	Amount        float64   `gorm:"type:decimal(10,4);not null" json:"amount"`         // 变动金额 (正数为充值, 负数为消费)
	BalanceBefore float64   `gorm:"type:decimal(10,4);not null" json:"balance_before"` // 变动前余额
	BalanceAfter  float64   `gorm:"type:decimal(10,4);not null" json:"balance_after"`  // 变动后余额
	Type          string    `gorm:"type:varchar(10);not null" json:"type"`             // 交易类型 (1:充值, 2:拉号码, 3:拉号-回退, 4:上分, 5:下分)
	ReferenceID   int64     `json:"reference_id"`                                      // 关联的业务ID, 例如phone_assignments.id
	Notes         string    `json:"notes"`                                             // 备注
	CreatedAt     time.Time `json:"created_at"`                                        // 创建时间
}

// TableName 指定表名
func (Transaction) TableName() string {
	return "sms_transactions"
}
