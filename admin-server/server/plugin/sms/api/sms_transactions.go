package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var SmsTransactions = new(smsTransactions)

type smsTransactions struct {}

// CreateSmsTransactions 创建交易记录
// @Tags SmsTransactions
// @Summary 创建交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsTransactions true "创建交易记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsTransactions/createSmsTransactions [post]
func (a *smsTransactions) CreateSmsTransactions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsTransactions
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsTransactions.CreateSmsTransactions(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSmsTransactions 删除交易记录
// @Tags SmsTransactions
// @Summary 删除交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsTransactions true "删除交易记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsTransactions/deleteSmsTransactions [delete]
func (a *smsTransactions) DeleteSmsTransactions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsTransactions.DeleteSmsTransactions(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteSmsTransactionsByIds 批量删除交易记录
// @Tags SmsTransactions
// @Summary 批量删除交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsTransactions/deleteSmsTransactionsByIds [delete]
func (a *smsTransactions) DeleteSmsTransactionsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsTransactions.DeleteSmsTransactionsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsTransactions 更新交易记录
// @Tags SmsTransactions
// @Summary 更新交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsTransactions true "更新交易记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsTransactions/updateSmsTransactions [put]
func (a *smsTransactions) UpdateSmsTransactions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.SmsTransactions
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsTransactions.UpdateSmsTransactions(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindSmsTransactions 用id查询交易记录
// @Tags SmsTransactions
// @Summary 用id查询交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询交易记录"
// @Success 200 {object} response.Response{data=model.SmsTransactions,msg=string} "查询成功"
// @Router /smsTransactions/findSmsTransactions [get]
func (a *smsTransactions) FindSmsTransactions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsTransactions, err := serviceSmsTransactions.GetSmsTransactions(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(resmsTransactions, c)
}
// GetSmsTransactionsList 分页获取交易记录列表
// @Tags SmsTransactions
// @Summary 分页获取交易记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsTransactionsSearch true "分页获取交易记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsTransactions/getSmsTransactionsList [get]
func (a *smsTransactions) GetSmsTransactionsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.SmsTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsTransactions.GetSmsTransactionsInfoList(ctx,pageInfo)
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
// GetSmsTransactionsPublic 不需要鉴权的交易记录接口
// @Tags SmsTransactions
// @Summary 不需要鉴权的交易记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsTransactions/getSmsTransactionsPublic [get]
func (a *smsTransactions) GetSmsTransactionsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceSmsTransactions.GetSmsTransactionsPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的交易记录接口信息"}, "获取成功", c)
}
