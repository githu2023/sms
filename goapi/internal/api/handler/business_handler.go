package handler

import (
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	businessService service.BusinessService
}

func NewBusinessHandler(businessService service.BusinessService) *BusinessHandler {
	return &BusinessHandler{businessService: businessService}
}

// GetBusinessTypes godoc
// @Summary Get all business types
// @Description Retrieves a list of all available business types.
// @Tags business
// @Produce  json
// @Success 200 {object} SuccessResponse{data=[]BusinessTypeResponse}
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/business_types [get]
// @Router /client/v1/business_types [get]
func (h *BusinessHandler) GetBusinessTypes(c *gin.Context) {
	businessTypes, err := h.businessService.ListBusinessTypes(c.Request.Context())
	if err != nil {
		common.RespondError(c, common.CodeInternalError)
		return
	}

	respData := make([]dto.BusinessTypeResponse, len(businessTypes))
	for i, bt := range businessTypes {
		respData[i] = dto.BusinessTypeResponse{
			ID:   bt.ID,
			Name: bt.Name,
			Code: bt.Code,
		}
	}

	common.RespondSuccess(c, respData)
}
