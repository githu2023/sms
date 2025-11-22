package dto

import "time"

// GetPhoneRequest 获取手机号请求
type GetPhoneRequest struct {
	BusinessType string `json:"business_type" binding:"required" validate:"oneof=verification social_media ecommerce other" example:"verification"`
	CardType     string `json:"card_type" binding:"required" validate:"oneof=any virtual physical temporary" example:"any"`
}

// GetPhoneResponse 获取手机号响应
type GetPhoneResponse struct {
	PhoneNumber      string    `json:"phone_number" example:"+15551234567"`
	CountryCode      string    `json:"country_code" example:"US"`
	Cost             float64   `json:"cost" example:"0.10"`
	ValidUntil       time.Time `json:"valid_until" example:"2024-01-01T12:30:00Z"`
	ProviderID       string    `json:"provider_id" example:"provider-1"`
	RemainingBalance float64   `json:"remaining_balance" example:"9.90"`
}

// GetCodeRequest 获取验证码请求
type GetCodeRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"phone" example:"+15551234567"`
	Timeout     int    `json:"timeout" validate:"min=1,max=300" example:"60"` // 秒数，最大5分钟
}

// GetCodeResponse 获取验证码响应
type GetCodeResponse struct {
	Code       string    `json:"code" example:"123456"`
	Message    string    `json:"message" example:"Your verification code is 123456"`
	ReceivedAt time.Time `json:"received_at" example:"2024-01-01T12:30:00Z"`
	ProviderID string    `json:"provider_id" example:"provider-1"`
}
