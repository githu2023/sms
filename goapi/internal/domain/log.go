package domain

import "time"

// APILog API请求日志表
type APILog struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   int64     `gorm:"index" json:"customer_id"`           // 客户ID
	RequestIP    string    `gorm:"size:45;not null" json:"request_ip"` // 请求来源IP
	RequestPath  string    `gorm:"not null" json:"request_path"`       // 请求的API路径
	RequestBody  string    `gorm:"type:text" json:"request_body"`      // 请求体内容
	ResponseCode int       `json:"response_code"`                      // HTTP响应状态码
	DurationMS   int       `json:"duration_ms"`                        // 请求处理耗时(毫秒)
	CreatedAt    time.Time `gorm:"index" json:"created_at"`            // 创建时间
}

// TableName 指定表名
func (APILog) TableName() string {
	return "sms_api_logs"
}
