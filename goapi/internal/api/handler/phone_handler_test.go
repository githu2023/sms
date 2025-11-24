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

func (m *MockPhoneService) GetPhone(ctx context.Context, customerID int64, businessType, cardType string, count int) ([]*service.GetPhoneResult, common.ErrorCode) {
	args := m.Called(ctx, customerID, businessType, cardType, count)
	if args.Get(0) == nil {
		return nil, args.Get(1).(common.ErrorCode)
	}
	return args.Get(0).([]*service.GetPhoneResult), args.Get(1).(common.ErrorCode)
}

func (m *MockPhoneService) GetCode(ctx context.Context, customerID int64, phoneNumber string) ([]*service.GetCodeResult, error) {
	args := m.Called(ctx, customerID, phoneNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*service.GetCodeResult), args.Error(1)
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
	expectedResult := []*service.GetPhoneResult{
		{
			PhoneNumber: "+15551234567",
			CountryCode: "US",
			Cost:        0.10,
			ValidUntil:  time.Now().Add(30 * time.Minute),
			ProviderID:  "test-provider",
			Balance:     9.90,
		},
	}

	mockService.On("GetPhone", mock.Anything, int64(1), "verification", "any", 1).Return(expectedResult, common.CodeSuccess)

	// Setup route
	router.POST("/api/v1/get_phone", func(c *gin.Context) {
		c.Set("customer_id", int64(1)) // customerID is int64
		handler.GetPhone(c)
	})

	// Create request
	reqBody := dto.GetPhoneRequest{
		BusinessType: "verification",
		CardType:     "any",
		Count:        1, // 新增的count字段
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

	// Verify batch response format
	respData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, respData, "phones")
	assert.Contains(t, respData, "total_cost")
	assert.Contains(t, respData, "success_count")
	assert.Contains(t, respData, "failed_count")

	// Verify phone data
	phones, ok := respData["phones"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, phones, 1) // 应该有一个手机号

	mockService.AssertExpectations(t)
}

func TestPhoneHandler_GetCode_Success(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Setup mock expectations
	expectedResult := []*service.GetCodeResult{
		{
			Code:       "123456",
			Message:    "验证码获取成功",
			ReceivedAt: time.Now(),
			ProviderID: "test-provider",
		},
	}

	mockService.On("GetCode", mock.Anything, int64(1), "+15551234567").Return(expectedResult, nil)

	// Setup route
	router.POST("/api/v1/get_code", func(c *gin.Context) {
		c.Set("customer_id", int64(1)) // customerID is int64
		handler.GetCode(c)
	})

	// Create request
	reqBody := dto.GetCodeRequest{
		PhoneNumbers: []string{"+15551234567"}, // 改为数组格式
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

	// Verify batch response format
	respData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, respData, "codes")
	assert.Contains(t, respData, "success_count")
	assert.Contains(t, respData, "pending_count")
	assert.Contains(t, respData, "failed_count")

	// Verify code data
	codes, ok := respData["codes"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, codes, 1) // 应该有一个验证码

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

func TestPhoneHandler_GetPhone_BatchRequest(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup mock expectations for multiple calls
	expectedResult := []*service.GetPhoneResult{
		{
			PhoneNumber: "+15551234567",
			CountryCode: "US",
			Cost:        0.10,
			ValidUntil:  time.Now().Add(30 * time.Minute),
			ProviderID:  "test-provider",
			Balance:     9.70, // 余额会减少
		},
		{
			PhoneNumber: "+15551234568",
			CountryCode: "US",
			Cost:        0.10,
			ValidUntil:  time.Now().Add(30 * time.Minute),
			ProviderID:  "test-provider",
			Balance:     9.60,
		},
		{
			PhoneNumber: "+15551234569",
			CountryCode: "US",
			Cost:        0.10,
			ValidUntil:  time.Now().Add(30 * time.Minute),
			ProviderID:  "test-provider",
			Balance:     9.50,
		},
	}

	// Mock 一次调用返回3个结果
	mockService.On("GetPhone", mock.Anything, int64(1), "verification", "any", 3).Return(expectedResult, common.CodeSuccess)

	// Setup route
	router.POST("/api/v1/get_phone", func(c *gin.Context) {
		c.Set("customer_id", int64(1))
		handler.GetPhone(c)
	})

	// Create request for 3 phones
	reqBody := dto.GetPhoneRequest{
		BusinessType: "verification",
		CardType:     "any",
		Count:        3, // 批量获取3个
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

	// Verify batch response
	respData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(3), respData["success_count"])
	assert.Equal(t, float64(0), respData["failed_count"])
	assert.InDelta(t, 0.30, respData["total_cost"], 0.01) // 使用InDelta来处理浮点数精度问题

	phones, ok := respData["phones"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, phones, 3)

	mockService.AssertExpectations(t)
}

func TestPhoneHandler_GetCode_BatchRequest(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup mock expectations for multiple phone numbers
	expectedResult1 := []*service.GetCodeResult{
		{
			Code:       "123456",
			Message:    "验证码获取成功",
			ReceivedAt: time.Now(),
			ProviderID: "test-provider",
		},
	}
	expectedResult2 := []*service.GetCodeResult{
		{
			Code:       "234567",
			Message:    "验证码获取成功",
			ReceivedAt: time.Now(),
			ProviderID: "test-provider",
		},
	}
	expectedResult3 := []*service.GetCodeResult{
		{
			Code:       "345678",
			Message:    "验证码获取成功",
			ReceivedAt: time.Now(),
			ProviderID: "test-provider",
		},
	}

	mockService.On("GetCode", mock.Anything, int64(1), "+15551234567").Return(expectedResult1, nil)
	mockService.On("GetCode", mock.Anything, int64(1), "+15551234568").Return(expectedResult2, nil)
	mockService.On("GetCode", mock.Anything, int64(1), "+15551234569").Return(expectedResult3, nil)

	// Setup route
	router.POST("/api/v1/get_code", func(c *gin.Context) {
		c.Set("customer_id", int64(1))
		handler.GetCode(c)
	})

	// Create request for multiple phone numbers
	reqBody := dto.GetCodeRequest{
		PhoneNumbers: []string{"+15551234567", "+15551234568", "+15551234569"},
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

	// Verify batch response
	respData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(3), respData["success_count"])
	assert.Equal(t, float64(0), respData["pending_count"])
	assert.Equal(t, float64(0), respData["failed_count"])

	codes, ok := respData["codes"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, codes, 3)

	mockService.AssertExpectations(t)
}

func TestPhoneHandler_GetPhone_InvalidCount(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup route
	router.POST("/api/v1/get_phone", func(c *gin.Context) {
		c.Set("customer_id", int64(1))
		handler.GetPhone(c)
	})

	// Test with count = 0 (should default to 1)
	reqBody := dto.GetPhoneRequest{
		BusinessType: "verification",
		CardType:     "any",
		Count:        0,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/v1/get_phone", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Mock should be called once (default to 1)
	expectedResult := []*service.GetPhoneResult{
		{
			PhoneNumber: "+15551234567",
			CountryCode: "US",
			Cost:        0.10,
			ValidUntil:  time.Now().Add(30 * time.Minute),
			ProviderID:  "test-provider",
			Balance:     9.90,
		},
	}
	mockService.On("GetPhone", mock.Anything, int64(1), "verification", "any", 1).Return(expectedResult, common.CodeSuccess)

	// Test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should succeed with count defaulted to 1
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestPhoneHandler_GetCode_EmptyPhoneNumbers(t *testing.T) {
	// Setup
	mockService := &MockPhoneService{}
	handler := NewPhoneHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup route
	router.POST("/api/v1/get_code", func(c *gin.Context) {
		c.Set("customer_id", int64(1))
		handler.GetCode(c)
	})

	// Create request with empty phone numbers
	reqBody := dto.GetCodeRequest{
		PhoneNumbers: []string{}, // 空数组
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/v1/get_code", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return bad request
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
