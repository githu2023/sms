package domain

const (
	TransactionTypeTopUp          = "1" // 充值/入账
	TransactionTypeDeduct         = "2" // 拉号消费（直接扣除可用余额）
	TransactionTypeRefund         = "3" // 拉号回退（退款）
	TransactionTypeCredit         = "4" // 上分
	TransactionTypeDebit          = "5" // 下分
	TransactionTypeFreeze         = "6" // 预冻结（可用余额转冻结）
	TransactionTypeUnfreeze       = "7" // 解冻（冻结转回可用）
	TransactionTypeFreezeToCharge = "8" // 冻结转实扣（冻结金额结算为正式消费）
)
