package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsPlatformProviderBusinessMapping = new(smsPlatformProviderBusinessMapping)

type smsPlatformProviderBusinessMapping struct {}

// Init 初始化 平台子业务 路由信息
func (r *smsPlatformProviderBusinessMapping) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsPlatformProviderBusinessMapping").Use(middleware.OperationRecord())
		group.POST("createSmsPlatformProviderBusinessMapping", apiSmsPlatformProviderBusinessMapping.CreateSmsPlatformProviderBusinessMapping)   // 新建平台子业务
		group.DELETE("deleteSmsPlatformProviderBusinessMapping", apiSmsPlatformProviderBusinessMapping.DeleteSmsPlatformProviderBusinessMapping) // 删除平台子业务
		group.DELETE("deleteSmsPlatformProviderBusinessMappingByIds", apiSmsPlatformProviderBusinessMapping.DeleteSmsPlatformProviderBusinessMappingByIds) // 批量删除平台子业务
		group.PUT("updateSmsPlatformProviderBusinessMapping", apiSmsPlatformProviderBusinessMapping.UpdateSmsPlatformProviderBusinessMapping)    // 更新平台子业务
	}
	{
	    group := private.Group("smsPlatformProviderBusinessMapping")
		group.GET("findSmsPlatformProviderBusinessMapping", apiSmsPlatformProviderBusinessMapping.FindSmsPlatformProviderBusinessMapping)        // 根据ID获取平台子业务
		group.GET("getSmsPlatformProviderBusinessMappingList", apiSmsPlatformProviderBusinessMapping.GetSmsPlatformProviderBusinessMappingList)  // 获取平台子业务列表
	}
	{
	    group := public.Group("smsPlatformProviderBusinessMapping")
	    group.GET("getSmsPlatformProviderBusinessMappingPublic", apiSmsPlatformProviderBusinessMapping.GetSmsPlatformProviderBusinessMappingPublic)  // 平台子业务开放接口
	}
}
