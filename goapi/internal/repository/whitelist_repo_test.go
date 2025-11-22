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

func TestWhitelistRepository_Create(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	whitelist := &domain.IPWhitelist{
		CustomerID: 1,
		IPAddress:  "192.168.1.1",
		Notes:      "Office IP",
		CreatedAt:  time.Now(),
	}

	err := repo.Create(ctx, whitelist)
	assert.NoError(t, err)
	assert.NotZero(t, whitelist.ID)
}

func TestWhitelistRepository_FindByID(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	whitelist := &domain.IPWhitelist{
		CustomerID: 1,
		IPAddress:  "192.168.1.1",
		Notes:      "Office IP",
		CreatedAt:  time.Now(),
	}
	db.Create(whitelist)

	// Test FindByID
	found, err := repo.FindByID(ctx, whitelist.ID)
	assert.NoError(t, err)
	assert.Equal(t, whitelist.CustomerID, found.CustomerID)
	assert.Equal(t, whitelist.IPAddress, found.IPAddress)

	// Test FindByID with non-existent ID
	_, err = repo.FindByID(ctx, 999)
	assert.Error(t, err)
}

func TestWhitelistRepository_FindByCustomerID(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	for i := 1; i <= 5; i++ {
		whitelist := &domain.IPWhitelist{
			CustomerID: 1,
			IPAddress:  fmt.Sprintf("192.168.1.%d", i),
			Notes:      fmt.Sprintf("IP %d", i),
			CreatedAt:  time.Now().Add(-time.Duration(i) * time.Hour),
		}
		db.Create(whitelist)
	}

	// Test pagination
	whitelists, total, err := repo.FindByCustomerID(ctx, 1, 3, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), total)
	assert.Len(t, whitelists, 3)

	// Test second page
	whitelists, total, err = repo.FindByCustomerID(ctx, 1, 3, 3)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), total)
	assert.Len(t, whitelists, 2)
}

func TestWhitelistRepository_FindByCustomerIDAndIP(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	whitelist := &domain.IPWhitelist{
		CustomerID: 1,
		IPAddress:  "192.168.1.1",
		Notes:      "Office IP",
		CreatedAt:  time.Now(),
	}
	db.Create(whitelist)

	// Test FindByCustomerIDAndIP
	found, err := repo.FindByCustomerIDAndIP(ctx, 1, "192.168.1.1")
	assert.NoError(t, err)
	assert.Equal(t, whitelist.ID, found.ID)

	// Test with non-existent IP
	_, err = repo.FindByCustomerIDAndIP(ctx, 1, "10.0.0.1")
	assert.Error(t, err)
}

func TestWhitelistRepository_Update(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	whitelist := &domain.IPWhitelist{
		CustomerID: 1,
		IPAddress:  "192.168.1.1",
		Notes:      "Office IP",
		CreatedAt:  time.Now(),
	}
	db.Create(whitelist)

	// Update
	whitelist.Notes = "Updated Office IP"
	err := repo.Update(ctx, whitelist)
	assert.NoError(t, err)

	// Verify update
	found, _ := repo.FindByID(ctx, whitelist.ID)
	assert.Equal(t, "Updated Office IP", found.Notes)
}

func TestWhitelistRepository_Delete(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	whitelist := &domain.IPWhitelist{
		CustomerID: 1,
		IPAddress:  "192.168.1.1",
		Notes:      "Office IP",
		CreatedAt:  time.Now(),
	}
	db.Create(whitelist)

	// Delete
	err := repo.Delete(ctx, whitelist.ID)
	assert.NoError(t, err)

	// Verify deletion
	_, err = repo.FindByID(ctx, whitelist.ID)
	assert.Error(t, err)
}

func TestWhitelistRepository_DeleteByCustomerIDAndIP(t *testing.T) {
	db := setupWhitelistTestDB()
	repo := NewWhitelistRepository(db)
	ctx := context.Background()

	// Create test data
	whitelist := &domain.IPWhitelist{
		CustomerID: 1,
		IPAddress:  "192.168.1.1",
		Notes:      "Office IP",
		CreatedAt:  time.Now(),
	}
	db.Create(whitelist)

	// Delete by customer ID and IP
	err := repo.DeleteByCustomerIDAndIP(ctx, 1, "192.168.1.1")
	assert.NoError(t, err)

	// Verify deletion
	_, err = repo.FindByCustomerIDAndIP(ctx, 1, "192.168.1.1")
	assert.Error(t, err)
}
