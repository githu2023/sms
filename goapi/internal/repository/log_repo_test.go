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

func setupLogTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&domain.APILog{})
	return db
}

func TestLogRepository_Create(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	log := &domain.APILog{
		CustomerID:   1,
		RequestIP:    "192.168.1.1",
		RequestPath:  "/api/v1/get_phone",
		RequestBody:  "{\"business_type\": \"verification\"}",
		ResponseCode: 200,
		DurationMS:   150,
		CreatedAt:    time.Now(),
	}

	err := repo.Create(ctx, log)
	assert.NoError(t, err)
	assert.NotZero(t, log.ID)
}

func TestLogRepository_FindByID(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	// Create test data
	log := &domain.APILog{
		CustomerID:   1,
		RequestIP:    "192.168.1.1",
		RequestPath:  "/api/v1/get_phone",
		RequestBody:  "{\"business_type\": \"verification\"}",
		ResponseCode: 200,
		DurationMS:   150,
		CreatedAt:    time.Now(),
	}
	db.Create(log)

	// Test FindByID
	found, err := repo.FindByID(ctx, log.ID)
	assert.NoError(t, err)
	assert.Equal(t, log.CustomerID, found.CustomerID)
	assert.Equal(t, log.RequestPath, found.RequestPath)

	// Test FindByID with non-existent ID
	_, err = repo.FindByID(ctx, 999)
	assert.Error(t, err)
}

func TestLogRepository_FindByCustomerID(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	// Create test data
	for i := 1; i <= 5; i++ {
		log := &domain.APILog{
			CustomerID:   1,
			RequestIP:    fmt.Sprintf("192.168.1.%d", i),
			RequestPath:  "/api/v1/get_phone",
			ResponseCode: 200,
			DurationMS:   i * 100,
			CreatedAt:    time.Now().Add(-time.Duration(i) * time.Hour),
		}
		db.Create(log)
	}

	// Test pagination
	logs, total, err := repo.FindByCustomerID(ctx, 1, 3, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), total)
	assert.Len(t, logs, 3)

	// Test second page
	logs, total, err = repo.FindByCustomerID(ctx, 1, 3, 3)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), total)
	assert.Len(t, logs, 2)
}

func TestLogRepository_FindByPath(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	// Create test data with different paths
	paths := []string{"/api/v1/get_phone", "/api/v1/get_code", "/api/v1/balance"}
	for i, path := range paths {
		for j := 0; j < i+1; j++ {
			log := &domain.APILog{
				CustomerID:   int64(j + 1),
				RequestIP:    "192.168.1.1",
				RequestPath:  path,
				ResponseCode: 200,
				CreatedAt:    time.Now().Add(-time.Duration(j) * time.Hour),
			}
			db.Create(log)
		}
	}

	// Test finding logs by specific path
	logs, total, err := repo.FindByPath(ctx, "/api/v1/get_phone", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, logs, 1)

	logs, total, err = repo.FindByPath(ctx, "/api/v1/get_code", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, logs, 2)
}

func TestLogRepository_FindByDateRange(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	// Create test data
	logs := []*domain.APILog{
		{CustomerID: 1, RequestPath: "/api/v1/test1", CreatedAt: yesterday.Add(-1 * time.Hour)}, // before range
		{CustomerID: 1, RequestPath: "/api/v1/test2", CreatedAt: now},                           // in range
		{CustomerID: 1, RequestPath: "/api/v1/test3", CreatedAt: now.Add(1 * time.Hour)},        // in range
		{CustomerID: 1, RequestPath: "/api/v1/test4", CreatedAt: tomorrow.Add(1 * time.Hour)},   // after range
	}

	for _, log := range logs {
		db.Create(log)
	}

	// Test date range query
	result, total, err := repo.FindByDateRange(ctx, yesterday, tomorrow, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, result, 2)
}

func TestLogRepository_FindByCustomerIDAndPath(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	// Create test data
	logs := []*domain.APILog{
		{CustomerID: 1, RequestPath: "/api/v1/get_phone", CreatedAt: time.Now()},
		{CustomerID: 1, RequestPath: "/api/v1/get_code", CreatedAt: time.Now()},
		{CustomerID: 2, RequestPath: "/api/v1/get_phone", CreatedAt: time.Now()},
		{CustomerID: 1, RequestPath: "/api/v1/get_phone", CreatedAt: time.Now().Add(-1 * time.Hour)},
	}

	for _, log := range logs {
		db.Create(log)
	}

	// Test finding logs by customer ID and path
	result, total, err := repo.FindByCustomerIDAndPath(ctx, 1, "/api/v1/get_phone", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, result, 2)

	// Test with different customer
	result, total, err = repo.FindByCustomerIDAndPath(ctx, 2, "/api/v1/get_phone", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, result, 1)
}

func TestLogRepository_DeleteOldLogs(t *testing.T) {
	db := setupLogTestDB()
	repo := NewLogRepository(db)
	ctx := context.Background()

	now := time.Now()
	cutoffDate := now.Add(-7 * 24 * time.Hour) // 7 days ago

	// Create test data
	logs := []*domain.APILog{
		{CustomerID: 1, RequestPath: "/api/v1/old1", CreatedAt: cutoffDate.Add(-1 * time.Hour)}, // old
		{CustomerID: 1, RequestPath: "/api/v1/old2", CreatedAt: cutoffDate.Add(-2 * time.Hour)}, // old
		{CustomerID: 1, RequestPath: "/api/v1/new1", CreatedAt: cutoffDate.Add(1 * time.Hour)},  // new
		{CustomerID: 1, RequestPath: "/api/v1/new2", CreatedAt: now},                            // new
	}

	for _, log := range logs {
		db.Create(log)
	}

	// Delete old logs
	deleted, err := repo.DeleteOldLogs(ctx, cutoffDate)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), deleted)

	// Verify remaining logs
	remaining, total, err := repo.FindByCustomerID(ctx, 1, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, remaining, 2)

	// Verify the remaining logs are the new ones
	for _, log := range remaining {
		assert.True(t, log.CreatedAt.After(cutoffDate))
	}
}
