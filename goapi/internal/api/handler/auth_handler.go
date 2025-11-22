package handler

import (
	"errors"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService service.UserService
	jwtConfig   config.JWTConfig
}

func NewAuthHandler(userService service.UserService, jwtConfig config.JWTConfig) *AuthHandler {
	return &AuthHandler{userService: userService, jwtConfig: jwtConfig}
}

// GetAPIToken godoc
// @Summary Get API Token
// @Description Exchanges a secret key for a short-lived API token.
// @Tags api-auth
// @Accept  json
// @Produce  json
// @Param   tokenRequest body GetAPITokenRequest true "API Token Request Info"
// @Success 200 {object} SuccessResponse{data=GetAPITokenResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/get_token [post]
func (h *AuthHandler) GetAPIToken(c *gin.Context) {
	var req dto.GetAPITokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	token, err := h.userService.GenerateAPIToken(c.Request.Context(), req.Secret)
	if err != nil {
		if errors.Is(err, service.ErrInvalidAPISecret) {
			common.RespondError(c, common.CodeInvalidSecret)
			return
		}
		common.RespondError(c, common.CodeInternalError)
		return
	}

	response := dto.GetAPITokenResponse{
		Token:     token,
		ExpiresIn: h.jwtConfig.APITokenExpiry, // 使用配置的过期时间
	}
	common.RespondSuccess(c, response)
}
