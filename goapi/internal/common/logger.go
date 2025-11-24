package common

import (
	"context"
	"fmt"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/repository"
	"time"
)

// APILogger provides a centralized API logging service
type APILogger struct {
	logRepo repository.LogRepository
}

// NewAPILogger creates a new API logger instance
func NewAPILogger(logRepo repository.LogRepository) *APILogger {
	return &APILogger{
		logRepo: logRepo,
	}
}

// LogSuccess logs a successful API call
func (l *APILogger) LogSuccess(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusSuccess, details, StatusOK)
}

// LogError logs a failed API call with specific error code
func (l *APILogger) LogError(ctx context.Context, customerID int64, path, details string, statusCode int) {
	l.logAPICall(ctx, customerID, path, StatusFailed, details, statusCode)
}

// LogBadRequest logs a bad request error (400)
func (l *APILogger) LogBadRequest(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusFailed, details, StatusBadRequest)
}

// LogUnauthorized logs an unauthorized error (401)
func (l *APILogger) LogUnauthorized(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusFailed, details, StatusUnauthorized)
}

// LogInsufficientBalance logs a payment required error (402)
func (l *APILogger) LogInsufficientBalance(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusFailed, details, StatusPaymentRequired)
}

// LogNotFound logs a not found error (404)
func (l *APILogger) LogNotFound(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusFailed, details, StatusNotFound)
}

// LogInternalError logs an internal server error (500)
func (l *APILogger) LogInternalError(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusFailed, details, StatusInternalServerError)
}

// LogCreated logs a successful resource creation (201)
func (l *APILogger) LogCreated(ctx context.Context, customerID int64, path, details string) {
	l.logAPICall(ctx, customerID, path, StatusSuccess, details, StatusCreated)
}

// logAPICall is the core logging function - single line of implementation
func (l *APILogger) logAPICall(ctx context.Context, customerID int64, path, status, details string, statusCode int) {
	go l.createLogEntry(ctx, customerID, path, status, details, statusCode)
}

// createLogEntry creates and saves the log entry
func (l *APILogger) createLogEntry(ctx context.Context, customerID int64, path, status, details string, statusCode int) {
	logEntry := &domain.APILog{
		CustomerID:   customerID,
		RequestIP:    getIPFromContext(ctx),
		RequestPath:  path,
		RequestBody:  fmt.Sprintf("Status: %s, Details: %s", status, details),
		ResponseCode: statusCode,
		CreatedAt:    time.Now(),
	}

	// Best effort logging - ignore errors
	l.logRepo.Create(context.Background(), logEntry)
}

// getIPFromContext extracts IP address from context
func getIPFromContext(ctx context.Context) string {
	if ip, ok := ctx.Value("ip_address").(string); ok {
		return ip
	}
	return "unknown"
}

// getUserAgentFromContext extracts user agent from context
func getUserAgentFromContext(ctx context.Context) string {
	if ua, ok := ctx.Value("user_agent").(string); ok {
		return ua
	}
	return "unknown"
}
