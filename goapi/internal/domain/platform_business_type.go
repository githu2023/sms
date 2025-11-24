package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// PlatformBusinessType 平台业务类型表
type PlatformBusinessType struct {
	ID          int64          `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `gorm:"type:varchar(255);not null;comment:平台业务名称" json:"name"`
	Code        string         `gorm:"type:varchar(50);not null;comment:平台业务编码" json:"code"`
	Description *string        `gorm:"type:varchar(500);comment:业务描述" json:"description"`
	Status      *bool          `gorm:"comment:启用状态" json:"status"`
	Remark      *string        `gorm:"type:varchar(500);comment:备注" json:"remark"`
}

func (PlatformBusinessType) TableName() string {
	return "sms_platform_business_types"
}

// PlatformBusinessTypeRepository 平台业务类型Repository接口
type PlatformBusinessTypeRepository interface {
	FindByID(ctx context.Context, id int64) (*PlatformBusinessType, error)
	FindByCode(ctx context.Context, code string) (*PlatformBusinessType, error)
	FindAll(ctx context.Context) ([]*PlatformBusinessType, error)
}
