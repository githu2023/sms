package service

import (
	"context"
	"fmt"
	"log"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/global"
	"sms-platform/goapi/internal/repository"
	"sms-platform/goapi/pkg/provider"
	"time"

	"gorm.io/gorm"
)

// SchedulerService 定时任务服务
type SchedulerService struct {
	cfg             config.SchedulerConfig
	assignmentRepo  repository.PhoneAssignmentRepository
	providerRepo    domain.ProviderRepository
	transactionRepo repository.TransactionRepository
	customerRepo    domain.CustomerRepository
	db              *gorm.DB
	stopChan        chan struct{}
	codeCheckTicker *time.Ticker
	cleanupTicker   *time.Ticker
}

// NewSchedulerService 创建新的定时任务服务
func NewSchedulerService(
	cfg config.SchedulerConfig,
	assignmentRepo repository.PhoneAssignmentRepository,
	providerRepo domain.ProviderRepository,
	transactionRepo repository.TransactionRepository,
	customerRepo domain.CustomerRepository,
	db *gorm.DB,
) *SchedulerService {
	return &SchedulerService{
		cfg:             cfg,
		assignmentRepo:  assignmentRepo,
		providerRepo:    providerRepo,
		transactionRepo: transactionRepo,
		customerRepo:    customerRepo,
		db:              db,
		stopChan:        make(chan struct{}),
	}
}

// Start 启动定时任务
func (s *SchedulerService) Start() {
	log.Println("Starting scheduler service...")

	// 启动验证码检查定时器
	s.codeCheckTicker = time.NewTicker(time.Duration(s.cfg.CodeCheckInterval) * time.Second)
	go s.runCodeCheckTask()

	// 启动过期分配清理定时器
	s.cleanupTicker = time.NewTicker(time.Duration(s.cfg.AssignmentCleanupInterval) * time.Second)
	go s.runCleanupTask()

	log.Printf("Scheduler service started - Code check interval: %ds, Cleanup interval: %ds",
		s.cfg.CodeCheckInterval, s.cfg.AssignmentCleanupInterval)
}

// Stop 停止定时任务
func (s *SchedulerService) Stop() {
	log.Println("Stopping scheduler service...")

	close(s.stopChan)

	if s.codeCheckTicker != nil {
		s.codeCheckTicker.Stop()
	}

	if s.cleanupTicker != nil {
		s.cleanupTicker.Stop()
	}

	log.Println("Scheduler service stopped")
}

// runCodeCheckTask 运行验证码检查任务
func (s *SchedulerService) runCodeCheckTask() {
	for {
		select {
		case <-s.codeCheckTicker.C:
			s.checkPendingCodes()
		case <-s.stopChan:
			return
		}
	}
}

// runCleanupTask 运行过期分配清理任务
func (s *SchedulerService) runCleanupTask() {
	for {
		select {
		case <-s.cleanupTicker.C:
			s.cleanupExpiredAssignments()
		case <-s.stopChan:
			return
		}
	}
}

// checkPendingCodes 检查待获取验证码的手机号
func (s *SchedulerService) checkPendingCodes() {
	ctx := context.Background()

	// 查找所有状态为 "pending" 的分配记录
	assignments, err := s.findPendingAssignments(ctx)
	if err != nil {
		log.Printf("Error finding pending assignments: %v", err)
		return
	}

	if len(assignments) == 0 {
		return
	}

	log.Printf("Found %d pending assignments to check for codes", len(assignments))

	for _, assignment := range assignments {
		s.processAssignment(ctx, assignment)
	}
}

// findPendingAssignments 查找所有待处理的分配记录
func (s *SchedulerService) findPendingAssignments(ctx context.Context) ([]*domain.PhoneAssignment, error) {
	var assignments []*domain.PhoneAssignment

	status := "pending"
	err := s.db.WithContext(ctx).Where("status = ? AND phone_number IS NOT NULL", status).Find(&assignments).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find pending assignments: %w", err)
	}

	return assignments, nil
}

// processAssignment 处理单个分配记录
func (s *SchedulerService) processAssignment(ctx context.Context, assignment *domain.PhoneAssignment) {
	// 检查是否超时
	if s.isAssignmentExpired(assignment) {
		s.markAssignmentAsExpired(ctx, assignment)
		return
	}

	// 尝试获取验证码
	if assignment.PhoneNumber != nil {
		s.tryGetCode(ctx, assignment)
	}
}

// isAssignmentExpired 检查分配是否已过期
func (s *SchedulerService) isAssignmentExpired(assignment *domain.PhoneAssignment) bool {
	timeout := time.Duration(s.cfg.CodeTimeout) * time.Second
	expiryTime := assignment.CreatedAt.Add(timeout)
	return time.Now().After(expiryTime)
}

// markAssignmentAsExpired 标记分配为已过期，并处理退款
func (s *SchedulerService) markAssignmentAsExpired(ctx context.Context, assignment *domain.PhoneAssignment) {
	expiredStatus := "expired"
	assignment.Status = &expiredStatus
	assignment.UpdatedAt = time.Now()

	if assignment.Remark == nil {
		remark := fmt.Sprintf("验证码获取超时，已过期。超时时间：%d秒", s.cfg.CodeTimeout)
		assignment.Remark = &remark
	}

	// 开始数据库事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新分配记录状态
	err := s.assignmentRepo.Update(ctx, tx, assignment)
	if err != nil {
		tx.Rollback()
		log.Printf("Error marking assignment %d as expired: %v", assignment.ID, err)
		return
	}

	// 如果已扣费，需要退款
	if assignment.CustomerID != nil && assignment.MerchantFee != nil && *assignment.MerchantFee > 0 {
		customerID := *assignment.CustomerID
		refundAmount := *assignment.MerchantFee

		// 获取当前余额
		currentBalance, err := s.transactionRepo.GetBalance(ctx, customerID)
		if err != nil {
			tx.Rollback()
			log.Printf("Error getting balance for customer %d: %v", customerID, err)
			return
		}

		// 创建退款交易记录（类型为 "3:拉号-回退"）
		refundAmountFloat := float32(refundAmount)
		balanceBeforeFloat := float32(currentBalance)
		balanceAfterFloat := float32(currentBalance + refundAmount)
		transactionType := "3" // 拉号-回退
		assignmentID := assignment.ID
		notes := fmt.Sprintf("手机号过期退款: %s (分配ID: %d)", 
			getPhoneNumber(assignment.PhoneNumber), assignment.ID)

		refundTx := &domain.Transaction{
			CustomerID:    customerID,
			Amount:        &refundAmountFloat,
			BalanceBefore: &balanceBeforeFloat,
			BalanceAfter:  &balanceAfterFloat,
			Type:          &transactionType,
			ReferenceID:   &assignmentID,
			Notes:         &notes,
			CreatedAt:     time.Now(),
		}

		err = s.transactionRepo.Create(ctx, refundTx)
		if err != nil {
			tx.Rollback()
			log.Printf("Error creating refund transaction for assignment %d: %v", assignment.ID, err)
			return
		}

		// 更新客户余额
		customer, err := s.customerRepo.FindByID(ctx, customerID)
		if err != nil {
			tx.Rollback()
			log.Printf("Error finding customer %d: %v", customerID, err)
			return
		}
		customer.Balance = currentBalance + refundAmount
		err = s.customerRepo.Update(ctx, customer)
		if err != nil {
			tx.Rollback()
			log.Printf("Error updating customer balance for customer %d: %v", customerID, err)
			return
		}

		log.Printf("Assignment %d marked as expired and refunded %.2f to customer %d", 
			assignment.ID, refundAmount, customerID)
	} else {
		log.Printf("Assignment %d marked as expired (no refund needed)", assignment.ID)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction for expired assignment %d: %v", assignment.ID, err)
	} else {
		log.Printf("Assignment %d marked as expired due to timeout", assignment.ID)
	}

	// 释放手机号
	if assignment.PhoneNumber != nil {
		s.releasePhoneNumber(ctx, assignment)
	}
}

// getPhoneNumber 安全获取手机号字符串
func getPhoneNumber(phone *string) string {
	if phone == nil {
		return "N/A"
	}
	return *phone
}

// tryGetCode 尝试获取验证码
func (s *SchedulerService) tryGetCode(ctx context.Context, assignment *domain.PhoneAssignment) {
	// 根据assignment中的ProviderID获取provider对象
	if assignment.ProviderID == nil {
		log.Printf("Assignment %d has no provider ID, cannot get code", assignment.ID)
		return
	}

	// 通过ProviderID查询Provider信息获取ProviderCode
	providerInfo, err := s.providerRepo.FindByID(ctx, int(*assignment.ProviderID))
	if err != nil {
		log.Printf("Error finding provider %d for assignment %d: %v", *assignment.ProviderID, assignment.ID, err)
		return
	}

	if providerInfo.Code == nil {
		log.Printf("Provider %d has no code", *assignment.ProviderID)
		return
	}

	// 从全局ProviderManager获取provider对象
	providerManager := global.GetProviderManager()
	providerInstance, err := providerManager.GetProviderByCode(*providerInfo.Code)
	if err != nil {
		log.Printf("Error getting provider instance for code %s: %v", *providerInfo.Code, err)
		return
	}

	// 设置较短的超时时间避免阻塞
	codeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 调用provider接口获取验证码
	// 如果数据库中有 extId，优先使用 extId
	var codeResponse *provider.CodeResponse
	if assignment.ExtId != nil && *assignment.ExtId != "" {
		// 使用 extId 获取验证码
		codeResponse, err = providerInstance.GetCode(codeCtx, *assignment.PhoneNumber, 10*time.Second, *assignment.ExtId)
	} else {
		// 使用 phoneNumber 获取验证码（兼容旧数据）
		codeResponse, err = providerInstance.GetCode(codeCtx, *assignment.PhoneNumber, 10*time.Second)
	}
	if err != nil {
		// 记录错误但不标记为失败，等待下次检查
		if assignment.FetchCount == nil {
			fetchCount := 1
			assignment.FetchCount = &fetchCount
		} else {
			*assignment.FetchCount++
		}

		// 更新获取次数
		err = s.assignmentRepo.Update(ctx, s.db, assignment)
		if err != nil {
			log.Printf("Error updating fetch count for assignment %d: %v", assignment.ID, err)
		}

		log.Printf("Failed to get code for assignment %d (phone: %s), attempt %d: %v",
			assignment.ID, *assignment.PhoneNumber, *assignment.FetchCount, err)
		return
	}

	// 成功获取验证码，更新记录
	s.updateAssignmentWithCode(ctx, assignment, codeResponse)
}

// updateAssignmentWithCode 更新分配记录的验证码信息
func (s *SchedulerService) updateAssignmentWithCode(ctx context.Context, assignment *domain.PhoneAssignment, codeResponse *provider.CodeResponse) {
	completedStatus := "completed"
	assignment.Status = &completedStatus
	assignment.VerificationCode = &codeResponse.Code
	assignment.UpdatedAt = time.Now()

	if assignment.FetchCount == nil {
		fetchCount := 1
		assignment.FetchCount = &fetchCount
	}

	// 更新备注
	remark := fmt.Sprintf("验证码获取成功，提供商：%s，获取时间：%s",
		codeResponse.ProviderID, codeResponse.ReceivedAt.Format("2006-01-02 15:04:05"))
	assignment.Remark = &remark

	err := s.assignmentRepo.Update(ctx, s.db, assignment)
	if err != nil {
		log.Printf("Error updating assignment %d with code: %v", assignment.ID, err)
	} else {
		log.Printf("Assignment %d successfully updated with verification code", assignment.ID)
	}

	// 释放手机号
	if assignment.PhoneNumber != nil {
		s.releasePhoneNumber(ctx, assignment)
	}
}

// releasePhoneNumber 释放手机号
func (s *SchedulerService) releasePhoneNumber(ctx context.Context, assignment *domain.PhoneAssignment) {
	if assignment.PhoneNumber == nil {
		return
	}

	// 获取provider实例
	if assignment.ProviderID == nil {
		log.Printf("Assignment %d has no provider ID, cannot release phone", assignment.ID)
		return
	}

	// 通过ProviderID查询Provider信息获取ProviderCode
	providerInfo, err := s.providerRepo.FindByID(ctx, int(*assignment.ProviderID))
	if err != nil {
		log.Printf("Error finding provider %d for assignment %d to release phone: %v", *assignment.ProviderID, assignment.ID, err)
		return
	}

	if providerInfo.Code == nil {
		log.Printf("Provider %d has no code, cannot release phone for assignment %d", *assignment.ProviderID, assignment.ID)
		return
	}

	// 从全局ProviderManager获取provider对象
	providerManager := global.GetProviderManager()
	providerInstance, err := providerManager.GetProviderByCode(*providerInfo.Code)
	if err != nil {
		log.Printf("Error getting provider instance for code %s to release phone: %v", *providerInfo.Code, err)
		return
	}

	// 调用ReleasePhone释放手机号
	// 如果数据库中有 extId，优先使用 extId（对于 BigBus666 等需要 extId 的运营商）
	releaseCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// ReleasePhone 接口只接收 phoneNumber，但内部会查找 extId
	// 如果 extId 在数据库中，Provider 应该能够从数据库中恢复映射关系
	// 或者我们需要修改 ReleasePhone 接口支持 extId 参数
	// 目前先使用 phoneNumber，Provider 内部会处理
	if err := providerInstance.ReleasePhone(releaseCtx, *assignment.PhoneNumber); err != nil {
		log.Printf("Error releasing phone %s for assignment %d: %v", *assignment.PhoneNumber, assignment.ID, err)
	} else {
		log.Printf("Successfully released phone %s for assignment %d", *assignment.PhoneNumber, assignment.ID)
	}
}

// cleanupExpiredAssignments 清理过期的分配记录
func (s *SchedulerService) cleanupExpiredAssignments() {
	ctx := context.Background()

	// 查找已过期但状态仍为 pending 的记录
	cutoffTime := time.Now().Add(-time.Duration(s.cfg.CodeTimeout) * time.Second)
	var expiredAssignments []*domain.PhoneAssignment

	status := "pending"
	err := s.db.WithContext(ctx).Where("status = ? AND created_at < ?", status, cutoffTime).Find(&expiredAssignments).Error
	if err != nil {
		log.Printf("Error finding expired assignments for cleanup: %v", err)
		return
	}

	if len(expiredAssignments) == 0 {
		return
	}

	log.Printf("Found %d expired assignments to cleanup", len(expiredAssignments))

	for _, assignment := range expiredAssignments {
		s.markAssignmentAsExpired(ctx, assignment)
	}
}
