package handler

import (
	"net/http"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"sms-platform/goapi/internal/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// TransactionHandler 处理交易相关的HTTP请求
type TransactionHandler struct {
	transactionService service.TransactionService
}

// NewTransactionHandler 创建交易处理器
func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

// GetTransactionHistory 获取交易历史记录
// @Summary 获取交易历史记录
// @Description 分页查询当前登录用户的交易历史记录
// @Tags Transaction
// @Accept json
// @Produce json
// @Param limit query int false "每页数量 (默认20, 最大100)" default(20)
// @Param offset query int false "偏移量 (默认0)" default(0)
// @Success 200 {object} dto.TransactionListResponse "交易记录列表"
// @Failure 400 {object} common.APIResponse "参数错误"
// @Failure 401 {object} common.APIResponse "未授权"
// @Failure 500 {object} common.APIResponse "服务器错误"
// @Security ApiKeyAuth
// @Router /api/v1/transactions [get]
func (h *TransactionHandler) GetTransactionHistory(c *gin.Context) {
	// 从上下文获取客户ID (由认证中间件设置)
	customerID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// 解析分页参数
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// 获取交易记录
	transactions, total, err := h.transactionService.GetTransactionHistory(
		c.Request.Context(),
		customerID,
		limit,
		offset,
	)
	if err != nil {
		common.RespondError(c, common.CodeInternalError)
		return
	}

	// 转换为响应DTO
	items := make([]dto.TransactionItem, len(transactions))
	for i, tx := range transactions {
		items[i] = dto.TransactionItem{
			ID:            tx.ID,
			Amount:        tx.Amount,
			BalanceBefore: tx.BalanceBefore,
			BalanceAfter:  tx.BalanceAfter,
			FrozenBefore:  tx.FrozenBefore,
			FrozenAfter:   tx.FrozenAfter,
			Type:          tx.Type,
			ReferenceID:   tx.ReferenceID,
			Notes:         tx.Notes,
			CreatedAt:     tx.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.TransactionListResponse{
		Total:        total,
		Limit:        limit,
		Offset:       offset,
		Transactions: items,
	})
}

// GetTransactionsByType 按类型获取交易记录
// @Summary 按类型获取交易记录
// @Description 分页查询指定类型的交易记录
// @Tags Transaction
// @Accept json
// @Produce json
// @Param type query int true "交易类型 (1:充值, 2:消费)" Enums(1, 2)
// @Param limit query int false "每页数量 (默认20, 最大100)" default(20)
// @Param offset query int false "偏移量 (默认0)" default(0)
// @Success 200 {object} dto.TransactionListResponse "交易记录列表"
// @Failure 400 {object} dto.ErrorResponse "参数错误"
// @Failure 401 {object} dto.ErrorResponse "未授权"
// @Failure 500 {object} dto.ErrorResponse "服务器错误"
// @Security ApiKeyAuth
// @Router /api/v1/transactions/by-type [get]
func (h *TransactionHandler) GetTransactionsByType(c *gin.Context) {
	// 从上下文获取客户ID
	customerID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// 解析交易类型
	transactionType, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	// 解析分页参数
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// 获取交易记录
	transactions, total, err := h.transactionService.GetTransactionsByType(
		c.Request.Context(),
		customerID,
		transactionType,
		limit,
		offset,
	)
	if err != nil {
		common.RespondError(c, common.CodeInternalError)
		return
	}

	// 转换为响应DTO
	items := make([]dto.TransactionItem, len(transactions))
	for i, tx := range transactions {
		items[i] = dto.TransactionItem{
			ID:            tx.ID,
			Amount:        tx.Amount,
			BalanceBefore: tx.BalanceBefore,
			BalanceAfter:  tx.BalanceAfter,
			FrozenBefore:  tx.FrozenBefore,
			FrozenAfter:   tx.FrozenAfter,
			Type:          tx.Type,
			ReferenceID:   tx.ReferenceID,
			Notes:         tx.Notes,
			CreatedAt:     tx.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.TransactionListResponse{
		Total:        total,
		Limit:        limit,
		Offset:       offset,
		Transactions: items,
	})
}

// GetTransactionsByDateRange 按日期范围获取交易记录
// @Summary 按日期范围获取交易记录
// @Description 分页查询指定日期范围内的交易记录
// @Tags Transaction
// @Accept json
// @Produce json
// @Param start_date query string true "开始日期 (格式: 2006-01-02)" format(date)
// @Param end_date query string true "结束日期 (格式: 2006-01-02)" format(date)
// @Param limit query int false "每页数量 (默认20, 最大100)" default(20)
// @Param offset query int false "偏移量 (默认0)" default(0)
// @Success 200 {object} dto.TransactionListResponse "交易记录列表"
// @Failure 400 {object} dto.ErrorResponse "参数错误"
// @Failure 401 {object} dto.ErrorResponse "未授权"
// @Failure 500 {object} dto.ErrorResponse "服务器错误"
// @Security ApiKeyAuth
// @Router /api/v1/transactions/by-date [get]
func (h *TransactionHandler) GetTransactionsByDateRange(c *gin.Context) {
	// 从上下文获取客户ID
	customerID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// 解析日期参数
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	// 设置时间为一天的结束
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	// 解析分页参数
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// 获取交易记录
	transactions, total, err := h.transactionService.GetTransactionsByDateRange(
		c.Request.Context(),
		customerID,
		startDate,
		endDate,
		limit,
		offset,
	)
	if err != nil {
		common.RespondError(c, common.CodeInternalError)
		return
	}

	// 转换为响应DTO
	items := make([]dto.TransactionItem, len(transactions))
	for i, tx := range transactions {
		items[i] = dto.TransactionItem{
			ID:            tx.ID,
			Amount:        tx.Amount,
			BalanceBefore: tx.BalanceBefore,
			BalanceAfter:  tx.BalanceAfter,
			FrozenBefore:  tx.FrozenBefore,
			FrozenAfter:   tx.FrozenAfter,
			Type:          tx.Type,
			ReferenceID:   tx.ReferenceID,
			Notes:         tx.Notes,
			CreatedAt:     tx.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.TransactionListResponse{
		Total:        total,
		Limit:        limit,
		Offset:       offset,
		Transactions: items,
	})
}
