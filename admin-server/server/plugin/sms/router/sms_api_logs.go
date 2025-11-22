package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsApiLogs = new(smsApiLogs)

type smsApiLogs struct {}

// Init 初始化 访问日志 路由信息
func (r *smsApiLogs) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsApiLogs").Use(middleware.OperationRecord())
		group.POST("createSmsApiLogs", apiSmsApiLogs.CreateSmsApiLogs)   // 新建访问日志
		group.DELETE("deleteSmsApiLogs", apiSmsApiLogs.DeleteSmsApiLogs) // 删除访问日志
		group.DELETE("deleteSmsApiLogsByIds", apiSmsApiLogs.DeleteSmsApiLogsByIds) // 批量删除访问日志
		group.PUT("updateSmsApiLogs", apiSmsApiLogs.UpdateSmsApiLogs)    // 更新访问日志
	}
	{
	    group := private.Group("smsApiLogs")
		group.GET("findSmsApiLogs", apiSmsApiLogs.FindSmsApiLogs)        // 根据ID获取访问日志
		group.GET("getSmsApiLogsList", apiSmsApiLogs.GetSmsApiLogsList)  // 获取访问日志列表
	}
	{
	    group := public.Group("smsApiLogs")
	    group.GET("getSmsApiLogsPublic", apiSmsApiLogs.GetSmsApiLogsPublic)  // 访问日志开放接口
	}
}
