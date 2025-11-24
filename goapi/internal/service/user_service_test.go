package service

import (
	"context"
	"errors"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/domain"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Define test JWT configuration for consistent testing
var testJWTConfig = config.JWTConfig{
	Secret:            "test_super_secret_key",
	ClientTokenExpiry: 259200, // 72 hours
	APITokenExpiry:    1800,   // 30 minutes
}

// Helper function for creating string pointers
func stringPtr(s string) *string {
	return &s
}

// MockCustomerRepository is a mock implementation of the CustomerRepository interface.
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) FindByUsername(ctx context.Context, username string) (*domain.Customer, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerRepository) FindByID(ctx context.Context, id int64) (*domain.Customer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerRepository) FindByAPISecretKey(ctx context.Context, apiSecretKey string) (*domain.Customer, error) { // New
	args := m.Called(ctx, apiSecretKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerRepository) FindByMerchantNoAndAPISecret(ctx context.Context, merchantNo, apiSecretKey string) (*domain.Customer, error) {
	args := m.Called(ctx, merchantNo, apiSecretKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	userService := NewUserService(mockRepo, testJWTConfig)
	ctx := context.Background()

	username := "testuser"
	email := "test@example.com"
	password := "password123"

	// Mock FindByUsername to return nil (user doesn't exist)
	mockRepo.On("FindByUsername", ctx, username).Return(nil, nil).Once()
	// We use mock.Anything for the customer argument because the password hash and secret key are generated inside the service
	mockRepo.On("Create", ctx, mock.AnythingOfType("*domain.Customer")).Return(nil).Once()

	customer, err := userService.Register(ctx, username, email, password)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, username, *customer.Username)
	assert.Equal(t, email, *customer.Email)
	assert.NotEmpty(t, customer.PasswordHash)
	err = bcrypt.CompareHashAndPassword([]byte(*customer.PasswordHash), []byte(password))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	userService := NewUserService(mockRepo, testJWTConfig)
	ctx := context.Background()

	username := "testuser"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockCustomer := &domain.Customer{
		ID:           1,
		Username:     &username,
		PasswordHash: stringPtr(string(hashedPassword)),
	}

	mockRepo.On("FindByUsername", ctx, username).Return(mockCustomer, nil)

	token, err := userService.Login(ctx, username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	userService := NewUserService(mockRepo, testJWTConfig)
	ctx := context.Background()

	username := "testuser"
	password := "password123"
	wrongPassword := "wrongpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockCustomer := &domain.Customer{
		ID:           1,
		Username:     &username,
		PasswordHash: stringPtr(string(hashedPassword)),
	}

	mockRepo.On("FindByUsername", ctx, username).Return(mockCustomer, nil)

	token, err := userService.Login(ctx, username, wrongPassword)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, ErrInvalidCredentials, err)

	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	userService := NewUserService(mockRepo, testJWTConfig)
	ctx := context.Background()

	username := "nonexistent"
	password := "password123"

	mockRepo.On("FindByUsername", ctx, username).Return(nil, errors.New("not found"))

	token, err := userService.Login(ctx, username, password)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, ErrUserNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestGenerateAPIToken_Success(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	userService := NewUserService(mockRepo, testJWTConfig)
	ctx := context.Background()

	apiSecret := "test_api_secret"
	merchantNo := "123456"
	mockCustomer := &domain.Customer{
		ID:           1,
		MerchantNo:   &merchantNo,
		APISecretKey: apiSecret,
	}

	mockRepo.On("FindByMerchantNoAndAPISecret", ctx, merchantNo, apiSecret).Return(mockCustomer, nil)

	token, err := userService.GenerateAPIToken(ctx, merchantNo, apiSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token claims manually for jwt/v5
	parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(testJWTConfig.Secret), nil
	})
	claims := parsedToken.Claims.(jwt.MapClaims)
	assert.Equal(t, float64(mockCustomer.ID), claims["user_id"])
	assert.Equal(t, "api", claims["aud"])

	// Manually verify expiration (within a small tolerance)
	exp := int64(claims["exp"].(float64))
	assert.True(t, exp > time.Now().Unix())
	assert.True(t, exp <= time.Now().Add(time.Minute*31).Unix()) // Allow for a small buffer

	mockRepo.AssertExpectations(t)
}

func TestGenerateAPIToken_InvalidSecret(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	userService := NewUserService(mockRepo, testJWTConfig)
	ctx := context.Background()

	apiSecret := "invalid_secret"
	merchantNo := "123456"

	mockRepo.On("FindByMerchantNoAndAPISecret", ctx, merchantNo, apiSecret).Return(nil, errors.New("not found"))

	token, err := userService.GenerateAPIToken(ctx, merchantNo, apiSecret)
	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, ErrInvalidAPISecret, err)
	mockRepo.AssertExpectations(t)
}
