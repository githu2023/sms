package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// BusinessConfigItem 业务配置项
type BusinessConfigItem struct {
	PlatformBusinessTypeID int64   `json:"platformBusinessTypeId" binding:"required"`
	BusinessCode           string  `json:"businessCode" binding:"required"`
	BusinessName           string  `json:"businessName" binding:"required"`
	Cost                   float64 `json:"cost" binding:"required,min=0"`
	Weight                 int32   `json:"weight" binding:"min=0"`
	Status                 int     `json:"status" binding:"required,oneof=0 1"`
}

type SmsCustomersSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Username       *string     `json:"username" form:"username"`
	MerchantNo     string      `json:"merchantNo" form:"merchantNo"`
	MerchantName   string      `json:"merchantName" form:"merchantName"`
	request.PageInfo
}

// CreateSmsCustomersReq 创建商户请求结构体(包含明文密码)
type CreateSmsCustomersReq struct {
	MerchantName   string   `json:"merchantName" form:"merchantName" binding:"required"`
	MerchantNo     string   `json:"merchantNo" form:"merchantNo" binding:"required"`
	Username       *string  `json:"username" form:"username" binding:"required"`
	Email          *string  `json:"email" form:"email" binding:"required,email"`
	Password       string   `json:"password" form:"password" binding:"required,min=6"`
	ApiSecretKey   *string  `json:"apiSecretKey" form:"apiSecretKey"`
	Balance        *float64 `json:"balance" form:"balance"`
	ParentID       *int64   `json:"parentId" form:"parentId"`
	Status         *bool    `json:"status" form:"status"`
	RegistrationIp *string  `json:"registrationIp" form:"registrationIp"`
	Remark         string   `json:"remark" form:"remark"`
}

// ConfigureBusinessReq 配置商户业务请求
type ConfigureBusinessReq struct {
	CustomerID     int64                `json:"customerId" binding:"required"`
	BusinessConfig []BusinessConfigItem `json:"businessConfig" binding:"required"`
}

// AdjustFrozenAmountReq 调整冻结金额请求
type AdjustFrozenAmountReq struct {
	ID           uint    `json:"ID" binding:"required"`
	FrozenAmount float64 `json:"frozenAmount" binding:"required,min=0"`
	Remark       string  `json:"remark"`
}
