package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sms-platform/goapi/internal/api/handler/testutils" // Import common test utils
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPhoneService for testing
type MockPhoneService struct {
	mock.Mock
}

func (m *MockPhoneService) GetPhone(ctx context.Context, customerID uint, businessType, cardType string) (*service.GetPhoneResult, error) {
	args := m.Called(ctx, customerID, businessType, cardType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.GetPhoneResult), args.Error(1)
}

func (m *MockPhoneService) GetCode(ctx context.Context, customerID uint, phoneNumber string, timeout time.Duration) (*service.GetCodeResult, error) {
	args := m.Called(ctx, customerID, phoneNumber, timeout)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.GetCodeResult), args.Error(1)
}

/*
// Moved to testutils.SetupTestRouter() to avoid redeclaration.
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}
*/

func TestPhoneHandler_GetPhone_Success(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Setup mock expectations
	expectedResult := &service.GetPhoneResult{
		PhoneNumber: "+15551234567",
		CountryCode: "US",
		Cost:        0.10,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  "test-provider",
		Balance:     9.90,
	}

	mockService.On("GetPhone", mock.Anything, uint(1), "verification", "any").Return(expectedResult, nil) // customerID is uint

	// Setup route
	router.POST("/api/v1/get_phone", func(c *gin.Context) {
		c.Set("customer_id", int64(1)) // customerID is int64
		handler.GetPhone(c)
	})

	// Create request
	reqBody := dto.GetPhoneRequest{
		BusinessType: "verification",
		CardType:     "any",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/v1/get_phone", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeSuccess, response.Code)

	mockService.AssertExpectations(t)
}

func TestPhoneHandler_GetCode_Success(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Setup mock expectations
	expectedResult := &service.GetCodeResult{
		Code:       "123456",
		Message:    "Your verification code is 123456",
		ReceivedAt: time.Now(),
		ProviderID: "test-provider",
	}

	mockService.On("GetCode", mock.Anything, uint(1), "+15551234567", 60*time.Second).Return(expectedResult, nil) // customerID is uint

	// Setup route
	router.POST("/api/v1/get_code", func(c *gin.Context) {
		c.Set("customer_id", int64(1)) // customerID is int64
		handler.GetCode(c)
	})

	// Create request
	reqBody := dto.GetCodeRequest{
		PhoneNumber: "+15551234567",
		Timeout:     60,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/v1/get_code", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeSuccess, response.Code)

	mockService.AssertExpectations(t)
}

func TestPhoneHandler_GetPhoneStatus_Success(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Setup route
	router.GET("/api/v1/phone_status", func(c *gin.Context) {
		c.Set("customer_id", int64(1)) // customerID is int64
		handler.GetPhoneStatus(c)
	})

	// Test
	req, _ := http.NewRequest("GET", "/api/v1/phone_status?phone_number=+15551234567", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeSuccess, response.Code)
}
