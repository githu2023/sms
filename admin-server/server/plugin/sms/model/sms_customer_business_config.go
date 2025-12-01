package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsCustomerBusinessConfig 商户业务配置 结构体
type SmsCustomerBusinessConfig struct {
	global.GVA_MODEL
	CustomerID             int64   `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:商户ID;not null;"`
	PlatformBusinessTypeID int64   `json:"platformBusinessTypeId" form:"platformBusinessTypeId" gorm:"column:platform_business_type_id;comment:平台业务类型ID;not null;"`
	BusinessCode           string  `json:"businessCode" form:"businessCode" gorm:"column:business_code;comment:业务编码;size:50;not null;"`
	BusinessName           string  `json:"businessName" form:"businessName" gorm:"column:business_name;comment:业务名称;size:255;not null;"`
	Cost                   float64 `json:"cost" form:"cost" gorm:"column:cost;comment:业务成本（单价）;type:decimal(10,4);not null;default:0.0000;"`
	Weight                 int32   `json:"weight" form:"weight" gorm:"column:weight;comment:权重（用于随机选择，权重越高被选中概率越大）;not null;default:1;"`
	Status                 int     `json:"status" form:"status" gorm:"column:status;comment:状态(1:启用,0:禁用);type:tinyint;not null;default:1;"`
}

// TableName 商户业务配置 SmsCustomerBusinessConfig自定义表名 sms_customer_business_config
func (SmsCustomerBusinessConfig) TableName() string {
	return "sms_customer_business_config"
}
