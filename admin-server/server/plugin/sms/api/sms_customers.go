package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var SmsCustomers = new(smsCustomers)

type smsCustomers struct{}

// CreateSmsCustomers 创建商户
// @Tags SmsCustomers
// @Summary 创建商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.CreateSmsCustomersReq true "创建商户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /smsCustomers/createSmsCustomers [post]
func (a *smsCustomers) CreateSmsCustomers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req request.CreateSmsCustomersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取客户端IP地址
	clientIP := c.ClientIP()
	if req.RegistrationIp == nil || *req.RegistrationIp == "" {
		req.RegistrationIp = &clientIP
	}

	err = serviceSmsCustomers.CreateSmsCustomers(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSmsCustomers 删除商户
// @Tags SmsCustomers
// @Summary 删除商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsCustomers true "删除商户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /smsCustomers/deleteSmsCustomers [delete]
func (a *smsCustomers) DeleteSmsCustomers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSmsCustomers.DeleteSmsCustomers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSmsCustomersByIds 批量删除商户
// @Tags SmsCustomers
// @Summary 批量删除商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /smsCustomers/deleteSmsCustomersByIds [delete]
func (a *smsCustomers) DeleteSmsCustomersByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSmsCustomers.DeleteSmsCustomersByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSmsCustomers 更新商户
// @Tags SmsCustomers
// @Summary 更新商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.UpdateSmsCustomersReq true "更新商户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /smsCustomers/updateSmsCustomers [put]
func (a *smsCustomers) UpdateSmsCustomers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req request.UpdateSmsCustomersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSmsCustomers.UpdateSmsCustomers(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSmsCustomers 用id查询商户
// @Tags SmsCustomers
// @Summary 用id查询商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商户"
// @Success 200 {object} response.Response{data=model.SmsCustomers,msg=string} "查询成功"
// @Router /smsCustomers/findSmsCustomers [get]
func (a *smsCustomers) FindSmsCustomers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	resmsCustomers, err := serviceSmsCustomers.GetSmsCustomers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resmsCustomers, c)
}

// GetSmsCustomersList 分页获取商户列表
// @Tags SmsCustomers
// @Summary 分页获取商户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsCustomersSearch true "分页获取商户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /smsCustomers/getSmsCustomersList [get]
func (a *smsCustomers) GetSmsCustomersList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.SmsCustomersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSmsCustomers.GetSmsCustomersInfoList(ctx, pageInfo)
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

// GetSmsCustomersPublic 不需要鉴权的商户接口
// @Tags SmsCustomers
// @Summary 不需要鉴权的商户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsCustomers/getSmsCustomersPublic [get]
func (a *smsCustomers) GetSmsCustomersPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceSmsCustomers.GetSmsCustomersPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的商户接口信息"}, "获取成功", c)
}

// CreditDebitSmsCustomers 充值/扣费
// @Tags SmsCustomers
// @Summary 充值/扣费
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.CreditDebitSmsCustomersReq true "充值/扣费"
// @Success 200 {object} response.Response{msg=string} "操作成功"
// @Router /smsCustomers/creditDebit [post]
func (a *smsCustomers) CreditDebitSmsCustomers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req request.CreditDebitSmsCustomersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = serviceSmsCustomers.CreditDebitSmsCustomers(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("充值/扣费失败!", zap.Error(err))
		response.FailWithMessage("操作失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// ConfigureBusiness 配置商户业务
// @Tags SmsCustomers
// @Summary 配置商户业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.ConfigureBusinessReq true "配置商户业务"
// @Success 200 {object} response.Response{msg=string} "配置成功"
// @Router /smsCustomers/configureBusiness [post]
func (a *smsCustomers) ConfigureBusiness(c *gin.Context) {
	ctx := c.Request.Context()

	var req request.ConfigureBusinessReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("配置业务参数绑定失败!", zap.Error(err), zap.Any("request", c.Request.Body))
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	global.GVA_LOG.Info("收到业务配置请求",
		zap.Int64("customerId", req.CustomerID),
		zap.Int("businessCount", len(req.BusinessConfig)),
		zap.Any("businessConfig", req.BusinessConfig))

	err = serviceSmsCustomers.ConfigureBusiness(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("配置业务失败!", zap.Error(err))
		response.FailWithMessage("配置失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("配置成功", c)
}

// AdjustFrozenAmount 调整冻结金额
// @Tags SmsCustomers
// @Summary 调整冻结金额
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.AdjustFrozenAmountReq true "调整冻结金额"
// @Success 200 {object} response.Response{msg=string} "调整成功"
// @Router /smsCustomers/adjustFrozenAmount [post]
func (a *smsCustomers) AdjustFrozenAmount(c *gin.Context) {
	ctx := c.Request.Context()

	var req request.AdjustFrozenAmountReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = serviceSmsCustomers.AdjustFrozenAmount(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("调整冻结金额失败!", zap.Error(err))
		response.FailWithMessage("调整失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("调整成功", c)
}

// GetBusinessConfig 获取商户业务配置
// @Tags SmsCustomers
// @Summary 获取商户业务配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param customerId query int true "商户ID"
// @Success 200 {object} response.Response{data=[]model.SmsCustomerBusinessConfig} "获取成功"
// @Router /smsCustomers/getBusinessConfig [get]
func (a *smsCustomers) GetBusinessConfig(c *gin.Context) {
	customerID := c.Query("customerId")
	if customerID == "" {
		response.FailWithMessage("商户ID不能为空", c)
		return
	}

	var configs []model.SmsCustomerBusinessConfig
	err := global.GVA_DB.Where("customer_id = ?", customerID).Find(&configs).Error
	if err != nil {
		global.GVA_LOG.Error("获取业务配置失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	// 构造响应数据，确保字段类型正确
	type BusinessConfigResponse struct {
		PlatformBusinessTypeID int64   `json:"platformBusinessTypeId"`
		BusinessCode           string  `json:"businessCode"`
		BusinessName           string  `json:"businessName"`
		Cost                   float64 `json:"cost"`
		Weight                 int32   `json:"weight"`
		Status                 int     `json:"status"`
	}

	result := make([]BusinessConfigResponse, len(configs))
	for i, config := range configs {
		result[i] = BusinessConfigResponse{
			PlatformBusinessTypeID: config.PlatformBusinessTypeID,
			BusinessCode:           config.BusinessCode,
			BusinessName:           config.BusinessName,
			Cost:                   config.Cost,
			Weight:                 config.Weight,
			Status:                 config.Status,
		}
	}

	response.OkWithData(result, c)
}
