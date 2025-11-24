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

	name := "TestProvider"
	apiConfig := `{"url":"http://test.com"}`
	isEnabled := true
	provider := &domain.Provider{
		Name:      &name,
		APIConfig: &apiConfig,
		IsEnabled: &isEnabled,
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

	name := "TestProvider2"
	apiConfig := `{"url":"http://test2.com"}`
	isEnabled := true
	provider := &domain.Provider{
		Name:      &name,
		APIConfig: &apiConfig,
		IsEnabled: &isEnabled,
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

	name1 := "P1"
	isEnabled1 := true
	name2 := "P2"
	isEnabled2 := false

	provider1 := &domain.Provider{Name: &name1, IsEnabled: &isEnabled1}
	provider2 := &domain.Provider{Name: &name2, IsEnabled: &isEnabled2}
	repo.Create(ctx, provider1)
	repo.Create(ctx, provider2)

	providers, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, providers, 2)

	// Compare by checking IDs and names instead of full struct equality
	var foundNames []string
	var foundIDs []int
	for _, p := range providers {
		foundNames = append(foundNames, *p.Name)
		foundIDs = append(foundIDs, p.ID)
	}

	assert.Contains(t, foundNames, "P1")
	assert.Contains(t, foundNames, "P2")
	assert.Contains(t, foundIDs, provider1.ID)
	assert.Contains(t, foundIDs, provider2.ID)
	assert.True(t, *providers[0].IsEnabled || *providers[1].IsEnabled)   // P1 enabled
	assert.True(t, !*providers[0].IsEnabled || !*providers[1].IsEnabled) // P2 disabled
}

func TestProviderRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	oldName := "OldName"
	oldEnabled := true
	provider := &domain.Provider{Name: &oldName, IsEnabled: &oldEnabled}
	repo.Create(ctx, provider)

	newName := "NewName"
	newEnabled := false
	provider.Name = &newName
	provider.IsEnabled = &newEnabled
	err := repo.Update(ctx, provider)
	assert.NoError(t, err)

	foundProvider, _ := repo.FindByID(ctx, provider.ID)
	assert.Equal(t, "NewName", *foundProvider.Name)
	assert.False(t, *foundProvider.IsEnabled)
}

func TestProviderRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)
	ctx := context.Background()

	name := "To_Delete"
	isEnabled := true
	provider := &domain.Provider{Name: &name, IsEnabled: &isEnabled}
	repo.Create(ctx, provider)

	err := repo.Delete(ctx, provider.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(ctx, provider.ID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
