package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var SmsPhoneAssignments = new(smsPhoneAssignments)

type smsPhoneAssignments struct {}

// Init 初始化 号码记录 路由信息
func (r *smsPhoneAssignments) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("smsPhoneAssignments").Use(middleware.OperationRecord())
		group.POST("createSmsPhoneAssignments", apiSmsPhoneAssignments.CreateSmsPhoneAssignments)   // 新建号码记录
		group.DELETE("deleteSmsPhoneAssignments", apiSmsPhoneAssignments.DeleteSmsPhoneAssignments) // 删除号码记录
		group.DELETE("deleteSmsPhoneAssignmentsByIds", apiSmsPhoneAssignments.DeleteSmsPhoneAssignmentsByIds) // 批量删除号码记录
		group.PUT("updateSmsPhoneAssignments", apiSmsPhoneAssignments.UpdateSmsPhoneAssignments)    // 更新号码记录
	}
	{
	    group := private.Group("smsPhoneAssignments")
		group.GET("findSmsPhoneAssignments", apiSmsPhoneAssignments.FindSmsPhoneAssignments)        // 根据ID获取号码记录
		group.GET("getSmsPhoneAssignmentsList", apiSmsPhoneAssignments.GetSmsPhoneAssignmentsList)  // 获取号码记录列表
	}
	{
	    group := public.Group("smsPhoneAssignments")
	    group.GET("getSmsPhoneAssignmentsPublic", apiSmsPhoneAssignments.GetSmsPhoneAssignmentsPublic)  // 号码记录开放接口
	}
}
