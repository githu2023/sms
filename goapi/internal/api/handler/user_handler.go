package handler

import (
	"errors"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
	jwtConfig   config.JWTConfig
}

func NewUserHandler(userService service.UserService, jwtConfig config.JWTConfig) *UserHandler {
	return &UserHandler{userService: userService, jwtConfig: jwtConfig}
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account for the client application.
// @Tags client-auth
// @Accept  json
// @Produce  json
// @Param   register body registerRequest true "Registration Info"
// @Success 201 {object} SuccessResponse{data=registerResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /client/v1/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	// TODO: Get Registration IP from middleware
	customer, err := h.userService.Register(c.Request.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		// Check if user already exists
		if err.Error() == "user already exists" {
			common.RespondError(c, common.CodeBadRequest)
			return
		}
		// Other errors
		common.RespondError(c, common.CodeInternalError)
		return
	}

	username := ""
	if customer.Username != nil {
		username = *customer.Username
	}

	response := dto.RegisterResponse{
		UserID:   customer.ID,
		Username: username,
	}

	c.JSON(201, common.SuccessResponse(response))
}

// Login godoc
// @Summary Login a user
// @Description Authenticates a user and returns a JWT.
// @Tags client-auth
// @Accept  json
// @Produce  json
// @Param   login body loginRequest true "Login Credentials"
// @Success 200 {object} SuccessResponse{data=loginResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /client/v1/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	token, err := h.userService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) || errors.Is(err, service.ErrInvalidCredentials) {
			common.RespondError(c, common.CodeInvalidCredentials)
			return
		}
		common.RespondError(c, common.CodeInternalError)
		return
	}

	response := dto.LoginResponse{
		Token:     token,
		ExpiresIn: h.jwtConfig.ClientTokenExpiry, // 使用配置的过期时间
	}
	common.RespondSuccess(c, response)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieves the profile of the logged-in user.
// @Tags client-account
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} SuccessResponse{data=ProfileResponse}
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /client/v1/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	// Assuming user ID is set in context by a JWT middleware
	userIDVal, exists := c.Get("customer_id")
	if !exists {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	// JWT middleware sets customer_id as int64 directly
	userID, ok := userIDVal.(int64)
	if !ok {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	customer, err := h.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			common.RespondError(c, common.CodeNotFound)
			return
		}
		common.RespondError(c, common.CodeInternalError)
		return
	}

	lastLoginAt := ""
	if customer.LastLoginAt != nil {
		lastLoginAt = customer.LastLoginAt.Format(time.RFC3339)
	}

	username := ""
	if customer.Username != nil {
		username = *customer.Username
	}

	email := ""
	if customer.Email != nil {
		email = *customer.Email
	}

	registrationIP := ""
	if customer.RegistrationIP != nil {
		registrationIP = *customer.RegistrationIP
	}

	response := dto.ProfileResponse{
		UserID:         customer.ID,
		Username:       username,
		Email:          email,
		Balance:        customer.Balance,
		FrozenAmount:   customer.FrozenAmount,
		APISecretKey:   customer.APISecretKey,
		RegistrationIP: registrationIP,
		LastLoginAt:    lastLoginAt,
	}
	common.RespondSuccess(c, response)
}

// UpdatePassword 修改密码
// POST /client/v1/change_password
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		common.RespondError(c, common.CodeUnauthorized)
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.RespondError(c, common.CodeBadRequest)
		return
	}

	uid := userID.(int64)
	err := h.userService.UpdatePassword(c.Request.Context(), uid, req.OldPassword, req.NewPassword)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			common.RespondError(c, common.CodeInvalidCredentials)
		case errors.Is(err, service.ErrUserNotFound):
			common.RespondError(c, common.CodeNotFound)
		default:
			common.RespondError(c, common.CodeInternalError)
		}
		return
	}

	common.RespondSuccess(c, gin.H{
		"message": "密码修改成功",
	})
}
