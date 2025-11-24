package dto

// AuthDTO 认证相关的数据传输对象

// GetAPITokenRequest API Token 请求
type GetAPITokenRequest struct {
	MerchantNo string `json:"merchant_no" binding:"required"` // 商户号
	Secret     string `json:"secret" binding:"required"`      // API密钥
}

// GetAPITokenResponse API Token 响应
type GetAPITokenResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}
