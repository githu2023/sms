package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new customer repository.
func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

func (r *customerRepository) FindByUsername(ctx context.Context, username string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) FindByID(ctx context.Context, id int64) (*domain.Customer, error) {
	var customer domain.Customer
	if err := r.db.WithContext(ctx).First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) FindByAPISecretKey(ctx context.Context, apiSecretKey string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := r.db.WithContext(ctx).Where("api_secret_key = ?", apiSecretKey).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) FindByMerchantNoAndAPISecret(ctx context.Context, merchantNo, apiSecretKey string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := r.db.WithContext(ctx).Where("merchant_no = ? AND api_secret_key = ?", merchantNo, apiSecretKey).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}
