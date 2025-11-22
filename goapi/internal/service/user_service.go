package service

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrInvalidAPISecret     = errors.New("invalid API secret")
	ErrPasswordUpdateFailed = errors.New("failed to update password")
)

// UserService defines the interface for user related business logic.
type UserService interface {
	Register(ctx context.Context, username, email, password string) (*domain.Customer, error)
	Login(ctx context.Context, username, password string) (string, error)
	GetProfile(ctx context.Context, userID int64) (*domain.Customer, error)
	GenerateAPIToken(ctx context.Context, apiSecretKey string) (string, error)
	UpdatePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error
}

type userService struct {
	repo      domain.CustomerRepository
	jwtConfig config.JWTConfig
}

// NewUserService creates a new user service.
func NewUserService(repo domain.CustomerRepository, jwtConfig config.JWTConfig) UserService {
	return &userService{repo: repo, jwtConfig: jwtConfig}
}

// Register creates a new user.
func (s *userService) Register(ctx context.Context, username, email, password string) (*domain.Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Simple secret key generation for now
	apiSecretKey := "secret_" + username

	customer := &domain.Customer{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		APISecretKey: apiSecretKey,
		Status:       1, // Active
	}

	err = s.repo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// Login authenticates a user and returns a JWT.
func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	customer, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": customer.ID,
		"exp":     time.Now().Add(time.Duration(s.jwtConfig.ClientTokenExpiry) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetProfile retrieves a user's profile.
func (s *userService) GetProfile(ctx context.Context, userID int64) (*domain.Customer, error) {
	return s.repo.FindByID(ctx, userID)
}

// GenerateAPIToken generates a JWT for API access based on the API secret key.
func (s *userService) GenerateAPIToken(ctx context.Context, apiSecretKey string) (string, error) {
	customer, err := s.repo.FindByAPISecretKey(ctx, apiSecretKey)
	if err != nil {
		return "", ErrInvalidAPISecret // Specific error for invalid secret
	}

	// Create JWT token for API usage
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": customer.ID,
		"exp":     time.Now().Add(time.Duration(s.jwtConfig.APITokenExpiry) * time.Second).Unix(),
		"aud":     "api", // Audience claim to differentiate from client tokens
	})

	tokenString, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// UpdatePassword updates a user's password after verifying the old password.
func (s *userService) UpdatePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	// Get user by ID
	customer, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return ErrUserNotFound
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(oldPassword)); err != nil {
		return ErrInvalidCredentials
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return ErrPasswordUpdateFailed
	}

	// Update password hash
	customer.PasswordHash = string(hashedPassword)
	if err := s.repo.Update(ctx, customer); err != nil {
		return ErrPasswordUpdateFailed
	}

	return nil
}
