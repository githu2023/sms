package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

// CustomerBusinessConfigRepository defines the interface for customer business config data operations
type CustomerBusinessConfigRepository interface {
	Create(ctx context.Context, config *domain.CustomerBusinessConfig) error
	FindByCustomerID(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error)
	FindByCustomerIDAndBusinessCode(ctx context.Context, customerID int64, businessCode string) (*domain.CustomerBusinessConfig, error)
	Update(ctx context.Context, config *domain.CustomerBusinessConfig) error
	Delete(ctx context.Context, id int64) error
	FindByCustomerIDAndEnabled(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error)
}

type customerBusinessConfigRepository struct {
	db *gorm.DB
}

// NewCustomerBusinessConfigRepository creates a new customer business config repository
func NewCustomerBusinessConfigRepository(db *gorm.DB) CustomerBusinessConfigRepository {
	return &customerBusinessConfigRepository{db: db}
}

func (r *customerBusinessConfigRepository) Create(ctx context.Context, config *domain.CustomerBusinessConfig) error {
	return r.db.WithContext(ctx).Table("sms_customer_business_config").Create(config).Error
}

func (r *customerBusinessConfigRepository) FindByCustomerID(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error) {
	var configs []*domain.CustomerBusinessConfig
	err := r.db.WithContext(ctx).Table("sms_customer_business_config").Where("customer_id = ?", customerID).Find(&configs).Error
	return configs, err
}

func (r *customerBusinessConfigRepository) FindByCustomerIDAndBusinessCode(ctx context.Context, customerID int64, businessCode string) (*domain.CustomerBusinessConfig, error) {
	var config domain.CustomerBusinessConfig
	err := r.db.WithContext(ctx).Table("sms_customer_business_config").Where("customer_id = ? AND business_code = ?", customerID, businessCode).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *customerBusinessConfigRepository) Update(ctx context.Context, config *domain.CustomerBusinessConfig) error {
	return r.db.WithContext(ctx).Table("sms_customer_business_config").Save(config).Error
}

func (r *customerBusinessConfigRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Table("sms_customer_business_config").Delete(&domain.CustomerBusinessConfig{}, id).Error
}

func (r *customerBusinessConfigRepository) FindByCustomerIDAndEnabled(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error) {
	var configs []*domain.CustomerBusinessConfig
	err := r.db.WithContext(ctx).Table("sms_customer_business_config").Where("customer_id = ? AND status = ?", customerID, true).Find(&configs).Error
	return configs, err
}