package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsPlatformBusinessTypes = new(smsPlatformBusinessTypes)

type smsPlatformBusinessTypes struct {}

// Init 初始化 平台业务 路由信息
func (r *smsPlatformBusinessTypes) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsPlatformBusinessTypes").Use(middleware.OperationRecord())
		group.POST("createSmsPlatformBusinessTypes", apiSmsPlatformBusinessTypes.CreateSmsPlatformBusinessTypes)   // 新建平台业务
		group.DELETE("deleteSmsPlatformBusinessTypes", apiSmsPlatformBusinessTypes.DeleteSmsPlatformBusinessTypes) // 删除平台业务
		group.DELETE("deleteSmsPlatformBusinessTypesByIds", apiSmsPlatformBusinessTypes.DeleteSmsPlatformBusinessTypesByIds) // 批量删除平台业务
		group.PUT("updateSmsPlatformBusinessTypes", apiSmsPlatformBusinessTypes.UpdateSmsPlatformBusinessTypes)    // 更新平台业务
	}
	{
	    group := private.Group("smsPlatformBusinessTypes")
		group.GET("findSmsPlatformBusinessTypes", apiSmsPlatformBusinessTypes.FindSmsPlatformBusinessTypes)        // 根据ID获取平台业务
		group.GET("getSmsPlatformBusinessTypesList", apiSmsPlatformBusinessTypes.GetSmsPlatformBusinessTypesList)  // 获取平台业务列表
	}
	{
	    group := public.Group("smsPlatformBusinessTypes")
	    group.GET("getSmsPlatformBusinessTypesPublic", apiSmsPlatformBusinessTypes.GetSmsPlatformBusinessTypesPublic)  // 平台业务开放接口
	}
}
