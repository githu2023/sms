package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"
	"time"

	"gorm.io/gorm"
)

// LogRepository defines the interface for log data operations
type LogRepository interface {
	Create(ctx context.Context, log *domain.APILog) error
	FindByID(ctx context.Context, id int64) (*domain.APILog, error)
	FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.APILog, int64, error)
	FindByPath(ctx context.Context, path string, limit, offset int) ([]*domain.APILog, int64, error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]*domain.APILog, int64, error)
	FindByCustomerIDAndPath(ctx context.Context, customerID int64, path string, limit, offset int) ([]*domain.APILog, int64, error)
	DeleteOldLogs(ctx context.Context, beforeDate time.Time) (int64, error)
}

// logRepository is the implementation of LogRepository
type logRepository struct {
	db *gorm.DB
}

// NewLogRepository creates a new instance of LogRepository
func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

// Create adds a new log entry
func (r *logRepository) Create(ctx context.Context, log *domain.APILog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// FindByID finds a log entry by ID
func (r *logRepository) FindByID(ctx context.Context, id int64) (*domain.APILog, error) {
	var log domain.APILog
	err := r.db.WithContext(ctx).First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// FindByCustomerID finds log entries by customer ID with pagination
func (r *logRepository) FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.APILog, int64, error) {
	var logs []*domain.APILog
	var total int64

	// Count total records
	err := r.db.WithContext(ctx).Model(&domain.APILog{}).Where("customer_id = ?", customerID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = r.db.WithContext(ctx).Where("customer_id = ?", customerID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// FindByPath finds log entries by request path with pagination
func (r *logRepository) FindByPath(ctx context.Context, path string, limit, offset int) ([]*domain.APILog, int64, error) {
	var logs []*domain.APILog
	var total int64

	// Count total records
	err := r.db.WithContext(ctx).Model(&domain.APILog{}).Where("request_path = ?", path).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = r.db.WithContext(ctx).Where("request_path = ?", path).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// FindByDateRange finds log entries within a date range
func (r *logRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]*domain.APILog, int64, error) {
	var logs []*domain.APILog
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.APILog{}).Where("created_at BETWEEN ? AND ?", startDate, endDate)

	// Count total records
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = query.Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// FindByCustomerIDAndPath finds log entries by customer ID and request path
func (r *logRepository) FindByCustomerIDAndPath(ctx context.Context, customerID int64, path string, limit, offset int) ([]*domain.APILog, int64, error) {
	var logs []*domain.APILog
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.APILog{}).Where("customer_id = ? AND request_path = ?", customerID, path)

	// Count total records
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = query.Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// DeleteOldLogs deletes log entries older than the specified date
func (r *logRepository) DeleteOldLogs(ctx context.Context, beforeDate time.Time) (int64, error) {
	result := r.db.WithContext(ctx).Where("created_at < ?", beforeDate).Delete(&domain.APILog{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}