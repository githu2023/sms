package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsProviders 三方渠道管理（上游服务商）结构体
type SmsProviders struct {
	global.GVA_MODEL
	Name        *string `json:"name" form:"name" gorm:"comment:三方名称;column:name;size:255;"`                           // 三方名称
	Code        *string `json:"code" form:"code" gorm:"comment:三方编码;column:code;size:50;"`                            // 三方编码
	ApiGateway  *string `json:"apiGateway" form:"apiGateway" gorm:"comment:三方API网关地址;column:api_gateway;type:text;"`  // 三方API网关地址
	MerchantId  *string `json:"merchantId" form:"merchantId" gorm:"comment:三方商户号;column:merchant_id;size:255;"`       // 三方商户号
	MerchantKey *string `json:"merchantKey" form:"merchantKey" gorm:"comment:三方商户key;column:merchant_key;type:text;"` // 三方商户key
	Status      *bool   `json:"status" form:"status" gorm:"comment:启用状态;column:status;"`                              // 启用状态
	Remark      *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:500;"`                       // 备注
}

// TableName 三方渠道管理 SmsProviders自定义表名 sms_providers
func (SmsProviders) TableName() string {
	return "sms_providers"
}
