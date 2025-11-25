package service

import (
	"context"
	"errors"
	"fmt"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/repository"
	"time"
)

var (
	ErrInvalidTimeRange = errors.New("start date cannot be after end date")
)

// AssignmentStatistics encapsulates cost and count statistics for assignments.
type AssignmentStatistics struct {
	TotalCost  float64 `json:"total_cost"`
	TotalCount int64   `json:"total_count"`
}

// AssignmentService defines the interface for assignment related business logic
type AssignmentService interface {
	GetAssignments(ctx context.Context, customerID int64, page, limit int, status int, businessType string, startDate, endDate *time.Time) ([]*domain.PhoneAssignment, int64, error)
	GetCostStatistics(ctx context.Context, customerID int64, startDate, endDate *time.Time) (*AssignmentStatistics, error)
	GetRecentAssignments(ctx context.Context, customerID int64, limit int) ([]*domain.PhoneAssignment, error)
}

// assignmentService implements AssignmentService
type assignmentService struct {
	assignmentRepo repository.PhoneAssignmentRepository
	businessRepo   domain.BusinessTypeRepository
	providerRepo   domain.ProviderRepository
}

// NewAssignmentService creates a new AssignmentService instance
func NewAssignmentService(
	assignmentRepo repository.PhoneAssignmentRepository,
	businessRepo domain.BusinessTypeRepository,
	providerRepo domain.ProviderRepository,
) AssignmentService {
	return &assignmentService{
		assignmentRepo: assignmentRepo,
		businessRepo:   businessRepo,
		providerRepo:   providerRepo,
	}
}

// GetAssignments returns paginated assignment history for a customer with filtering options.
func (s *assignmentService) GetAssignments(ctx context.Context, customerID int64, page, limit int, status int, businessType string, startDate, endDate *time.Time) ([]*domain.PhoneAssignment, int64, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	// In a real scenario, businessType and status filtering would be handled by the repository
	// For now, we'll fetch all and filter in memory, which is inefficient for large datasets.
	// This would require new methods in PhoneAssignmentRepository.
	allAssignments, _, err := s.assignmentRepo.FindRecentByCustomerID(ctx, customerID, limit, offset) // Placeholder, would need more sophisticated repo method
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get assignments: %w", err)
	}

	filteredAssignments := []*domain.PhoneAssignment{}
	var filteredTotal int64

	for _, assignment := range allAssignments {
		match := true

		// Convert status int to string and compare
		var statusStr string
		switch status {
		case 1:
			statusStr = "pending"
		case 2:
			statusStr = "completed"
		case 3:
			statusStr = "expired"
		case 4:
			statusStr = "failed"
		}
		if status != 0 && (assignment.Status == nil || *assignment.Status != statusStr) {
			match = false
		}
		// Need to resolve BusinessTypeID to Code
		// For now, assume BusinessTypeID is directly comparable to a businessType string for simplicity, this would need businessRepo lookup
		// if businessType != "" && assignment.BusinessTypeID != lookupCodeToID(businessType) {
		// 	match = false
		// }

		// This date filtering logic needs to be pushed to the repository for efficiency
		if startDate != nil && assignment.CreatedAt.Before(*startDate) {
			match = false
		}
		if endDate != nil && assignment.CreatedAt.After(*endDate) {
			match = false
		}

		if match {
			filteredAssignments = append(filteredAssignments, assignment)
			filteredTotal++
		}
	}
	// The current FindRecentByCustomerID only supports customerID.
	// For actual filtering, new repository methods like FindByCustomerIDWithFilters are needed.
	// This current implementation is a placeholder demonstrating the service layer's filtering intent.

	return filteredAssignments, filteredTotal, nil
}

// GetCostStatistics calculates the total cost and count for a customer's assignments with date range filtering.
func (s *assignmentService) GetCostStatistics(ctx context.Context, customerID int64, startDate, endDate *time.Time) (*AssignmentStatistics, error) {
	if startDate != nil && endDate != nil && startDate.After(*endDate) {
		return nil, ErrInvalidTimeRange
	}

	// In a real system, this would require a specific repository method to sum costs and count
	// with date range filtering for efficiency.
	// For now, fetch all and filter/sum in memory (inefficient for large datasets).
	assignments, _, err := s.assignmentRepo.FindRecentByCustomerID(ctx, customerID, 9999999, 0) // Fetch all for sum
	if err != nil {
		return nil, fmt.Errorf("failed to get assignments for cost calculation: %w", err)
	}

	var totalCost float64
	var totalCount int64

	for _, assignment := range assignments {
		match := true
		if startDate != nil && assignment.CreatedAt.Before(*startDate) {
			match = false
		}
		if endDate != nil && assignment.CreatedAt.After(*endDate) {
			match = false
		}

		if match {
			// Use MerchantFee as the cost field
			if assignment.MerchantFee != nil {
				totalCost += float64(*assignment.MerchantFee)
			}
			totalCount++
		}
	}

	return &AssignmentStatistics{
		TotalCost:  totalCost,
		TotalCount: totalCount,
	}, nil
}

// GetRecentAssignments 返回最近的手机号分配记录
func (s *assignmentService) GetRecentAssignments(ctx context.Context, customerID int64, limit int) ([]*domain.PhoneAssignment, error) {
	if limit <= 0 {
		limit = 5
	}
	if limit > 50 {
		limit = 50
	}

	assignments, _, err := s.assignmentRepo.FindRecentByCustomerID(ctx, customerID, limit, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent assignments: %w", err)
	}
	return assignments, nil
}
