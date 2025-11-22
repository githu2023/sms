package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsIpWhitelist = new(smsIpWhitelist)

type smsIpWhitelist struct {}

// Init 初始化 白名单 路由信息
func (r *smsIpWhitelist) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsIpWhitelist").Use(middleware.OperationRecord())
		group.POST("createSmsIpWhitelist", apiSmsIpWhitelist.CreateSmsIpWhitelist)   // 新建白名单
		group.DELETE("deleteSmsIpWhitelist", apiSmsIpWhitelist.DeleteSmsIpWhitelist) // 删除白名单
		group.DELETE("deleteSmsIpWhitelistByIds", apiSmsIpWhitelist.DeleteSmsIpWhitelistByIds) // 批量删除白名单
		group.PUT("updateSmsIpWhitelist", apiSmsIpWhitelist.UpdateSmsIpWhitelist)    // 更新白名单
	}
	{
	    group := private.Group("smsIpWhitelist")
		group.GET("findSmsIpWhitelist", apiSmsIpWhitelist.FindSmsIpWhitelist)        // 根据ID获取白名单
		group.GET("getSmsIpWhitelistList", apiSmsIpWhitelist.GetSmsIpWhitelistList)  // 获取白名单列表
	}
	{
	    group := public.Group("smsIpWhitelist")
	    group.GET("getSmsIpWhitelistPublic", apiSmsIpWhitelist.GetSmsIpWhitelistPublic)  // 白名单开放接口
	}
}
