package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sms-platform/goapi/internal/api/handler/testutils" // Import common test utils
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/domain" // Added domain import
	"sms-platform/goapi/internal/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAssignmentService for testing
type MockAssignmentService struct {
	mock.Mock
}

// These methods are from the richer AssignmentService that the test expects.
// They will be implemented in the actual service as per TODO, but for now, we define them in mock.

func (m *MockAssignmentService) GetAssignments(ctx context.Context, customerID int64, page, limit int, status int, businessType string, startDate, endDate *time.Time) ([]*domain.PhoneAssignment, int64, error) {
	args := m.Called(ctx, customerID, page, limit, status, businessType, startDate, endDate)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]*domain.PhoneAssignment), args.Get(1).(int64), args.Error(2)
}

func (m *MockAssignmentService) GetCostStatistics(ctx context.Context, customerID int64, startDate, endDate *time.Time) (*service.AssignmentStatistics, error) {
	args := m.Called(ctx, customerID, startDate, endDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.AssignmentStatistics), args.Error(1)
}

/*
// Commenting out unimplemented/misaligned mock methods
func (m *MockAssignmentService) GetAssignmentHistory(ctx context.Context, customerID int64, page, limit int) (*service.AssignmentHistoryResult, error) {
	args := m.Called(ctx, customerID, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.AssignmentHistoryResult), args.Error(1)
}

func (m *MockAssignmentService) GetAssignmentHistoryByStatus(ctx context.Context, customerID int64, status int, page, limit int) (*service.AssignmentHistoryResult, error) {
	args := m.Called(ctx, customerID, status, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.AssignmentHistoryResult), args.Error(1)
}

func (m *MockAssignmentService) GetAssignmentHistoryByBusinessType(ctx context.Context, customerID int64, businessTypeID int, page, limit int) (*service.AssignmentHistoryResult, error) {
	args := m.Called(ctx, customerID, businessTypeID, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.AssignmentHistoryResult), args.Error(1)
}

func (m *MockAssignmentService) GetAssignmentHistoryByDateRange(ctx context.Context, customerID int64, startDate, endDate time.Time, page, limit int) (*service.AssignmentHistoryResult, error) {
	args := m.Called(ctx, customerID, startDate, endDate, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.AssignmentHistoryResult), args.Error(1)
}

func (m *MockAssignmentService) UpdateAssignmentStatus(ctx context.Context, assignmentID int64, status int, verificationCode string) error {
	args := m.Called(ctx, assignmentID, status, verificationCode)
	return args.Error(0)
}

func (m *MockAssignmentService) ExpireOldAssignments(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *MockAssignmentService) CreateAssignment(ctx context.Context, assignment *domain.PhoneAssignment) error {
	args := m.Called(ctx, assignment)
	return args.Error(0)
}
*/

func TestAssignmentHandler_GetAssignments_Success(t *testing.T) {
	// Setup
	mockService := &MockAssignmentService{}
	handler := NewAssignmentHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Mock data - use domain.PhoneAssignment as service.GetAssignments returns this
	now := time.Now()
	assignments := []*domain.PhoneAssignment{
		{
			ID:               1,
			ProviderID:       "provider-1",
			BusinessTypeID:   1, // Assuming this exists or is mocked
			CardType:         "virtual",
			PhoneNumber:      "+15551234567",
			VerificationCode: "123456",
			Cost:             0.10,
			Status:           2,
			ExpiresAt:        &now,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}
	total := int64(1)

	// Setup mock expectations for the updated GetAssignments signature
	mockService.On("GetAssignments", mock.Anything, int64(1), 1, 20, 0, "", (*time.Time)(nil), (*time.Time)(nil)).Return(assignments, total, nil)

	// Setup route
	router.GET("/client/v1/assignments", handler.GetAssignments)

	// Test
	req, _ := http.NewRequest("GET", "/client/v1/assignments?page=1&limit=20", nil)
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
	assert.Equal(t, float64(1), data["pagination"].(map[string]interface{})["total"])
	// Further assertions on items would require more complex DTO conversion logic which is handled in the handler.

	mockService.AssertExpectations(t)
}

func TestAssignmentHandler_GetAssignments_WithFilter_Success(t *testing.T) {
	// Setup
	mockService := &MockAssignmentService{}
	handler := NewAssignmentHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Mock data
	now := time.Now()
	assignments := []*domain.PhoneAssignment{
		{
			ID:               2,
			ProviderID:       "provider-2",
			BusinessTypeID:   2,
			CardType:         "physical",
			PhoneNumber:      "+15557654321",
			VerificationCode: "987654",
			Cost:             0.20,
			Status:           1,
			ExpiresAt:        &now,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}
	total := int64(1)

	// Setup mock expectations for filtering
	startDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC) // End of day

	mockService.On("GetAssignments", mock.Anything, int64(1), 1, 20, 1, "test_business", &startDate, &endDate).Return(assignments, total, nil)

	// Setup route
	router.GET("/client/v1/assignments", handler.GetAssignments)

	// Test
	req, _ := http.NewRequest("GET", "/client/v1/assignments?status=1&business_type=test_business&start_date=2023-01-01&end_date=2023-12-31", nil)
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

func TestAssignmentHandler_GetAssignments_Unauthorized(t *testing.T) {
	// Setup
	mockService := &MockAssignmentService{}
	handler := NewAssignmentHandler(mockService)
	router := gin.New() // Don't use setupTestRouter to avoid mock customer_id

	// Setup route without setting customer_id
	router.GET("/client/v1/assignments", handler.GetAssignments)

	// Test
	req, _ := http.NewRequest("GET", "/client/v1/assignments", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeUnauthorized, response.Code)
}

func TestAssignmentHandler_GetAssignments_ServiceError(t *testing.T) {
	// Setup
	mockService := &MockAssignmentService{}
	handler := NewAssignmentHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Setup mock expectations for the updated GetAssignments signature
	mockService.On("GetAssignments", mock.Anything, int64(1), 1, 20, 0, "", (*time.Time)(nil), (*time.Time)(nil)).Return(nil, int64(0), errors.New("database error"))

	// Setup route
	router.GET("/client/v1/assignments", handler.GetAssignments)

	// Test
	req, _ := http.NewRequest("GET", "/client/v1/assignments", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeInternalError, response.Code)

	mockService.AssertExpectations(t)
}

func TestAssignmentHandler_GetCostStatistics_Success(t *testing.T) {
	// Setup
	mockService := &MockAssignmentService{}
	handler := NewAssignmentHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Mock data
	startDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC) // End of day

	stats := &service.AssignmentStatistics{ // Use service.AssignmentStatistics
		TotalCost:  10.50,
		TotalCount: 25,
	}

	// Setup mock expectations for the updated GetCostStatistics signature
	mockService.On("GetCostStatistics", mock.Anything, int64(1), &startDate, &endDate).Return(stats, nil)

	// Setup route
	router.GET("/client/v1/assignments/statistics", handler.GetCostStatistics)

	// Test
	req, _ := http.NewRequest("GET", "/client/v1/assignments/statistics?start_date=2023-01-01&end_date=2023-12-31", nil)
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
	assert.Equal(t, 10.5, data["total_cost"])
	assert.Equal(t, float64(25), data["total_count"])

	mockService.AssertExpectations(t)
}

func TestAssignmentHandler_GetCostStatistics_InvalidDateFormat(t *testing.T) {
	// Setup
	mockService := &MockAssignmentService{}
	handler := NewAssignmentHandler(mockService)
	router := testutils.SetupTestRouter() // Use common setupTestRouter

	// Setup route
	router.GET("/client/v1/assignments/statistics", handler.GetCostStatistics)

	// Test with invalid date format
	req, _ := http.NewRequest("GET", "/client/v1/assignments/statistics?start_date=invalid&end_date=2023-12-31", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response common.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, common.CodeBadRequest, response.Code)
	assert.Contains(t, response.Msg, "请求参数错误") // Changed from "Invalid date format" to match common.RespondErrorWithMsg
}
