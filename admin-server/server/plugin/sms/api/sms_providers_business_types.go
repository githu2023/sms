package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var SmsProvidersBusinessTypes = new(smsProvidersBusinessTypes)

type smsProvidersBusinessTypes struct{}

// CreateSmsProvidersBusinessTypes 创建三方业务
// @Tags SmsProvidersBusinessTypes
// @Summary 创建三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProvidersBusinessTypes true "创建三方业务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsProvidersBusinessTypes/createSmsProvidersBusinessTypes [post]
func (a *smsProvidersBusinessTypes) CreateSmsProvidersBusinessTypes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.SmsProvidersBusinessTypes
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsProvidersBusinessTypes.CreateSmsProvidersBusinessTypes(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSmsProvidersBusinessTypes 删除三方业务
// @Tags SmsProvidersBusinessTypes
// @Summary 删除三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProvidersBusinessTypes true "删除三方业务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypes [delete]
func (a *smsProvidersBusinessTypes) DeleteSmsProvidersBusinessTypes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsProvidersBusinessTypes.DeleteSmsProvidersBusinessTypes(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSmsProvidersBusinessTypesByIds 批量删除三方业务
// @Tags SmsProvidersBusinessTypes
// @Summary 批量删除三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除三方业务"
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypesByIds [delete]
func (a *smsProvidersBusinessTypes) DeleteSmsProvidersBusinessTypesByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsProvidersBusinessTypes.DeleteSmsProvidersBusinessTypesByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsProvidersBusinessTypes 更新三方业务
// @Tags SmsProvidersBusinessTypes
// @Summary 更新三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProvidersBusinessTypes true "更新三方业务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsProvidersBusinessTypes/updateSmsProvidersBusinessTypes [put]
func (a *smsProvidersBusinessTypes) UpdateSmsProvidersBusinessTypes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.SmsProvidersBusinessTypes
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsProvidersBusinessTypes.UpdateSmsProvidersBusinessTypes(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSmsProvidersBusinessTypes 用id查询三方业务
// @Tags SmsProvidersBusinessTypes
// @Summary 用id查询三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsProvidersBusinessTypes true "用id查询三方业务"
// @Success 200 {object} response.Response{data=model.SmsProvidersBusinessTypes,msg=string} "查询成功"
// @Router /smsProvidersBusinessTypes/findSmsProvidersBusinessTypes [get]
func (a *smsProvidersBusinessTypes) FindSmsProvidersBusinessTypes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsProvidersBusinessTypes, err := serviceSmsProvidersBusinessTypes.GetSmsProvidersBusinessTypes(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resmsProvidersBusinessTypes, c)
}

// GetSmsProvidersBusinessTypesList 分页获取三方业务列表
// @Tags SmsProvidersBusinessTypes
// @Summary 分页获取三方业务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsProvidersBusinessTypesSearch true "分页获取三方业务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsProvidersBusinessTypes/getSmsProvidersBusinessTypesList [get]
func (a *smsProvidersBusinessTypes) GetSmsProvidersBusinessTypesList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.SmsProvidersBusinessTypesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsProvidersBusinessTypes.GetSmsProvidersBusinessTypesInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSmsProvidersBusinessTypesPublic 不需要鉴权的三方业务接口
// @Tags SmsProvidersBusinessTypes
// @Summary 不需要鉴权的三方业务接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsProvidersBusinessTypesSearch true "分页获取三方业务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsProvidersBusinessTypes/getSmsProvidersBusinessTypesPublic [get]
func (a *smsProvidersBusinessTypes) GetSmsProvidersBusinessTypesPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceSmsProvidersBusinessTypes.GetSmsProvidersBusinessTypesPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的三方业务接口信息",
	}, "获取成功", c)
}
