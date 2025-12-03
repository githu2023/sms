package service

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MockTransactionRepository is a mock implementation of the TransactionRepository interface
type MockTransactionRepository struct {
	mock.Mock
}

func createServiceTestTransaction(id int64, customerID int64, amount float32, balanceBefore, balanceAfter float32, txType, notes string) *domain.Transaction {
	return &domain.Transaction{
		ID:            id,
		CustomerID:    customerID,
		Amount:        &amount,
		BalanceBefore: &balanceBefore,
		BalanceAfter:  &balanceAfter,
		Type:          &txType,
		Notes:         &notes,
		CreatedAt:     time.Now(),
	}
}

func createTestTx(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	tx := db.Begin()
	require.NoError(t, tx.Error)
	return tx
}

func (m *MockTransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, transaction *domain.Transaction) error {
	args := m.Called(ctx, tx, transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindByID(ctx context.Context, id int64) (*domain.Transaction, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindByCustomerID(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *MockTransactionRepository) FindByCustomerIDAndType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, transactionType, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *MockTransactionRepository) FindByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, startDate, endDate, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *MockTransactionRepository) GetBalance(ctx context.Context, customerID int64) (float64, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockTransactionRepository) GetBalanceDetail(ctx context.Context, customerID int64) (float64, float64, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).(float64), args.Get(1).(float64), args.Error(2)
}

func (m *MockTransactionRepository) GetBalanceForUpdate(ctx context.Context, tx *gorm.DB, customerID int64) (float64, error) {
	return m.GetBalance(ctx, customerID)
}

func (m *MockTransactionRepository) ReserveBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*repository.BalanceChangeSnapshot, error) {
	args := m.Called(ctx, tx, customerID, amount)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repository.BalanceChangeSnapshot), args.Error(1)
}

func (m *MockTransactionRepository) CommitReservedBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*repository.BalanceChangeSnapshot, error) {
	args := m.Called(ctx, tx, customerID, amount)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repository.BalanceChangeSnapshot), args.Error(1)
}

func (m *MockTransactionRepository) ReleaseReservedBalance(ctx context.Context, tx *gorm.DB, customerID int64, amount float64) (*repository.BalanceChangeSnapshot, error) {
	args := m.Called(ctx, tx, customerID, amount)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repository.BalanceChangeSnapshot), args.Error(1)
}

func (m *MockTransactionRepository) Update(ctx context.Context, transaction *domain.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) BeginTx(ctx context.Context) (*gorm.DB, error) {
	args := m.Called(ctx)
	return args.Get(0).(*gorm.DB), args.Error(1)
}

func TestTransactionService_TopUp_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 100.0
	currentBalance := 50.0
	notes := "Initial deposit"

	// Setup expectations
	mockRepo.On("GetBalance", ctx, customerID).Return(currentBalance, nil)
	mockRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(nil)

	// Execute TopUp
	transaction, err := service.TopUp(ctx, customerID, amount, notes)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, customerID, transaction.CustomerID)
	assert.Equal(t, float32(amount), *transaction.Amount)
	assert.Equal(t, float32(currentBalance), *transaction.BalanceBefore)
	assert.Equal(t, float32(currentBalance+amount), *transaction.BalanceAfter)
	assert.Equal(t, domain.TransactionTypeTopUp, *transaction.Type)
	assert.Equal(t, notes, *transaction.Notes)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_TopUp_InvalidAmount(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	// Test with negative amount
	transaction, err := service.TopUp(ctx, 1, -100.0, "Invalid")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAmount, err)
	assert.Nil(t, transaction)

	// Test with zero amount
	transaction, err = service.TopUp(ctx, 1, 0, "Invalid")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAmount, err)
	assert.Nil(t, transaction)
}

func TestTransactionService_TopUp_TransactionBeginFailed(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	mockRepo.On("GetBalance", ctx, int64(1)).Return(0.0, errors.New("db error"))

	transaction, err := service.TopUp(ctx, 1, 100.0, "Test")
	assert.Error(t, err)
	assert.Nil(t, transaction)
	assert.Contains(t, err.Error(), "failed to get current balance")

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_Deduct_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 30.0
	currentBalance := 100.0
	referenceID := int64(123)
	notes := "SMS charge"

	// Setup expectations
	mockRepo.On("GetBalance", ctx, customerID).Return(currentBalance, nil)
	mockRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(nil)

	// Execute Deduct
	transaction, err := service.Deduct(ctx, customerID, amount, referenceID, notes)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, customerID, transaction.CustomerID)
	assert.Equal(t, float32(-amount), *transaction.Amount) // Deduct amounts should be negative
	assert.Equal(t, float32(currentBalance), *transaction.BalanceBefore)
	assert.Equal(t, float32(currentBalance-amount), *transaction.BalanceAfter)
	assert.Equal(t, domain.TransactionTypeDeduct, *transaction.Type)
	assert.Equal(t, referenceID, *transaction.ReferenceID)
	assert.Equal(t, notes, *transaction.Notes)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_Deduct_InsufficientBalance(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 150.0 // More than current balance
	currentBalance := 100.0

	// Setup expectations
	mockRepo.On("GetBalance", ctx, customerID).Return(currentBalance, nil)

	// Execute Deduct
	transaction, err := service.Deduct(ctx, customerID, amount, 0, "Test")

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, ErrInsufficientBalance, err)
	assert.Nil(t, transaction)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_Deduct_InvalidAmount(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	// Test with negative amount
	transaction, err := service.Deduct(ctx, 1, -50.0, 0, "Invalid")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAmount, err)
	assert.Nil(t, transaction)

	// Test with zero amount
	transaction, err = service.Deduct(ctx, 1, 0, 0, "Invalid")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAmount, err)
	assert.Nil(t, transaction)
}

func TestTransactionService_ReserveFunds_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 20.0
	referenceID := int64(99)
	notes := "reserve"

	tx := createTestTx(t)
	snapshot := &repository.BalanceChangeSnapshot{
		BalanceBefore: 100,
		BalanceAfter:  80,
		FrozenBefore:  10,
		FrozenAfter:   30,
	}

	mockRepo.On("BeginTx", ctx).Return(tx, nil)
	mockRepo.On("ReserveBalance", ctx, tx, customerID, amount).Return(snapshot, nil)
	mockRepo.On("CreateWithTx", ctx, tx, mock.AnythingOfType("*domain.Transaction")).Return(nil)

	transaction, err := service.ReserveFunds(ctx, customerID, amount, referenceID, notes)
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, domain.TransactionTypeFreeze, *transaction.Type)
	assert.Equal(t, float32(-amount), *transaction.Amount)
	assert.Equal(t, float32(snapshot.BalanceAfter), *transaction.BalanceAfter)
	assert.Equal(t, float32(snapshot.FrozenAfter), *transaction.FrozenAfter)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_ReserveFunds_InsufficientBalance(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 50.0
	tx := createTestTx(t)

	mockRepo.On("BeginTx", ctx).Return(tx, nil)
	mockRepo.On("ReserveBalance", ctx, tx, customerID, amount).Return((*repository.BalanceChangeSnapshot)(nil), repository.ErrInsufficientBalance)

	transaction, err := service.ReserveFunds(ctx, customerID, amount, 0, "reserve")
	assert.ErrorIs(t, err, ErrInsufficientBalance)
	assert.Nil(t, transaction)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_CommitReservedFunds_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 10.0
	tx := createTestTx(t)
	snapshot := &repository.BalanceChangeSnapshot{
		BalanceBefore: 80,
		BalanceAfter:  80,
		FrozenBefore:  25,
		FrozenAfter:   15,
	}

	mockRepo.On("BeginTx", ctx).Return(tx, nil)
	mockRepo.On("CommitReservedBalance", ctx, tx, customerID, amount).Return(snapshot, nil)
	mockRepo.On("CreateWithTx", ctx, tx, mock.AnythingOfType("*domain.Transaction")).Return(nil)

	transaction, err := service.CommitReservedFunds(ctx, customerID, amount, 10, "commit")
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, domain.TransactionTypeFreezeToCharge, *transaction.Type)
	assert.Equal(t, float32(-amount), *transaction.Amount) // 修正：应该是负的金额，表示扣款
	assert.Equal(t, float32(snapshot.FrozenAfter), *transaction.FrozenAfter)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_CommitReservedFunds_InsufficientFrozen(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 30.0
	tx := createTestTx(t)

	mockRepo.On("BeginTx", ctx).Return(tx, nil)
	mockRepo.On("CommitReservedBalance", ctx, tx, customerID, amount).Return((*repository.BalanceChangeSnapshot)(nil), repository.ErrInsufficientFrozenFunds)

	transaction, err := service.CommitReservedFunds(ctx, customerID, amount, 0, "commit")
	assert.ErrorIs(t, err, ErrInsufficientFrozen)
	assert.Nil(t, transaction)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_ReleaseReservedFunds_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 15.0
	tx := createTestTx(t)
	snapshot := &repository.BalanceChangeSnapshot{
		BalanceBefore: 60,
		BalanceAfter:  75,
		FrozenBefore:  20,
		FrozenAfter:   5,
	}

	mockRepo.On("BeginTx", ctx).Return(tx, nil)
	mockRepo.On("ReleaseReservedBalance", ctx, tx, customerID, amount).Return(snapshot, nil)
	mockRepo.On("CreateWithTx", ctx, tx, mock.AnythingOfType("*domain.Transaction")).Return(nil)

	transaction, err := service.ReleaseReservedFunds(ctx, customerID, amount, 0, "release")
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, domain.TransactionTypeUnfreeze, *transaction.Type)
	assert.Equal(t, float32(amount), *transaction.Amount)
	assert.Equal(t, float32(snapshot.BalanceAfter), *transaction.BalanceAfter)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_ReleaseReservedFunds_InsufficientFrozen(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	amount := 25.0
	tx := createTestTx(t)

	mockRepo.On("BeginTx", ctx).Return(tx, nil)
	mockRepo.On("ReleaseReservedBalance", ctx, tx, customerID, amount).Return((*repository.BalanceChangeSnapshot)(nil), repository.ErrInsufficientFrozenFunds)

	transaction, err := service.ReleaseReservedFunds(ctx, customerID, amount, 0, "release")
	assert.ErrorIs(t, err, ErrInsufficientFrozen)
	assert.Nil(t, transaction)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetBalance_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	expectedBalance := 250.75

	mockRepo.On("GetBalance", ctx, customerID).Return(expectedBalance, nil)

	balance, err := service.GetBalance(ctx, customerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, balance)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetBalance_Error(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)

	mockRepo.On("GetBalance", ctx, customerID).Return(0.0, errors.New("db error"))

	balance, err := service.GetBalance(ctx, customerID)

	assert.Error(t, err)
	assert.Equal(t, 0.0, balance)
	assert.Contains(t, err.Error(), "failed to get balance")

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetBalanceDetail_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	mockRepo.On("GetBalanceDetail", ctx, customerID).Return(200.0, 50.0, nil)

	detail, err := service.GetBalanceDetail(ctx, customerID)
	assert.NoError(t, err)
	assert.NotNil(t, detail)
	assert.Equal(t, 200.0, detail.Balance)
	assert.Equal(t, 50.0, detail.FrozenAmount)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetBalanceDetail_Error(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	mockRepo.On("GetBalanceDetail", ctx, customerID).Return(0.0, 0.0, errors.New("db error"))

	detail, err := service.GetBalanceDetail(ctx, customerID)
	assert.Error(t, err)
	assert.Nil(t, detail)
	assert.Contains(t, err.Error(), "failed to get balance detail")

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionHistory_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	limit := 10
	offset := 0

	expectedTransactions := []*domain.Transaction{
		createServiceTestTransaction(1, customerID, 100, 0, 100, domain.TransactionTypeTopUp, ""),
		createServiceTestTransaction(2, customerID, 50, 100, 50, domain.TransactionTypeDeduct, ""),
	}
	expectedTotal := int64(25)

	mockRepo.On("FindByCustomerID", ctx, customerID, limit, offset).Return(expectedTransactions, expectedTotal, nil)

	transactions, total, err := service.GetTransactionHistory(ctx, customerID, limit, offset)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
	assert.Equal(t, expectedTotal, total)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionHistory_DefaultLimits(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)

	// Test with invalid limit (too high) - should default to 20
	mockRepo.On("FindByCustomerID", ctx, customerID, 20, 0).Return([]*domain.Transaction{}, int64(0), nil)
	_, _, err := service.GetTransactionHistory(ctx, customerID, 150, 0)
	assert.NoError(t, err)

	// Test with invalid offset (negative) - should default to 0
	mockRepo.On("FindByCustomerID", ctx, customerID, 20, 0).Return([]*domain.Transaction{}, int64(0), nil)
	_, _, err = service.GetTransactionHistory(ctx, customerID, 20, -5)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionsByType_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	transactionType := 1 // TopUp
	limit := 10
	offset := 0

	expectedTransactions := []*domain.Transaction{
		createServiceTestTransaction(1, customerID, 100, 0, 100, domain.TransactionTypeTopUp, ""),
		createServiceTestTransaction(3, customerID, 200, 100, 300, domain.TransactionTypeTopUp, ""),
	}
	expectedTotal := int64(15)

	mockRepo.On("FindByCustomerIDAndType", ctx, customerID, transactionType, limit, offset).Return(expectedTransactions, expectedTotal, nil)

	transactions, total, err := service.GetTransactionsByType(ctx, customerID, transactionType, limit, offset)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
	assert.Equal(t, expectedTotal, total)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionsByType_InvalidType(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	// Test with invalid transaction type
	transactions, total, err := service.GetTransactionsByType(ctx, 1, 5, 10, 0)

	assert.Error(t, err)
	assert.Nil(t, transactions)
	assert.Equal(t, int64(0), total)
	assert.Contains(t, err.Error(), "invalid transaction type")
}

func TestTransactionService_GetTransactionsByDateRange_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	customerID := int64(1)
	startDate := time.Now().Add(-7 * 24 * time.Hour) // 7 days ago
	endDate := time.Now()
	limit := 10
	offset := 0

	tx1 := createServiceTestTransaction(1, customerID, 100, 0, 100, domain.TransactionTypeTopUp, "")
	tx1.CreatedAt = startDate.Add(1 * time.Hour)
	tx2 := createServiceTestTransaction(2, customerID, 50, 100, 50, domain.TransactionTypeDeduct, "")
	tx2.CreatedAt = endDate.Add(-1 * time.Hour)

	expectedTransactions := []*domain.Transaction{tx1, tx2}
	expectedTotal := int64(8)

	mockRepo.On("FindByDateRange", ctx, customerID, startDate, endDate, limit, offset).Return(expectedTransactions, expectedTotal, nil)

	transactions, total, err := service.GetTransactionsByDateRange(ctx, customerID, startDate, endDate, limit, offset)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
	assert.Equal(t, expectedTotal, total)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionsByDateRange_InvalidDateRange(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	startDate := time.Now()
	endDate := startDate.Add(-1 * time.Hour) // End before start

	// Test with invalid date range
	transactions, total, err := service.GetTransactionsByDateRange(ctx, 1, startDate, endDate, 10, 0)

	assert.Error(t, err)
	assert.Nil(t, transactions)
	assert.Equal(t, int64(0), total)
	assert.Contains(t, err.Error(), "end date must be after start date")
}
