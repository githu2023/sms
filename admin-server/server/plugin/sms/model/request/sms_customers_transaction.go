package request

// CreditDebitSmsCustomersReq 充值/扣费请求
type CreditDebitSmsCustomersReq struct {
	CustomerId int64   `json:"customerId" binding:"required" form:"customerId"`         // 客户ID
	Amount     float64 `json:"amount" binding:"required,gt=0" form:"amount"`            // 金额（正数）
	Type       string  `json:"type" binding:"required,oneof=1 2 3 4 5 6 7" form:"type"` // 操作类型: 1=充值, 2=拉号码, 3=拉号-回退, 4=上分, 5=下分, 6=冻结金额, 7=冻结金额返回
	Notes      *string `json:"notes" form:"notes"`                                      // 备注
}

// TransactionType 交易类型常量
const (
	TransactionTypeCredit       = "1" // 1 = 充值
	TransactionTypePullNumber   = "2" // 2 = 拉号码
	TransactionTypePullRollback = "3" // 3 = 拉号-回退
	TransactionTypeAddPoints    = "4" // 4 = 上分
	TransactionTypeDeductPoints = "5" // 5 = 下分
	TransactionTypeFreeze       = "6" // 6 = 冻结金额
	TransactionTypeUnfreeze     = "7" // 7 = 冻结金额返回
)

// TransactionTypeText 交易类型文本描述
var TransactionTypeText = map[string]string{
	TransactionTypeCredit:       "充值",
	TransactionTypePullNumber:   "拉号码",
	TransactionTypePullRollback: "拉号-回退",
	TransactionTypeAddPoints:    "上分",
	TransactionTypeDeductPoints: "下分",
	TransactionTypeFreeze:       "冻结金额",
	TransactionTypeUnfreeze:     "冻结金额返回",
}
