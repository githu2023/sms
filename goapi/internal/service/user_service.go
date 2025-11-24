package service

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	GenerateAPIToken(ctx context.Context, merchantNo, apiSecretKey string) (string, error)
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
	// Check if user already exists
	existingUser, _ := s.repo.FindByUsername(ctx, username)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Generate random API secret key (64 characters)
	apiSecretKey := utils.GenerateRandomString(64)
	passwordHashStr := utils.BcryptHash(password)
	status := true // Active
	merchanNo := utils.GenerateRandomMerchantNo()
	parentID := int64(0)

	customer := &domain.Customer{
		MerchantName: &username,  // 默认商户名是用户名
		MerchantNo:   &merchanNo, // 随机6位数字
		Username:     &username,
		Email:        &email,
		PasswordHash: &passwordHashStr,
		APISecretKey: apiSecretKey, // 随机64位字符串
		ParentID:     &parentID,    // 默认为0
		Status:       &status,
	}

	err := s.repo.Create(ctx, customer)
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

	if customer.PasswordHash == nil {
		return "", ErrInvalidCredentials
	}

	if !utils.BcryptCheck(password, *customer.PasswordHash) {
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

// GenerateAPIToken generates a JWT for API access based on merchant number and API secret key.
func (s *userService) GenerateAPIToken(ctx context.Context, merchantNo, apiSecretKey string) (string, error) {
	customer, err := s.repo.FindByMerchantNoAndAPISecret(ctx, merchantNo, apiSecretKey)
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
	if customer.PasswordHash == nil {
		return ErrInvalidCredentials
	}

	if !utils.BcryptCheck(oldPassword, *customer.PasswordHash) {
		return ErrInvalidCredentials
	}

	// Hash new password
	passwordHashStr := utils.BcryptHash(newPassword)
	// Update password hash
	customer.PasswordHash = &passwordHashStr
	if err := s.repo.Update(ctx, customer); err != nil {
		return ErrPasswordUpdateFailed
	}

	return nil
}
