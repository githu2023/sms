package service

import (
	"context"
	"errors"
	"fmt"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/repository"
	"time"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrTransactionFailed   = errors.New("transaction failed")
	ErrInsufficientFrozen  = errors.New("insufficient frozen funds")
)

// TransactionService defines the interface for transaction business logic
type TransactionService interface {
	// TopUp adds funds to customer account
	TopUp(ctx context.Context, customerID int64, amount float64, notes string) (*domain.Transaction, error)

	// Deduct subtracts funds from customer account
	Deduct(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// ReserveFunds moves funds from balance to frozen balance
	ReserveFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// CommitReservedFunds consumes previously frozen funds without touching balance
	CommitReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// ReleaseReservedFunds returns frozen funds back to available balance
	ReleaseReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// GetBalance returns current customer balance
	GetBalance(ctx context.Context, customerID int64) (float64, error)

	// GetBalanceDetail returns available balance and frozen amount
	GetBalanceDetail(ctx context.Context, customerID int64) (*BalanceDetail, error)

	// GetTransactionHistory returns paginated transaction history
	GetTransactionHistory(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error)

	// GetTransactionsByType returns transactions filtered by type
	GetTransactionsByType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error)

	// GetTransactionsByDateRange returns transactions within date range
	GetTransactionsByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error)
}

// transactionService implements TransactionService interface
type transactionService struct {
	transactionRepo repository.TransactionRepository
}

// BalanceDetail represents detailed balance info
type BalanceDetail struct {
	Balance      float64
	FrozenAmount float64
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

// TopUp adds funds to customer account with database transaction
func (s *transactionService) TopUp(ctx context.Context, customerID int64, amount float64, notes string) (*domain.Transaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Get current balance
	currentBalance, err := s.transactionRepo.GetBalance(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %w", err)
	}

	// Create transaction record
	amountFloat := float32(amount)
	balanceBeforeFloat := float32(currentBalance)
	balanceAfterFloat := float32(currentBalance + amount)
	transactionType := domain.TransactionTypeTopUp

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amountFloat,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		Type:          &transactionType, // 1 = topup
		Notes:         &notes,
		CreatedAt:     time.Now(),
	}

	// Create transaction
	err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return transaction, nil
}

// Deduct subtracts funds from customer account with balance check
func (s *transactionService) Deduct(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Get current balance
	currentBalance, err := s.transactionRepo.GetBalance(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %w", err)
	}

	// Check sufficient balance
	if currentBalance < amount {
		return nil, ErrInsufficientBalance
	}

	// Create transaction record
	amountFloat := -float32(amount) // Negative for deduction
	balanceBeforeFloat := float32(currentBalance)
	balanceAfterFloat := float32(currentBalance - amount)
	transactionType := domain.TransactionTypeDeduct

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amountFloat,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		Type:          &transactionType, // 2 = deduct
		ReferenceID:   &referenceID,
		Notes:         &notes,
		CreatedAt:     time.Now(),
	}

	// Create transaction
	err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return transaction, nil
}

func (s *transactionService) ReserveFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin reserve transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	snapshot, err := s.transactionRepo.ReserveBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientBalance) {
			return nil, ErrInsufficientBalance
		}
		return nil, err
	}

	amountFloat := -float32(amount)
	balanceBeforeFloat := float32(snapshot.BalanceBefore)
	balanceAfterFloat := float32(snapshot.BalanceAfter)
	frozenBeforeFloat := float32(snapshot.FrozenBefore)
	frozenAfterFloat := float32(snapshot.FrozenAfter)
	transactionType := domain.TransactionTypeFreeze
	refID := referenceID
	noteCopy := notes

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amountFloat,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		FrozenBefore:  &frozenBeforeFloat,
		FrozenAfter:   &frozenAfterFloat,
		Type:          &transactionType,
		ReferenceID:   &refID,
		Notes:         &noteCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.CreateWithTx(ctx, tx, transaction); err != nil {
		return nil, fmt.Errorf("failed to create reserve transaction: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit reserve transaction: %w", err)
	}
	committed = true
	return transaction, nil
}

func (s *transactionService) CommitReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin commit transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	snapshot, err := s.transactionRepo.CommitReservedBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientFrozenFunds) {
			return nil, ErrInsufficientFrozen
		}
		return nil, err
	}

	zeroAmount := float32(0)
	balanceBeforeFloat := float32(snapshot.BalanceBefore)
	balanceAfterFloat := float32(snapshot.BalanceAfter)
	frozenBeforeFloat := float32(snapshot.FrozenBefore)
	frozenAfterFloat := float32(snapshot.FrozenAfter)
	transactionType := domain.TransactionTypeFreezeToCharge
	refID := referenceID
	noteCopy := notes

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &zeroAmount,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		FrozenBefore:  &frozenBeforeFloat,
		FrozenAfter:   &frozenAfterFloat,
		Type:          &transactionType,
		ReferenceID:   &refID,
		Notes:         &noteCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.CreateWithTx(ctx, tx, transaction); err != nil {
		return nil, fmt.Errorf("failed to create commit transaction: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit commit transaction: %w", err)
	}
	committed = true
	return transaction, nil
}

func (s *transactionService) ReleaseReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin release transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	snapshot, err := s.transactionRepo.ReleaseReservedBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientFrozenFunds) {
			return nil, ErrInsufficientFrozen
		}
		return nil, err
	}

	amountFloat := float32(amount)
	balanceBeforeFloat := float32(snapshot.BalanceBefore)
	balanceAfterFloat := float32(snapshot.BalanceAfter)
	frozenBeforeFloat := float32(snapshot.FrozenBefore)
	frozenAfterFloat := float32(snapshot.FrozenAfter)
	transactionType := domain.TransactionTypeUnfreeze
	refID := referenceID
	noteCopy := notes

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amountFloat,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		FrozenBefore:  &frozenBeforeFloat,
		FrozenAfter:   &frozenAfterFloat,
		Type:          &transactionType,
		ReferenceID:   &refID,
		Notes:         &noteCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.CreateWithTx(ctx, tx, transaction); err != nil {
		return nil, fmt.Errorf("failed to create release transaction: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit release transaction: %w", err)
	}
	committed = true
	return transaction, nil
}

// GetBalance returns current customer balance
func (s *transactionService) GetBalance(ctx context.Context, customerID int64) (float64, error) {
	balance, err := s.transactionRepo.GetBalance(ctx, customerID)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}
	return balance, nil
}

// GetBalanceDetail returns both available balance and frozen amount
func (s *transactionService) GetBalanceDetail(ctx context.Context, customerID int64) (*BalanceDetail, error) {
	balance, frozen, err := s.transactionRepo.GetBalanceDetail(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance detail: %w", err)
	}
	return &BalanceDetail{
		Balance:      balance,
		FrozenAmount: frozen,
	}, nil
}

// GetTransactionHistory returns paginated transaction history
func (s *transactionService) GetTransactionHistory(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error) {
	// Validate pagination parameters
	if limit <= 0 || limit > 100 {
		limit = 20 // Default limit
	}
	if offset < 0 {
		offset = 0
	}

	transactions, total, err := s.transactionRepo.FindByCustomerID(ctx, customerID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get transaction history: %w", err)
	}

	return transactions, total, nil
}

// GetTransactionsByType returns transactions filtered by type
func (s *transactionService) GetTransactionsByType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error) {
	// Validate transaction type
	if transactionType != 1 && transactionType != 2 {
		return nil, 0, fmt.Errorf("invalid transaction type: %d", transactionType)
	}

	// Validate pagination parameters
	if limit <= 0 || limit > 100 {
		limit = 20 // Default limit
	}
	if offset < 0 {
		offset = 0
	}

	transactions, total, err := s.transactionRepo.FindByCustomerIDAndType(ctx, customerID, transactionType, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get transactions by type: %w", err)
	}

	return transactions, total, nil
}

// GetTransactionsByDateRange returns transactions within date range
func (s *transactionService) GetTransactionsByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error) {
	// Validate date range
	if endDate.Before(startDate) {
		return nil, 0, errors.New("end date must be after start date")
	}

	// Validate pagination parameters
	if limit <= 0 || limit > 100 {
		limit = 20 // Default limit
	}
	if offset < 0 {
		offset = 0
	}

	transactions, total, err := s.transactionRepo.FindByDateRange(ctx, customerID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get transactions by date range: %w", err)
	}

	return transactions, total, nil
}
