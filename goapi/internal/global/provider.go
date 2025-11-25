package global

import (
	"context"
	"errors"
	"fmt"
	"sms-platform/goapi/internal/domain"
	"sms-platform/goapi/pkg/provider"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ProviderRepository 使用domain中定义的接口，避免重复定义
type ProviderRepository = domain.ProviderRepository

// ProviderManager 全局运营商管理器
// 负责管理所有运营商实例，提供根据CODE获取运营商对象的功能
type ProviderManager struct {
	providers    map[string]provider.SMSProvider // key: provider code
	mutex        sync.RWMutex
	db           *gorm.DB
	providerRepo ProviderRepository
}

var (
	providerManager     *ProviderManager
	providerManagerOnce sync.Once
)

// InitProviderManager 初始化全局运营商管理器
// 从数据库加载所有启用的运营商，并注册到管理器中
func InitProviderManager(db *gorm.DB, providerRepo ProviderRepository, providerBusinessTypeRepo domain.ProviderBusinessTypeRepository) error {
	providerManagerOnce.Do(func() {
		providerManager = &ProviderManager{
			providers:    make(map[string]provider.SMSProvider),
			db:           db,
			providerRepo: providerRepo,
		}
	})

	// 从数据库加载所有启用的运营商
	ctx := context.Background()
	providers, err := providerRepo.FindAll(ctx)
	if err != nil {
		zap.S().Errorf("Failed to load providers from database: %v", err)
		return fmt.Errorf("failed to load providers from database: %w", err)
	}

	zap.S().Infof("Loading providers from database, total: %d", len(providers))

	// 注册每个运营商
	registeredCount := 0
	for _, p := range providers {
		// 允许 is_enabled 为 NULL 时也注册，或者 is_enabled 为 true
		if p.Code != nil && (p.IsEnabled == nil || *p.IsEnabled) {
			providerCode := *p.Code
			zap.S().Infof("Processing provider: code=%s, name=%s, api_gateway=%s, is_enabled=%v",
				providerCode,
				getStringValue(p.Name, "Unknown"),
				getStringValue(p.APIGateway, "N/A"),
				p.IsEnabled != nil && *p.IsEnabled)
			// 从数据库读取该运营商支持的业务类型
			businessTypes, err := providerBusinessTypeRepo.FindByProviderCode(ctx, providerCode)
			if err != nil {
				// 如果查询失败，记录日志但继续使用默认配置
				zap.S().Warnf("Failed to load business types for provider %s: %v", providerCode, err)
			} else {
				zap.S().Infof("Loaded %d business types for provider %s", len(businessTypes), providerCode)
			}

			// 转换为 provider.BusinessTypeConfig
			var businessTypeConfigs []provider.BusinessTypeConfig
			for _, bt := range businessTypes {
				if bt.BusinessCode != nil && (bt.Status == nil || *bt.Status) {
					businessName := ""
					if bt.BusinessName != nil {
						businessName = *bt.BusinessName
					}
					price := 0.0
					if bt.Price != nil {
						price = *bt.Price
					}
					businessTypeConfigs = append(businessTypeConfigs, provider.BusinessTypeConfig{
						BusinessCode: *bt.BusinessCode,
						BusinessName: businessName,
						Price:        price,
					})
				}
			}

			// 根据运营商类型创建对应的provider实例
			var providerInstance provider.SMSProvider

			// 判断是否为 MQTT Provider（根据 API 网关地址或 code 判断）
			isMQTTProvider := false
			if p.APIGateway != nil {
				apiGateway := *p.APIGateway
				// 如果 API 网关包含 "jczl70.com" 或 "mqtt/msg"，则认为是 MQTT Provider
				if contains(apiGateway, "jczl70.com") || contains(apiGateway, "mqtt/msg") {
					isMQTTProvider = true
					zap.S().Infof("Provider %s identified as MQTT provider (by API gateway: %s)", providerCode, apiGateway)
				}
			}
			// 或者根据 code 判断（如果 code 包含 "mqtt"）
			if contains(providerCode, "mqtt") {
				isMQTTProvider = true
				zap.S().Infof("Provider %s identified as MQTT provider (by code)", providerCode)
			}

			if isMQTTProvider {
				zap.S().Infof("Creating MQTT provider: code=%s", providerCode)
				// 创建 MQTTProvider
				// 从数据库字段读取配置：
				// - APIGateway: api_gateway 字段
				// - ProviderID: merchant_id 字段（作为 id 参数）
				// - ProviderKey: merchant_key 字段（作为 key 参数）
				apiGateway := ""
				if p.APIGateway != nil {
					apiGateway = *p.APIGateway
				}
				providerID := ""
				if p.MerchantID != nil {
					providerID = sanitizeCredential(*p.MerchantID)
				}
				providerKey := ""
				if p.MerchantKey != nil {
					providerKey = sanitizeCredential(*p.MerchantKey)
				}

				// 验证必要配置
				if apiGateway == "" || providerID == "" || providerKey == "" {
					zap.S().Warnf("MQTT provider %s missing required config, skipping",
						zap.String("provider_code", providerCode),
						zap.String("api_gateway", apiGateway),
						zap.String("provider_id", providerID),
						zap.Bool("has_key", providerKey != ""))
					continue
				}

				// 获取成本（从业务类型配置中获取，或使用默认值）
				costPerSMS := 1.0
				if len(businessTypeConfigs) > 0 && businessTypeConfigs[0].Price > 0 {
					costPerSMS = businessTypeConfigs[0].Price
				}

				providerInstance = provider.NewMQTTProvider(provider.MQTTConfig{
					ID:                 providerCode,
					Name:               getStringValue(p.Name, "MQTT"),
					APIGateway:         apiGateway,
					ProviderID:         providerID,
					ProviderKey:        providerKey,
					Priority:           100,
					CostPerSMS:         costPerSMS,
					SupportedCountries: []string{"CN"},
					Timeout:            30 * time.Second,
				})
				zap.S().Infof("MQTT provider created successfully: code=%s, api_gateway=%s, business_types=%d",
					providerCode, apiGateway, len(businessTypeConfigs))
			} else if providerCode == "bigbus666" || contains(providerCode, "bigbus") {
				zap.S().Infof("Creating BigBus666 provider: code=%s", providerCode)
				// 创建 BigBus666Provider
				// 从数据库字段读取配置：
				// - APIGateway: api_gateway 字段
				// - CustomerOutNumber: merchant_id 字段
				// - EncryptKey: merchant_key 字段
				// - ProjectName: extra_config JSON 字段中的 projectName
				apiGateway := ""
				if p.APIGateway != nil {
					apiGateway = *p.APIGateway
				}
				customerOutNumber := ""
				if p.MerchantID != nil {
					customerOutNumber = sanitizeCredential(*p.MerchantID)
				}
				encryptKey := ""
				if p.MerchantKey != nil {
					encryptKey = sanitizeCredential(*p.MerchantKey)
				}

				// 从 ExtraConfig 中读取特殊配置
				projectName := "hema" // 默认值
				if p.ExtraConfig != nil {
					if pn, ok := (*p.ExtraConfig)["projectName"].(string); ok && pn != "" {
						projectName = pn
					}
				}

				// 验证必要配置
				if apiGateway == "" || customerOutNumber == "" || encryptKey == "" {
					zap.S().Warnf("BigBus666 provider %s missing required config (api_gateway, merchant_id, or merchant_key), skipping", providerCode)
					continue
				}

				// 获取成本（从业务类型配置中获取，或使用默认值）
				costPerSMS := 1.0
				if len(businessTypeConfigs) > 0 && businessTypeConfigs[0].Price > 0 {
					costPerSMS = businessTypeConfigs[0].Price
				}

				providerInstance = provider.NewBigBus666Provider(provider.BigBus666Config{
					ID:                 providerCode,
					Name:               getStringValue(p.Name, "BigBus666"),
					APIGateway:         apiGateway,
					CustomerOutNumber:  customerOutNumber,
					EncryptKey:         encryptKey,
					ProjectName:        projectName, // 从 ExtraConfig 读取
					Priority:           100,
					CostPerSMS:         costPerSMS,
					SupportedCountries: []string{"CN"},
					Timeout:            30 * time.Second,
				})
			} else {
				// 使用本地Provider实现，从数据库读取配置
				zap.S().Infof("Creating Local provider: code=%s", providerCode)
				providerInstance = provider.NewLocalProvider(providerCode, getStringValue(p.Name, "Local"), 100, businessTypeConfigs)
			}

			if err := providerManager.RegisterProvider(providerCode, providerInstance); err != nil {
				zap.S().Errorf("Failed to register provider %s: %v", providerCode, err)
				continue
			}
			registeredCount++
			zap.S().Infof("Provider registered successfully: code=%s, type=%s",
				providerCode, providerInstance.GetProviderInfo().Type)
		} else {
			zap.S().Debugf("Skipping provider: code=%s, is_enabled=%v",
				getStringValue(p.Code, "N/A"),
				p.IsEnabled != nil && *p.IsEnabled)
		}
	}

	zap.S().Infof("Provider manager initialization completed: registered %d providers", registeredCount)
	return nil
}

// RegisterProvider 注册运营商
func (pm *ProviderManager) RegisterProvider(code string, providerInstance provider.SMSProvider) error {
	if code == "" {
		return errors.New("provider code cannot be empty")
	}
	if providerInstance == nil {
		return errors.New("provider instance cannot be nil")
	}

	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	pm.providers[code] = providerInstance
	return nil
}

// GetProviderByCode 根据运营商CODE获取运营商对象
// 任何地方都可以直接调用 global.GetProviderManager().GetProviderByCode("code")
func (pm *ProviderManager) GetProviderByCode(code string) (provider.SMSProvider, error) {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()

	providerInstance, exists := pm.providers[code]
	if !exists {
		zap.S().Warnf("Provider not found: code=%s, available providers: %v", code, pm.getProviderCodes())
		return nil, fmt.Errorf("provider with code %s not found", code)
	}

	zap.S().Debugf("Provider retrieved: code=%s, type=%s", code, providerInstance.GetProviderInfo().Type)
	return providerInstance, nil
}

// getProviderCodes 获取所有已注册的provider codes（用于日志）
func (pm *ProviderManager) getProviderCodes() []string {
	codes := make([]string, 0, len(pm.providers))
	for code := range pm.providers {
		codes = append(codes, code)
	}
	return codes
}

// GetAllProviders 获取所有注册的运营商
func (pm *ProviderManager) GetAllProviders() map[string]provider.SMSProvider {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()

	// 返回副本，避免外部修改
	result := make(map[string]provider.SMSProvider)
	for k, v := range pm.providers {
		result[k] = v
	}
	return result
}

// GetProviderFromDB 从数据库获取运营商信息并返回对应的provider对象
func (pm *ProviderManager) GetProviderFromDB(ctx context.Context, providerCode string) (provider.SMSProvider, *domain.Provider, error) {
	// 先从内存中查找
	if prov, err := pm.GetProviderByCode(providerCode); err == nil {
		// 同时返回数据库中的信息
		providers, err := pm.providerRepo.FindAll(ctx)
		if err == nil {
			for _, p := range providers {
				if p.Code != nil && *p.Code == providerCode {
					return prov, p, nil
				}
			}
		}
		return prov, nil, nil
	}

	// 如果内存中没有，从数据库加载
	providers, err := pm.providerRepo.FindAll(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load providers: %w", err)
	}

	for _, p := range providers {
		if p.Code != nil && *p.Code == providerCode {
			if p.IsEnabled != nil && *p.IsEnabled {
				// 创建provider实例并注册（暂时使用空业务类型列表，因为这里没有 providerBusinessTypeRepo）
				providerInstance := provider.NewLocalProvider(*p.Code, *p.Name, 100, nil)
				pm.RegisterProvider(providerCode, providerInstance)
				return providerInstance, p, nil
			}
			return nil, p, fmt.Errorf("provider %s is disabled", providerCode)
		}
	}

	return nil, nil, fmt.Errorf("provider with code %s not found", providerCode)
}

// GetProviderManager 获取全局运营商管理器实例
// 任何地方都可以直接调用 global.GetProviderManager() 获取管理器
func GetProviderManager() *ProviderManager {
	if providerManager == nil {
		panic("ProviderManager not initialized. Call InitProviderManager first.")
	}
	return providerManager
}

// contains 检查字符串是否包含子字符串（不区分大小写）
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// getStringValue 安全获取字符串指针的值
func getStringValue(ptr *string, defaultValue string) string {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func sanitizeCredential(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "’", "")
	value = strings.ReplaceAll(value, "'", "")
	return value
}
