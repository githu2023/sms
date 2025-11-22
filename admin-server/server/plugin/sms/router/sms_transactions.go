package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsTransactions = new(smsTransactions)

type smsTransactions struct {}

// Init 初始化 交易记录 路由信息
func (r *smsTransactions) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsTransactions").Use(middleware.OperationRecord())
		group.POST("createSmsTransactions", apiSmsTransactions.CreateSmsTransactions)   // 新建交易记录
		group.DELETE("deleteSmsTransactions", apiSmsTransactions.DeleteSmsTransactions) // 删除交易记录
		group.DELETE("deleteSmsTransactionsByIds", apiSmsTransactions.DeleteSmsTransactionsByIds) // 批量删除交易记录
		group.PUT("updateSmsTransactions", apiSmsTransactions.UpdateSmsTransactions)    // 更新交易记录
	}
	{
	    group := private.Group("smsTransactions")
		group.GET("findSmsTransactions", apiSmsTransactions.FindSmsTransactions)        // 根据ID获取交易记录
		group.GET("getSmsTransactionsList", apiSmsTransactions.GetSmsTransactionsList)  // 获取交易记录列表
	}
	{
	    group := public.Group("smsTransactions")
	    group.GET("getSmsTransactionsPublic", apiSmsTransactions.GetSmsTransactionsPublic)  // 交易记录开放接口
	}
}
