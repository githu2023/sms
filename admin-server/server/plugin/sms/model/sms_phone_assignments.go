package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsPhoneAssignments 号码记录 结构体
type SmsPhoneAssignments struct {
	global.GVA_MODEL
	BusinessName     string   `json:"businessName" form:"businessName" gorm:"comment:业务名称;column:business_name;size:255;"`                                                //业务名称
	BusinessCode     string   `json:"businessCode" form:"businessCode" gorm:"comment:业务编码;column:business_code;size:50;"`                                                 //业务编码
	MerchantNo       string   `json:"merchantNo" form:"merchantNo" gorm:"comment:商户号;column:merchant_no;size:50;"`                                                        //商户号
	MerchantName     string   `json:"merchantName" form:"merchantName" gorm:"comment:商户名称;column:merchant_name;size:255;"`                                                //商户名称
	PhoneNumber      *string  `json:"phoneNumber" form:"phoneNumber" gorm:"comment:获取到的手机号;column:phone_number;size:50;"`                                                 //获取到的手机号
	VerificationCode *string  `json:"verificationCode" form:"verificationCode" gorm:"comment:获取到的验证码;column:verification_code;size:50;"`                                  //获取到的验证码
	FetchCount       *int32   `json:"fetchCount" form:"fetchCount" gorm:"comment:获取验证码次数;column:fetch_count;default:0;"`                                                  //获取验证码次数
	Status           *string  `json:"status" form:"status" gorm:"comment:状态 (pending:待取码, completed:已完成, expired:已过期, failed:失败);column:status;size:20;default:pending;"` //状态
	ProviderCost     *float64 `json:"providerCost" form:"providerCost" gorm:"comment:渠道成本;column:provider_cost;type:decimal(10,4);"`                                      //渠道成本
	MerchantFee      *float64 `json:"merchantFee" form:"merchantFee" gorm:"comment:商户费用;column:merchant_fee;type:decimal(10,4);"`                                         //商户费用
	AgentFee         *float64 `json:"agentFee" form:"agentFee" gorm:"comment:代理费用;column:agent_fee;type:decimal(10,4);"`                                                  //代理费用
	Profit           *float64 `json:"profit" form:"profit" gorm:"comment:利润;column:profit;type:decimal(10,4);"`                                                           //利润
	Remark           *string  `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:500;"`                                                                     //备注

	// 内部关联字段
	CustomerId             *int64 `json:"customerId" form:"customerId" gorm:"comment:客户ID, 关联到sms_customers.id;column:customer_id;"`                      //客户ID
	ProviderId             *int64 `json:"providerId" form:"providerId" gorm:"comment:服务商ID, 关联到sms_providers.id;column:provider_id;"`                     //服务商ID
	PlatformBusinessTypeId *int64 `json:"platformBusinessTypeId" form:"platformBusinessTypeId" gorm:"comment:平台业务类型ID;column:platform_business_type_id;"` //平台业务类型ID
}

// TableName 号码记录 SmsPhoneAssignments自定义表名 sms_phone_assignments
func (SmsPhoneAssignments) TableName() string {
	return "sms_phone_assignments"
}
