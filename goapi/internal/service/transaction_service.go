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

	// Refund returns funds to customer account (used for expired phone number refunds)
	Refund(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// Deduct subtracts funds from customer account
	Deduct(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// ReserveFunds moves funds from balance to frozen balance
	ReserveFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// CommitReservedFunds consumes previously frozen funds without touching balance
	CommitReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// ReleaseReservedFunds returns frozen funds back to available balance
	ReleaseReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// ReserveAndCommitFunds atomically reserves and commits funds in a single transaction
	// Returns two transaction records: [0]=reserve, [1]=commit
	ReserveAndCommitFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, reserveNotes, commitNotes string) ([]*domain.Transaction, error)

	// ReserveAndCommitFundsWithoutRecord atomically reserves and commits funds without creating transaction records
	// Only updates balance and frozen amount in customer table
	ReserveAndCommitFundsWithoutRecord(ctx context.Context, customerID int64, amount float64) error

	// CreateDeductRecord creates a deduct transaction record without actually deducting balance
	// Used for recording deductions that have already been completed by other means
	CreateDeductRecord(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

	// ReserveAndCommitWithSingleRecord atomically reserves and commits funds, but only creates a reserve transaction record
	// The commit step (frozen to charged) does not create a record because balance doesn't change
	ReserveAndCommitWithSingleRecord(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error)

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

// Refund returns funds to customer account (Type 3: 拉号回退/退款)
func (s *transactionService) Refund(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
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
	transactionType := domain.TransactionTypeRefund
	refID := referenceID

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amountFloat,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		Type:          &transactionType, // 3 = refund
		ReferenceID:   &refID,
		Notes:         &notes,
		CreatedAt:     time.Now(),
	}

	// Create transaction
	err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create refund transaction: %w", err)
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

	// 冻结转实扣：Amount应该记录负的金额（表示扣款），FrozenBefore和FrozenAfter记录冻结金额的变化
	negativeAmount := float32(-amount) // 负数表示扣款
	balanceBeforeFloat := float32(snapshot.BalanceBefore)
	balanceAfterFloat := float32(snapshot.BalanceAfter)
	frozenBeforeFloat := float32(snapshot.FrozenBefore)
	frozenAfterFloat := float32(snapshot.FrozenAfter)
	transactionType := domain.TransactionTypeFreezeToCharge
	refID := referenceID
	noteCopy := notes

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &negativeAmount, // 记录负的金额，表示从冻结金额中扣除
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

// ReserveAndCommitFunds 原子操作：预冻结 + 立即扣款
// 这个方法将两个操作合并到一个事务中，避免并发问题
// 返回两个交易记录：[0]=预冻结记录, [1]=冻结转实扣记录
func (s *transactionService) ReserveAndCommitFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, reserveNotes, commitNotes string) ([]*domain.Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// 开启一个事务，将预冻结和扣款合并
	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	// 步骤1: 预冻结（余额 -> 冻结金额）
	reserveSnapshot, err := s.transactionRepo.ReserveBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientBalance) {
			return nil, ErrInsufficientBalance
		}
		return nil, fmt.Errorf("reserve balance failed: %w", err)
	}

	// 创建预冻结交易记录
	reserveAmountFloat := -float32(amount)
	reserveBalanceBefore := float32(reserveSnapshot.BalanceBefore)
	reserveBalanceAfter := float32(reserveSnapshot.BalanceAfter)
	reserveFrozenBefore := float32(reserveSnapshot.FrozenBefore)
	reserveFrozenAfter := float32(reserveSnapshot.FrozenAfter)
	reserveType := domain.TransactionTypeFreeze
	reserveRefID := referenceID
	reserveNotesCopy := reserveNotes

	reserveTx := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &reserveAmountFloat,
		BalanceBefore: &reserveBalanceBefore,
		BalanceAfter:  &reserveBalanceAfter,
		FrozenBefore:  &reserveFrozenBefore,
		FrozenAfter:   &reserveFrozenAfter,
		Type:          &reserveType,
		ReferenceID:   &reserveRefID,
		Notes:         &reserveNotesCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.CreateWithTx(ctx, tx, reserveTx); err != nil {
		return nil, fmt.Errorf("failed to create reserve transaction: %w", err)
	}

	// 步骤2: 冻结转实扣（冻结金额减少，余额不变）
	commitSnapshot, err := s.transactionRepo.CommitReservedBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientFrozenFunds) {
			return nil, ErrInsufficientFrozen
		}
		return nil, fmt.Errorf("commit reserved balance failed: %w", err)
	}

	// 创建冻结转实扣交易记录
	commitAmountFloat := -float32(amount) // 负数表示扣款
	commitBalanceBefore := float32(commitSnapshot.BalanceBefore)
	commitBalanceAfter := float32(commitSnapshot.BalanceAfter)
	commitFrozenBefore := float32(commitSnapshot.FrozenBefore)
	commitFrozenAfter := float32(commitSnapshot.FrozenAfter)
	commitType := domain.TransactionTypeFreezeToCharge
	commitRefID := referenceID
	commitNotesCopy := commitNotes

	commitTx := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &commitAmountFloat,
		BalanceBefore: &commitBalanceBefore,
		BalanceAfter:  &commitBalanceAfter,
		FrozenBefore:  &commitFrozenBefore,
		FrozenAfter:   &commitFrozenAfter,
		Type:          &commitType,
		ReferenceID:   &commitRefID,
		Notes:         &commitNotesCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.CreateWithTx(ctx, tx, commitTx); err != nil {
		return nil, fmt.Errorf("failed to create commit transaction: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	return []*domain.Transaction{reserveTx, commitTx}, nil
}

// ReserveAndCommitFundsWithoutRecord 原子操作：预冻结 + 冻结转实扣，但不创建交易记录
// 这个方法用于简化交易记录，只更新余额和冻结金额，不写入交易表
func (s *transactionService) ReserveAndCommitFundsWithoutRecord(ctx context.Context, customerID int64, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	// 开启一个事务，将预冻结和扣款合并
	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	// 步骤1: 预冻结（余额 -> 冻结金额）
	_, err = s.transactionRepo.ReserveBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientBalance) {
			return ErrInsufficientBalance
		}
		return fmt.Errorf("reserve balance failed: %w", err)
	}

	// 步骤2: 冻结转实扣（冻结金额减少，余额不变）
	_, err = s.transactionRepo.CommitReservedBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientFrozenFunds) {
			return ErrInsufficientFrozen
		}
		return fmt.Errorf("commit reserved balance failed: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	return nil
}

// CreateDeductRecord 只创建扣款交易记录，不实际操作余额
// 用于记录已经通过其他方式完成的扣款操作
func (s *transactionService) CreateDeductRecord(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// 获取当前余额（已经扣款后的余额）
	currentBalance, err := s.transactionRepo.GetBalance(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current balance: %w", err)
	}

	// 创建交易记录（记录扣款前后的余额）
	amountFloat := -float32(amount)
	balanceBeforeFloat := float32(currentBalance + amount) // 扣款前的余额
	balanceAfterFloat := float32(currentBalance)           // 扣款后的余额（当前余额）
	transactionType := domain.TransactionTypeDeduct
	refID := referenceID
	notesCopy := notes

	transaction := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amountFloat,
		BalanceBefore: &balanceBeforeFloat,
		BalanceAfter:  &balanceAfterFloat,
		Type:          &transactionType,
		ReferenceID:   &refID,
		Notes:         &notesCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.Create(ctx, transaction); err != nil {
		return nil, fmt.Errorf("failed to create deduct record: %w", err)
	}

	return transaction, nil
}

// ReserveAndCommitWithSingleRecord 原子操作：预冻结 + 冻结转实扣，但只为预冻结创建交易记录
// 冻结转实扣不创建记录（因为余额没变化，只是冻结金额变化）
func (s *transactionService) ReserveAndCommitWithSingleRecord(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// 开启一个事务，将预冻结和扣款合并
	tx, err := s.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	// 步骤1: 预冻结（余额 -> 冻结金额）
	reserveSnapshot, err := s.transactionRepo.ReserveBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientBalance) {
			return nil, ErrInsufficientBalance
		}
		return nil, fmt.Errorf("reserve balance failed: %w", err)
	}

	// 创建预冻结交易记录（余额变动了）
	reserveAmountFloat := -float32(amount)
	reserveBalanceBefore := float32(reserveSnapshot.BalanceBefore)
	reserveBalanceAfter := float32(reserveSnapshot.BalanceAfter)
	reserveFrozenBefore := float32(reserveSnapshot.FrozenBefore)
	reserveFrozenAfter := float32(reserveSnapshot.FrozenAfter)
	reserveType := domain.TransactionTypeFreeze
	reserveRefID := referenceID
	reserveNotesCopy := notes

	reserveTx := &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &reserveAmountFloat,
		BalanceBefore: &reserveBalanceBefore,
		BalanceAfter:  &reserveBalanceAfter,
		FrozenBefore:  &reserveFrozenBefore,
		FrozenAfter:   &reserveFrozenAfter,
		Type:          &reserveType,
		ReferenceID:   &reserveRefID,
		Notes:         &reserveNotesCopy,
		CreatedAt:     time.Now(),
	}

	if err := s.transactionRepo.CreateWithTx(ctx, tx, reserveTx); err != nil {
		return nil, fmt.Errorf("failed to create reserve transaction: %w", err)
	}

	// 步骤2: 冻结转实扣（冻结金额减少，余额不变）
	// 这一步不创建交易记录，因为余额没有变化，只是冻结金额减少
	_, err = s.transactionRepo.CommitReservedBalance(ctx, tx, customerID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientFrozenFunds) {
			return nil, ErrInsufficientFrozen
		}
		return nil, fmt.Errorf("commit reserved balance failed: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	return reserveTx, nil
}
