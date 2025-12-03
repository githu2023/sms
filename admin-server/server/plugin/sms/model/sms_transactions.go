package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsTransactions 交易记录 结构体
type SmsTransactions struct {
	global.GVA_MODEL
	CustomerId    *int64   `json:"customerId" form:"customerId" gorm:"comment:客户ID;column:customer_id;"`                                   //客户ID
	Amount        *float64 `json:"amount" form:"amount" gorm:"comment:变动金额 (正数为充值, 负数为消费);column:amount;size:10;"`                         //变动金额 (正数为充值, 负数为消费)
	BalanceBefore *float64 `json:"balanceBefore" form:"balanceBefore" gorm:"comment:变动前余额;column:balance_before;size:10;"`                 //变动前余额
	BalanceAfter  *float64 `json:"balanceAfter" form:"balanceAfter" gorm:"comment:变动后余额;column:balance_after;size:10;"`                    //变动后余额
	FrozenBefore  *float64 `json:"frozenBefore" form:"frozenBefore" gorm:"comment:变动前冻结金额;column:frozen_before;size:10;"`                  //变动前冻结金额
	FrozenAfter   *float64 `json:"frozenAfter" form:"frozenAfter" gorm:"comment:变动后冻结金额;column:frozen_after;size:10;"`                     //变动后冻结金额
	Type          *string  `json:"type" form:"type" gorm:"comment:交易类型 (1:充值, 2:拉号码, 3:拉号-回退, 4:上分, 5:下分);column:type;size:10;"`           //交易类型 (1:充值, 2:拉号码, 3:拉号-回退, 4:上分, 5:下分)
	ReferenceId   *int64   `json:"referenceId" form:"referenceId" gorm:"comment:关联的业务ID, 例如sms_phone_assignments.id;column:reference_id;"` //关联的业务ID, 例如sms_phone_assignments.id
	Notes         *string  `json:"notes" form:"notes" gorm:"comment:备注;column:notes;size:255;"`                                            //备注
}

// TableName 交易记录 SmsTransactions自定义表名 sms_transactions
func (SmsTransactions) TableName() string {
	return "sms_transactions"
}
