package dto

import "time"

// GetAssignmentsRequest is the request DTO for fetching assignment history.
type GetAssignmentsRequest struct {
	Page         int    `form:"page" binding:"omitempty,min=1" example:"1"`
	Limit        int    `form:"limit" binding:"omitempty,min=1,max=100" example:"20"`
	Status       int    `form:"status" binding:"omitempty,min=1,max=4" example:"1"` // 1:pending, 2:completed, 3:expired, 4:failed
	BusinessType string `form:"business_type" example:"qq"`
	StartDate    string `form:"start_date" binding:"omitempty,datetime=2006-01-02" example:"2024-01-01"` // YYYY-MM-DD
	EndDate      string `form:"end_date" binding:"omitempty,datetime=2006-01-02" example:"2024-01-31"`   // YYYY-MM-DD
}

// AssignmentHistoryItem represents a single assignment record in the response.
type AssignmentHistoryItem struct {
	ID               int64      `json:"id" example:"1"`
	PhoneNumber      string     `json:"phone_number" example:"+15551234567"`
	BusinessType     string     `json:"business_type" example:"qq"`
	CardType         string     `json:"card_type" example:"virtual"`
	VerificationCode string     `json:"verification_code,omitempty" example:"123456"`
	Cost             float64    `json:"cost" example:"0.10"`
	Status           int        `json:"status" example:"2"` // 1:pending, 2:completed, 3:expired, 4:failed
	ExpiresAt        *time.Time `json:"expires_at,omitempty" example:"2024-01-01T12:30:00Z"`
	CreatedAt        time.Time  `json:"created_at" example:"2024-01-01T12:00:00Z"`
	ProviderName     string     `json:"provider_name" example:"MockSMS"` // Name of the provider
}

// AssignmentHistoryResponse is the response DTO for fetching assignment history.
type AssignmentHistoryResponse struct {
	Items      []AssignmentHistoryItem `json:"items"`
	Pagination struct {
		Total int64 `json:"total"`
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
	} `json:"pagination"`
}

// RecentAssignmentsResponse 列出最近的拉号记录
type RecentAssignmentsResponse struct {
	Items []AssignmentHistoryItem `json:"items"`
}

// CostStatisticsResponse is the response DTO for cost statistics.
type CostStatisticsResponse struct {
	TotalCost  float64 `json:"total_cost" example:"12.50"`
	TotalCount int64   `json:"total_count" example:"125"`
}
