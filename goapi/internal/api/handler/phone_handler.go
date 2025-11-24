package handler

import (
	"context"
	"fmt"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"sms-platform/goapi/internal/utils"
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

// GetPhone 批量获取手机号
// @Summary 批量获取手机号
// @Description 为指定的业务类型批量获取手机号码（1-10个）
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

	// 设置默认值，如果没有指定count则默认1
	if req.Count <= 0 {
		req.Count = 1
	}
	// 限制最多10个
	if req.Count > 10 {
		req.Count = 10
	}

	// 获取用户ID
	userID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// 添加请求上下文信息
	ctx := context.WithValue(c.Request.Context(), "ip_address", c.ClientIP())
	ctx = context.WithValue(ctx, "user_agent", c.GetHeader("User-Agent"))

	// 初始化响应结构
	response := &dto.GetPhoneResponse{
		Phones:       make([]dto.PhoneInfo, 0, req.Count),
		TotalCost:    0,
		SuccessCount: 0,
		FailedCount:  0,
	}

	result, errCode := h.phoneService.GetPhone(ctx, userID, req.BusinessType, req.CardType, req.Count)
	if errCode != common.CodeSuccess {
		common.RespondError(c, errCode)
		return
	}

	for _, phone := range result {
		// 	// 添加成功获取的手机号
		phoneInfo := dto.PhoneInfo{
			PhoneNumber: phone.PhoneNumber,
			CountryCode: phone.CountryCode,
			Cost:        phone.Cost,
			ValidUntil:  phone.ValidUntil,
			ProviderID:  phone.ProviderID,
		}
		response.Phones = append(response.Phones, phoneInfo)
		response.TotalCost += phone.Cost
		response.SuccessCount++
		response.RemainingBalance = phone.Balance // 使用最新的余额
	}

	common.RespondSuccess(c, response)
}

// GetCode 批量获取验证码
// @Summary 批量获取验证码
// @Description 批量获取指定手机号的短信验证码。如果验证码还未获取到，返回等待状态，客户端需要再次请求。
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

	// 验证手机号数量
	if len(req.PhoneNumbers) == 0 {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "手机号列表不能为空")
		return
	}
	if len(req.PhoneNumbers) > 10 {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "手机号数量不能超过10个")
		return
	}

	// 获取用户ID
	userID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// 添加请求上下文信息
	ctx := context.WithValue(c.Request.Context(), "ip_address", c.ClientIP())
	ctx = context.WithValue(ctx, "user_agent", c.GetHeader("User-Agent"))

	// 初始化响应结构
	response := &dto.GetCodeResponse{
		Codes:        make([]dto.CodeInfo, 0, len(req.PhoneNumbers)),
		SuccessCount: 0,
		PendingCount: 0,
		FailedCount:  0,
	}

	// 使用 channel 和 goroutine 并发获取验证码
	type codeResult struct {
		phone  string
		result interface{} // 可能是 CodeInfo 或 error
	}

	resultChan := make(chan codeResult, len(req.PhoneNumbers))

	// 启动并发获取
	for _, phone := range req.PhoneNumbers {
		go func(phoneNumber string) {
			results, err := h.phoneService.GetCode(ctx, userID, phoneNumber)
			if err != nil {
				resultChan <- codeResult{phone: phoneNumber, result: err}
			} else if len(results) > 0 {
				// 取第一个结果
				result := results[0]
				codeInfo := dto.CodeInfo{
					PhoneNumber: phoneNumber,
					Code:        result.Code,
					Message:     result.Message,
					ReceivedAt:  result.ReceivedAt,
					ProviderID:  result.ProviderID,
				}

				// 根据 code 是否为空判断状态
				if result.Code != "" {
					codeInfo.Status = "success"
				} else {
					// code 为空表示等待中，客户端需要再次请求
					codeInfo.Status = "pending"
				}

				resultChan <- codeResult{phone: phoneNumber, result: codeInfo}
			} else {
				resultChan <- codeResult{phone: phoneNumber, result: fmt.Errorf("no code result")}
			}
		}(phone)
	}

	// 收集结果
	for i := 0; i < len(req.PhoneNumbers); i++ {
		result := <-resultChan
		switch r := result.result.(type) {
		case dto.CodeInfo:
			response.Codes = append(response.Codes, r)
			if r.Status == "success" {
				response.SuccessCount++
			} else if r.Status == "pending" {
				response.PendingCount++
			}
		case error:
			// 处理错误情况
			codeInfo := dto.CodeInfo{
				PhoneNumber: result.phone,
				Code:        "",
				Message:     r.Error(),
				ReceivedAt:  time.Now(),
				ProviderID:  "",
				Status:      "failed",
			}
			response.FailedCount++
			response.Codes = append(response.Codes, codeInfo)
		}
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
	customerID, ok := utils.RequireCustomerID(c)
	if !ok {
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
