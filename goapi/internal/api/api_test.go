package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sms-platform/goapi/internal/api/handler"
	"sms-platform/goapi/internal/api/middleware"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5" // Add this import
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Define a test JWT secret for consistent testing
const testJWTSecret = "test_jwt_secret"

// 测试用的JWT配置
var testJWTConfig = config.JWTConfig{
	Secret:            testJWTSecret,
	ClientTokenExpiry: 259200, // 72 hours
	APITokenExpiry:    1800,   // 30 minutes
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

func (m *MockCustomerRepository) FindByAPISecretKey(ctx context.Context, apiSecretKey string) (*domain.Customer, error) {
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

// MockBusinessTypeRepository is a mock implementation of the BusinessTypeRepository interface.
type MockBusinessTypeRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeRepository) Create(ctx context.Context, businessType *domain.BusinessType) error {
	args := m.Called(ctx, businessType)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) FindByCode(ctx context.Context, code string) (*domain.BusinessType, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindByID(ctx context.Context, id int) (*domain.BusinessType, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindAll(ctx context.Context) ([]*domain.BusinessType, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) Update(ctx context.Context, businessType *domain.BusinessType) error {
	args := m.Called(ctx, businessType)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockCustomerBusinessConfigRepository is a mock implementation of the CustomerBusinessConfigRepository interface.
type MockCustomerBusinessConfigRepository struct {
	mock.Mock
}

func (m *MockCustomerBusinessConfigRepository) Create(ctx context.Context, config *domain.CustomerBusinessConfig) error {
	args := m.Called(ctx, config)
	return args.Error(0)
}

func (m *MockCustomerBusinessConfigRepository) FindByCustomerIDAndBusinessCode(ctx context.Context, customerID int64, businessCode string) (*domain.CustomerBusinessConfig, error) {
	args := m.Called(ctx, customerID, businessCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.CustomerBusinessConfig), args.Error(1)
}

func (m *MockCustomerBusinessConfigRepository) FindByCustomerID(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.CustomerBusinessConfig), args.Error(1)
}

func (m *MockCustomerBusinessConfigRepository) Update(ctx context.Context, config *domain.CustomerBusinessConfig) error {
	args := m.Called(ctx, config)
	return args.Error(0)
}

func (m *MockCustomerBusinessConfigRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCustomerBusinessConfigRepository) FindByCustomerIDAndEnabled(ctx context.Context, customerID int64) ([]*domain.CustomerBusinessConfig, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.CustomerBusinessConfig), args.Error(1)
}

// SetupTestRouter sets up a Gin router for testing.
func SetupTestRouter(mockCustomerRepo *MockCustomerRepository, mockBusinessTypeRepo *MockBusinessTypeRepository) (*gin.Engine, *MockCustomerBusinessConfigRepository) {
	gin.SetMode(gin.TestMode) // Set Gin mode for testing

	// Create mock CustomerBusinessConfigRepository
	mockCustomerBusinessConfigRepo := &MockCustomerBusinessConfigRepository{}

	// Initialize Services with mock repositories
	userService := service.NewUserService(mockCustomerRepo, testJWTConfig)
	businessService := service.NewBusinessService(mockBusinessTypeRepo, mockCustomerBusinessConfigRepo)

	// Initialize Handlers
	userHandler := handler.NewUserHandler(userService, testJWTConfig)
	businessHandler := handler.NewBusinessHandler(businessService)
	authHandler := handler.NewAuthHandler(userService, testJWTConfig) // New AuthHandler

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())

	// Client API Group
	clientV1 := router.Group("/client/v1")
	{
		clientV1.POST("/register", userHandler.Register)
		clientV1.POST("/login", userHandler.Login)
		clientV1.GET("/business_types", businessHandler.GetBusinessTypes)

		clientAuth := clientV1.Group("/")
		clientAuth.Use(middleware.JWTAuthMiddleware(testJWTSecret))
		{
			clientAuth.GET("/profile", userHandler.GetProfile)
		}
	}

	// Programmatic API Group
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/get_token", authHandler.GetAPIToken) // New

		// Authenticated routes for apiV1
		apiAuth := apiV1.Group("/")
		apiAuth.Use(middleware.APITokenAuthMiddleware(testJWTSecret))
		{
			apiAuth.GET("/business_types", businessHandler.GetBusinessTypes)
		}
	}

	return router, mockCustomerBusinessConfigRepo
}

func TestRegisterHandler(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository) // Not used, but passed to setup
	router, _ := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "password123",
	})
	req, _ := http.NewRequest("POST", "/client/v1/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Mock repository behavior
	// First check if username exists (should return nil, nil for new user)
	mockCustomerRepo.On("FindByUsername", mock.Anything, "testuser").Return(nil, nil).Once()
	// Then create the new user
	mockCustomerRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(nil).Once()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp common.APIResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeSuccess, resp.Code)
	assert.NotNil(t, resp.Data)
	data := resp.Data.(map[string]interface{})
	assert.NotNil(t, data["user_id"])
	assert.Equal(t, "testuser", data["username"])
	mockCustomerRepo.AssertExpectations(t)
}

func TestLoginHandler_Success(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository) // Not used, but passed to setup
	router, _ := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo)

	username := "newuser"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	mockCustomer := &domain.Customer{
		ID:           2,
		Username:     &username,
		PasswordHash: &hashedPasswordStr,
	}

	reqBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	req, _ := http.NewRequest("POST", "/client/v1/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockCustomerRepo.On("FindByUsername", mock.Anything, username).Return(mockCustomer, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp common.APIResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeSuccess, resp.Code)
	assert.NotNil(t, resp.Data)
	data := resp.Data.(map[string]interface{})
	assert.NotNil(t, data["token"])
	mockCustomerRepo.AssertExpectations(t)
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository) // Not used, but passed to setup
	router, _ := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo)

	username := "testuser"
	password := "password123"
	wrongPassword := "wrongpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	mockCustomer := &domain.Customer{
		ID:           1,
		Username:     &username,
		PasswordHash: &hashedPasswordStr,
	}

	reqBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": wrongPassword,
	})
	req, _ := http.NewRequest("POST", "/client/v1/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockCustomerRepo.On("FindByUsername", mock.Anything, username).Return(mockCustomer, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp common.APIResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeInvalidCredentials, resp.Code)
	assert.Equal(t, "用户名或密码错误", resp.Msg)
	mockCustomerRepo.AssertExpectations(t)
}

func TestGetProfileHandler_Success(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository)              // Not used, but passed to setup
	router, _ := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo) // Setup router once

	userID := int64(1)
	username := "profileuser"
	email := "profile@example.com"
	balance := 100.50
	apiSecret := "secret_profile"
	regIP := "192.168.1.1"
	mockCustomer := &domain.Customer{
		ID:             userID,
		Username:       &username,
		Email:          &email,
		Balance:        balance,
		APISecretKey:   apiSecret,
		RegistrationIP: &regIP,
		LastLoginAt:    &time.Time{},
	}
	// This mock is for userService.GetProfile called by the handler after successful authentication
	mockCustomerRepo.On("FindByID", mock.Anything, userID).Return(mockCustomer, nil).Once()

	// Manually create a valid JWT token for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID, // Explicitly set as int64, not float64
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	signedToken, err := token.SignedString([]byte(testJWTSecret))
	assert.NoError(t, err)
	assert.NotEmpty(t, signedToken)

	req, _ := http.NewRequest("GET", "/client/v1/profile", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signedToken))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp common.APIResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeSuccess, resp.Code)
	assert.NotNil(t, resp.Data)
	data := resp.Data.(map[string]interface{})
	assert.Equal(t, float64(userID), data["user_id"])
	assert.Equal(t, "profileuser", data["username"])
	mockCustomerRepo.AssertExpectations(t)
}

func TestGetProfileHandler_Unauthorized(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository) // Not used, but passed to setup
	router, _ := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo)

	req, _ := http.NewRequest("GET", "/client/v1/profile", nil)
	// No Authorization header
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var resp common.APIResponse // Expect the standard error response
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeUnauthorized, resp.Code) // Expect the code from middleware
	assert.Equal(t, "未授权，token错误或过期", resp.Msg)         // Expect the message from middleware
	mockCustomerRepo.AssertNotCalled(t, "FindByID")
}

func TestGetBusinessTypesHandler_ProgrammaticAPI_Unauthorized(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	router, _ := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo)

	req, _ := http.NewRequest("GET", "/api/v1/business_types", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var resp common.APIResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeUnauthorized, resp.Code) // Assuming middleware returns this code for missing token
	assert.Equal(t, "未授权，token错误或过期", resp.Msg)
}

func TestGetBusinessTypesHandler_ProgrammaticAPI_Authenticated(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	router, mockCustomerBusinessConfigRepo := SetupTestRouter(mockCustomerRepo, mockBusinessTypeRepo)

	// --- Step 1: Mock data and behavior for API Token generation ---
	apiSecretKey := "secret_apitestuser"
	testUsername := "apitestuser"
	merchantNo := "123456"
	mockCustomer := &domain.Customer{
		ID:           101,
		Username:     &testUsername,
		MerchantNo:   &merchantNo,
		APISecretKey: apiSecretKey,
	}

	// Mock FindByMerchantNoAndAPISecret for GenerateAPIToken
	// merchantNo is already defined above
	mockCustomerRepo.On("FindByMerchantNoAndAPISecret", mock.Anything, merchantNo, apiSecretKey).Return(mockCustomer, nil).Once()

	// We create a separate userService instance for generating the token,
	// so it uses the mockRepo and the testJWTSecret.
	tokenUserService := service.NewUserService(mockCustomerRepo, testJWTConfig)

	// Request to get API Token
	tokenReqBody, _ := json.Marshal(dto.GetAPITokenRequest{Secret: apiSecretKey})
	tokenReq, _ := http.NewRequest("POST", "/api/v1/get_token", bytes.NewBuffer(tokenReqBody))
	tokenReq.Header.Set("Content-Type", "application/json")

	// Directly call the authHandler GetAPIToken method with a test context
	// to get the token without involving the router's middleware.
	// This avoids potential issues with router's JWT secret setup.
	// Alternatively, we could run the router and make an actual request to /api/v1/get_token.
	// For simplicity in testing a specific component (userService token generation),
	// directly calling the service is often better.
	generatedToken, err := tokenUserService.GenerateAPIToken(context.Background(), merchantNo, apiSecretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, generatedToken)

	apiToken := generatedToken

	// --- Step 2: Mock data and behavior for Business Types API access ---
	// GetBusinessTypes returns CustomerBusinessConfig, not BusinessType
	mockCustomerBusinessConfigs := []*domain.CustomerBusinessConfig{
		{
			ID:           1,
			BusinessCode: "api_qq",
			BusinessName: "API_QQ",
			Weight:       50,
		},
		{
			ID:           2,
			BusinessCode: "api_wechat",
			BusinessName: "API_WeChat",
			Weight:       50,
		},
	}

	// Mock FindByCustomerIDAndEnabled for businessService.GetBusinessTypesForCustomer
	mockCustomerBusinessConfigRepo.On("FindByCustomerIDAndEnabled", mock.Anything, int64(101)).Return(mockCustomerBusinessConfigs, nil).Once()

	// Request to access protected business types endpoint
	req, _ := http.NewRequest("GET", "/api/v1/business_types", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp common.APIResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, common.CodeSuccess, resp.Code)
	assert.NotNil(t, resp.Data)

	data := resp.Data.([]interface{})
	assert.Len(t, data, 2)
	assert.Equal(t, "API_QQ", data[0].(map[string]interface{})["business_name"])
	assert.Equal(t, "api_qq", data[0].(map[string]interface{})["business_code"])

	mockCustomerRepo.AssertExpectations(t)
	mockBusinessTypeRepo.AssertExpectations(t)
}
