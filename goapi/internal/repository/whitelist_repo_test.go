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

func setupWhitelistTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&domain.IPWhitelist{})
	return db
}

func createTestIPWhitelist(customerID int64, ip, remark string) *domain.IPWhitelist {
	status := true
	return &domain.IPWhitelist{
		CustomerID: customerID,
		IPAddress:  &ip,
		Status:     &status,
		Remark:     &remark,
		CreatedAt:  time.Now(),
	}
}

func TestWhitelistRepository_Create(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	whitelist := createTestIPWhitelist(1, "192.168.1.1", "Office IP")

	err := repo.Create(ctx, whitelist)
	assert.NoError(t, err)
	assert.NotZero(t, whitelist.ID)
}

func TestWhitelistRepository_FindByID(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	whitelist := createTestIPWhitelist(1, "192.168.1.1", "Office IP")
	db.Create(whitelist)

	found, err := repo.FindByID(ctx, whitelist.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, whitelist.ID, found.ID)
	assert.Equal(t, "192.168.1.1", *found.IPAddress)
}

func TestWhitelistRepository_FindByCustomerID(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	for i := 1; i <= 5; i++ {
		whitelist := createTestIPWhitelist(1, fmt.Sprintf("192.168.1.%d", i), fmt.Sprintf("IP %d", i))
		db.Create(whitelist)
	}

	// Test pagination
	whitelists, total, err := repo.FindByCustomerID(ctx, 1, 3, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), total)
	assert.Len(t, whitelists, 3)
}

func TestWhitelistRepository_Update(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	whitelist := createTestIPWhitelist(1, "192.168.1.1", "Office IP")
	db.Create(whitelist)

	// Update the whitelist
	newRemark := "Updated Office IP"
	whitelist.Remark = &newRemark

	err := repo.Update(ctx, whitelist)
	assert.NoError(t, err)

	// Verify update
	found, _ := repo.FindByID(ctx, whitelist.ID)
	assert.Equal(t, "Updated Office IP", *found.Remark)
}

func TestWhitelistRepository_Delete(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	whitelist := createTestIPWhitelist(1, "192.168.1.1", "Office IP")
	db.Create(whitelist)

	err := repo.Delete(ctx, whitelist.ID)
	assert.NoError(t, err)

	// Verify deletion
	found, err := repo.FindByID(ctx, whitelist.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}
