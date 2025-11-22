package repository

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&domain.Provider{})
	assert.NoError(t, err)
	return db
}

func TestProviderRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	provider := &domain.Provider{
		Name:      "TestProvider",
		APIConfig: `{"url":"http://test.com"}`,
		IsEnabled: true,
	}

	err := repo.Create(ctx, provider)
	assert.NoError(t, err)
	assert.NotZero(t, provider.ID)

	foundProvider, err := repo.FindByID(ctx, provider.ID)
	assert.NoError(t, err)
	assert.Equal(t, provider.Name, foundProvider.Name)
}

func TestProviderRepository_FindByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	provider := &domain.Provider{
		Name:      "TestProvider2",
		APIConfig: `{"url":"http://test2.com"}`,
		IsEnabled: true,
	}
	repo.Create(ctx, provider)

	foundProvider, err := repo.FindByID(ctx, provider.ID)
	assert.NoError(t, err)
	assert.Equal(t, provider.ID, foundProvider.ID)
	assert.Equal(t, provider.Name, foundProvider.Name)

	_, err = repo.FindByID(ctx, 999) // Non-existent ID
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestProviderRepository_FindAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	provider1 := &domain.Provider{Name: "P1", IsEnabled: true}
	provider2 := &domain.Provider{Name: "P2", IsEnabled: false}
	repo.Create(ctx, provider1)
	repo.Create(ctx, provider2)

	providers, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, providers, 2)
	assert.ElementsMatch(t, []*domain.Provider{provider1, provider2}, providers)
}

func TestProviderRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	provider := &domain.Provider{Name: "OldName", IsEnabled: true}
	repo.Create(ctx, provider)

	provider.Name = "NewName"
	provider.IsEnabled = false
	err := repo.Update(ctx, provider)
	assert.NoError(t, err)

	foundProvider, _ := repo.FindByID(ctx, provider.ID)
	assert.Equal(t, "NewName", foundProvider.Name)
	assert.False(t, foundProvider.IsEnabled)
}

func TestProviderRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	provider := &domain.Provider{Name: "To_Delete", IsEnabled: true}
	repo.Create(ctx, provider)

	err := repo.Delete(ctx, provider.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(ctx, provider.ID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
