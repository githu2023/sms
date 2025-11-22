package handler

import (
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/service"

	"github.com/gin-gonic/gin"
)

// BalanceHandler handles balance related API requests
type BalanceHandler struct {
	transactionService service.TransactionService
}

// NewBalanceHandler creates a new BalanceHandler instance
func NewBalanceHandler(transactionService service.TransactionService) *BalanceHandler {
	return &BalanceHandler{
		transactionService: transactionService,
	}
}

// BalanceResponse represents the response structure for balance queries
type BalanceResponse struct {
	Balance  float64 `json:"balance" example:"123.45"`
	Currency string  `json:"currency" example:"USD"`
}

// GetBalance handles GET /api/v1/balance and GET /client/v1/balance
// @Summary Get user balance
// @Description Get current balance for authenticated user
// @Tags Balance
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse{data=BalanceResponse} "Success"
// @Failure 401 {object} common.APIResponse "Unauthorized"
// @Failure 500 {object} common.APIResponse "Internal Server Error"
// @Router /api/v1/balance [get]
// @Router /client/v1/balance [get]
func (h *BalanceHandler) GetBalance(c *gin.Context) {
	// Get customer_id from context (set by authentication middleware)
	customerIDInterface, exists := c.Get("customer_id")
	if !exists {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	customerID, ok := customerIDInterface.(uint)
	if !ok {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	// Get balance from transaction service
	balance, err := h.transactionService.GetBalance(c.Request.Context(), int64(customerID))
	if err != nil {
		common.RespondErrorWithMsg(c, common.CodeInternalError, "Failed to get balance: "+err.Error())
		return
	}

	// Prepare response
	response := BalanceResponse{
		Balance:  balance,
		Currency: "USD", // Hardcoded for now, could be configurable
	}

	common.RespondSuccess(c, response)
}
