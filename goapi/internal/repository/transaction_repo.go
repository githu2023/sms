package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"
	"time"

	"gorm.io/gorm"
)

// TransactionRepository defines the interface for transaction data operations
type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction) error
	FindByID(ctx context.Context, id int64) (*domain.Transaction, error)
	FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error)
	FindByCustomerIDAndType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error)
	FindByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error)
	GetBalance(ctx context.Context, customerID int64) (float64, error)
	Update(ctx context.Context, transaction *domain.Transaction) error
	BeginTx(ctx context.Context) (*gorm.DB, error)
}

// transactionRepository is the implementation of TransactionRepository
type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// Create adds a new transaction
func (r *transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

// FindByID finds a transaction by ID
func (r *transactionRepository) FindByID(ctx context.Context, id int64) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.WithContext(ctx).First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// FindByCustomerID finds transactions by customer ID with pagination
func (r *transactionRepository) FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error) {
	var transactions []*domain.Transaction
	var total int64

	// Count total records
	err := r.db.WithContext(ctx).Model(&domain.Transaction{}).Where("customer_id = ?", customerID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = r.db.WithContext(ctx).Where("customer_id = ?", customerID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&transactions).Error
	if err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// FindByCustomerIDAndType finds transactions by customer ID and type
func (r *transactionRepository) FindByCustomerIDAndType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error) {
	var transactions []*domain.Transaction
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Transaction{}).Where("customer_id = ? AND type = ?", customerID, transactionType)

	// Count total records
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = query.Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&transactions).Error
	if err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// FindByDateRange finds transactions within a date range
func (r *transactionRepository) FindByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error) {
	var transactions []*domain.Transaction
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Transaction{}).Where("customer_id = ? AND created_at BETWEEN ? AND ?", customerID, startDate, endDate)

	// Count total records
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Find records with pagination
	err = query.Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&transactions).Error
	if err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// GetBalance calculates the current balance for a customer
func (r *transactionRepository) GetBalance(ctx context.Context, customerID int64) (float64, error) {
	var result struct {
		Balance float64
	}

	err := r.db.WithContext(ctx).Model(&domain.Transaction{}).
		Select("COALESCE(SUM(CASE WHEN type = '1' THEN amount WHEN type = '2' THEN -amount ELSE 0 END), 0) as balance").
		Where("customer_id = ?", customerID).
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return result.Balance, nil
}

// Update updates a transaction
func (r *transactionRepository) Update(ctx context.Context, transaction *domain.Transaction) error {
	return r.db.WithContext(ctx).Save(transaction).Error
}

// BeginTx starts a database transaction
func (r *transactionRepository) BeginTx(ctx context.Context) (*gorm.DB, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}
