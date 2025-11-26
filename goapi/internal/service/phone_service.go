package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sms-platform/goapi/internal/common"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/internal/global"
	"sms-platform/goapi/internal/repository"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PhoneService handles phone number and verification code operations
type PhoneService struct {
	transactionRepo                     repository.TransactionRepository
	transactionService                  TransactionService
	logRepo                             repository.LogRepository
	assignmentRepo                      repository.PhoneAssignmentRepository
	customerBusinessConfigRepo          repository.CustomerBusinessConfigRepository
	businessTypeRepo                    domain.BusinessTypeRepository
	platformBusinessTypeRepo            domain.PlatformBusinessTypeRepository
	platformProviderBusinessMappingRepo domain.PlatformProviderBusinessMappingRepository
	providerBusinessTypeRepo            domain.ProviderBusinessTypeRepository
	providerRepo                        domain.ProviderRepository
	customerRepo                        domain.CustomerRepository
	db                                  *gorm.DB
	apiLogger                           *common.APILogger
}

// PhoneServiceInterface defines the phone service methods
type PhoneServiceInterface interface {
	GetPhone(ctx context.Context, customerID int64, businessType, cardType string, count int) ([]*GetPhoneResult, common.ErrorCode)
	GetCode(ctx context.Context, customerID int64, phoneNumber string) ([]*GetCodeResult, error)
}

// GetPhoneResult represents the result of getting a phone number
type GetPhoneResult struct {
	PhoneNumber string    `json:"phone_number"`
	CountryCode string    `json:"country_code"`
	Cost        float64   `json:"cost"`
	ValidUntil  time.Time `json:"valid_until"`
	ProviderID  string    `json:"provider_id"`
	Balance     float64   `json:"remaining_balance"`
}

// GetCodeResult represents the result of getting a verification code
type GetCodeResult struct {
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	ReceivedAt time.Time `json:"received_at"`
	ProviderID string    `json:"provider_id"`
}

// NewPhoneService creates a new phone service instance
func NewPhoneService(
	transactionRepo repository.TransactionRepository,
	transactionService TransactionService,
	logRepo repository.LogRepository,
	assignmentRepo repository.PhoneAssignmentRepository,
	customerBusinessConfigRepo repository.CustomerBusinessConfigRepository,
	businessTypeRepo domain.BusinessTypeRepository,
	platformBusinessTypeRepo domain.PlatformBusinessTypeRepository,
	platformProviderBusinessMappingRepo domain.PlatformProviderBusinessMappingRepository,
	providerBusinessTypeRepo domain.ProviderBusinessTypeRepository,
	providerRepo domain.ProviderRepository,
	customerRepo domain.CustomerRepository,
	db *gorm.DB,
) *PhoneService {
	return &PhoneService{
		transactionRepo:                     transactionRepo,
		transactionService:                  transactionService,
		logRepo:                             logRepo,
		assignmentRepo:                      assignmentRepo,
		customerBusinessConfigRepo:          customerBusinessConfigRepo,
		businessTypeRepo:                    businessTypeRepo,
		platformBusinessTypeRepo:            platformBusinessTypeRepo,
		platformProviderBusinessMappingRepo: platformProviderBusinessMappingRepo,
		providerBusinessTypeRepo:            providerBusinessTypeRepo,
		providerRepo:                        providerRepo,
		customerRepo:                        customerRepo,
		db:                                  db,
		apiLogger:                           common.NewAPILogger(logRepo),
	}
}

// GetPhone requests a phone number for receiving SMS
func (s *PhoneService) GetPhone(ctx context.Context, customerID int64, businessType, cardType string, count int) ([]*GetPhoneResult, common.ErrorCode) {
	// 1. 查询客户的业务配置（平台业务类型）
	customerConfig, err := s.customerBusinessConfigRepo.FindByCustomerIDAndBusinessCode(ctx, customerID, businessType)
	if err != nil {
		s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("Customer business config not found for business type: %s", businessType))
		return nil, common.CodeInvalidBusinessType
	}

	// 检查配置是否启用
	if customerConfig.Status != nil && !*customerConfig.Status {
		s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("Business type %s is disabled for this customer", businessType))
		return nil, common.CodeInvalidBusinessType
	}

	customer, err := s.customerRepo.FindByID(ctx, customerID)
	if err != nil {
		s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("Failed to load customer info: %v", err))
		return nil, common.CodeInternalError
	}

	if customer.MerchantNo == nil || *customer.MerchantNo == "" ||
		customer.MerchantName == nil || *customer.MerchantName == "" {
		msg := "Customer merchant_no or merchant_name not configured"
		s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_phone", msg)
		return nil, common.CodeInternalError
	}

	merchantNo := *customer.MerchantNo
	merchantName := *customer.MerchantName

	// 2. 根据平台业务类型ID，查询出所有子业务类型映射（运营商业务类型）
	mappings, err := s.platformProviderBusinessMappingRepo.FindByPlatformBusinessTypeID(ctx, customerConfig.PlatformBusinessTypeID)
	if err != nil {
		s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("Failed to find mappings for platform business type: %d", customerConfig.PlatformBusinessTypeID))
		return nil, common.CodeInternalError
	}

	if len(mappings) == 0 {
		s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("No provider business mappings found for platform business type: %s", businessType))
		return nil, common.CodeInvalidBusinessType
	}

	// 3. 根据权重随机选择一个运营商的业务子类型
	selectedMapping := s.selectByWeight(mappings)
	if selectedMapping == nil {
		return nil, common.CodeInternalError
	}

	// 4. 根据子业务类型，找到对应的运营商对象
	providerManager := global.GetProviderManager()
	global.LogInfo("查找运营商",
		zap.Int64("customer_id", customerID),
		zap.String("business_type", businessType),
		zap.String("provider_code", *selectedMapping.ProviderCode),
		zap.String("business_code", *selectedMapping.BusinessCode))

	providerInstance, err := providerManager.GetProviderByCode(*selectedMapping.ProviderCode)
	if err != nil {
		errorMsg := fmt.Sprintf("服务商未找到: code=%s, 错误=%v", *selectedMapping.ProviderCode, err)
		global.LogError("获取手机号失败：服务商未找到",
			zap.Int64("customer_id", customerID),
			zap.String("business_type", businessType),
			zap.String("provider_code", *selectedMapping.ProviderCode),
			zap.Error(err))
		s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone", errorMsg)
		return nil, common.CodeProviderNotFound
	}

	providerInfo := providerInstance.GetProviderInfo()
	global.LogInfo("运营商获取成功",
		zap.Int64("customer_id", customerID),
		zap.String("provider_code", *selectedMapping.ProviderCode),
		zap.String("provider_id", providerInfo.ID),
		zap.String("provider_name", providerInfo.Name),
		zap.String("provider_type", providerInfo.Type))

	// 5. 查询运营商业务类型，获取价格
	providerBusinessType, err := s.providerBusinessTypeRepo.FindByProviderCodeAndBusinessCode(
		ctx, *selectedMapping.ProviderCode, *selectedMapping.BusinessCode)
	if err != nil {
		errorMsg := fmt.Sprintf("服务商业务类型未配置: provider=%s, business_code=%s, 错误=%v",
			*selectedMapping.ProviderCode, *selectedMapping.BusinessCode, err)
		global.LogError("获取手机号失败：服务商业务类型未配置",
			zap.Int64("customer_id", customerID),
			zap.String("business_type", businessType),
			zap.String("provider_code", *selectedMapping.ProviderCode),
			zap.String("business_code", *selectedMapping.BusinessCode),
			zap.Error(err))
		s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone", errorMsg)
		return nil, common.CodeProviderBusinessNotFound
	}

	costPerPhone := providerInfo.CostPerSMS
	if providerBusinessType.Price != nil && *providerBusinessType.Price > 0 {
		costPerPhone = *providerBusinessType.Price
	}
	if costPerPhone <= 0 {
		costPerPhone = 0.01
	}

	// 6. 判断 COST * count 是否超过了客户端的余额
	totalCost := costPerPhone * float64(count)
	balance, err := s.transactionRepo.GetBalance(ctx, customerID)
	if err != nil {
		s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone", fmt.Sprintf("Balance check error: %v", err))
		return nil, common.CodeInternalError
	}

	if balance < totalCost {
		s.apiLogger.LogInsufficientBalance(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("Insufficient balance: %.2f < %.2f", balance, totalCost))
		return nil, common.CodeInsufficientBalance
	}

	// 7. 调用运营商的接口，获取手机号（循环count次）
	results := make([]*GetPhoneResult, 0, count)
	successCount := 0
	failedCount := 0
	currentBalance := balance

	for i := 0; i < count; i++ {
		reserveNotes := fmt.Sprintf("预冻结手机号: 业务=%s, card=%s, attempt=%d/%d",
			customerConfig.BusinessName, cardType, i+1, count)

		reserveTx, err := s.transactionService.ReserveFunds(ctx, customerID, costPerPhone, 0, reserveNotes)
		if err != nil {
			if errors.Is(err, ErrInsufficientBalance) {
				s.apiLogger.LogInsufficientBalance(ctx, customerID, "/api/phone/get_phone",
					fmt.Sprintf("Insufficient balance during reservation (attempt %d/%d)", i+1, count))
				if successCount == 0 {
					return nil, common.CodeInsufficientBalance
				}
				break
			}

			s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone",
				fmt.Sprintf("Reserve funds error: %v", err))
			if successCount == 0 {
				return nil, common.CodeInternalError
			}
			break
		}

		if reserveTx != nil && reserveTx.BalanceAfter != nil {
			currentBalance = float64(*reserveTx.BalanceAfter)
		}

		// 调用运营商接口获取手机号
		global.LogInfo("调用运营商接口获取手机号",
			zap.Int64("customer_id", customerID),
			zap.String("provider_code", *selectedMapping.ProviderCode),
			zap.String("provider_id", providerInfo.ID),
			zap.String("business_code", *selectedMapping.BusinessCode),
			zap.String("card_type", cardType),
			zap.Int("attempt", i+1),
			zap.Int("total_count", count))

		phoneResponse, err := providerInstance.GetPhone(ctx, *selectedMapping.BusinessCode, cardType)
		if err != nil {
			failedCount++
			errorMsg := fmt.Sprintf("服务商获取手机号失败: provider=%s, business_code=%s, card_type=%s, 错误=%v",
				*selectedMapping.ProviderCode, *selectedMapping.BusinessCode, cardType, err)
			global.LogError("服务商获取手机号失败",
				zap.Int64("customer_id", customerID),
				zap.String("provider_code", *selectedMapping.ProviderCode),
				zap.String("provider_id", providerInfo.ID),
				zap.String("business_code", *selectedMapping.BusinessCode),
				zap.String("card_type", cardType),
				zap.Int("attempt", i+1),
				zap.String("error_type", fmt.Sprintf("%T", err)),
				zap.Error(err))
			s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone", errorMsg)

			releaseNotes := fmt.Sprintf("释放预冻结（服务商失败）: provider=%s, error=%v",
				*selectedMapping.ProviderCode, err)
			releaseTx, releaseErr := s.transactionService.ReleaseReservedFunds(ctx, customerID, costPerPhone, 0, releaseNotes)
			if releaseErr != nil {
				global.LogError("释放预冻结失败",
					zap.Int64("customer_id", customerID),
					zap.Error(releaseErr))
			} else if releaseTx != nil && releaseTx.BalanceAfter != nil {
				currentBalance = float64(*releaseTx.BalanceAfter)
			}
			continue
		}

		global.LogInfo("运营商接口调用成功",
			zap.Int64("customer_id", customerID),
			zap.String("provider_code", *selectedMapping.ProviderCode),
			zap.String("phone_number", phoneResponse.PhoneNumber),
			zap.String("ext_id", phoneResponse.ExtId),
			zap.Float64("cost", phoneResponse.Cost))

		// 8. 如果获取到手机号成功，记录到数据库
		providerCost := costPerPhone
		merchantFee := providerCost
		profit := 0.0
		pendingStatus := "pending"
		remark := fmt.Sprintf("业务：%s (%s)，提供商：%s，子业务：%s",
			customerConfig.BusinessName, customerConfig.BusinessCode, *selectedMapping.ProviderCode, *selectedMapping.BusinessCode)

		var providerDBID *int64
		if s.providerRepo != nil && selectedMapping.ProviderCode != nil {
			if providerRecord, err := s.providerRepo.FindByCode(ctx, *selectedMapping.ProviderCode); err == nil {
				id := int64(providerRecord.ID)
				providerDBID = &id
			}
		}

		assignment := &domain.PhoneAssignment{
			BusinessName:           customerConfig.BusinessName,
			BusinessCode:           customerConfig.BusinessCode,
			MerchantNo:             merchantNo,
			MerchantName:           merchantName,
			PhoneNumber:            &phoneResponse.PhoneNumber,
			Status:                 &pendingStatus,
			ProviderCost:           &providerCost,
			MerchantFee:            &merchantFee,
			Profit:                 &profit,
			Remark:                 &remark,
			CustomerID:             &customerID,
			PlatformBusinessTypeID: &customerConfig.PlatformBusinessTypeID,
			CreatedAt:              time.Now(),
			UpdatedAt:              time.Now(),
		}

		// 保存 extId（如果 Provider 返回了 extId）
		if phoneResponse.ExtId != "" {
			assignment.ExtId = &phoneResponse.ExtId
		}

		// 如果有ProviderID，设置它
		if providerDBID != nil {
			assignment.ProviderID = providerDBID
		} else if selectedMapping.ProviderBusinessTypeID != nil {
			providerID := int64(providerBusinessType.ProviderID)
			assignment.ProviderID = &providerID
		}

		if err := s.assignmentRepo.Create(ctx, nil, assignment); err != nil {
			s.transactionService.ReleaseReservedFunds(ctx, customerID, costPerPhone, 0,
				fmt.Sprintf("释放预冻结（保存号码失败）: %v", err))
			s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone",
				fmt.Sprintf("Assignment creation error: %v", err))
			return nil, common.CodeInternalError
		}

		if reserveTx != nil {
			s.attachTransactionReference(ctx, reserveTx, assignment.ID)
		}

		commitNotes := fmt.Sprintf("冻结转消费: assignment_id=%d", assignment.ID)
		if _, err := s.transactionService.CommitReservedFunds(ctx, customerID, providerCost, assignment.ID, commitNotes); err != nil {
			global.LogError("冻结转实扣失败",
				zap.Int64("customer_id", customerID),
				zap.Int64("assignment_id", assignment.ID),
				zap.Error(err))
			s.transactionService.ReleaseReservedFunds(ctx, customerID, providerCost, assignment.ID,
				fmt.Sprintf("释放预冻结（转消费失败）: %v", err))
			failStatus := "failed"
			assignment.Status = &failStatus
			assignment.UpdatedAt = time.Now()
			_ = s.assignmentRepo.Update(ctx, nil, assignment)
			s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone",
				fmt.Sprintf("Commit frozen funds error: %v", err))
			return nil, common.CodeInternalError
		}

		// 添加到结果
		results = append(results, &GetPhoneResult{
			PhoneNumber: phoneResponse.PhoneNumber,
			CountryCode: phoneResponse.CountryCode,
			Cost:        providerCost,
			ValidUntil:  phoneResponse.ValidUntil,
			ProviderID:  phoneResponse.ProviderID,
			Balance:     currentBalance,
		})

		successCount++
	}

	// 如果有失败的，记录日志
	if failedCount > 0 {
		s.apiLogger.LogError(ctx, customerID, "/api/phone/get_phone",
			fmt.Sprintf("Partial success: %d succeeded, %d failed", successCount, failedCount), 200)
	}

	// 如果全部失败，返回错误
	if successCount == 0 {
		errorMsg := fmt.Sprintf("所有获取手机号尝试都失败: 尝试次数=%d, 失败次数=%d, provider=%s, business_code=%s",
			count, failedCount, *selectedMapping.ProviderCode, *selectedMapping.BusinessCode)
		global.LogError("获取手机号失败：所有尝试都失败",
			zap.Int64("customer_id", customerID),
			zap.String("business_type", businessType),
			zap.String("card_type", cardType),
			zap.Int("count", count),
			zap.Int("failed_count", failedCount),
			zap.String("provider_code", *selectedMapping.ProviderCode),
			zap.String("business_code", *selectedMapping.BusinessCode))
		s.apiLogger.LogInternalError(ctx, customerID, "/api/phone/get_phone", errorMsg)
		return nil, common.CodeThirdPartyError
	}

	s.apiLogger.LogSuccess(ctx, customerID, "/api/phone/get_phone",
		fmt.Sprintf("Successfully got %d phone numbers, failed %d", successCount, failedCount))

	return results, common.CodeSuccess
}

func (s *PhoneService) attachTransactionReference(ctx context.Context, txRecord *domain.Transaction, referenceID int64) {
	if txRecord == nil || txRecord.ID == 0 {
		return
	}

	txRecord.ReferenceID = &referenceID
	if err := s.transactionRepo.Update(ctx, txRecord); err != nil {
		global.LogError("更新交易引用失败",
			zap.Int64("transaction_id", txRecord.ID),
			zap.Int64("reference_id", referenceID),
			zap.Error(err))
	}
}

// selectByWeight 根据权重随机选择一个映射
func (s *PhoneService) selectByWeight(mappings []*domain.PlatformProviderBusinessMapping) *domain.PlatformProviderBusinessMapping {
	if len(mappings) == 0 {
		return nil
	}

	// 计算总权重
	totalWeight := 0
	for _, m := range mappings {
		if m.Weight != nil && m.Status != nil && *m.Status {
			totalWeight += *m.Weight
		}
	}

	if totalWeight == 0 {
		// 如果所有权重都是0，随机选择一个
		enabledMappings := make([]*domain.PlatformProviderBusinessMapping, 0)
		for _, m := range mappings {
			if m.Status != nil && *m.Status {
				enabledMappings = append(enabledMappings, m)
			}
		}
		if len(enabledMappings) == 0 {
			return nil
		}
		return enabledMappings[rand.Intn(len(enabledMappings))]
	}

	// 根据权重随机选择
	random := rand.Intn(totalWeight)
	currentWeight := 0
	for _, m := range mappings {
		if m.Weight != nil && m.Status != nil && *m.Status {
			currentWeight += *m.Weight
			if random < currentWeight {
				return m
			}
		}
	}

	// 如果没找到（理论上不应该发生），返回第一个启用的
	for _, m := range mappings {
		if m.Status != nil && *m.Status {
			return m
		}
	}

	return nil
}

// GetCode retrieves SMS verification code for a specific phone number
// 直接查询数据库，因为验证码是由定时器异步获取并更新到数据库的
// GetCode 获取验证码
// 此函数直接查询数据库，如果验证码还未获取到，返回等待状态，客户端需要再次请求。
func (s *PhoneService) GetCode(ctx context.Context, customerID int64, phoneNumber string) ([]*GetCodeResult, error) {

	// 查询数据库记录 - 直接通过手机号查询，然后验证customerID
	assignment, err := s.assignmentRepo.FindByPhone(ctx, nil, phoneNumber)
	if err != nil {
		s.apiLogger.LogNotFound(ctx, customerID, "/api/phone/get_code",
			fmt.Sprintf("Phone: %s not found", phoneNumber))
		return nil, common.ErrPhoneNumberNotFound
	}

	// 验证customerID是否匹配
	if assignment.CustomerID == nil || *assignment.CustomerID != customerID {
		s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_code",
			fmt.Sprintf("Phone: %s does not belong to customer", phoneNumber))
		return nil, fmt.Errorf("手机号不属于该客户")
	}

	// 检查状态是否为expired或failed
	if assignment.Status != nil {
		status := *assignment.Status
		if status == "expired" {
			s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_code",
				fmt.Sprintf("Phone: %s assignment expired", phoneNumber))
			return nil, common.ErrPhoneNumberExpired
		}
		if status == "failed" {
			s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_code",
				fmt.Sprintf("Phone: %s code fetch failed", phoneNumber))
			return nil, fmt.Errorf("验证码获取失败")
		}
	}

	// 检查是否过期（使用配置的超时时间，默认300秒）
	codeTimeout := 300 * time.Second // 默认5分钟，可以从配置读取
	expiryTime := assignment.CreatedAt.Add(codeTimeout)
	if time.Now().After(expiryTime) {
		s.apiLogger.LogBadRequest(ctx, customerID, "/api/phone/get_code",
			fmt.Sprintf("Phone: %s assignment expired", phoneNumber))
		return nil, common.ErrPhoneNumberExpired
	}

	// 如果已经有验证码，直接返回
	if assignment.VerificationCode != nil && *assignment.VerificationCode != "" {
		status := assignment.Status
		if status != nil && *status == "completed" {
			// 获取ProviderID（如果有）
			providerID := ""
			if assignment.ProviderID != nil {
				providerID = fmt.Sprintf("%d", *assignment.ProviderID)
			}

			s.apiLogger.LogSuccess(ctx, customerID, "/api/phone/get_code",
				fmt.Sprintf("Phone: %s, Code retrieved from database", phoneNumber))

			return []*GetCodeResult{
				{
					Code:       *assignment.VerificationCode,
					Message:    "验证码获取成功",
					ReceivedAt: assignment.UpdatedAt, // 使用更新时间作为接收时间
					ProviderID: providerID,
				},
			}, nil
		}
	}

	// 如果没有验证码，直接返回等待状态，客户端会再次请求
	// 返回空code，表示等待中
	providerID := ""
	if assignment.ProviderID != nil {
		providerID = fmt.Sprintf("%d", *assignment.ProviderID)
	}

	s.apiLogger.LogSuccess(ctx, customerID, "/api/phone/get_code",
		fmt.Sprintf("Phone: %s, Code waiting (client will retry)", phoneNumber))

	return []*GetCodeResult{
		{
			Code:       "", // 空code表示等待中
			Message:    "验证码获取中，请稍后重试",
			ReceivedAt: time.Now(),
			ProviderID: providerID,
		},
	}, nil
}
