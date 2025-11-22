package service

import (
	"context"
	"errors"
	"net"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/repository"
)

var (
	ErrInvalidIPFormat   = errors.New("invalid IP address format")
	ErrWhitelistNotFound = errors.New("whitelist entry not found")
)

// WhitelistService defines the interface for IP whitelist business logic
type WhitelistService interface {
	AddWhitelist(ctx context.Context, customerID int64, ipAddress, notes string) error
	DeleteWhitelist(ctx context.Context, customerID int64, ipAddress string) error
	ListWhitelists(ctx context.Context, customerID int64, page, limit int) ([]*domain.IPWhitelist, int64, error)
	GetWhitelist(ctx context.Context, customerID int64, ipAddress string) (*domain.IPWhitelist, error)
}

type whitelistService struct {
	repo repository.WhitelistRepository
}

// NewWhitelistService creates a new WhitelistService
func NewWhitelistService(repo repository.WhitelistRepository) WhitelistService {
	return &whitelistService{repo: repo}
}

// AddWhitelist adds a new IP whitelist entry
func (s *whitelistService) AddWhitelist(ctx context.Context, customerID int64, ipAddress, notes string) error {
	if !isValidIPOrCIDR(ipAddress) {
		return ErrInvalidIPFormat
	}
	entry := &domain.IPWhitelist{
		CustomerID: customerID,
		IPAddress:  ipAddress,
		Notes:      notes,
	}
	return s.repo.Create(ctx, entry)
}

// DeleteWhitelist deletes an IP whitelist entry
func (s *whitelistService) DeleteWhitelist(ctx context.Context, customerID int64, ipAddress string) error {
	return s.repo.DeleteByCustomerIDAndIP(ctx, customerID, ipAddress)
}

// ListWhitelists lists all IP whitelist entries for a customer
func (s *whitelistService) ListWhitelists(ctx context.Context, customerID int64, page, limit int) ([]*domain.IPWhitelist, int64, error) {
	offset := (page - 1) * limit
	return s.repo.FindByCustomerID(ctx, customerID, limit, offset)
}

// GetWhitelist gets a specific IP whitelist entry
func (s *whitelistService) GetWhitelist(ctx context.Context, customerID int64, ipAddress string) (*domain.IPWhitelist, error) {
	return s.repo.FindByCustomerIDAndIP(ctx, customerID, ipAddress)
}

// isValidIPOrCIDR checks if the input is a valid IP or CIDR
func isValidIPOrCIDR(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	_, _, err := net.ParseCIDR(ip)
	return err == nil
}
