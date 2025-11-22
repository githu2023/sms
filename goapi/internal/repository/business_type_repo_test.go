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

func setupTestDBForBusinessType(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&domain.BusinessType{})
	assert.NoError(t, err)
	return db
}

func TestBusinessTypeRepository_Create(t *testing.T) {
	db := setupTestDBForBusinessType(t)
	repo := NewBusinessTypeRepository(db)
	ctx := context.Background()

	bt := &domain.BusinessType{
		Name:      "QQ",
		Code:      "qq",
		IsEnabled: true,
	}

	err := repo.Create(ctx, bt)
	assert.NoError(t, err)
	assert.NotZero(t, bt.ID)

	foundBT, err := repo.FindByID(ctx, bt.ID)
	assert.NoError(t, err)
	assert.Equal(t, bt.Name, foundBT.Name)
	assert.Equal(t, bt.Code, foundBT.Code)
}

func TestBusinessTypeRepository_FindByCode(t *testing.T) {
	db := setupTestDBForBusinessType(t)
	repo := NewBusinessTypeRepository(db)
	ctx := context.Background()

	bt := &domain.BusinessType{
		Name:      "WeChat",
		Code:      "wechat",
		IsEnabled: true,
	}
	repo.Create(ctx, bt)

	foundBT, err := repo.FindByCode(ctx, "wechat")
	assert.NoError(t, err)
	assert.Equal(t, bt.ID, foundBT.ID)
	assert.Equal(t, bt.Name, foundBT.Name)

	_, err = repo.FindByCode(ctx, "nonexistent")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestBusinessTypeRepository_FindAll(t *testing.T) {
	db := setupTestDBForBusinessType(t)
	repo := NewBusinessTypeRepository(db)
	ctx := context.Background()

	bt1 := &domain.BusinessType{Name: "Alipay", Code: "alipay", IsEnabled: true}
	bt2 := &domain.BusinessType{Name: "Taobao", Code: "taobao", IsEnabled: false}
	repo.Create(ctx, bt1)
	repo.Create(ctx, bt2)

	bts, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, bts, 2)
	assert.ElementsMatch(t, []*domain.BusinessType{bt1, bt2}, bts)
}

func TestBusinessTypeRepository_Update(t *testing.T) {
	db := setupTestDBForBusinessType(t)
	repo := NewBusinessTypeRepository(db)
	ctx := context.Background()

	bt := &domain.BusinessType{Name: "OldName", Code: "old_code", IsEnabled: true}
	repo.Create(ctx, bt)

	bt.Name = "NewName"
	bt.IsEnabled = false
	err := repo.Update(ctx, bt)
	assert.NoError(t, err)

	foundBT, _ := repo.FindByID(ctx, bt.ID)
	assert.Equal(t, "NewName", foundBT.Name)
	assert.False(t, foundBT.IsEnabled)
}

func TestBusinessTypeRepository_Delete(t *testing.T) {
	db := setupTestDBForBusinessType(t)
	repo := NewBusinessTypeRepository(db)
	ctx := context.Background()

	bt := &domain.BusinessType{Name: "To_Delete", Code: "delete_code", IsEnabled: true}
	repo.Create(ctx, bt)

	err := repo.Delete(ctx, bt.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(ctx, bt.ID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
