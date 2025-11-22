package domain

import (
	"context"
	"time"
)

// Customer 客户信息表
type Customer struct {
	ID             int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username       string     `gorm:"unique" json:"username"`                     // 客户端登录用户名
	Email          string     `gorm:"unique" json:"email"`                        // 客户端登录邮箱
	PasswordHash   string     `json:"-"`                                          // 客户端登录用的密码哈希
	APISecretKey   string     `gorm:"unique;not null" json:"api_secret_key"`      // 用于生成API Token的唯一密钥
	Balance        float64    `gorm:"type:decimal(10,4);not null" json:"balance"` // 客户余额
	Status         int        `gorm:"not null;default:1" json:"status"`           // 客户状态 (1:正常, 2:冻结, 0:已删除)
	RegistrationIP string     `gorm:"size:45" json:"registration_ip"`             // 注册时的IP地址
	LastLoginIP    string     `gorm:"size:45" json:"last_login_ip"`               // 最后一次登录的IP地址
	LastLoginAt    *time.Time `json:"last_login_at"`                              // 最后一次登录的时间
	CreatedAt      time.Time  `json:"created_at"`                                 // 创建时间
	UpdatedAt      time.Time  `json:"updated_at"`                                 // 更新时间
}

// CustomerRepository defines the interface for customer data operations.
type CustomerRepository interface {
	Create(ctx context.Context, customer *Customer) error
	FindByUsername(ctx context.Context, username string) (*Customer, error)
	FindByID(ctx context.Context, id int64) (*Customer, error)
	FindByAPISecretKey(ctx context.Context, apiSecretKey string) (*Customer, error) // New
	Update(ctx context.Context, customer *Customer) error
}

// TableName 指定表名
func (Customer) TableName() string {
	return "sms_customers"
}
