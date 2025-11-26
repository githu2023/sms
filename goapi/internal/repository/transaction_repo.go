package repository

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/domain"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	// ErrInsufficientBalance 表示可用余额不足，无法完成预冻结
	ErrInsufficientBalance = errors.New("insufficient balance")
	// ErrInsufficientFrozenFunds 表示冻结余额不足，无法完成解冻或结算
	ErrInsufficientFrozenFunds = errors.New("insufficient frozen funds")
)

// BalanceChangeSnapshot 记录一次原子更新前后的余额/冻结金额
type BalanceChangeSnapshot struct {
	BalanceBefore float64
	BalanceAfter  float64
	FrozenBefore  float64
	FrozenAfter   float64
}

// TransactionRepository defines the interface for transaction data operations
type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, transaction *domain.Transaction) error
	FindByID(ctx context.Context, id int64) (*domain.Transaction, error)
	FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error)
	FindByCustomerIDAndType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error)
	FindByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error)
	GetBalance(ctx context.Context, customerID int64) (float64, error)
	GetBalanceDetail(ctx context.Context, customerID int64) (float64, float64, error)
	GetBalanceForUpdate(ctx context.Context, tx *gorm.DB, customerID int64) (float64, error)
	ReserveBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*BalanceChangeSnapshot, error)
	CommitReservedBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*BalanceChangeSnapshot, error)
	ReleaseReservedBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*BalanceChangeSnapshot, error)
	Update(ctx context.Context, transaction *domain.Transaction) error
	BeginTx(ctx context.Context) (*gorm.DB, error)
}

// transactionRepository is the implementation of TransactionRepository
type transactionRepository struct {
	db           *gorm.DB
	customerRepo domain.CustomerRepository
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB, customerRepo domain.CustomerRepository) TransactionRepository {
	return &transactionRepository{
		db:           db,
		customerRepo: customerRepo,
	}
}

// Create adds a new transaction
func (r *transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	return r.CreateWithTx(ctx, nil, transaction)
}

// CreateWithTx adds a new transaction within provided transaction (if any)
func (r *transactionRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, transaction *domain.Transaction) error {
	db := r.selectDB(tx)
	return db.WithContext(ctx).Create(transaction).Error
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

// GetBalance gets the current balance from customer repository
func (r *transactionRepository) GetBalance(ctx context.Context, customerID int64) (float64, error) {
	customer, err := r.customerRepo.FindByID(ctx, customerID)
	if err != nil {
		return 0, err
	}

	return customer.Balance, nil
}

// GetBalanceForUpdate locks the customer row and returns current balance
func (r *transactionRepository) GetBalanceForUpdate(ctx context.Context, tx *gorm.DB, customerID int64) (float64, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	var customer domain.Customer
	if err := db.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		return 0, err
	}
	return customer.Balance, nil
}

// GetBalanceDetail returns both balance and frozen_amount
func (r *transactionRepository) GetBalanceDetail(ctx context.Context, customerID int64) (float64, float64, error) {
	customer, err := r.customerRepo.FindByID(ctx, customerID)
	if err != nil {
		return 0, 0, err
	}
	return customer.Balance, customer.FrozenAmount, nil
}

func (r *transactionRepository) selectDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *transactionRepository) lockCustomer(ctx context.Context, db *gorm.DB, customerID int64) (*domain.Customer, error) {
	var customer domain.Customer
	if err := db.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *transactionRepository) ReserveBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*BalanceChangeSnapshot, error) {
	db := r.selectDB(tx)
	customer, err := r.lockCustomer(ctx, db, customerID)
	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, errors.New("reserve amount must be positive")
	}
	if customer.Balance < amount {
		return nil, ErrInsufficientBalance
	}

	snapshot := &BalanceChangeSnapshot{
		BalanceBefore: customer.Balance,
		FrozenBefore:  customer.FrozenAmount,
		BalanceAfter:  customer.Balance - amount,
		FrozenAfter:   customer.FrozenAmount + amount,
	}

	if err := db.WithContext(ctx).Model(&domain.Customer{}).
		Where("id = ?", customerID).
		Updates(map[string]interface{}{
			"balance":       snapshot.BalanceAfter,
			"frozen_amount": snapshot.FrozenAfter,
			"updated_at":    time.Now(),
		}).Error; err != nil {
		return nil, err
	}

	return snapshot, nil
}

func (r *transactionRepository) CommitReservedBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*BalanceChangeSnapshot, error) {
	db := r.selectDB(tx)
	customer, err := r.lockCustomer(ctx, db, customerID)
	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, errors.New("commit amount must be positive")
	}
	if customer.FrozenAmount < amount {
		return nil, ErrInsufficientFrozenFunds
	}

	snapshot := &BalanceChangeSnapshot{
		BalanceBefore: customer.Balance,
		BalanceAfter:  customer.Balance,
		FrozenBefore:  customer.FrozenAmount,
		FrozenAfter:   customer.FrozenAmount - amount,
	}

	if err := db.WithContext(ctx).Model(&domain.Customer{}).
		Where("id = ?", customerID).
		Updates(map[string]interface{}{
			"frozen_amount": snapshot.FrozenAfter,
			"updated_at":    time.Now(),
		}).Error; err != nil {
		return nil, err
	}

	return snapshot, nil
}

func (r *transactionRepository) ReleaseReservedBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*BalanceChangeSnapshot, error) {
	db := r.selectDB(tx)
	customer, err := r.lockCustomer(ctx, db, customerID)
	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, errors.New("release amount must be positive")
	}
	if customer.FrozenAmount < amount {
		return nil, ErrInsufficientFrozenFunds
	}

	snapshot := &BalanceChangeSnapshot{
		BalanceBefore: customer.Balance,
		FrozenBefore:  customer.FrozenAmount,
		BalanceAfter:  customer.Balance + amount,
		FrozenAfter:   customer.FrozenAmount - amount,
	}

	if err := db.WithContext(ctx).Model(&domain.Customer{}).
		Where("id = ?", customerID).
		Updates(map[string]interface{}{
			"balance":       snapshot.BalanceAfter,
			"frozen_amount": snapshot.FrozenAfter,
			"updated_at":    time.Now(),
		}).Error; err != nil {
		return nil, err
	}

	return snapshot, nil
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
