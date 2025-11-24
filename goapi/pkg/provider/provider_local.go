package provider

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// LocalProvider 本地SMS运营商实现
// 实现了SMSProvider接口，从数据库读取配置，提供完整的SMS服务功能
type LocalProvider struct {
	info              *ProviderInfo
	healthy           bool
	phoneNumbersInUse map[string]*AssignedPhone // 已分配的手机号
	codesReceived     map[string]*ReceivedCode  // 收到的验证码
	mu                sync.RWMutex

	// 配置参数
	successRate       int                      // 成功率 (0-100)
	minDelayMs        int                      // 最小延迟(毫秒)
	maxDelayMs        int                      // 最大延迟(毫秒)
	supportedBusiness map[string]*BusinessRule // 支持的业务类型
	supportedCards    map[string]bool          // 支持的卡片类型
	countries         []string                 // 支持的国家
}

// AssignedPhone 已分配的手机号信息
type AssignedPhone struct {
	PhoneNumber  string
	CountryCode  string
	BusinessType string
	CardType     string
	AssignedAt   time.Time
	ExpiresAt    time.Time
	Cost         float64
}

// ReceivedCode 接收到的验证码信息
type ReceivedCode struct {
	Code         string
	PhoneNumber  string
	BusinessType string
	ReceivedAt   time.Time
	ExpiresAt    time.Time
}

// BusinessRule 业务类型规则
type BusinessRule struct {
	Name        string
	SuccessRate float32
	CostFactor  float32
	Delay       time.Duration
	Countries   []string
}

// NewLocalProvider 创建新的本地运营商实例
// businessTypes: 从数据库读取的业务类型列表，如果为nil则使用默认配置
func NewLocalProvider(id, name string, priority int, businessTypes []BusinessTypeConfig) *LocalProvider {
	return NewLocalProviderWithConfig(id, name, priority, LocalProviderConfig{
		SuccessRate:   95,
		MinDelayMs:    100,
		MaxDelayMs:    500,
		BusinessTypes: businessTypes,
	})
}

// BusinessTypeConfig 业务类型配置（从数据库读取）
type BusinessTypeConfig struct {
	BusinessCode string
	BusinessName string
	Price        float64
}

// LocalProviderConfig 本地运营商配置
type LocalProviderConfig struct {
	SuccessRate   int                  // 成功率 (0-100)
	MinDelayMs    int                  // 最小延迟
	MaxDelayMs    int                  // 最大延迟
	BusinessTypes []BusinessTypeConfig // 从数据库读取的业务类型列表
}

// NewLocalProviderWithConfig 使用自定义配置创建本地运营商
func NewLocalProviderWithConfig(id, name string, priority int, config LocalProviderConfig) *LocalProvider {
	// 验证配置参数
	if config.SuccessRate < 0 || config.SuccessRate > 100 {
		config.SuccessRate = 95
	}
	if config.MinDelayMs < 0 {
		config.MinDelayMs = 100
	}
	if config.MaxDelayMs < config.MinDelayMs {
		config.MaxDelayMs = config.MinDelayMs + 100
	}

	lp := &LocalProvider{
		info: &ProviderInfo{
			ID:                 id,
			Name:               name,
			Type:               "local",
			Priority:           priority,
			CostPerSMS:         0.01, // 1美分基础费用
			SupportedCountries: []string{"US", "CN", "UK", "DE", "FR"},
			RateLimit:          1000, // 每分钟1000个请求
			Timeout:            10 * time.Second,
			Metadata: map[string]string{
				"version":     "1.0",
				"environment": "local",
				"features":    "database-driven,configurable",
			},
		},
		healthy:           true,
		phoneNumbersInUse: make(map[string]*AssignedPhone),
		codesReceived:     make(map[string]*ReceivedCode),
		successRate:       config.SuccessRate,
		minDelayMs:        config.MinDelayMs,
		maxDelayMs:        config.MaxDelayMs,
		countries:         []string{"US", "CN", "UK", "DE", "FR"},
	}

	// 初始化支持的业务类型（从数据库读取或使用默认配置）
	lp.initSupportedBusiness(config.BusinessTypes)
	// 初始化支持的卡片类型
	lp.initSupportedCards()

	return lp
}

// initSupportedBusiness 初始化支持的业务类型
// 如果 businessTypes 不为空，从数据库读取；否则使用默认配置
func (lp *LocalProvider) initSupportedBusiness(businessTypes []BusinessTypeConfig) {
	lp.supportedBusiness = make(map[string]*BusinessRule)

	// 如果提供了从数据库读取的业务类型，使用它们
	if len(businessTypes) > 0 {
		for _, bt := range businessTypes {
			if bt.BusinessCode == "" {
				continue
			}
			// 根据业务类型设置默认规则
			rule := &BusinessRule{
				Name:        bt.BusinessName,
				SuccessRate: 0.95, // 默认成功率
				CostFactor:  1.0,   // 默认成本因子
				Delay:       2 * time.Second,
				Countries:   []string{"US", "CN", "UK", "DE", "FR"},
			}
			// 根据业务代码设置特定规则
			switch bt.BusinessCode {
			case "qq":
				rule.SuccessRate = 0.95
				rule.CostFactor = 1.0
				rule.Delay = 2 * time.Second
				rule.Countries = []string{"CN", "US"}
			case "wx", "wechat":
				rule.SuccessRate = 0.98
				rule.CostFactor = 1.2
				rule.Delay = 1 * time.Second
				rule.Countries = []string{"CN"}
			case "whatsapp":
				rule.SuccessRate = 0.90
				rule.CostFactor = 1.5
				rule.Delay = 3 * time.Second
				rule.Countries = []string{"US", "UK", "DE", "FR"}
			case "telegram":
				rule.SuccessRate = 0.92
				rule.CostFactor = 1.1
				rule.Delay = 2 * time.Second
				rule.Countries = []string{"US", "UK", "DE", "FR"}
			}
			lp.supportedBusiness[bt.BusinessCode] = rule
		}
	} else {
		// 使用默认配置（向后兼容）
		lp.supportedBusiness = map[string]*BusinessRule{
			"qq": {
				Name:        "QQ注册",
				SuccessRate: 0.95,
				CostFactor:  1.0,
				Delay:       2 * time.Second,
				Countries:   []string{"CN", "US"},
			},
			"wechat": {
				Name:        "微信注册",
				SuccessRate: 0.98,
				CostFactor:  1.2,
				Delay:       1 * time.Second,
				Countries:   []string{"CN"},
			},
			"whatsapp": {
				Name:        "WhatsApp注册",
				SuccessRate: 0.90,
				CostFactor:  1.5,
				Delay:       3 * time.Second,
				Countries:   []string{"US", "UK", "DE", "FR"},
			},
			"telegram": {
				Name:        "Telegram注册",
				SuccessRate: 0.92,
				CostFactor:  1.1,
				Delay:       2 * time.Second,
				Countries:   []string{"US", "UK", "DE", "FR"},
			},
			"test": {
				Name:        "测试服务",
				SuccessRate: 0.99,
				CostFactor:  0.5,
				Delay:       500 * time.Millisecond,
				Countries:   []string{"US", "CN", "UK", "DE", "FR"},
			},
		}
	}
}

// initSupportedCards 初始化支持的卡片类型
func (lp *LocalProvider) initSupportedCards() {
	lp.supportedCards = map[string]bool{
		"physical": true, // 支持物理卡
		"virtual":  true, // 支持虚拟卡
	}
}

// GetProviderInfo 返回运营商信息
func (lp *LocalProvider) GetProviderInfo() *ProviderInfo {
	return lp.info
}

// IsHealthy 检查运营商健康状态
func (lp *LocalProvider) IsHealthy(ctx context.Context) bool {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return lp.healthy
}

// SetHealthy 设置健康状态（主要用于测试）
func (lp *LocalProvider) SetHealthy(healthy bool) {
	lp.mu.Lock()
	defer lp.mu.Unlock()
	lp.healthy = healthy
}

// GetPhone 获取手机号码
func (lp *LocalProvider) GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error) {
	// 模拟网络延迟
	if err := lp.simulateLatency(ctx); err != nil {
		return nil, err
	}

	// 模拟失败
	if lp.simulateFailure() {
		return nil, NewProviderError("TEMPORARY_FAILURE", "模拟的临时失败")
	}

	// 验证业务类型
	businessRule, ok := lp.supportedBusiness[businessType]
	if !ok {
		return nil, NewProviderError("UNSUPPORTED_SERVICE",
			fmt.Sprintf("不支持的业务类型: %s", businessType))
	}

	// 验证卡片类型
	if !lp.supportedCards[cardType] {
		return nil, NewProviderError("UNSUPPORTED_CARD_TYPE",
			fmt.Sprintf("不支持的卡片类型: %s", cardType))
	}

	// 生成手机号
	phoneNumber := lp.generatePhoneNumber()
	countryCode := lp.selectCountryCode(businessRule.Countries)

	// 计算成本
	cost := 0.01 * float64(businessRule.CostFactor)

	// 记录已分配的手机号
	lp.mu.Lock()
	lp.phoneNumbersInUse[phoneNumber] = &AssignedPhone{
		PhoneNumber:  phoneNumber,
		CountryCode:  countryCode,
		BusinessType: businessType,
		CardType:     cardType,
		AssignedAt:   time.Now(),
		ExpiresAt:    time.Now().Add(30 * time.Minute),
		Cost:         cost,
	}
	lp.mu.Unlock()

	return &PhoneResponse{
		PhoneNumber: phoneNumber,
		CountryCode: countryCode,
		Cost:        cost,
		ValidUntil:  time.Now().Add(30 * time.Minute),
		ProviderID:  lp.info.ID,
	}, nil
}

// GetCode 获取验证码
// 注意：对于本地Provider，如果手机号不在内存中，我们仍然可以生成验证码
// 因为这是一个模拟实现，实际生产环境应该从数据库或其他持久化存储中恢复状态
// extId 参数：本地 Provider 不使用 extId，保留参数以兼容接口
func (lp *LocalProvider) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration, extId ...string) (*CodeResponse, error) {
	// 先检查内存中是否有该手机号
	lp.mu.RLock()
	assignedPhone, exists := lp.phoneNumbersInUse[phoneNumber]
	lp.mu.RUnlock()

	var businessType string
	var businessRule *BusinessRule

	if exists {
		// 检查是否过期
		if time.Now().After(assignedPhone.ExpiresAt) {
			lp.mu.Lock()
			delete(lp.phoneNumbersInUse, phoneNumber)
			lp.mu.Unlock()
			return nil, NewProviderError("PHONE_EXPIRED", "手机号已过期")
		}
		businessType = assignedPhone.BusinessType
		businessRule = lp.supportedBusiness[businessType]
	} else {
		// 如果手机号不在内存中，尝试从默认业务类型生成验证码
		// 这是一个fallback机制，用于处理定时器调用时手机号不在内存中的情况
		// 实际生产环境应该从数据库恢复状态
		businessType = "wx" // 默认业务类型，可以根据实际情况调整
		businessRule = lp.supportedBusiness[businessType]
		if businessRule == nil {
			// 如果默认业务类型也不存在，使用第一个可用的业务类型
			for bt, rule := range lp.supportedBusiness {
				businessType = bt
				businessRule = rule
				break
			}
		}
		if businessRule == nil {
			return nil, NewProviderError("INVALID_PHONE", "手机号未分配且无可用业务类型")
		}
	}

	// 模拟验证码接收延迟
	smsDelay := businessRule.Delay
	if timeout < smsDelay {
		smsDelay = timeout / 2 // 确保在超时前到达
	}

	// 创建带超时的上下文
	codeCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// 等待验证码到达（模拟）
	select {
	case <-codeCtx.Done():
		return nil, NewProviderError("CODE_TIMEOUT", "验证码获取超时")
	case <-time.After(smsDelay):
		// 生成验证码
		code := lp.generateVerificationCode()

		// 记录验证码
		lp.mu.Lock()
		lp.codesReceived[phoneNumber] = &ReceivedCode{
			Code:         code,
			PhoneNumber:  phoneNumber,
			BusinessType: businessType,
			ReceivedAt:   time.Now(),
			ExpiresAt:    time.Now().Add(5 * time.Minute),
		}
		lp.mu.Unlock()

		return &CodeResponse{
			Code:       code,
			Message:    fmt.Sprintf("验证码已接收 - %s", businessType),
			ReceivedAt: time.Now(),
			ProviderID: lp.info.ID,
		}, nil
	}
}

// ReleasePhone 释放手机号
func (lp *LocalProvider) ReleasePhone(ctx context.Context, phoneNumber string) error {
	lp.mu.Lock()
	defer lp.mu.Unlock()

	if _, exists := lp.phoneNumbersInUse[phoneNumber]; !exists {
		return NewProviderError("INVALID_PHONE", "手机号未分配")
	}

	delete(lp.phoneNumbersInUse, phoneNumber)
	delete(lp.codesReceived, phoneNumber)

	return nil
}

// GetAssignedPhones 获取已分配的手机号(用于调试和监控)
func (lp *LocalProvider) GetAssignedPhones() map[string]*AssignedPhone {
	lp.mu.RLock()
	defer lp.mu.RUnlock()

	// 返回副本
	result := make(map[string]*AssignedPhone)
	for k, v := range lp.phoneNumbersInUse {
		result[k] = v
	}
	return result
}

// GetReceivedCodes 获取接收到的验证码(用于调试和监控)
func (lp *LocalProvider) GetReceivedCodes() map[string]*ReceivedCode {
	lp.mu.RLock()
	defer lp.mu.RUnlock()

	// 返回副本
	result := make(map[string]*ReceivedCode)
	for k, v := range lp.codesReceived {
		result[k] = v
	}
	return result
}

// CleanupExpired 清理过期的手机号和验证码
func (lp *LocalProvider) CleanupExpired() {
	lp.mu.Lock()
	defer lp.mu.Unlock()

	now := time.Now()

	// 清理过期手机号
	for phone, assigned := range lp.phoneNumbersInUse {
		if now.After(assigned.ExpiresAt) {
			delete(lp.phoneNumbersInUse, phone)
		}
	}

	// 清理过期验证码
	for phone, code := range lp.codesReceived {
		if now.After(code.ExpiresAt) {
			delete(lp.codesReceived, phone)
		}
	}
}

// simulateLatency 模拟网络延迟
func (lp *LocalProvider) simulateLatency(ctx context.Context) error {
	if lp.minDelayMs == 0 && lp.maxDelayMs == 0 {
		return nil
	}

	delay := time.Duration(lp.minDelayMs+rand.Intn(lp.maxDelayMs-lp.minDelayMs+1)) * time.Millisecond
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		return nil
	}
}

// simulateFailure 模拟失败（根据成功率）
func (lp *LocalProvider) simulateFailure() bool {
	if lp.successRate >= 100 {
		return false
	}
	return rand.Intn(100) >= lp.successRate
}

// generatePhoneNumber 生成手机号
func (lp *LocalProvider) generatePhoneNumber() string {
	// 生成随机手机号（示例格式）
	prefixes := []string{"86", "1", "44", "33", "49"}
	prefix := prefixes[rand.Intn(len(prefixes))]
	number := fmt.Sprintf("%s%d%09d", prefix, rand.Intn(10), rand.Intn(1000000000))
	return number
}

// selectCountryCode 选择国家代码
func (lp *LocalProvider) selectCountryCode(countries []string) string {
	if len(countries) == 0 {
		return "US"
	}
	return countries[rand.Intn(len(countries))]
}

// generateVerificationCode 生成验证码
func (lp *LocalProvider) generateVerificationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// GetStats 获取统计信息
func (lp *LocalProvider) GetStats() map[string]interface{} {
	lp.mu.RLock()
	defer lp.mu.RUnlock()

	return map[string]interface{}{
		"provider_id":          lp.info.ID,
		"provider_name":        lp.info.Name,
		"phones_in_use":        len(lp.phoneNumbersInUse),
		"codes_received":       len(lp.codesReceived),
		"success_rate":         lp.successRate,
		"supported_business":   len(lp.supportedBusiness),
		"supported_countries":  len(lp.countries),
		"healthy":              lp.healthy,
	}
}

// GetSupportedBusiness 获取支持的业务类型
func (lp *LocalProvider) GetSupportedBusiness() map[string]*BusinessRule {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return lp.supportedBusiness
}

// GetSupportedCards 获取支持的卡片类型
func (lp *LocalProvider) GetSupportedCards() map[string]bool {
	return lp.supportedCards
}

