package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"
	"gorm.io/gorm"
	"time"
)

// PhoneAssignmentRepository defines the interface for phone assignment data operations.
type PhoneAssignmentRepository interface {
	Create(ctx context.Context, tx *gorm.DB, assignment *domain.PhoneAssignment) error
	FindByPhone(ctx context.Context, tx *gorm.DB, phone string) (*domain.PhoneAssignment, error)
	FindActiveByCustomerIDAndPhone(ctx context.Context, tx *gorm.DB, customerID int64, phone string) (*domain.PhoneAssignment, error)
	FindRecentByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.PhoneAssignment, int64, error)
	Update(ctx context.Context, tx *gorm.DB, assignment *domain.PhoneAssignment) error
	UpdateVerificationCode(ctx context.Context, tx *gorm.DB, id int64, code string, receivedAt time.Time) error
	UpdateStatus(ctx context.Context, tx *gorm.DB, id int64, status int) error
	FindExpiredAssignments(ctx context.Context, limit int) ([]*domain.PhoneAssignment, error)
	CountByCustomerID(ctx context.Context, customerID int64) (int64, error)
}

// phoneAssignmentRepository is the implementation of PhoneAssignmentRepository
type phoneAssignmentRepository struct {
	db *gorm.DB
}

// NewPhoneAssignmentRepository creates a new instance of PhoneAssignmentRepository
func NewPhoneAssignmentRepository(db *gorm.DB) PhoneAssignmentRepository {
	return &phoneAssignmentRepository{db: db}
}

// Create adds a new phone assignment entry, optionally within a provided GORM transaction
func (r *phoneAssignmentRepository) Create(ctx context.Context, tx *gorm.DB, assignment *domain.PhoneAssignment) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.WithContext(ctx).Create(assignment).Error
}

// FindByPhone finds a phone assignment entry by phone number, optionally within a provided GORM transaction
func (r *phoneAssignmentRepository) FindByPhone(ctx context.Context, tx *gorm.DB, phone string) (*domain.PhoneAssignment, error) {
	db := r.db
	if tx != nil {
		db = tx
	}
	var assignment domain.PhoneAssignment
	err := db.WithContext(ctx).Where("phone_number = ?", phone).First(&assignment).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

// FindActiveByCustomerIDAndPhone finds an active phone assignment entry by customer ID and phone number
func (r *phoneAssignmentRepository) FindActiveByCustomerIDAndPhone(ctx context.Context, tx *gorm.DB, customerID int64, phone string) (*domain.PhoneAssignment, error) {
	db := r.db
	if tx != nil {
		db = tx
	}
	var assignment domain.PhoneAssignment
	// Assuming status 1 means active/awaiting code
	err := db.WithContext(ctx).Where("customer_id = ? AND phone_number = ? AND status = ?", customerID, phone, 1).First(&assignment).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

// FindRecentByCustomerID finds recent phone assignment entries by customer ID with pagination
func (r *phoneAssignmentRepository) FindRecentByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.PhoneAssignment, int64, error) {
	var assignments []*domain.PhoneAssignment
	var total int64

	// Count total records
	err := r.db.WithContext(ctx).Model(&domain.PhoneAssignment{}).Where("customer_id = ?", customerID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = r.db.WithContext(ctx).Where("customer_id = ?", customerID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&assignments).Error
	if err != nil {
		return nil, 0, err
	}

	return assignments, total, nil
}

// Update updates a phone assignment entry, optionally within a provided GORM transaction
func (r *phoneAssignmentRepository) Update(ctx context.Context, tx *gorm.DB, assignment *domain.PhoneAssignment) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.WithContext(ctx).Save(assignment).Error
}

// UpdateVerificationCode updates the verification code for a phone assignment
func (r *phoneAssignmentRepository) UpdateVerificationCode(ctx context.Context, tx *gorm.DB, id int64, code string, receivedAt time.Time) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.WithContext(ctx).Model(&domain.PhoneAssignment{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verification_code": code,
			"status":            2, // Status 2: completed
			"updated_at":        receivedAt,
		}).Error
}

// UpdateStatus updates the status of a phone assignment
func (r *phoneAssignmentRepository) UpdateStatus(ctx context.Context, tx *gorm.DB, id int64, status int) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.WithContext(ctx).Model(&domain.PhoneAssignment{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// FindExpiredAssignments finds assignments that have expired and are still awaiting code
func (r *phoneAssignmentRepository) FindExpiredAssignments(ctx context.Context, limit int) ([]*domain.PhoneAssignment, error) {
	var assignments []*domain.PhoneAssignment
	err := r.db.WithContext(ctx).Where("status = ? AND expires_at < ?", 1, time.Now()).Limit(limit).Find(&assignments).Error
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

// CountByCustomerID counts assignments for a specific customer
func (r *phoneAssignmentRepository) CountByCustomerID(ctx context.Context, customerID int64) (int64, error) {
	var total int64
	err := r.db.WithContext(ctx).Model(&domain.PhoneAssignment{}).Where("customer_id = ?", customerID).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
