package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsProviders = new(smsProviders)

type smsProviders struct {}

// Init 初始化 服务端 路由信息
func (r *smsProviders) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsProviders").Use(middleware.OperationRecord())
		group.POST("createSmsProviders", apiSmsProviders.CreateSmsProviders)   // 新建服务端
		group.DELETE("deleteSmsProviders", apiSmsProviders.DeleteSmsProviders) // 删除服务端
		group.DELETE("deleteSmsProvidersByIds", apiSmsProviders.DeleteSmsProvidersByIds) // 批量删除服务端
		group.PUT("updateSmsProviders", apiSmsProviders.UpdateSmsProviders)    // 更新服务端
	}
	{
	    group := private.Group("smsProviders")
		group.GET("findSmsProviders", apiSmsProviders.FindSmsProviders)        // 根据ID获取服务端
		group.GET("getSmsProvidersList", apiSmsProviders.GetSmsProvidersList)  // 获取服务端列表
	}
	{
	    group := public.Group("smsProviders")
	    group.GET("getSmsProvidersPublic", apiSmsProviders.GetSmsProvidersPublic)  // 服务端开放接口
	}
}
