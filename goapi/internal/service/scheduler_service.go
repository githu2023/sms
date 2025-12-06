package service

import (
	"context"
	"errors"
	"fmt"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/global"
	"sms-platform/goapi/internal/repository"
	"sms-platform/goapi/pkg/provider"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// SchedulerService 定时任务服务
type SchedulerService struct {
	cfg             config.SchedulerConfig
	assignmentRepo  repository.PhoneAssignmentRepository
	providerRepo    domain.ProviderRepository
	transactionSvc  TransactionService
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
	transactionSvc TransactionService,
	customerRepo domain.CustomerRepository,
	db *gorm.DB,
) *SchedulerService {
	return &SchedulerService{
		cfg:            cfg,
		assignmentRepo: assignmentRepo,
		providerRepo:   providerRepo,
		transactionSvc: transactionSvc,
		customerRepo:   customerRepo,
		db:             db,
		stopChan:       make(chan struct{}),
	}
}

// Start 启动定时任务
func (s *SchedulerService) Start() {
	zap.S().Info("Starting scheduler service...")

	// 启动验证码检查定时器
	s.codeCheckTicker = time.NewTicker(time.Duration(s.cfg.CodeCheckInterval) * time.Second)
	go s.runCodeCheckTask()

	// 启动过期分配清理定时器
	s.cleanupTicker = time.NewTicker(time.Duration(s.cfg.AssignmentCleanupInterval) * time.Second)
	go s.runCleanupTask()

	zap.S().Infof("Scheduler service started - Code check interval: %ds, Cleanup interval: %ds",
		s.cfg.CodeCheckInterval, s.cfg.AssignmentCleanupInterval)
}

// Stop 停止定时任务
func (s *SchedulerService) Stop() {
	zap.S().Info("Stopping scheduler service...")

	close(s.stopChan)

	if s.codeCheckTicker != nil {
		s.codeCheckTicker.Stop()
	}

	if s.cleanupTicker != nil {
		s.cleanupTicker.Stop()
	}

	zap.S().Info("Scheduler service stopped")
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
		zap.S().Errorf("Error finding pending assignments: %v", err)
		return
	}

	if len(assignments) == 0 {
		return
	}

	zap.S().Infof("Found %d pending assignments to check for codes", len(assignments))

	// 使用 errgroup 并发处理，限制并发数
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(s.cfg.MaxConcurrentChecks) // 限制并发数

	for _, assignment := range assignments {
		// 避免闭包问题，复制一份
		assign := assignment
		g.Go(func() error {
			s.processAssignment(gctx, assign)
			return nil // 不返回错误，让所有 goroutine 都执行完
		})
	}

	// 等待所有并发任务完成
	if err := g.Wait(); err != nil {
		zap.S().Errorf("Error processing assignments concurrently: %v", err)
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

	// 开始数据库事务用于更新分配状态
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := s.assignmentRepo.Update(ctx, tx, assignment); err != nil {
		tx.Rollback()
		zap.S().Errorf("Error marking assignment %d as expired: %v", assignment.ID, err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		zap.S().Errorf("Error committing expired assignment %d: %v", assignment.ID, err)
		return
	}

	// 如果已扣费，需要退款到余额
	// 注意：使用 ReserveAndCommitWithSingleRecord 后，冻结金额已经转为实扣（冻结金额=0）
	// 所以过期时应该使用 Refund (Type 3: 拉号回退) 退款到余额
	if assignment.CustomerID != nil && assignment.MerchantFee != nil && *assignment.MerchantFee > 0 {
		customerID := *assignment.CustomerID
		refundAmount := *assignment.MerchantFee
		notes := fmt.Sprintf("手机号过期退款: %s (分配ID: %d)",
			getPhoneNumber(assignment.PhoneNumber), assignment.ID)

		if _, err := s.transactionSvc.Refund(ctx, customerID, refundAmount, assignment.ID, notes); err != nil {
			zap.S().Errorf("Error refunding for expired assignment %d: %v", assignment.ID, err)
		} else {
			zap.S().Infof("Assignment %d marked as expired and refunded %.2f to customer %d",
				assignment.ID, refundAmount, customerID)
		}
	} else {
		zap.S().Infof("Assignment %d marked as expired (no refund needed)", assignment.ID)
	}

	zap.S().Infof("Assignment %d marked as expired due to timeout", assignment.ID)

	// 释放手机号
	if assignment.PhoneNumber != nil {
		s.releasePhoneNumber(ctx, assignment)
	}
}

// markAssignmentAsExpiredWithoutRelease 标记分配为已过期，处理退款，但不释放手机号（因为运营商已释放）
func (s *SchedulerService) markAssignmentAsExpiredWithoutRelease(ctx context.Context, assignment *domain.PhoneAssignment) {
	expiredStatus := "expired"
	assignment.Status = &expiredStatus
	assignment.UpdatedAt = time.Now()

	if assignment.Remark == nil {
		remark := "手机号已被运营商释放"
		assignment.Remark = &remark
	}

	// 开始数据库事务用于更新分配状态
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := s.assignmentRepo.Update(ctx, tx, assignment); err != nil {
		tx.Rollback()
		zap.S().Errorf("Error marking assignment %d as expired: %v", assignment.ID, err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		zap.S().Errorf("Error committing expired assignment %d: %v", assignment.ID, err)
		return
	}

	// 如果已扣费，需要退款到余额
	if assignment.CustomerID != nil && assignment.MerchantFee != nil && *assignment.MerchantFee > 0 {
		customerID := *assignment.CustomerID
		refundAmount := *assignment.MerchantFee
		notes := fmt.Sprintf("手机号已被运营商释放退款: %s (分配ID: %d)",
			getPhoneNumber(assignment.PhoneNumber), assignment.ID)

		if _, err := s.transactionSvc.Refund(ctx, customerID, refundAmount, assignment.ID, notes); err != nil {
			zap.S().Errorf("Error refunding for expired assignment %d: %v", assignment.ID, err)
		} else {
			zap.S().Infof("Assignment %d marked as expired (already released by provider) and refunded %.2f to customer %d",
				assignment.ID, refundAmount, customerID)
		}
	} else {
		zap.S().Infof("Assignment %d marked as expired (already released by provider, no refund needed)", assignment.ID)
	}

	zap.S().Infof("Assignment %d marked as expired - phone already released by provider", assignment.ID)
	// 注意：不调用 releasePhoneNumber，因为运营商已经释放了
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
		zap.S().Warnf("Assignment %d has no provider ID, cannot get code", assignment.ID)
		return
	}

	// 通过ProviderID查询Provider信息获取ProviderCode
	providerInfo, err := s.providerRepo.FindByID(ctx, int(*assignment.ProviderID))
	if err != nil {
		zap.S().Errorf("Error finding provider %d for assignment %d: %v", *assignment.ProviderID, assignment.ID, err)
		return
	}

	if providerInfo.Code == nil {
		zap.S().Warnf("Provider %d has no code", *assignment.ProviderID)
		return
	}

	// 从全局ProviderManager获取provider对象
	providerManager := global.GetProviderManager()
	providerInstance, err := providerManager.GetProviderByCode(*providerInfo.Code)
	if err != nil {
		zap.S().Errorf("Error getting provider instance for code %s: %v", *providerInfo.Code, err)
		return
	}

	// 设置较短的超时时间避免阻塞（使用配置的单次请求超时）
	requestTimeout := time.Duration(s.cfg.ProviderRequestTimeout) * time.Second
	codeCtx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	// 调用provider接口获取验证码
	// 如果数据库中有 extId，优先使用 extId
	var codeResponse *provider.CodeResponse
	if assignment.ExtId != nil && *assignment.ExtId != "" {
		// 使用 extId 获取验证码
		codeResponse, err = providerInstance.GetCode(codeCtx, *assignment.PhoneNumber, requestTimeout, *assignment.ExtId)
	} else {
		// 使用 phoneNumber 获取验证码（兼容旧数据）
		codeResponse, err = providerInstance.GetCode(codeCtx, *assignment.PhoneNumber, requestTimeout)
	}
	if err != nil {
		// 判断是否是"暂未接收"错误
		if errors.Is(err, provider.ErrCodeNotReceived) {
			// 这是正常情况，验证码还没发送，等待下次检查
			zap.S().Debugf("Verification code not received yet for assignment %d (phone: %s)",
				assignment.ID, *assignment.PhoneNumber)
			return
		}

		// 判断是否是"已释放"错误
		if errors.Is(err, provider.ErrPhoneAlreadyReleased) {
			// 手机号已被运营商释放，直接标记为过期，不需要再释放
			zap.S().Infof("Phone number already released by provider for assignment %d, marking as expired without release",
				assignment.ID)
			s.markAssignmentAsExpiredWithoutRelease(ctx, assignment)
			return
		}

		// 其他错误，记录失败次数
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
			zap.S().Errorf("Error updating fetch count for assignment %d: %v", assignment.ID, err)
		}

		zap.S().Warnf("Failed to get code for assignment %d (phone: %s), attempt %d: %v",
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
		zap.S().Errorf("Error updating assignment %d with code: %v", assignment.ID, err)
	} else {
		zap.S().Infof("Assignment %d successfully updated with verification code", assignment.ID)
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
		zap.S().Warnf("Assignment %d has no provider ID, cannot release phone", assignment.ID)
		return
	}

	// 通过ProviderID查询Provider信息获取ProviderCode
	providerInfo, err := s.providerRepo.FindByID(ctx, int(*assignment.ProviderID))
	if err != nil {
		zap.S().Errorf("Error finding provider %d for assignment %d to release phone: %v", *assignment.ProviderID, assignment.ID, err)
		return
	}

	providerCode := "unknown"
	if providerInfo.Code != nil {
		providerCode = *providerInfo.Code
	} else {
		zap.S().Warnf("Provider %d has no code, cannot release phone for assignment %d", *assignment.ProviderID, assignment.ID)
		return
	}

	// 从全局ProviderManager获取provider对象
	providerManager := global.GetProviderManager()
	providerInstance, err := providerManager.GetProviderByCode(providerCode)
	if err != nil {
		zap.S().Errorf("Error getting provider instance (code=%s, id=%d) to release phone: %v",
			providerCode, *assignment.ProviderID, err)
		return
	}

	// 调用ReleasePhone释放手机号
	// 如果数据库中有 extId，优先使用 extId（对于 BigBus666、MQTT 等需要 extId 的运营商）
	releaseCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 如果数据库中有 extId，传递给 ReleasePhone
	var extId *string
	if assignment.ExtId != nil && *assignment.ExtId != "" {
		extId = assignment.ExtId
	}

	if extId != nil {
		if err := providerInstance.ReleasePhone(releaseCtx, *assignment.PhoneNumber, *extId); err != nil {
			zap.S().Warnf("Error releasing phone %s (extId: %s) for assignment %d via provider %s: %v",
				*assignment.PhoneNumber, *extId, assignment.ID, providerCode, err)
		} else {
			zap.S().Infof("Successfully released phone %s (extId: %s) for assignment %d via provider %s",
				*assignment.PhoneNumber, *extId, assignment.ID, providerCode)
		}
	} else {
		// 如果没有 extId，尝试使用 phoneNumber（某些 provider 可能支持）
		if err := providerInstance.ReleasePhone(releaseCtx, *assignment.PhoneNumber); err != nil {
			zap.S().Warnf("Error releasing phone %s (no extId) for assignment %d via provider %s: %v",
				*assignment.PhoneNumber, assignment.ID, providerCode, err)
		} else {
			zap.S().Infof("Successfully released phone %s (no extId) for assignment %d via provider %s",
				*assignment.PhoneNumber, assignment.ID, providerCode)
		}
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
		zap.S().Errorf("Error finding expired assignments for cleanup: %v", err)
		return
	}

	if len(expiredAssignments) == 0 {
		return
	}

	zap.S().Infof("Found %d expired assignments to cleanup", len(expiredAssignments))

	for _, assignment := range expiredAssignments {
		s.markAssignmentAsExpired(ctx, assignment)
	}
}
