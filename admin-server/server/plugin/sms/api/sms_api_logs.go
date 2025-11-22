package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var SmsApiLogs = new(smsApiLogs)

type smsApiLogs struct {}

// CreateSmsApiLogs 创建访问日志
// @Tags SmsApiLogs
// @Summary 创建访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsApiLogs true "创建访问日志"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsApiLogs/createSmsApiLogs [post]
func (a *smsApiLogs) CreateSmsApiLogs(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsApiLogs
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsApiLogs.CreateSmsApiLogs(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSmsApiLogs 删除访问日志
// @Tags SmsApiLogs
// @Summary 删除访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsApiLogs true "删除访问日志"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsApiLogs/deleteSmsApiLogs [delete]
func (a *smsApiLogs) DeleteSmsApiLogs(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsApiLogs.DeleteSmsApiLogs(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteSmsApiLogsByIds 批量删除访问日志
// @Tags SmsApiLogs
// @Summary 批量删除访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsApiLogs/deleteSmsApiLogsByIds [delete]
func (a *smsApiLogs) DeleteSmsApiLogsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsApiLogs.DeleteSmsApiLogsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsApiLogs 更新访问日志
// @Tags SmsApiLogs
// @Summary 更新访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsApiLogs true "更新访问日志"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsApiLogs/updateSmsApiLogs [put]
func (a *smsApiLogs) UpdateSmsApiLogs(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsApiLogs
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsApiLogs.UpdateSmsApiLogs(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindSmsApiLogs 用id查询访问日志
// @Tags SmsApiLogs
// @Summary 用id查询访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询访问日志"
// @Success 200 {object} response.Response{data=model.SmsApiLogs,msg=string} "查询成功"
// @Router /smsApiLogs/findSmsApiLogs [get]
func (a *smsApiLogs) FindSmsApiLogs(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsApiLogs, err := serviceSmsApiLogs.GetSmsApiLogs(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(resmsApiLogs, c)
}
// GetSmsApiLogsList 分页获取访问日志列表
// @Tags SmsApiLogs
// @Summary 分页获取访问日志列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsApiLogsSearch true "分页获取访问日志列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsApiLogs/getSmsApiLogsList [get]
func (a *smsApiLogs) GetSmsApiLogsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.SmsApiLogsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsApiLogs.GetSmsApiLogsInfoList(ctx,pageInfo)
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
// GetSmsApiLogsPublic 不需要鉴权的访问日志接口
// @Tags SmsApiLogs
// @Summary 不需要鉴权的访问日志接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsApiLogs/getSmsApiLogsPublic [get]
func (a *smsApiLogs) GetSmsApiLogsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceSmsApiLogs.GetSmsApiLogsPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的访问日志接口信息"}, "获取成功", c)
}
