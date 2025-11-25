package handler

import (
	"errors"
	"strconv"
	"time"

	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"sms-platform/goapi/internal/utils"

	"github.com/gin-gonic/gin"
)

// AssignmentHandler 手机号分配记录处理器
type AssignmentHandler struct {
	assignmentService service.AssignmentService
}

// NewAssignmentHandler 创建新的手机号分配记录处理器
func NewAssignmentHandler(assignmentService service.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{
		assignmentService: assignmentService,
	}
}

// GetAssignments 获取手机号分配记录
// @Summary 获取手机号分配记录
// @Description 分页获取当前用户的手机号分配历史记录，支持按状态、业务类型和时间范围筛选。
// @Tags Assignment
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param limit query int false "每页数量" default(20)
// @Param status query int false "状态筛选 (1:待取码, 2:已完成, 3:已过期, 4:失败)"
// @Param business_type query string false "业务类型代码" example:"qq"
// @Param start_date query string false "开始日期 (YYYY-MM-DD)" example:"2024-01-01"
// @Param end_date query string false "结束日期 (YYYY-MM-DD)" example:"2024-01-31"
// @Success 200 {object} common.Response{data=dto.AssignmentHistoryResponse} "成功获取分配记录"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 401 {object} common.Response "未授权"
// @Failure 500 {object} common.Response "服务器内部错误"
// @Router /client/v1/assignments [get]
func (h *AssignmentHandler) GetAssignments(c *gin.Context) {
	var req dto.GetAssignmentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "请求参数错误: "+err.Error())
		return
	}

	// Set default values
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 20
	}

	customerIDInt64, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	var startDate *time.Time
	if req.StartDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, "开始日期格式错误，应为 YYYY-MM-DD")
			return
		}
		startDate = &parsedDate
	}

	var endDate *time.Time
	if req.EndDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, "结束日期格式错误，应为 YYYY-MM-DD")
			return
		}
		// 结束日期设置为当天的最后一刻
		endOfDay := parsedDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		endDate = &endOfDay
	}

	assignments, total, err := h.assignmentService.GetAssignments(
		c.Request.Context(),
		customerIDInt64,
		req.Page,
		req.Limit,
		req.Status,
		req.BusinessType,
		startDate,
		endDate,
	)
	if err != nil {
		if errors.Is(err, common.ErrInvalidTimeRange) {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, err.Error())
			return
		}
		common.RespondErrorWithMsg(c, common.CodeInternalError, "获取分配记录失败: "+err.Error()) // Changed c.Request.Context() to c
		return
	}

	// 转换领域模型为DTO
	items := make([]dto.AssignmentHistoryItem, len(assignments))
	for i, assignment := range assignments {
		items[i] = buildAssignmentHistoryItem(assignment)
	}

	common.RespondSuccess(c, dto.AssignmentHistoryResponse{
		Items: items,
		Pagination: struct {
			Total int64 "json:\"total\""
			Page  int   "json:\"page\""
			Limit int   "json:\"limit\""
		}{
			Total: total,
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
}

// GetCostStatistics 获取手机号分配成本统计
// @Summary 获取手机号分配成本统计
// @Description 获取当前用户的手机号分配总成本和总数量，支持按时间范围筛选。
// @Tags Assignment
// @Accept json
// @Produce json
// @Param start_date query string false "开始日期 (YYYY-MM-DD)" example:"2024-01-01"
// @Param end_date query string false "结束日期 (YYYY-MM-DD)" example:"2024-01-31"
// @Success 200 {object} common.Response{data=dto.CostStatisticsResponse} "成功获取成本统计"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 401 {object} common.Response "未授权"
// @Failure 500 {object} common.Response "服务器内部错误"
// @Router /client/v1/assignments/statistics [get]
func (h *AssignmentHandler) GetCostStatistics(c *gin.Context) {
	var req struct {
		StartDate string `form:"start_date" binding:"omitempty,datetime=2006-01-02" example:"2024-01-01"`
		EndDate   string `form:"end_date" binding:"omitempty,datetime=2006-01-02" example:"2024-01-31"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		common.RespondErrorWithMsg(c, common.CodeBadRequest, "请求参数错误: "+err.Error())
		return
	}

	customerIDInt64, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	var startDate *time.Time
	if req.StartDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, "开始日期格式错误，应为 YYYY-MM-DD")
			return
		}
		startDate = &parsedDate
	}

	var endDate *time.Time
	if req.EndDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, "结束日期格式错误，应为 YYYY-MM-DD")
			return
		}
		endOfDay := parsedDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		endDate = &endOfDay
	}

	stats, err := h.assignmentService.GetCostStatistics(c.Request.Context(), customerIDInt64, startDate, endDate)
	if err != nil {
		if errors.Is(err, common.ErrInvalidTimeRange) {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, err.Error())
			return
		}
		common.RespondErrorWithMsg(c, common.CodeInternalError, "获取成本统计失败: "+err.Error()) // Changed c.Request.Context() to c
		return
	}

	common.RespondSuccess(c, dto.CostStatisticsResponse{
		TotalCost:  stats.TotalCost,
		TotalCount: stats.TotalCount,
	})
}

// GetRecentAssignments 返回最近的手机号获取记录
func (h *AssignmentHandler) GetRecentAssignments(c *gin.Context) {
	limit := 5
	if limitStr := c.Query("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 || parsedLimit > 50 {
			common.RespondErrorWithMsg(c, common.CodeBadRequest, "limit 需要是 1-50 的整数")
			return
		}
		limit = parsedLimit
	}

	customerIDInt64, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	assignments, err := h.assignmentService.GetRecentAssignments(c.Request.Context(), customerIDInt64, limit)
	if err != nil {
		common.RespondErrorWithMsg(c, common.CodeInternalError, "获取最近分配记录失败: "+err.Error())
		return
	}

	items := make([]dto.AssignmentHistoryItem, len(assignments))
	for i, assignment := range assignments {
		items[i] = buildAssignmentHistoryItem(assignment)
	}

	common.RespondSuccess(c, dto.RecentAssignmentsResponse{
		Items: items,
	})
}

func buildAssignmentHistoryItem(assignment *domain.PhoneAssignment) dto.AssignmentHistoryItem {
	phoneNumber := ""
	if assignment.PhoneNumber != nil {
		phoneNumber = *assignment.PhoneNumber
	}

	verificationCode := ""
	if assignment.VerificationCode != nil {
		verificationCode = *assignment.VerificationCode
	}

	status := 1 // default pending
	if assignment.Status != nil {
		switch *assignment.Status {
		case "pending":
			status = 1
		case "completed":
			status = 2
		case "expired":
			status = 3
		case "failed":
			status = 4
		}
	}

	cost := 0.0
	if assignment.MerchantFee != nil {
		cost = float64(*assignment.MerchantFee)
	}

	businessType := assignment.BusinessCode
	providerName := ""
	if assignment.Provider != nil && assignment.Provider.Name != nil {
		providerName = *assignment.Provider.Name
	}

	return dto.AssignmentHistoryItem{
		ID:               assignment.ID,
		PhoneNumber:      phoneNumber,
		BusinessType:     businessType,
		CardType:         "virtual",
		VerificationCode: verificationCode,
		Cost:             cost,
		Status:           status,
		CreatedAt:        assignment.CreatedAt,
		ProviderName:     providerName,
	}
}
