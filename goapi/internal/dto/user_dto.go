package dto

// UserDTO 用户相关的数据传输对象

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// RegisterResponse 注册响应
type RegisterResponse struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

// ProfileResponse 用户资料响应
type ProfileResponse struct {
	UserID         int64   `json:"user_id"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	Balance        float64 `json:"balance"`
	FrozenAmount   float64 `json:"frozen_amount"`
	APISecretKey   string  `json:"api_secret_key"`
	RegistrationIP string  `json:"registration_ip"`
	LastLoginAt    string  `json:"last_login_at"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
