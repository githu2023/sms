package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsCustomers = new(smsCustomers)

type smsCustomers struct{}

// Init 初始化 商户 路由信息
func (r *smsCustomers) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("smsCustomers").Use(middleware.OperationRecord())
		group.POST("createSmsCustomers", apiSmsCustomers.CreateSmsCustomers)             // 新建商户
		group.DELETE("deleteSmsCustomers", apiSmsCustomers.DeleteSmsCustomers)           // 删除商户
		group.DELETE("deleteSmsCustomersByIds", apiSmsCustomers.DeleteSmsCustomersByIds) // 批量删除商户
		group.PUT("updateSmsCustomers", apiSmsCustomers.UpdateSmsCustomers)              // 更新商户
		group.POST("creditDebit", apiSmsCustomers.CreditDebitSmsCustomers)               // 充值/扣费
		group.POST("configureBusiness", apiSmsCustomers.ConfigureBusiness)               // 配置商户业务
		group.POST("adjustFrozenAmount", apiSmsCustomers.AdjustFrozenAmount)             // 调整冻结金额
	}
	{
		group := private.Group("smsCustomers")
		group.GET("findSmsCustomers", apiSmsCustomers.FindSmsCustomers)       // 根据ID获取商户
		group.GET("getSmsCustomersList", apiSmsCustomers.GetSmsCustomersList) // 获取商户列表
		group.GET("getBusinessConfig", apiSmsCustomers.GetBusinessConfig)     // 获取商户业务配置
	}
	{
		group := public.Group("smsCustomers")
		group.GET("getSmsCustomersPublic", apiSmsCustomers.GetSmsCustomersPublic) // 商户开放接口
	}
}
