package dto

import "time"

// GetPhoneRequest 获取手机号请求
type GetPhoneRequest struct {
	BusinessType string `json:"business_type" binding:"required" validate:"oneof=verification social_media ecommerce other" example:"verification"`
	CardType     string `json:"card_type" binding:"required" validate:"oneof=any virtual physical temporary" example:"any"`
	Count        int    `json:"count" validate:"min=1,max=10" example:"1"` // 批量获取数量，默认1，最大10
}

// PhoneInfo 单个手机号信息
type PhoneInfo struct {
	PhoneNumber string    `json:"phone_number" example:"+15551234567"`
	CountryCode string    `json:"country_code" example:"US"`
	Cost        float64   `json:"cost" example:"0.10"`
	ValidUntil  time.Time `json:"valid_until" example:"2024-01-01T12:30:00Z"`
	ProviderID  string    `json:"provider_id" example:"provider-1"`
}

// GetPhoneResponse 获取手机号响应
type GetPhoneResponse struct {
	Phones           []PhoneInfo `json:"phones"`            // 支持批量返回
	TotalCost        float64     `json:"total_cost"`        // 总费用
	RemainingBalance float64     `json:"remaining_balance"` // 剩余余额
	SuccessCount     int         `json:"success_count"`     // 成功获取数量
	FailedCount      int         `json:"failed_count"`      // 失败数量
}

// GetCodeRequest 获取验证码请求
type GetCodeRequest struct {
	PhoneNumbers []string `json:"phone_numbers" binding:"required,min=1,max=10" example:"[\"+15551234567\", \"+15551234568\"]"`
}

// CodeInfo 单个验证码信息
type CodeInfo struct {
	PhoneNumber string    `json:"phone_number" example:"+15551234567"`
	Code        string    `json:"code" example:"123456"`
	Message     string    `json:"message" example:"Your verification code is 123456"`
	ReceivedAt  time.Time `json:"received_at" example:"2024-01-01T12:30:00Z"`
	ProviderID  string    `json:"provider_id" example:"provider-1"`
	Status      string    `json:"status" example:"success"` // success, pending, failed
}

// GetCodeResponse 获取验证码响应
type GetCodeResponse struct {
	Codes        []CodeInfo `json:"codes"`         // 支持批量返回
	SuccessCount int        `json:"success_count"` // 成功获取数量
	PendingCount int        `json:"pending_count"` // 等待中数量（验证码还未获取到）
	FailedCount  int        `json:"failed_count"`  // 失败数量
}
