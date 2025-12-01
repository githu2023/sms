package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

var SmsCustomers = new(smsCustomers)

type smsCustomers struct{}

// CreateSmsCustomers 创建商户记录
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) CreateSmsCustomers(ctx context.Context, req *request.CreateSmsCustomersReq) (err error) {
	// 生成密码哈希
	passwordHash := utils.BcryptHash(req.Password)

	// 生成API密钥（如果未提供）
	var apiSecretKey string
	if req.ApiSecretKey != nil && *req.ApiSecretKey != "" {
		apiSecretKey = *req.ApiSecretKey
	} else {
		// 生成32字节的随机API密钥
		bytes := make([]byte, 32)
		if _, err := rand.Read(bytes); err != nil {
			return err
		}
		apiSecretKey = hex.EncodeToString(bytes)
	}

	// 创建商户模型
	smsCustomers := &model.SmsCustomers{
		MerchantName:   req.MerchantName,
		MerchantNo:     req.MerchantNo,
		Username:       req.Username,
		Email:          req.Email,
		PasswordHash:   &passwordHash,
		ApiSecretKey:   &apiSecretKey,
		Balance:        req.Balance,
		ParentID:       req.ParentID,
		Status:         req.Status,
		RegistrationIp: req.RegistrationIp,
		Remark:         req.Remark,
	}

	err = global.GVA_DB.Create(smsCustomers).Error
	return err
}

// DeleteSmsCustomers 删除商户记录
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) DeleteSmsCustomers(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsCustomers{}, "id = ?", ID).Error
	return err
}

// DeleteSmsCustomersByIds 批量删除商户记录
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) DeleteSmsCustomersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsCustomers{}, "id in ?", IDs).Error
	return err
}

// UpdateSmsCustomers 更新商户记录
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) UpdateSmsCustomers(ctx context.Context, smsCustomers model.SmsCustomers) (err error) {
	err = global.GVA_DB.Model(&model.SmsCustomers{}).Where("id = ?", smsCustomers.ID).Updates(&smsCustomers).Error
	return err
}

// GetSmsCustomers 根据ID获取商户记录
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) GetSmsCustomers(ctx context.Context, ID string) (smsCustomers model.SmsCustomers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsCustomers).Error
	return
}

// GetSmsCustomersInfoList 分页获取商户记录
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) GetSmsCustomersInfoList(ctx context.Context, info request.SmsCustomersSearch) (list []model.SmsCustomers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmsCustomers{})
	var smsCustomerss []model.SmsCustomers
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Username != nil && *info.Username != "" {
		db = db.Where("username LIKE ?", "%"+*info.Username+"%")
	}

	if info.MerchantNo != "" {
		db = db.Where("merchant_no LIKE ?", "%"+info.MerchantNo+"%")
	}

	if info.MerchantName != "" {
		db = db.Where("merchant_name LIKE ?", "%"+info.MerchantName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&smsCustomerss).Error
	return smsCustomerss, total, err
}

func (s *smsCustomers) GetSmsCustomersPublic(ctx context.Context) {

}

// CreditDebitSmsCustomers 充值/扣费
// Author [yourname](https://github.com/yourname)
func (s *smsCustomers) CreditDebitSmsCustomers(ctx context.Context, req *request.CreditDebitSmsCustomersReq) error {
	// 开始事务
	tx := global.GVA_DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 先获取当前客户信息
	var customer model.SmsCustomers
	if err := tx.Where("id = ?", req.CustomerId).First(&customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 初始化余额和冻结金额
	currentBalance := float64(0)
	if customer.Balance != nil {
		currentBalance = *customer.Balance
	}
	currentFrozenAmount := float64(0)
	if customer.FrozenAmount != nil {
		currentFrozenAmount = *customer.FrozenAmount
	}

	// 2. 根据交易类型处理余额和冻结金额
	var amount float64
	var newBalance, newFrozenAmount float64

	switch req.Type {
	case request.TransactionTypeCredit, request.TransactionTypeAddPoints:
		// 充值、上分：增加余额
		amount = req.Amount
		newBalance = currentBalance + amount
		newFrozenAmount = currentFrozenAmount

	case request.TransactionTypePullNumber, request.TransactionTypePullRollback, request.TransactionTypeDeductPoints:
		// 拉号码、拉号-回退、下分：减少余额
		if currentBalance < req.Amount {
			tx.Rollback()
			return errors.New("余额不足")
		}
		amount = -req.Amount
		newBalance = currentBalance + amount
		newFrozenAmount = currentFrozenAmount

	case request.TransactionTypeFreeze:
		// 冻结金额：从余额转到冻结金额
		if currentBalance < req.Amount {
			tx.Rollback()
			return errors.New("可用余额不足")
		}
		amount = -req.Amount // 余额减少
		newBalance = currentBalance - req.Amount
		newFrozenAmount = currentFrozenAmount + req.Amount

	case request.TransactionTypeUnfreeze:
		// 冻结金额返回：从冻结金额返回到余额
		if currentFrozenAmount < req.Amount {
			tx.Rollback()
			return errors.New("冻结金额不足")
		}
		amount = req.Amount // 余额增加
		newBalance = currentBalance + req.Amount
		newFrozenAmount = currentFrozenAmount - req.Amount

	default:
		tx.Rollback()
		return errors.New("无效的交易类型")
	}

	// 3. 更新客户余额和冻结金额
	updateData := map[string]interface{}{
		"balance":       newBalance,
		"frozen_amount": newFrozenAmount,
	}
	if err := tx.Model(&customer).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 4. 创建交易记录
	transactionType := req.Type
	transaction := &model.SmsTransactions{
		CustomerId:    &req.CustomerId,
		Amount:        &amount,
		BalanceBefore: &currentBalance,
		BalanceAfter:  &newBalance,
		Type:          &transactionType,
		Notes:         req.Notes,
	}

	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 记录规范化日志
	typeText, ok := request.TransactionTypeText[transactionType]
	if !ok {
		typeText = fmt.Sprintf("未知类型(%s)", transactionType)
	}
	global.GVA_LOG.Info(
		"商户余额变动",
		zap.Int64("customerId", req.CustomerId),
		zap.String("transactionType", typeText),
		zap.Float64("amountChange", amount),
		zap.Float64("balanceBefore", currentBalance),
		zap.Float64("balanceAfter", newBalance),
		zap.Float64("frozenAmountBefore", currentFrozenAmount),
		zap.Float64("frozenAmountAfter", newFrozenAmount),
	)

	// 提交事务
	return tx.Commit().Error
}

// ConfigureBusiness 配置商户业务
func (s *smsCustomers) ConfigureBusiness(ctx context.Context, req *request.ConfigureBusinessReq) error {
	tx := global.GVA_DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 获取当前已配置的业务ID列表
	var existingConfigs []model.SmsCustomerBusinessConfig
	if err := tx.Where("customer_id = ?", req.CustomerID).Find(&existingConfigs).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建已配置业务的映射表（platformBusinessTypeId -> config）
	existingMap := make(map[int64]*model.SmsCustomerBusinessConfig)
	for i := range existingConfigs {
		existingMap[existingConfigs[i].PlatformBusinessTypeID] = &existingConfigs[i]
	}

	// 创建新配置的映射表
	newConfigMap := make(map[int64]bool)
	for _, item := range req.BusinessConfig {
		newConfigMap[item.PlatformBusinessTypeID] = true
	}

	// 处理每个业务配置：更新已存在的，插入新的
	for _, item := range req.BusinessConfig {
		if existing, found := existingMap[item.PlatformBusinessTypeID]; found {
			// 已存在，更新
			updates := map[string]interface{}{
				"business_code": item.BusinessCode,
				"business_name": item.BusinessName,
				"cost":          item.Cost,
				"weight":        item.Weight,
				"status":        item.Status,
			}
			if err := tx.Model(existing).Updates(updates).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// 不存在，插入
			config := &model.SmsCustomerBusinessConfig{
				CustomerID:             req.CustomerID,
				PlatformBusinessTypeID: item.PlatformBusinessTypeID,
				BusinessCode:           item.BusinessCode,
				BusinessName:           item.BusinessName,
				Cost:                   item.Cost,
				Weight:                 item.Weight,
				Status:                 item.Status,
			}
			if err := tx.Create(config).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	global.GVA_LOG.Info(
		"商户业务配置更新",
		zap.Int64("customerId", req.CustomerID),
		zap.Int("configCount", len(req.BusinessConfig)),
	)

	return tx.Commit().Error
}

// AdjustFrozenAmount 调整冻结金额
func (s *smsCustomers) AdjustFrozenAmount(ctx context.Context, req *request.AdjustFrozenAmountReq) error {
	var customer model.SmsCustomers
	if err := global.GVA_DB.WithContext(ctx).Where("id = ?", req.ID).First(&customer).Error; err != nil {
		return err
	}

	oldFrozenAmount := float64(0)
	if customer.FrozenAmount != nil {
		oldFrozenAmount = *customer.FrozenAmount
	}

	if err := global.GVA_DB.WithContext(ctx).Model(&customer).Update("frozen_amount", req.FrozenAmount).Error; err != nil {
		return err
	}

	global.GVA_LOG.Info(
		"商户冻结金额调整",
		zap.Uint("customerId", req.ID),
		zap.Float64("oldFrozenAmount", oldFrozenAmount),
		zap.Float64("newFrozenAmount", req.FrozenAmount),
		zap.String("remark", req.Remark),
	)

	return nil
}
