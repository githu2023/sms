package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsCustomers 商户 结构体
type SmsCustomers struct {
	global.GVA_MODEL
	MerchantName   string     `json:"merchantName" form:"merchantName" gorm:"comment:商户名称;column:merchant_name;size:255;"`
	MerchantNo     string     `json:"merchantNo" form:"merchantNo" gorm:"comment:商户号;column:merchant_no;size:50;uniqueIndex;"`
	Username       *string    `json:"username" form:"username" gorm:"comment:客户端登录用户名;column:username;size:255;"`
	Email          *string    `json:"email" form:"email" gorm:"comment:客户端登录邮箱;column:email;size:255;"`
	PasswordHash   *string    `json:"passwordHash" form:"passwordHash" gorm:"comment:客户端登录用的密码哈希;column:password_hash;size:255;"`
	ApiSecretKey   *string    `json:"apiSecretKey" form:"apiSecretKey" gorm:"comment:用于生成API Token的唯一密钥;column:api_secret_key;size:255;"`
	Balance        *float64   `json:"balance" form:"balance" gorm:"comment:客户余额;column:balance;type:decimal(10,2);default:0.00;"`
	ParentID       *int64     `json:"parentId" form:"parentId" gorm:"comment:上级商户ID;column:parent_id;"`
	FrozenAmount   *float64   `json:"frozenAmount" form:"frozenAmount" gorm:"comment:冻结金额;column:frozen_amount;type:decimal(10,2);default:0.00;"`
	Status         *bool      `json:"status" form:"status" gorm:"comment:客户状态 (1:正常, 0:冻结);column:status;"`
	RegistrationIp *string    `json:"registrationIp" form:"registrationIp" gorm:"comment:注册时的IP地址;column:registration_ip;size:45;"`
	LastLoginIp    *string    `json:"lastLoginIp" form:"lastLoginIp" gorm:"comment:最后一次登录的IP地址;column:last_login_ip;size:45;"`
	LastLoginAt    *time.Time `json:"lastLoginAt" form:"lastLoginAt" gorm:"comment:最后一次登录的时间;column:last_login_at;"`
	Remark         string     `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:500;"`
}

// TableName 商户 SmsCustomers自定义表名 sms_customers
func (SmsCustomers) TableName() string {
	return "sms_customers"
}
