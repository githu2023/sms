package dto

import "time"

// TransactionItem 交易记录项
type TransactionItem struct {
	ID            int64     `json:"id" example:"1"`
	Amount        *float32  `json:"amount" example:"100.50"`                   // 变动金额 (正数为增加, 负数为减少)
	BalanceBefore *float32  `json:"balance_before" example:"500.00"`           // 变动前余额
	BalanceAfter  *float32  `json:"balance_after" example:"600.50"`            // 变动后余额
	FrozenBefore  *float32  `json:"frozen_before,omitempty" example:"0.00"`    // 变动前冻结金额
	FrozenAfter   *float32  `json:"frozen_after,omitempty" example:"0.00"`     // 变动后冻结金额
	Type          *string   `json:"type" example:"1"`                          // 交易类型 (1:充值, 2:拉号消费, 3:拉号回退, 4:上分, 5:下分, 6:预冻结, 7:解冻, 8:冻结转实扣)
	ReferenceID   *int64    `json:"reference_id,omitempty" example:"123"`      // 关联的业务ID
	Notes         *string   `json:"notes,omitempty" example:"用户充值"`            // 备注
	CreatedAt     time.Time `json:"created_at" example:"2024-01-01T12:00:00Z"` // 创建时间
}

// TransactionListResponse 交易记录列表响应
type TransactionListResponse struct {
	Total        int64             `json:"total" example:"100"` // 总记录数
	Limit        int               `json:"limit" example:"20"`  // 每页数量
	Offset       int               `json:"offset" example:"0"`  // 偏移量
	Transactions []TransactionItem `json:"transactions"`        // 交易记录列表
}

// TransactionTypeDescription 交易类型说明
var TransactionTypeDescription = map[string]string{
	"1": "充值",
	"2": "拉号消费",
	"3": "拉号回退",
	"4": "上分",
	"5": "下分",
	"6": "预冻结",
	"7": "解冻",
	"8": "冻结转实扣",
}
