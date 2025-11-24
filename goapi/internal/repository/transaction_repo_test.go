package repository

import (
	"context"
	"fmt"
	"sms-platform/goapi/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTransactionTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&domain.Transaction{}, &domain.Customer{})
	return db
}

func setupTransactionRepo(db *gorm.DB) TransactionRepository {
	customerRepo := NewCustomerRepository(db)
	return NewTransactionRepository(db, customerRepo)
}

// Helper function to create Transaction with pointer fields
func createTestTransaction(customerID int64, amount float32, balanceBefore, balanceAfter float32, txType, notes string, createdAt time.Time) *domain.Transaction {
	return &domain.Transaction{
		CustomerID:    customerID,
		Amount:        &amount,
		BalanceBefore: &balanceBefore,
		BalanceAfter:  &balanceAfter,
		Type:          &txType,
		Notes:         &notes,
		CreatedAt:     createdAt,
	}
}

func TestTransactionRepository_Create(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	transaction := createTestTransaction(1, 100.50, 0, 100.50, "1", "Initial deposit", time.Now())

	err := repo.Create(ctx, transaction)
	assert.NoError(t, err)
	assert.NotZero(t, transaction.ID)
}

func TestTransactionRepository_FindByID(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	// Create test data
	transaction := createTestTransaction(1, 100.50, 0, 100.50, "1", "Test transaction", time.Now())
	db.Create(transaction)

	// Test FindByID
	found, err := repo.FindByID(ctx, transaction.ID)
	assert.NoError(t, err)
	assert.Equal(t, transaction.CustomerID, found.CustomerID)
	assert.Equal(t, transaction.Amount, found.Amount)

	// Test FindByID with non-existent ID
	_, err = repo.FindByID(ctx, 999)
	assert.Error(t, err)
}

func TestTransactionRepository_FindByCustomerID(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	// Create test data
	for i := 1; i <= 5; i++ {
		transaction := createTestTransaction(1, float32(i*10), 0, float32(i*10), "1", fmt.Sprintf("Transaction %d", i), time.Now().Add(-time.Duration(i)*time.Hour))
		db.Create(transaction)
	}

	// Test pagination
	transactions, total, err := repo.FindByCustomerID(ctx, 1, 3, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), total)
	assert.Len(t, transactions, 3)
}

func TestTransactionRepository_FindByCustomerIDAndType(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	// Create test data with different types
	for i := 1; i <= 3; i++ {
		transaction := createTestTransaction(1, float32(i*10), 0, 0, "1", "", time.Now().Add(-time.Duration(i)*time.Hour))
		db.Create(transaction)
	}

	for i := 1; i <= 2; i++ {
		transaction := createTestTransaction(1, float32(i*5), 0, 0, "2", "", time.Now().Add(-time.Duration(i)*time.Hour))
		db.Create(transaction)
	}

	// Test finding topup transactions
	transactions, total, err := repo.FindByCustomerIDAndType(ctx, 1, 1, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), total)
	assert.Len(t, transactions, 3)

	// Test finding deduct transactions
	transactions, total, err = repo.FindByCustomerIDAndType(ctx, 1, 2, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, transactions, 2)
}

func TestTransactionRepository_GetBalance(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	// Create test customers with balances
	customer1 := &domain.Customer{
		ID:           1,
		Balance:      120.0,
		APISecretKey: "test-key-1",
	}
	customer2 := &domain.Customer{
		ID:           2,
		Balance:      200.0,
		APISecretKey: "test-key-2",
	}
	db.Create(customer1)
	db.Create(customer2)

	// Test balance for customer 1
	balance, err := repo.GetBalance(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 120.0, balance)

	// Test balance for customer 2
	balance, err = repo.GetBalance(ctx, 2)
	assert.NoError(t, err)
	assert.Equal(t, 200.0, balance)

	// Test balance for non-existent customer
	balance, err = repo.GetBalance(ctx, 999)
	assert.Error(t, err) // Should return error for non-existent customer
}

func TestTransactionRepository_FindByDateRange(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	// Create test data
	transactions := []*domain.Transaction{
		createTestTransaction(1, 100, 0, 100, "1", "", yesterday.Add(-1*time.Hour)), // before range
		createTestTransaction(1, 50, 100, 150, "1", "", now),                        // in range
		createTestTransaction(1, 20, 150, 130, "2", "", now.Add(1*time.Hour)),       // in range
		createTestTransaction(1, 30, 130, 160, "1", "", tomorrow.Add(1*time.Hour)),  // after range
	}

	for _, tx := range transactions {
		db.Create(tx)
	}

	// Test date range query
	result, total, err := repo.FindByDateRange(ctx, 1, yesterday, tomorrow, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, result, 2)
}

func TestTransactionRepository_Update(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	// Create test data
	transaction := createTestTransaction(1, 100.50, 0, 100.50, "1", "Initial notes", time.Now())
	db.Create(transaction)

	// Update
	notes := "Updated notes"
	transaction.Notes = &notes
	err := repo.Update(ctx, transaction)
	assert.NoError(t, err)

	// Verify update
	found, _ := repo.FindByID(ctx, transaction.ID)
	assert.Equal(t, "Updated notes", *found.Notes)
}

func TestTransactionRepository_BeginTx(t *testing.T) {
	db := setupTransactionTestDB()
	repo := setupTransactionRepo(db)
	ctx := context.Background()

	// Test transaction begin
	tx, err := repo.BeginTx(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	// Test rollback
	tx.Rollback()
}
