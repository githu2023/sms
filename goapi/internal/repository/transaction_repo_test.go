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
	db.AutoMigrate(&domain.Transaction{})
	return db
}

func TestTransactionRepository_Create(t *testing.T) {
	db := setupTransactionTestDB()
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	transaction := &domain.Transaction{
		CustomerID:    1,
		Amount:        100.50,
		BalanceBefore: 0,
		BalanceAfter:  100.50,
		Type:          "1", // topup
		Notes:         "Initial deposit",
		CreatedAt:     time.Now(),
	}

	err := repo.Create(ctx, transaction)
	assert.NoError(t, err)
	assert.NotZero(t, transaction.ID)
}

func TestTransactionRepository_FindByID(t *testing.T) {
	db := setupTransactionTestDB()
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	// Create test data
	transaction := &domain.Transaction{
		CustomerID:    1,
		Amount:        100.50,
		BalanceBefore: 0,
		BalanceAfter:  100.50,
		Type:          "1",
		Notes:         "Test transaction",
		CreatedAt:     time.Now(),
	}
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
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	// Create test data
	for i := 1; i <= 5; i++ {
		transaction := &domain.Transaction{
			CustomerID:    1,
			Amount:        float64(i * 10),
			BalanceBefore: 0,
			BalanceAfter:  float64(i * 10),
			Type:          "1",
			Notes:         fmt.Sprintf("Transaction %d", i),
			CreatedAt:     time.Now().Add(-time.Duration(i) * time.Hour),
		}
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
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	// Create test data with different types
	for i := 1; i <= 3; i++ {
		transaction := &domain.Transaction{
			CustomerID: 1,
			Amount:     float64(i * 10),
			Type:       "1", // topup
			CreatedAt:  time.Now().Add(-time.Duration(i) * time.Hour),
		}
		db.Create(transaction)
	}

	for i := 1; i <= 2; i++ {
		transaction := &domain.Transaction{
			CustomerID: 1,
			Amount:     float64(i * 5),
			Type:       "2", // deduct
			CreatedAt:  time.Now().Add(-time.Duration(i) * time.Hour),
		}
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
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	// Create test transactions
	transactions := []*domain.Transaction{
		{CustomerID: 1, Amount: 100, Type: "1", CreatedAt: time.Now()}, // topup +100
		{CustomerID: 1, Amount: 20, Type: "2", CreatedAt: time.Now()},  // deduct -20
		{CustomerID: 1, Amount: 50, Type: "1", CreatedAt: time.Now()},  // topup +50
		{CustomerID: 1, Amount: 10, Type: "2", CreatedAt: time.Now()},  // deduct -10
		{CustomerID: 2, Amount: 200, Type: "1", CreatedAt: time.Now()}, // different customer
	}

	for _, tx := range transactions {
		db.Create(tx)
	}

	// Test balance calculation for customer 1
	balance, err := repo.GetBalance(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 120.0, balance) // 100 - 20 + 50 - 10 = 120

	// Test balance for customer 2
	balance, err = repo.GetBalance(ctx, 2)
	assert.NoError(t, err)
	assert.Equal(t, 200.0, balance)

	// Test balance for non-existent customer
	balance, err = repo.GetBalance(ctx, 999)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, balance)
}

func TestTransactionRepository_FindByDateRange(t *testing.T) {
	db := setupTransactionTestDB()
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	// Create test data
	transactions := []*domain.Transaction{
		{CustomerID: 1, Amount: 100, Type: "1", CreatedAt: yesterday.Add(-1 * time.Hour)}, // before range
		{CustomerID: 1, Amount: 50, Type: "1", CreatedAt: now},                            // in range
		{CustomerID: 1, Amount: 20, Type: "2", CreatedAt: now.Add(1 * time.Hour)},         // in range
		{CustomerID: 1, Amount: 30, Type: "1", CreatedAt: tomorrow.Add(1 * time.Hour)},    // after range
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
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	// Create test data
	transaction := &domain.Transaction{
		CustomerID: 1,
		Amount:     100.50,
		Type:       "1",
		Notes:      "Initial notes",
		CreatedAt:  time.Now(),
	}
	db.Create(transaction)

	// Update
	transaction.Notes = "Updated notes"
	err := repo.Update(ctx, transaction)
	assert.NoError(t, err)

	// Verify update
	found, _ := repo.FindByID(ctx, transaction.ID)
	assert.Equal(t, "Updated notes", found.Notes)
}

func TestTransactionRepository_BeginTx(t *testing.T) {
	db := setupTransactionTestDB()
	repo := NewTransactionRepository(db)
	ctx := context.Background()

	// Test transaction begin
	tx, err := repo.BeginTx(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	// Test rollback
	tx.Rollback()
}
