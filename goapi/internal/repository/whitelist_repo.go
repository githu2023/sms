package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

// WhitelistRepository defines the interface for whitelist data operations
type WhitelistRepository interface {
	Create(ctx context.Context, whitelist *domain.IPWhitelist) error
	FindByID(ctx context.Context, id int64) (*domain.IPWhitelist, error)
	FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.IPWhitelist, int64, error)
	FindByCustomerIDAndIP(ctx context.Context, customerID int64, ipAddress string) (*domain.IPWhitelist, error)
	Update(ctx context.Context, whitelist *domain.IPWhitelist) error
	Delete(ctx context.Context, id int64) error
	DeleteByCustomerIDAndIP(ctx context.Context, customerID int64, ipAddress string) error
}

// whitelistRepository is the implementation of WhitelistRepository
type whitelistRepository struct {
	db *gorm.DB
}

// NewWhitelistRepository creates a new instance of WhitelistRepository
func NewWhitelistRepository(db *gorm.DB) WhitelistRepository {
	return &whitelistRepository{db: db}
}

// Create adds a new whitelist entry
func (r *whitelistRepository) Create(ctx context.Context, whitelist *domain.IPWhitelist) error {
	return r.db.WithContext(ctx).Create(whitelist).Error
}

// FindByID finds a whitelist entry by ID
func (r *whitelistRepository) FindByID(ctx context.Context, id int64) (*domain.IPWhitelist, error) {
	var whitelist domain.IPWhitelist
	err := r.db.WithContext(ctx).First(&whitelist, id).Error
	if err != nil {
		return nil, err
	}
	return &whitelist, nil
}

// FindByCustomerID finds whitelist entries by customer ID with pagination
func (r *whitelistRepository) FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.IPWhitelist, int64, error) {
	var whitelists []*domain.IPWhitelist
	var total int64

	// Count total records
	err := r.db.WithContext(ctx).Model(&domain.IPWhitelist{}).Where("customer_id = ?", customerID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = r.db.WithContext(ctx).Where("customer_id = ?", customerID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&whitelists).Error
	if err != nil {
		return nil, 0, err
	}

	return whitelists, total, nil
}

// FindByCustomerIDAndIP finds a specific whitelist entry by customer ID and IP
func (r *whitelistRepository) FindByCustomerIDAndIP(ctx context.Context, customerID int64, ipAddress string) (*domain.IPWhitelist, error) {
	var whitelist domain.IPWhitelist
	err := r.db.WithContext(ctx).Where("customer_id = ? AND ip_address = ?", customerID, ipAddress).First(&whitelist).Error
	if err != nil {
		return nil, err
	}
	return &whitelist, nil
}

// Update updates a whitelist entry
func (r *whitelistRepository) Update(ctx context.Context, whitelist *domain.IPWhitelist) error {
	return r.db.WithContext(ctx).Save(whitelist).Error
}

// Delete deletes a whitelist entry by ID
func (r *whitelistRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&domain.IPWhitelist{}, id).Error
}

// DeleteByCustomerIDAndIP deletes a whitelist entry by customer ID and IP
func (r *whitelistRepository) DeleteByCustomerIDAndIP(ctx context.Context, customerID int64, ipAddress string) error {
	return r.db.WithContext(ctx).Where("customer_id = ? AND ip_address = ?", customerID, ipAddress).Delete(&domain.IPWhitelist{}).Error
}
