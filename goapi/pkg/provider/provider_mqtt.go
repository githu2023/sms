package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"go.uber.org/zap"
)

// MQTTProvider MQTT运营商实现
// API文档: http://szb.jczl70.com:6086/mqtt/msg/
type MQTTProvider struct {
	info       *ProviderInfo
	healthy    bool
	mu         sync.RWMutex
	apiGateway string // API网关地址 (例如: http://szb.jczl70.com:6086)
	id         string // 运营商ID (从 merchant_id 读取)
	key        string // 运营商KEY (从 merchant_key 读取)
	httpClient *http.Client
}

// MQTTConfig MQTT配置
type MQTTConfig struct {
	ID                 string // Provider ID
	Name               string
	APIGateway         string // API网关地址
	ProviderID         string // 运营商ID (从 merchant_id 读取)
	ProviderKey        string // 运营商KEY (从 merchant_key 读取)
	Priority           int
	CostPerSMS         float64
	SupportedCountries []string
	Timeout            time.Duration
}

// NewMQTTProvider 创建MQTT运营商实例
func NewMQTTProvider(config MQTTConfig) *MQTTProvider {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &MQTTProvider{
		info: &ProviderInfo{
			ID:                 config.ID,
			Name:               config.Name,
			Type:               "http",
			Priority:           config.Priority,
			CostPerSMS:         config.CostPerSMS,
			SupportedCountries: config.SupportedCountries,
			Timeout:            config.Timeout,
			Metadata: map[string]string{
				"api_gateway": config.APIGateway,
				"provider_id": config.ProviderID,
			},
		},
		healthy:    true,
		apiGateway: config.APIGateway,
		id:         config.ProviderID,
		key:        config.ProviderKey,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// GetProviderInfo 返回运营商信息
func (p *MQTTProvider) GetProviderInfo() *ProviderInfo {
	return p.info
}

// IsHealthy 检查运营商是否健康
func (p *MQTTProvider) IsHealthy(ctx context.Context) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.healthy
}

// SetHealthy 设置健康状态
func (p *MQTTProvider) SetHealthy(healthy bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.healthy = healthy
}

// GetPhone 获取手机号码
// API: http://szb.jczl70.com:6086/mqtt/msg/getNumber?id=%d&key=%s
// 返回: {"number":"1888888888","extId":"2025111215435966402734","id":1}
func (p *MQTTProvider) GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error) {
	if !p.IsHealthy(ctx) {
		return nil, ErrProviderUnavailable
	}

	zap.S().Infof("[MQTTProvider] ========== 获取手机号 - 开始 ==========")
	zap.S().Infof("[MQTTProvider] 请求参数: businessType=%s, cardType=%s, id=%s", businessType, cardType, p.id)

	// 构建URL
	apiURL := fmt.Sprintf("%s/mqtt/msg/getNumber?id=%s&key=%s", p.apiGateway, url.QueryEscape(businessType), url.QueryEscape(p.key))
	zap.S().Infof("[MQTTProvider] 请求URL: %s", apiURL)

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		zap.S().Errorf("[MQTTProvider] 创建请求失败: %v", err)
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.1")

	// 发送请求
	zap.S().Infof("[MQTTProvider] 发送GET请求到: %s", apiURL)
	resp, err := p.httpClient.Do(req)
	if err != nil {
		zap.S().Errorf("[MQTTProvider] 请求失败: %v", err)
		zap.S().Errorf("[MQTTProvider] ========== 获取手机号 - 失败 ==========")
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	zap.S().Infof("[MQTTProvider] HTTP状态码: %d", resp.StatusCode)

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.S().Errorf("[MQTTProvider] 读取响应失败: %v", err)
		zap.S().Errorf("[MQTTProvider] ========== 获取手机号 - 失败 ==========")
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	zap.S().Infof("[MQTTProvider] 原始响应数据: %s", string(respBody))

	if resp.StatusCode != http.StatusOK {
		zap.S().Warnf("[MQTTProvider] HTTP错误: 状态码=%d, 响应内容=%s", resp.StatusCode, string(respBody))
		zap.S().Warnf("[MQTTProvider] ========== 获取手机号 - 失败 ==========")
		return nil, fmt.Errorf("HTTP错误: %d, 响应: %s", resp.StatusCode, string(respBody))
	}

	// 解析响应
	// 注意：运营商可能返回成功响应 {"number":"...","extId":"...","id":1}
	// 也可能返回错误响应 {"id":-1,"msg":"error key"}
	var apiResponse struct {
		Number string `json:"number"`
		ExtId  string `json:"extId"`
		ID     int    `json:"id"`
		Msg    string `json:"msg"` // 错误消息字段
	}

	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		zap.S().Errorf("[MQTTProvider] 解析响应失败: error=%v, 响应内容=%s", err, string(respBody))
		zap.S().Errorf("[MQTTProvider] ========== 获取手机号 - 失败 ==========")
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	// 先检查是否是错误响应（id < 0 或 msg 不为空）
	if apiResponse.ID < 0 || apiResponse.Msg != "" {
		errorMsg := apiResponse.Msg
		if errorMsg == "" {
			errorMsg = fmt.Sprintf("运营商返回错误，id=%d", apiResponse.ID)
		}
		zap.S().Errorf("[MQTTProvider] ========== 获取手机号 - 运营商返回错误 ==========")
		zap.S().Errorf("[MQTTProvider] 错误详情: id=%d, msg=%s", apiResponse.ID, errorMsg)
		zap.S().Errorf("[MQTTProvider] 完整响应数据(JSON): %s", string(respBody))
		zap.S().Errorf("[MQTTProvider] ========== 错误信息结束 ==========")
		return nil, NewProviderError("API_ERROR", errorMsg)
	}

	// 检查成功响应是否包含必要字段
	if apiResponse.Number == "" || apiResponse.ExtId == "" {
		zap.S().Warnf("[MQTTProvider] 返回数据不完整: number=%s, extId=%s, id=%d", apiResponse.Number, apiResponse.ExtId, apiResponse.ID)
		zap.S().Warnf("[MQTTProvider] 完整响应数据(JSON): %s", string(respBody))
		zap.S().Warnf("[MQTTProvider] ========== 获取手机号 - 失败 ==========")
		return nil, NewProviderError("INVALID_RESPONSE", "返回数据不完整")
	}

	zap.S().Infof("[MQTTProvider] 获取手机号成功: phone=%s, ext_id=%s", apiResponse.Number, apiResponse.ExtId)
	zap.S().Infof("[MQTTProvider] ========== 获取手机号 - 成功 ==========")

	return &PhoneResponse{
		PhoneNumber: apiResponse.Number,
		CountryCode: "CN", // 默认中国
		Cost:        p.info.CostPerSMS,
		ValidUntil:  time.Now().Add(30 * time.Minute), // 默认30分钟有效期
		ProviderID:  p.info.ID,
		ExtId:       apiResponse.ExtId, // 保存 extId，用于后续获取验证码和释放手机号
	}, nil
}

// GetCode 获取验证码
// API: http://szb.jczl70.com:6086/mqtt/msg/getCode?extId=%s
// 返回: {"code":0,"message":"成功","data":{"receiveStatus":1,"message":"456039"}}
// extId 参数：如果提供了 extId，直接使用；否则从 phoneNumber 查找（需要从数据库读取）
func (p *MQTTProvider) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration, extId ...string) (*CodeResponse, error) {
	if !p.IsHealthy(ctx) {
		return nil, ErrProviderUnavailable
	}

	var extIdValue string
	// 如果提供了 extId 参数，直接使用
	if len(extId) > 0 && extId[0] != "" {
		extIdValue = extId[0]
	} else {
		// 如果没有提供 extId，需要从数据库读取（这里无法直接访问数据库）
		// 所以返回错误，提示需要提供 extId
		zap.S().Errorf("[MQTTProvider] 获取验证码失败: 手机号 %s 未提供 extId 参数", phoneNumber)
		return nil, NewProviderError("INVALID_PHONE", fmt.Sprintf("手机号 %s 需要提供 extId 参数", phoneNumber))
	}

	zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 开始 ==========")
	zap.S().Infof("[MQTTProvider] 请求参数: phoneNumber=%s, extId=%s, timeout=%v", phoneNumber, extIdValue, timeout)

	// 创建带超时的上下文（单次请求）
	codeCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// 构建URL - 单次查询，不轮询
	apiURL := fmt.Sprintf("%s/mqtt/msg/getCode?extId=%s", p.apiGateway, url.QueryEscape(extIdValue))
	zap.S().Infof("[MQTTProvider] 获取验证码 - URL: %s", apiURL)

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(codeCtx, "GET", apiURL, nil)
	if err != nil {
		zap.S().Errorf("[MQTTProvider] 创建请求失败: %v", err)
		zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 创建请求失败 ==========")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.1")

	// 发送请求
	resp, err := p.httpClient.Do(req)
	if err != nil {
		zap.S().Warnf("[MQTTProvider] 获取验证码API调用失败: provider=%s, phone=%s, ext_id=%s, error=%v",
			p.info.ID, phoneNumber, extIdValue, err)
		zap.S().Infof("[MQTTProvider] ========== 获取验证码 - API调用失败 ==========")
		return nil, err
	}

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		zap.S().Errorf("[MQTTProvider] 读取响应失败: %v", err)
		zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 读取响应失败 ==========")
		return nil, err
	}

	zap.S().Infof("[MQTTProvider] 获取验证码 - 响应: %s", string(respBody))

	// 解析响应
	var apiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			ReceiveStatus int    `json:"receiveStatus"` // 0失败 1成功
			Message       string `json:"message"`       // 短信内容
			Error         string `json:"error"`         // 错误信息
		} `json:"data"`
	}

	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		zap.S().Warnf("[MQTTProvider] 解析验证码响应失败: provider=%s, error=%v", p.info.ID, err)
		zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 解析失败 ==========")
		return nil, err
	}

	// 输出完整响应数据
	responseJSON, _ := json.Marshal(apiResponse)
	zap.S().Infof("[MQTTProvider] 获取验证码 - 响应详情: code=%d, message=%s, receiveStatus=%d, message=%s",
		apiResponse.Code, apiResponse.Message, apiResponse.Data.ReceiveStatus, apiResponse.Data.Message)
	zap.S().Infof("[MQTTProvider] 获取验证码 - 完整响应数据(JSON): %s", string(responseJSON))

	if apiResponse.Code != 0 {
		// 检查是否是 "already release" 错误
		if apiResponse.Message == "already release" {
			zap.S().Warnf("[MQTTProvider] 手机号已被运营商释放: phoneNumber=%s, extId=%s", phoneNumber, extIdValue)
			zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 已释放 ==========")
			return nil, ErrPhoneAlreadyReleased
		}
		zap.S().Errorf("[MQTTProvider] ========== 获取验证码 - 失败 ==========")
		zap.S().Infof("[MQTTProvider] 错误详情: code=%d, message=%s",
			apiResponse.Code, apiResponse.Message)
		return nil, NewProviderError("API_ERROR", apiResponse.Message)
	}

	// 检查接收状态
	if apiResponse.Data.ReceiveStatus == 1 {
		// 成功接收到短信
		code := apiResponse.Data.Message
		zap.S().Infof("[MQTTProvider] 获取验证码成功: phoneNumber=%s, extId=%s, code=%s",
			phoneNumber, extIdValue, code)
		zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 成功 ==========")

		return &CodeResponse{
			Code:       code,
			Message:    "验证码接收成功",
			ReceivedAt: time.Now(),
			ProviderID: p.info.ID,
		}, nil
	} else if apiResponse.Data.ReceiveStatus == 0 {
		// 还未接收到短信
		zap.S().Infof("[MQTTProvider] 验证码暂未接收: phoneNumber=%s, extId=%s", phoneNumber, extIdValue)
		zap.S().Infof("[MQTTProvider] ========== 获取验证码 - 暂未接收 ==========")
		return nil, ErrCodeNotReceived
	}

	// 未知状态
	zap.S().Warnf("[MQTTProvider] 未知的接收状态: %d", apiResponse.Data.ReceiveStatus)
	return nil, NewProviderError("UNKNOWN_STATUS", "未知的接收状态")
}

// ReleasePhone 释放手机号
// API: http://szb.jczl70.com:6086/mqtt/msg/release?extId=%s&status=%d
// status: 1=注册成功, 2=超时, 3=已注册, 4=其它问题
// extId 参数：如果提供了 extId，直接使用；否则返回错误（因为API需要 extId）
func (p *MQTTProvider) ReleasePhone(ctx context.Context, phoneNumber string, extId ...string) error {
	if !p.IsHealthy(ctx) {
		return ErrProviderUnavailable
	}

	var extIdValue string
	// 如果提供了 extId 参数，直接使用
	if len(extId) > 0 && extId[0] != "" {
		extIdValue = extId[0]
	} else {
		// 如果没有提供 extId，返回错误
		zap.S().Errorf("[MQTTProvider] 释放手机号失败: 手机号 %s 需要提供 extId 参数", phoneNumber)
		zap.S().Infof("[MQTTProvider] 提示: ReleasePhone 接口需要 extId，请从数据库读取 ext_id 字段")
		return NewProviderError("INVALID_PHONE", fmt.Sprintf("手机号 %s 需要提供 extId 参数", phoneNumber))
	}

	// 默认状态为 4（其它问题），调用者可以根据实际情况传入不同的状态
	// 但接口只接收 phoneNumber 和 extId，status 需要从外部传入
	// 这里我们使用默认值 4，或者可以通过 ExtraConfig 配置
	status := 4 // 默认：其它问题
	return p.ReleasePhoneWithExtId(ctx, extIdValue, status)
}

// ReleasePhoneWithExtId 释放手机号（使用 extId）
// 这是辅助方法，用于在知道 extId 的情况下释放手机号
func (p *MQTTProvider) ReleasePhoneWithExtId(ctx context.Context, extId string, status int) error {
	if !p.IsHealthy(ctx) {
		return ErrProviderUnavailable
	}

	zap.S().Infof("[MQTTProvider] ========== 释放手机号 - 开始 ==========")
	zap.S().Infof("[MQTTProvider] 释放参数: extId=%s, status=%d", extId, status)

	// 构建URL
	apiURL := fmt.Sprintf("%s/mqtt/msg/release?extId=%s&status=%d", p.apiGateway, url.QueryEscape(extId), status)
	zap.S().Infof("[MQTTProvider] 请求URL: %s", apiURL)

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		zap.S().Infof("[MQTTProvider] 创建请求失败: %v", err)
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.1")

	// 发送请求
	zap.S().Infof("[MQTTProvider] 发送GET请求到: %s", apiURL)
	resp, err := p.httpClient.Do(req)
	if err != nil {
		zap.S().Errorf("[MQTTProvider] 释放手机号API调用失败: provider=%s, ext_id=%s, status=%d, error=%v",
			p.info.ID, extId, status, err)
		zap.S().Errorf("[MQTTProvider] ========== 释放手机号 - 失败 ==========")
		return fmt.Errorf("释放手机号失败: %w", err)
	}
	defer resp.Body.Close()

	zap.S().Infof("[MQTTProvider] HTTP状态码: %d", resp.StatusCode)

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.S().Infof("[MQTTProvider] 读取响应失败: %v", err)
		zap.S().Errorf("[MQTTProvider] ========== 释放手机号 - 失败 ==========")
		return fmt.Errorf("读取响应失败: %w", err)
	}

	zap.S().Infof("[MQTTProvider] 原始响应数据: %s", string(respBody))

	if resp.StatusCode != http.StatusOK {
		zap.S().Infof("[MQTTProvider] HTTP错误: 状态码=%d, 响应内容=%s", resp.StatusCode, string(respBody))
		zap.S().Errorf("[MQTTProvider] ========== 释放手机号 - 失败 ==========")
		return fmt.Errorf("HTTP错误: %d, 响应: %s", resp.StatusCode, string(respBody))
	}

	// 解析响应（如果API返回JSON）
	var apiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(respBody, &apiResponse); err == nil {
		// 如果解析成功，检查 code
		if apiResponse.Code != 0 {
			zap.S().Errorf("[MQTTProvider] 释放手机号返回错误: provider=%s, code=%d, message=%s",
				p.info.ID, apiResponse.Code, apiResponse.Message)
			zap.S().Errorf("[MQTTProvider] ========== 释放手机号 - 失败 ==========")
			return NewProviderError("API_ERROR", apiResponse.Message)
		}
	}

	zap.S().Infof("[MQTTProvider] 释放手机号成功: provider=%s, ext_id=%s, status=%d", p.info.ID, extId, status)
	zap.S().Infof("[MQTTProvider] ========== 释放手机号 - 成功 ==========")

	return nil
}
