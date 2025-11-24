package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Customer SMS客户信息表
type Customer struct {
	ID             int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	MerchantName   *string        `gorm:"type:varchar(255);comment:商户名称" json:"merchant_name"`
	MerchantNo     *string        `gorm:"type:varchar(50);uniqueIndex;comment:商户号" json:"merchant_no"`
	Username       *string        `gorm:"type:varchar(255);comment:客户端登录用户名" json:"username"`
	Email          *string        `gorm:"type:varchar(255);comment:客户端登录邮箱" json:"email"`
	PasswordHash   *string        `gorm:"type:varchar(255);comment:客户端登录用的密码哈希" json:"password_hash"`
	APISecretKey   string         `gorm:"type:varchar(255);not null;comment:用于生成API Token的唯一密钥" json:"api_secret_key"`
	Balance        float64        `gorm:"type:decimal(10,2);default:0.00;comment:客户余额" json:"balance"`
	ParentID       *int64         `gorm:"comment:上级商户ID" json:"parent_id"`
	FrozenAmount   float64        `gorm:"type:decimal(10,2);default:0.00;comment:冻结金额" json:"frozen_amount"`
	Status         *bool          `gorm:"comment:客户状态 (true:正常, false:冻结)" json:"status"`
	RegistrationIP *string        `gorm:"type:varchar(45);comment:注册时的IP地址" json:"registration_ip"`
	LastLoginIP    *string        `gorm:"type:varchar(45);comment:最后一次登录的IP地址" json:"last_login_ip"`
	LastLoginAt    *time.Time     `gorm:"comment:最后一次登录的时间" json:"last_login_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Remark         *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`
}

// CustomerRepository defines the interface for customer data operations.
type CustomerRepository interface {
	Create(ctx context.Context, customer *Customer) error
	FindByUsername(ctx context.Context, username string) (*Customer, error)
	FindByID(ctx context.Context, id int64) (*Customer, error)
	FindByAPISecretKey(ctx context.Context, apiSecretKey string) (*Customer, error)                       // New
	FindByMerchantNoAndAPISecret(ctx context.Context, merchantNo, apiSecretKey string) (*Customer, error) // New
	Update(ctx context.Context, customer *Customer) error
}

// TableName 指定表名
func (Customer) TableName() string {
	return "sms_customers"
}
