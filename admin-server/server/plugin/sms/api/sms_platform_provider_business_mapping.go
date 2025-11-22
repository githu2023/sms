package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var SmsPlatformProviderBusinessMapping = new(smsPlatformProviderBusinessMapping)

type smsPlatformProviderBusinessMapping struct {}

// CreateSmsPlatformProviderBusinessMapping 创建平台子业务
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 创建平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformProviderBusinessMapping true "创建平台子业务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsPlatformProviderBusinessMapping/createSmsPlatformProviderBusinessMapping [post]
func (a *smsPlatformProviderBusinessMapping) CreateSmsPlatformProviderBusinessMapping(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsPlatformProviderBusinessMapping
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsPlatformProviderBusinessMapping.CreateSmsPlatformProviderBusinessMapping(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSmsPlatformProviderBusinessMapping 删除平台子业务
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 删除平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformProviderBusinessMapping true "删除平台子业务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMapping [delete]
func (a *smsPlatformProviderBusinessMapping) DeleteSmsPlatformProviderBusinessMapping(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsPlatformProviderBusinessMapping.DeleteSmsPlatformProviderBusinessMapping(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteSmsPlatformProviderBusinessMappingByIds 批量删除平台子业务
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 批量删除平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMappingByIds [delete]
func (a *smsPlatformProviderBusinessMapping) DeleteSmsPlatformProviderBusinessMappingByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsPlatformProviderBusinessMapping.DeleteSmsPlatformProviderBusinessMappingByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsPlatformProviderBusinessMapping 更新平台子业务
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 更新平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformProviderBusinessMapping true "更新平台子业务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsPlatformProviderBusinessMapping/updateSmsPlatformProviderBusinessMapping [put]
func (a *smsPlatformProviderBusinessMapping) UpdateSmsPlatformProviderBusinessMapping(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsPlatformProviderBusinessMapping
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsPlatformProviderBusinessMapping.UpdateSmsPlatformProviderBusinessMapping(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindSmsPlatformProviderBusinessMapping 用id查询平台子业务
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 用id查询平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询平台子业务"
// @Success 200 {object} response.Response{data=model.SmsPlatformProviderBusinessMapping,msg=string} "查询成功"
// @Router /smsPlatformProviderBusinessMapping/findSmsPlatformProviderBusinessMapping [get]
func (a *smsPlatformProviderBusinessMapping) FindSmsPlatformProviderBusinessMapping(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsPlatformProviderBusinessMapping, err := serviceSmsPlatformProviderBusinessMapping.GetSmsPlatformProviderBusinessMapping(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(resmsPlatformProviderBusinessMapping, c)
}
// GetSmsPlatformProviderBusinessMappingList 分页获取平台子业务列表
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 分页获取平台子业务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsPlatformProviderBusinessMappingSearch true "分页获取平台子业务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingList [get]
func (a *smsPlatformProviderBusinessMapping) GetSmsPlatformProviderBusinessMappingList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.SmsPlatformProviderBusinessMappingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsPlatformProviderBusinessMapping.GetSmsPlatformProviderBusinessMappingInfoList(ctx,pageInfo)
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
// GetSmsPlatformProviderBusinessMappingPublic 不需要鉴权的平台子业务接口
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 不需要鉴权的平台子业务接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingPublic [get]
func (a *smsPlatformProviderBusinessMapping) GetSmsPlatformProviderBusinessMappingPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceSmsPlatformProviderBusinessMapping.GetSmsPlatformProviderBusinessMappingPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的平台子业务接口信息"}, "获取成功", c)
}
