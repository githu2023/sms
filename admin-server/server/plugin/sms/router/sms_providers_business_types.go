package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsProvidersBusinessTypes = new(smsProvidersBusinessTypes)

type smsProvidersBusinessTypes struct{}

// Init 初始化 三方业务 路由信息
func (r *smsProvidersBusinessTypes) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("smsProvidersBusinessTypes").Use(middleware.OperationRecord())
		group.POST("createSmsProvidersBusinessTypes", apiSmsProvidersBusinessTypes.CreateSmsProvidersBusinessTypes)             // 新建三方业务
		group.DELETE("deleteSmsProvidersBusinessTypes", apiSmsProvidersBusinessTypes.DeleteSmsProvidersBusinessTypes)           // 删除三方业务
		group.DELETE("deleteSmsProvidersBusinessTypesByIds", apiSmsProvidersBusinessTypes.DeleteSmsProvidersBusinessTypesByIds) // 批量删除三方业务
		group.PUT("updateSmsProvidersBusinessTypes", apiSmsProvidersBusinessTypes.UpdateSmsProvidersBusinessTypes)              // 更新三方业务
	}
	{
		group := private.Group("smsProvidersBusinessTypes")
		group.GET("findSmsProvidersBusinessTypes", apiSmsProvidersBusinessTypes.FindSmsProvidersBusinessTypes)       // 根据ID获取三方业务
		group.GET("getSmsProvidersBusinessTypesList", apiSmsProvidersBusinessTypes.GetSmsProvidersBusinessTypesList) // 获取三方业务列表
	}
	{
		group := public.Group("smsProvidersBusinessTypes")
		group.GET("getSmsProvidersBusinessTypesPublic", apiSmsProvidersBusinessTypes.GetSmsProvidersBusinessTypesPublic) // 三方业务开放接口
	}
}
