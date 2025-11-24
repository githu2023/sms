package handler

import (
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"sms-platform/goapi/internal/utils"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	businessService service.BusinessService
}

func NewBusinessHandler(businessService service.BusinessService) *BusinessHandler {
	return &BusinessHandler{businessService: businessService}
}

// GetBusinessTypes godoc
// @Summary Get business types for current customer
// @Description Retrieves a list of business types assigned to the current customer.
// @Tags business
// @Produce  json
// @Success 200 {object} SuccessResponse{data=[]CustomerBusinessTypeResponse}
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/business_types [get]
// @Router /client/v1/business_types [get]
func (h *BusinessHandler) GetBusinessTypes(c *gin.Context) {
	// Get customer_id from context (set by authentication middleware)
	customerID, ok := utils.RequireCustomerID(c)
	if !ok {
		return
	}

	// Get business types assigned to this customer
	customerBusinessConfigs, err := h.businessService.GetBusinessTypesForCustomer(c.Request.Context(), customerID)
	if err != nil {
		common.RespondError(c, common.CodeInternalError)
		return
	}

	// Convert to response format
	respData := make([]dto.CustomerBusinessTypeResponse, len(customerBusinessConfigs))
	for i, config := range customerBusinessConfigs {
		respData[i] = dto.CustomerBusinessTypeResponse{
			ID:           config.ID,
			BusinessCode: config.BusinessCode,
			BusinessName: config.BusinessName,
			Weight:       config.Weight,
		}
	}

	common.RespondSuccess(c, respData)
}
