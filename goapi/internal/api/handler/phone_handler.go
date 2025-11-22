package handler

import (
	"context"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

// PhoneHandler 手机号处理器
type PhoneHandler struct {
	phoneService service.PhoneServiceInterface
}

// NewPhoneHandler 创建新的手机号处理器
func NewPhoneHandler(phoneService service.PhoneServiceInterface) *PhoneHandler {
	return &PhoneHandler{
		phoneService: phoneService,
	}
}

// GetPhone 获取手机号
// @Summary 获取手机号
// @Description 为指定的业务类型获取一个可用的手机号码
// @Tags Phone
// @Accept json
// @Produce json
// @Param request body dto.GetPhoneRequest true "获取手机号请求参数"
// @Success 200 {object} common.Response{data=dto.GetPhoneResponse} "成功获取手机号"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 401 {object} common.Response "未授权"
// @Failure 403 {object} common.Response "余额不足"
// @Failure 500 {object} common.Response "服务器内部错误"
// @Router /api/v1/get_phone [post]
// @Router /client/v1/get_phone [post]
func (h *PhoneHandler) GetPhone(c *gin.Context) {
	var req dto.GetPhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "请求参数错误: "+err.Error())
		return
	}

	// 获取用户ID
	customerID, exists := c.Get("customer_id")
	if !exists {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	// 类型断言
	var userID uint
	switch id := customerID.(type) {
	case uint:
		userID = id
	case int:
		userID = uint(id)
	case int64:
		userID = uint(id)
	case float64:
		userID = uint(id)
	default:
		common.RespondErrorWithMsg(c, common.CodeUnauthorized, "用户ID格式错误")
		return
	}

	// 添加请求上下文信息
	ctx := context.WithValue(c.Request.Context(), "ip_address", c.ClientIP())
	ctx = context.WithValue(ctx, "user_agent", c.GetHeader("User-Agent"))

	// 调用服务获取手机号
	result, err := h.phoneService.GetPhone(ctx, userID, req.BusinessType, req.CardType)
	if err != nil {
		// 根据不同错误类型返回对应的错误码
		switch err {
		case common.ErrInsufficientBalance:
			common.RespondError(c, common.CodeInsufficientBalance)
		default:
			if err.Error() == "no healthy providers available" {
				common.RespondError(c, common.CodeThirdPartyError)
			} else {
				common.RespondErrorWithMsg(c, common.CodeInternalError, "获取手机号失败: "+err.Error())
			}
		}
		return
	}

	// 构造响应
	response := &dto.GetPhoneResponse{
		PhoneNumber:      result.PhoneNumber,
		CountryCode:      result.CountryCode,
		Cost:             result.Cost,
		ValidUntil:       result.ValidUntil,
		ProviderID:       result.ProviderID,
		RemainingBalance: result.Balance,
	}

	common.RespondSuccess(c, response)
}

// GetCode 获取验证码
// @Summary 获取验证码
// @Description 获取指定手机号的短信验证码，支持长轮询
// @Tags Phone
// @Accept json
// @Produce json
// @Param request body dto.GetCodeRequest true "获取验证码请求参数"
// @Success 200 {object} common.Response{data=dto.GetCodeResponse} "成功获取验证码"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 401 {object} common.Response "未授权"
// @Failure 404 {object} common.Response "手机号不存在"
// @Failure 408 {object} common.Response "请求超时"
// @Failure 500 {object} common.Response "服务器内部错误"
// @Router /api/v1/get_code [post]
// @Router /client/v1/get_code [post]
func (h *PhoneHandler) GetCode(c *gin.Context) {
	var req dto.GetCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "请求参数错误: "+err.Error())
		return
	}

	// 获取用户ID
	customerID, exists := c.Get("customer_id")
	if !exists {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	// 类型断言
	var userID uint
	switch id := customerID.(type) {
	case uint:
		userID = id
	case int:
		userID = uint(id)
	case int64:
		userID = uint(id)
	case float64:
		userID = uint(id)
	default:
		common.RespondErrorWithMsg(c, common.CodeUnauthorized, "用户ID格式错误")
		return
	}

	// 设置默认超时时间
	timeout := time.Duration(req.Timeout) * time.Second
	if req.Timeout <= 0 {
		timeout = 60 * time.Second // 默认60秒
	}
	if timeout > 5*time.Minute {
		timeout = 5 * time.Minute // 最大5分钟
	}

	// 添加请求上下文信息
	ctx := context.WithValue(c.Request.Context(), "ip_address", c.ClientIP())
	ctx = context.WithValue(ctx, "user_agent", c.GetHeader("User-Agent"))

	// 调用服务获取验证码
	result, err := h.phoneService.GetCode(ctx, userID, req.PhoneNumber, timeout)
	if err != nil {
		// 根据不同错误类型返回对应的错误码
		switch err {
		case common.ErrCodeTimeout:
			common.RespondError(c, common.CodeCodeTimeout)
		default:
			if err.Error() == "no provider could retrieve code for phone" {
				common.RespondError(c, common.CodePhoneNotFound)
			} else {
				common.RespondErrorWithMsg(c, common.CodeGetCodeFailed, "获取验证码失败: "+err.Error())
			}
		}
		return
	}

	// 构造响应
	response := &dto.GetCodeResponse{
		Code:       result.Code,
		Message:    result.Message,
		ReceivedAt: result.ReceivedAt,
		ProviderID: result.ProviderID,
	}

	common.RespondSuccess(c, response)
}

// GetPhoneStatus 获取手机号状态 (可选的额外接口)
// @Summary 获取手机号状态
// @Description 查询指定手机号的当前状态和有效期
// @Tags Phone
// @Accept json
// @Produce json
// @Param phone_number query string true "手机号" example:"+15551234567"
// @Success 200 {object} common.Response{data=map[string]interface{}} "手机号状态信息"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 401 {object} common.Response "未授权"
// @Failure 404 {object} common.Response "手机号不存在"
// @Router /api/v1/phone_status [get]
// @Router /client/v1/phone_status [get]
func (h *PhoneHandler) GetPhoneStatus(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	if phoneNumber == "" {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "手机号参数不能为空")
		return
	}

	// 获取用户ID
	customerID, exists := c.Get("customer_id")
	if !exists {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	// 这里可以实现查询手机号状态的逻辑
	// 目前返回简单的模拟数据
	response := map[string]interface{}{
		"phone_number": phoneNumber,
		"status":       "active",
		"valid_until":  time.Now().Add(30 * time.Minute),
		"customer_id":  customerID,
	}

	common.RespondSuccess(c, response)
}
