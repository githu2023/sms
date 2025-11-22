package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var SmsProviders = new(smsProviders)

type smsProviders struct {}

// CreateSmsProviders 创建服务端
// @Tags SmsProviders
// @Summary 创建服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProviders true "创建服务端"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsProviders/createSmsProviders [post]
func (a *smsProviders) CreateSmsProviders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsProviders
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsProviders.CreateSmsProviders(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSmsProviders 删除服务端
// @Tags SmsProviders
// @Summary 删除服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProviders true "删除服务端"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsProviders/deleteSmsProviders [delete]
func (a *smsProviders) DeleteSmsProviders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsProviders.DeleteSmsProviders(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteSmsProvidersByIds 批量删除服务端
// @Tags SmsProviders
// @Summary 批量删除服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsProviders/deleteSmsProvidersByIds [delete]
func (a *smsProviders) DeleteSmsProvidersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsProviders.DeleteSmsProvidersByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsProviders 更新服务端
// @Tags SmsProviders
// @Summary 更新服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProviders true "更新服务端"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsProviders/updateSmsProviders [put]
func (a *smsProviders) UpdateSmsProviders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsProviders
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsProviders.UpdateSmsProviders(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindSmsProviders 用id查询服务端
// @Tags SmsProviders
// @Summary 用id查询服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询服务端"
// @Success 200 {object} response.Response{data=model.SmsProviders,msg=string} "查询成功"
// @Router /smsProviders/findSmsProviders [get]
func (a *smsProviders) FindSmsProviders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsProviders, err := serviceSmsProviders.GetSmsProviders(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(resmsProviders, c)
}
// GetSmsProvidersList 分页获取服务端列表
// @Tags SmsProviders
// @Summary 分页获取服务端列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsProvidersSearch true "分页获取服务端列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsProviders/getSmsProvidersList [get]
func (a *smsProviders) GetSmsProvidersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.SmsProvidersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsProviders.GetSmsProvidersInfoList(ctx,pageInfo)
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
// GetSmsProvidersPublic 不需要鉴权的服务端接口
// @Tags SmsProviders
// @Summary 不需要鉴权的服务端接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsProviders/getSmsProvidersPublic [get]
func (a *smsProviders) GetSmsProvidersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceSmsProviders.GetSmsProvidersPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的服务端接口信息"}, "获取成功", c)
}
