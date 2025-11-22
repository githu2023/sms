package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsProvidersBusinessTypes 三方业务 结构体
type SmsProvidersBusinessTypes struct {
	global.GVA_MODEL
	ProviderID   *int     `json:"providerId" form:"providerId" gorm:"column:provider_id;comment:三方渠道ID;"`
	ProviderCode string   `json:"providerCode" form:"providerCode" gorm:"column:provider_code;comment:三方编码;size:50;"`
	BusinessName string   `json:"businessName" form:"businessName" gorm:"column:business_name;comment:业务名称;size:255;"`
	BusinessCode string   `json:"businessCode" form:"businessCode" gorm:"column:business_code;comment:业务编码;size:50;"`
	Price        *float64 `json:"price" form:"price" gorm:"column:price;comment:该渠道该业务的价格;type:decimal(10,4);"`
	Status       *bool    `json:"status" form:"status" gorm:"column:status;comment:该渠道是否支持该业务;"`
	Remark       string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;"`
}

// TableName 三方业务 SmsProvidersBusinessTypes自定义表名 sms_providers_business_types
func (SmsProvidersBusinessTypes) TableName() string {
	return "sms_providers_business_types"
}
