package repository

import (
	"context"
	"sms-platform/goapi/internal/domain"

	"gorm.io/gorm"
)

type providerRepository struct {
	db *gorm.DB
}

// NewProviderRepository creates a new provider repository.
func NewProviderRepository(db *gorm.DB) domain.ProviderRepository {
	return &providerRepository{db: db}
}

func (r *providerRepository) Create(ctx context.Context, provider *domain.Provider) error {
	return r.db.WithContext(ctx).Create(provider).Error
}

func (r *providerRepository) FindByID(ctx context.Context, id int) (*domain.Provider, error) {
	var provider domain.Provider
	if err := r.db.WithContext(ctx).First(&provider, id).Error; err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *providerRepository) FindAll(ctx context.Context) ([]*domain.Provider, error) {
	var providers []*domain.Provider
	if err := r.db.WithContext(ctx).Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *providerRepository) Update(ctx context.Context, provider *domain.Provider) error {
	return r.db.WithContext(ctx).Save(provider).Error
}

func (r *providerRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.Provider{}, id).Error
}
