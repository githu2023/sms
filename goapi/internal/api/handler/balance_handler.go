package handler

import (
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/service"
	"sms-platform/goapi/internal/utils"

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
	Balance      float64 `json:"balance" example:"123.45"`
	FrozenAmount float64 `json:"frozen_amount" example:"10.00"`
	Currency     string  `json:"currency" example:"USD"`
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
	customerID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// Get balance from transaction service
	detail, err := h.transactionService.GetBalanceDetail(c.Request.Context(), customerID)
	if err != nil {
		common.RespondErrorWithMsg(c, common.CodeInternalError, "Failed to get balance: "+err.Error())
		return
	}

	// Prepare response
	response := BalanceResponse{
		Balance:      detail.Balance,
		FrozenAmount: detail.FrozenAmount,
		Currency:     "USD", // Hardcoded for now, could be configurable
	}

	common.RespondSuccess(c, response)
}
