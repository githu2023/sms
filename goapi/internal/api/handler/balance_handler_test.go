package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sms-platform/goapi/internal/api/handler/testutils"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransactionService for testing
type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) GetBalance(ctx context.Context, customerID int64) (float64, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockTransactionService) GetBalanceDetail(ctx context.Context, customerID int64) (*service.BalanceDetail, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.BalanceDetail), args.Error(1)
}

func (m *MockTransactionService) TopUp(ctx context.Context, customerID int64, amount float64, notes string) (*domain.Transaction, error) {
	args := m.Called(ctx, customerID, amount, notes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) Deduct(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	args := m.Called(ctx, customerID, amount, referenceID, notes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) ReserveFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	args := m.Called(ctx, customerID, amount, referenceID, notes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) CommitReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	args := m.Called(ctx, customerID, amount, referenceID, notes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) ReleaseReservedFunds(ctx context.Context, customerID int64, amount float64, referenceID int64, notes string) (*domain.Transaction, error) {
	args := m.Called(ctx, customerID, amount, referenceID, notes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) GetTransactionHistory(ctx context.Context, customerID int64, limit, offset int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *MockTransactionService) GetTransactionsByType(ctx context.Context, customerID int64, transactionType int, limit, offset int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, transactionType, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *MockTransactionService) GetTransactionsByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, limit, offset int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, startDate, endDate, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

func TestBalanceHandler_GetBalance_Success(t *testing.T) {
	// Setup
	mockService := &MockTransactionService{}
	handler := NewBalanceHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup mock expectations
	mockService.On("GetBalanceDetail", mock.Anything, int64(1)).Return(&service.BalanceDetail{
		Balance:      123.45,
		FrozenAmount: 10.0,
	}, nil)

	// Setup route
	router.GET("/api/v1/balance", func(c *gin.Context) {
		c.Set("customer_id", int64(1))
		handler.GetBalance(c)
	})

	// Test
	req, _ := http.NewRequest("GET", "/api/v1/balance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeSuccess, response.Code)

	// Check response data
	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, 123.45, data["balance"])
	assert.Equal(t, 10.0, data["frozen_amount"])
	assert.Equal(t, "USD", data["currency"])

	mockService.AssertExpectations(t)
}

func TestBalanceHandler_GetBalance_Unauthorized(t *testing.T) {
	// Setup
	mockService := &MockTransactionService{}
	handler := NewBalanceHandler(mockService)

	// Create router without the customer_id middleware
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Setup route without setting customer_id
	router.GET("/api/v1/balance", handler.GetBalance)

	// Test
	req, _ := http.NewRequest("GET", "/api/v1/balance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeUnauthorized, response.Code)
}

func TestBalanceHandler_GetBalance_InternalError(t *testing.T) {
	// Setup
	mockService := &MockTransactionService{}
	handler := NewBalanceHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup mock expectations
	mockService.On("GetBalanceDetail", mock.Anything, int64(1)).Return((*service.BalanceDetail)(nil), errors.New("database error"))

	// Setup route
	router.GET("/api/v1/balance", func(c *gin.Context) {
		c.Set("customer_id", int64(1))
		handler.GetBalance(c)
	})

	// Test
	req, _ := http.NewRequest("GET", "/api/v1/balance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeInternalError, response.Code)
	assert.Contains(t, response.Msg, "Failed to get balance")

	mockService.AssertExpectations(t)
}

func TestBalanceHandler_GetBalance_InvalidCustomerID(t *testing.T) {
	// Setup
	mockService := &MockTransactionService{}
	handler := NewBalanceHandler(mockService)
	router := testutils.SetupTestRouter()

	// Setup route with invalid customer_id type
	router.GET("/api/v1/balance", func(c *gin.Context) {
		c.Set("customer_id", "invalid") // String instead of uint
		handler.GetBalance(c)
	})

	// Test
	req, _ := http.NewRequest("GET", "/api/v1/balance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeUnauthorized, response.Code)
}
