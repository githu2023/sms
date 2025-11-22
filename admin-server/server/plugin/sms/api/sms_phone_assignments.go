package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var SmsPhoneAssignments = new(smsPhoneAssignments)

type smsPhoneAssignments struct {}

// CreateSmsPhoneAssignments 创建号码记录
// @Tags SmsPhoneAssignments
// @Summary 创建号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPhoneAssignments true "创建号码记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsPhoneAssignments/createSmsPhoneAssignments [post]
func (a *smsPhoneAssignments) CreateSmsPhoneAssignments(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsPhoneAssignments
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsPhoneAssignments.CreateSmsPhoneAssignments(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSmsPhoneAssignments 删除号码记录
// @Tags SmsPhoneAssignments
// @Summary 删除号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPhoneAssignments true "删除号码记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsPhoneAssignments/deleteSmsPhoneAssignments [delete]
func (a *smsPhoneAssignments) DeleteSmsPhoneAssignments(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsPhoneAssignments.DeleteSmsPhoneAssignments(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteSmsPhoneAssignmentsByIds 批量删除号码记录
// @Tags SmsPhoneAssignments
// @Summary 批量删除号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsPhoneAssignments/deleteSmsPhoneAssignmentsByIds [delete]
func (a *smsPhoneAssignments) DeleteSmsPhoneAssignmentsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsPhoneAssignments.DeleteSmsPhoneAssignmentsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsPhoneAssignments 更新号码记录
// @Tags SmsPhoneAssignments
// @Summary 更新号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPhoneAssignments true "更新号码记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsPhoneAssignments/updateSmsPhoneAssignments [put]
func (a *smsPhoneAssignments) UpdateSmsPhoneAssignments(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsPhoneAssignments
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsPhoneAssignments.UpdateSmsPhoneAssignments(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindSmsPhoneAssignments 用id查询号码记录
// @Tags SmsPhoneAssignments
// @Summary 用id查询号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询号码记录"
// @Success 200 {object} response.Response{data=model.SmsPhoneAssignments,msg=string} "查询成功"
// @Router /smsPhoneAssignments/findSmsPhoneAssignments [get]
func (a *smsPhoneAssignments) FindSmsPhoneAssignments(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsPhoneAssignments, err := serviceSmsPhoneAssignments.GetSmsPhoneAssignments(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(resmsPhoneAssignments, c)
}
// GetSmsPhoneAssignmentsList 分页获取号码记录列表
// @Tags SmsPhoneAssignments
// @Summary 分页获取号码记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsPhoneAssignmentsSearch true "分页获取号码记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsPhoneAssignments/getSmsPhoneAssignmentsList [get]
func (a *smsPhoneAssignments) GetSmsPhoneAssignmentsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.SmsPhoneAssignmentsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsPhoneAssignments.GetSmsPhoneAssignmentsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetSmsPhoneAssignmentsPublic 不需要鉴权的号码记录接口
// @Tags SmsPhoneAssignments
// @Summary 不需要鉴权的号码记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsPhoneAssignments/getSmsPhoneAssignmentsPublic [get]
func (a *smsPhoneAssignments) GetSmsPhoneAssignmentsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceSmsPhoneAssignments.GetSmsPhoneAssignmentsPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的号码记录接口信息"}, "获取成功", c)
}
