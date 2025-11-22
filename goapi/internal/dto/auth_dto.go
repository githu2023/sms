package dto

// AuthDTO 认证相关的数据传输对象

// GetAPITokenRequest API Token 请求
type GetAPITokenRequest struct {
	Secret string `json:"secret" binding:"required"`
}

// GetAPITokenResponse API Token 响应
type GetAPITokenResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}
