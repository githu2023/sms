package provider

import (
	"bytes"
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

// BigBus666Provider BigBus666运营商实现
// 文档: http://mhm1111z.bigbus666.top:2086/doc/doc_m_hema.html
type BigBus666Provider struct {
	info              *ProviderInfo
	healthy           bool
	mu                sync.RWMutex
	apiGateway        string            // API网关地址
	customerOutNumber string            // 客户外部数字
	encryptKey        string            // AES加密密钥
	httpClient        *http.Client      // HTTP客户端
	extIdToPhone      map[string]string // extId -> phoneNumber 映射
	phoneToExtId      map[string]string // phoneNumber -> extId 映射
}

// BigBus666Config BigBus666配置
type BigBus666Config struct {
	ID                 string
	Name               string
	APIGateway         string // API网关地址
	CustomerOutNumber  string // 客户外部数字
	EncryptKey         string // AES加密密钥
	ProjectName        string // 项目名称（从 ExtraConfig 读取，默认 "hema"）
	Priority           int
	CostPerSMS         float64
	SupportedCountries []string
	Timeout            time.Duration
}

// NewBigBus666Provider 创建BigBus666运营商实例
func NewBigBus666Provider(config BigBus666Config) *BigBus666Provider {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	// 设置默认 projectName
	if config.ProjectName == "" {
		config.ProjectName = "hema"
	}

	return &BigBus666Provider{
		info: &ProviderInfo{
			ID:                 config.ID,
			Name:               config.Name,
			Type:               "http",
			Priority:           config.Priority,
			CostPerSMS:         config.CostPerSMS,
			SupportedCountries: config.SupportedCountries,
			Timeout:            config.Timeout,
			Metadata: map[string]string{
				"api_gateway":         config.APIGateway,
				"customer_out_number": config.CustomerOutNumber,
				"project_name":        config.ProjectName,
			},
		},
		healthy:           true,
		apiGateway:        config.APIGateway,
		customerOutNumber: config.CustomerOutNumber,
		encryptKey:        config.EncryptKey,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		extIdToPhone: make(map[string]string),
		phoneToExtId: make(map[string]string),
	}
}

// GetProviderInfo 返回运营商信息
func (p *BigBus666Provider) GetProviderInfo() *ProviderInfo {
	return p.info
}

// IsHealthy 检查运营商是否健康
func (p *BigBus666Provider) IsHealthy(ctx context.Context) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.healthy
}

// SetHealthy 设置健康状态
func (p *BigBus666Provider) SetHealthy(healthy bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.healthy = healthy
}

// GetPhone 获取手机号码
func (p *BigBus666Provider) GetPhone(ctx context.Context, businessType, cardType string) (*PhoneResponse, error) {
	if !p.IsHealthy(ctx) {
		return nil, ErrProviderUnavailable
	}

	// 调用获取号码接口 /n/{customerOutNumber}
	// 请求参数: {"projectName": "项目名称"}
	// 优先使用配置中的 projectName，如果没有则使用 businessType
	projectName := p.info.Metadata["project_name"]
	if projectName == "" {
		projectName = businessType // 回退到使用业务类型
	}

	// 输出完整配置信息（用于对接）
	zap.S().Infof("[BigBus666] ========== 获取手机号 - 开始 ==========")
	// 安全显示密钥前4位
	keyPrefix := ""
	if len(p.encryptKey) >= 4 {
		keyPrefix = p.encryptKey[:4]
	} else {
		keyPrefix = p.encryptKey
	}
	zap.S().Infof("[BigBus666] Provider配置: provider=%s, apiGateway=%s, customerOutNumber=%s, encryptKey前4位=%s",
		p.info.ID, p.apiGateway, p.customerOutNumber, keyPrefix)
	zap.S().Infof("[BigBus666] 请求参数: projectName=%s, businessType=%s, cardType=%s",
		projectName, businessType, cardType)

	requestData := map[string]string{
		"projectName": projectName,
	}

	// 输出完整请求数据（JSON格式，用于对接）
	requestJSON, _ := json.Marshal(requestData)
	zap.S().Infof("[BigBus666] 请求数据(JSON): %s", string(requestJSON))

	// 转换为 map[string]interface{}
	requestDataInterface := make(map[string]interface{})
	for k, v := range requestData {
		requestDataInterface[k] = v
	}

	response, err := p.callAPI(ctx, "n", requestDataInterface)
	if err != nil {
		zap.S().Infof("[BigBus666] 获取手机号API调用失败: provider=%s, projectName=%s, business_type=%s, error=%v",
			p.info.ID, projectName, businessType, err)
		zap.S().Infof("[BigBus666] ========== 获取手机号 - 失败 ==========")
		return nil, fmt.Errorf("获取手机号失败: %w", err)
	}

	// 解析响应
	var apiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
		Data    struct {
			ExtId  string `json:"extId"`
			Mobile string `json:"mobile"`
		} `json:"data"`
	}

	if err := json.Unmarshal(response, &apiResponse); err != nil {
		zap.S().Infof("[BigBus666] 解析响应失败: provider=%s, error=%v", p.info.ID, err)
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	if apiResponse.Code != 0 || !apiResponse.Success {
		// 输出完整错误信息（用于对接）
		zap.S().Infof("[BigBus666] ========== 获取手机号 - 运营商返回错误 ==========")
		zap.S().Infof("[BigBus666] 错误详情: code=%d, message=%s, success=%v",
			apiResponse.Code, apiResponse.Message, apiResponse.Success)
		zap.S().Infof("[BigBus666] 请求信息: URL=%s/%s/%s, projectName=%s, customerOutNumber=%s",
			p.apiGateway, "n", p.customerOutNumber, projectName, p.customerOutNumber)
		zap.S().Infof("[BigBus666] 完整响应数据(JSON): %s", string(response))
		zap.S().Infof("[BigBus666] ========== 错误信息结束 ==========")
		return nil, NewProviderError("API_ERROR", apiResponse.Message)
	}

	extId := apiResponse.Data.ExtId
	phoneNumber := apiResponse.Data.Mobile

	if extId == "" || phoneNumber == "" {
		return nil, NewProviderError("INVALID_RESPONSE", "返回数据不完整")
	}

	// 保存映射关系
	p.mu.Lock()
	p.extIdToPhone[extId] = phoneNumber
	p.phoneToExtId[phoneNumber] = extId
	p.mu.Unlock()

	zap.S().Infof("[BigBus666] 获取手机号成功: provider=%s, phone=%s, ext_id=%s", p.info.ID, phoneNumber, extId)
	zap.S().Infof("[BigBus666] ========== 获取手机号 - 成功 ==========")

	return &PhoneResponse{
		PhoneNumber: phoneNumber,
		CountryCode: "CN", // 默认中国，可以根据实际情况调整
		Cost:        p.info.CostPerSMS,
		ValidUntil:  time.Now().Add(30 * time.Minute), // 默认30分钟有效期
		ProviderID:  p.info.ID,
		ExtId:       extId, // 保存 extId，用于后续获取验证码和释放手机号
	}, nil
}

// GetCode 获取验证码
// extId 参数：如果提供了 extId，直接使用；否则从 phoneNumber 查找
func (p *BigBus666Provider) GetCode(ctx context.Context, phoneNumber string, timeout time.Duration, extId ...string) (*CodeResponse, error) {
	if !p.IsHealthy(ctx) {
		return nil, ErrProviderUnavailable
	}

	var extIdValue string
	// 如果提供了 extId 参数，直接使用
	if len(extId) > 0 && extId[0] != "" {
		extIdValue = extId[0]
	} else {
		// 否则从内存映射中查找
		p.mu.RLock()
		var exists bool
		extIdValue, exists = p.phoneToExtId[phoneNumber]
		p.mu.RUnlock()

		if !exists {
			return nil, NewProviderError("INVALID_PHONE", fmt.Sprintf("手机号 %s 未找到对应的extId", phoneNumber))
		}
	}

	zap.S().Infof("[BigBus666] ========== 获取验证码 - 开始 ==========")
	zap.S().Infof("[BigBus666] 请求参数: phoneNumber=%s, extId=%s, timeout=%v", phoneNumber, extIdValue, timeout)

	// 创建带超时的上下文（单次请求）
	codeCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// 调用接收短信接口 /r/{customerOutNumber} - 单次查询，不轮询
	requestData := map[string]string{
		"extId": extIdValue,
	}

	// 输出请求数据
	requestJSON, _ := json.Marshal(requestData)
	zap.S().Infof("[BigBus666] 获取验证码 - 请求数据(JSON): %s", string(requestJSON))

	// 转换为 map[string]interface{}
	requestDataInterface := make(map[string]interface{})
	for k, v := range requestData {
		requestDataInterface[k] = v
	}

	response, err := p.callAPI(codeCtx, "r", requestDataInterface)
	if err != nil {
		zap.S().Warnf("[BigBus666] 获取验证码API调用失败: provider=%s, phone=%s, ext_id=%s, error=%v",
			p.info.ID, phoneNumber, extIdValue, err)
		zap.S().Infof("[BigBus666] ========== 获取验证码 - API调用失败 ==========")
		return nil, err
	}

	// 解析响应
	var apiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
		Data    struct {
			ReceiveStatus int    `json:"receiveStatus"` // 0失败 1成功
			Message       string `json:"message"`       // 短信内容
			Error         string `json:"error"`         // 错误信息
		} `json:"data"`
	}

	if err := json.Unmarshal(response, &apiResponse); err != nil {
		zap.S().Warnf("[BigBus666] 解析验证码响应失败: provider=%s, error=%v", p.info.ID, err)
		zap.S().Infof("[BigBus666] ========== 获取验证码 - 解析失败 ==========")
		return nil, err
	}

	// 输出完整响应数据
	responseJSON, _ := json.Marshal(apiResponse)
	zap.S().Infof("[BigBus666] 获取验证码 - 响应: code=%d, success=%v, receiveStatus=%d, message=%s",
		apiResponse.Code, apiResponse.Success, apiResponse.Data.ReceiveStatus, apiResponse.Data.Message)
	zap.S().Infof("[BigBus666] 获取验证码 - 完整响应数据(JSON): %s", string(responseJSON))

	if apiResponse.Code != 0 || !apiResponse.Success {
		// 检查是否是 "already release" 错误
		if apiResponse.Message == "already release" {
			zap.S().Warnf("[BigBus666] 手机号已被运营商释放: phoneNumber=%s, extId=%s", phoneNumber, extIdValue)
			zap.S().Infof("[BigBus666] ========== 获取验证码 - 已释放 ==========")
			return nil, ErrPhoneAlreadyReleased
		}
		zap.S().Infof("[BigBus666] ========== 获取验证码 - 失败 ==========")
		zap.S().Infof("[BigBus666] 错误详情: code=%d, message=%s, success=%v",
			apiResponse.Code, apiResponse.Message, apiResponse.Success)
		return nil, NewProviderError("API_ERROR", apiResponse.Message)
	}

	// 检查接收状态
	if apiResponse.Data.ReceiveStatus == 1 {
		// 成功接收到短信
		code := apiResponse.Data.Message
		zap.S().Infof("[BigBus666] 获取验证码成功: phoneNumber=%s, extId=%s, code=%s",
			phoneNumber, extIdValue, code)
		zap.S().Infof("[BigBus666] ========== 获取验证码 - 成功 ==========")

		return &CodeResponse{
			Code:       code,
			Message:    "验证码接收成功",
			ReceivedAt: time.Now(),
			ProviderID: p.info.ID,
		}, nil
	} else if apiResponse.Data.ReceiveStatus == 0 {
		// 还未接收到短信
		zap.S().Infof("[BigBus666] 验证码暂未接收: phoneNumber=%s, extId=%s", phoneNumber, extIdValue)
		zap.S().Infof("[BigBus666] ========== 获取验证码 - 暂未接收 ==========")
		return nil, ErrCodeNotReceived
	}

	// 未知状态
	zap.S().Warnf("[BigBus666] 未知的接收状态: %d", apiResponse.Data.ReceiveStatus)
	return nil, NewProviderError("UNKNOWN_STATUS", "未知的接收状态")
}

// callAPI 调用API接口
func (p *BigBus666Provider) callAPI(ctx context.Context, endpoint string, requestData map[string]interface{}) ([]byte, error) {
	// 1. 将请求数据转换为JSON字符串
	jsonBytes, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求数据失败: %w", err)
	}
	jsonStr := string(jsonBytes)

	zap.S().Infof("[BigBus666] API调用 - 请求数据(原始JSON): %s", jsonStr)

	// 2. AES加密
	encryptedBytes, err := p.encryptAES(jsonStr)
	if err != nil {
		zap.S().Infof("[BigBus666] API调用 - 加密失败: %v", err)
		return nil, fmt.Errorf("加密失败: %w", err)
	}

	// 输出加密后的原始字节数据（十六进制）
	zap.S().Infof("[BigBus666] API调用 - 加密后原始数据(HEX): %s", hex.EncodeToString(encryptedBytes))
	zap.S().Infof("[BigBus666] API调用 - 加密后数据长度: %d 字节", len(encryptedBytes))

	// 3. Base64编码
	base64Str := base64.StdEncoding.EncodeToString(encryptedBytes)

	// 4. 构建请求URL
	url := fmt.Sprintf("%s/%s/%s", p.apiGateway, endpoint, p.customerOutNumber)
	zap.S().Infof("[BigBus666] API调用 - 完整URL: %s", url)
	zap.S().Infof("[BigBus666] API调用 - 请求参数: endpoint=%s, customerOutNumber=%s", endpoint, p.customerOutNumber)

	// 5. 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(base64Str))
	if err != nil {
		zap.S().Infof("[BigBus666] API调用 - 创建请求失败: %v", err)
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("User-Agent", "Chrome/69.0.3497.81 Safari/537.36")

	// 输出加密后的Base64数据（用于对接）
	zap.S().Infof("[BigBus666] API调用 - 加密后Base64数据: %s", base64Str)
	zap.S().Infof("[BigBus666] API调用 - 请求头: Content-Type=%s, User-Agent=%s", req.Header.Get("Content-Type"), req.Header.Get("User-Agent"))

	// 6. 发送请求
	zap.S().Infof("[BigBus666] API调用 - 发送POST请求到: %s", url)
	resp, err := p.httpClient.Do(req)
	if err != nil {
		zap.S().Infof("[BigBus666] API调用 - 请求失败: %v", err)
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	zap.S().Infof("[BigBus666] API调用 - HTTP状态码: %d", resp.StatusCode)

	// 7. 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.S().Infof("[BigBus666] API调用 - 读取响应失败: %v", err)
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	zap.S().Infof("[BigBus666] API调用 - 原始响应数据(Base64): %s", string(respBody))
	zap.S().Infof("[BigBus666] API调用 - 响应长度: %d 字节", len(respBody))

	if resp.StatusCode != http.StatusOK {
		zap.S().Infof("[BigBus666] API调用 - HTTP错误: 状态码=%d, 响应内容=%s", resp.StatusCode, string(respBody))
		return nil, fmt.Errorf("HTTP错误: %d, 响应: %s", resp.StatusCode, string(respBody))
	}

	// 8. Base64解码
	encryptedResponse, err := base64.StdEncoding.DecodeString(string(respBody))
	if err != nil {
		zap.S().Infof("[BigBus666] API调用 - Base64解码失败: %v, 响应长度: %d", err, len(respBody))
		return nil, fmt.Errorf("Base64解码失败: %w", err)
	}

	// 输出解密前的加密数据（十六进制）
	zap.S().Infof("[BigBus666] API调用 - 解密前加密数据(HEX): %s", hex.EncodeToString(encryptedResponse))
	zap.S().Infof("[BigBus666] API调用 - 解密前加密数据长度: %d 字节", len(encryptedResponse))

	// 9. AES解密
	decryptedBytes, err := p.decryptAES(encryptedResponse)
	if err != nil {
		zap.S().Infof("[BigBus666] API调用 - AES解密失败: %v, 加密数据长度=%d", err, len(encryptedResponse))
		return nil, fmt.Errorf("解密失败: %w", err)
	}

	// 输出完整响应数据（用于对接）
	zap.S().Infof("[BigBus666] API调用 - 解密后响应数据(JSON): %s", string(decryptedBytes))
	return decryptedBytes, nil
}

// encryptAES AES加密 (AES/ECB/PKCS7Padding)
func (p *BigBus666Provider) encryptAES(plaintext string) ([]byte, error) {
	// 确保key长度为16字节 (AES-128)
	key := p.padKey(p.encryptKey, 16)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// PKCS7Padding
	plaintextBytes := []byte(plaintext)
	plaintextBytes = p.pkcs7Padding(plaintextBytes, aes.BlockSize)

	// ECB模式加密
	ciphertext := make([]byte, len(plaintextBytes))
	for i := 0; i < len(plaintextBytes); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], plaintextBytes[i:i+aes.BlockSize])
	}

	return ciphertext, nil
}

// decryptAES AES解密 (AES/ECB/PKCS7Padding)
func (p *BigBus666Provider) decryptAES(ciphertext []byte) ([]byte, error) {
	// 确保key长度为16字节 (AES-128)
	key := p.padKey(p.encryptKey, 16)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// ECB模式解密
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	// 去除PKCS7Padding
	plaintext = p.pkcs7Unpadding(plaintext)

	return plaintext, nil
}

// pkcs7Padding PKCS7填充
func (p *BigBus666Provider) pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7Unpadding 去除PKCS7填充
func (p *BigBus666Provider) pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return data
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return data
	}
	return data[:(length - unpadding)]
}

// ReleasePhone 释放手机号
// 根据 BigBus666 API 文档，释放接口应该是 /d/{customerOutNumber} (delete)
// 请求参数: {"extId": "xxx"}
// extId 参数：如果提供了 extId，直接使用；否则从内存映射中查找
func (p *BigBus666Provider) ReleasePhone(ctx context.Context, phoneNumber string, extId ...string) error {
	if !p.IsHealthy(ctx) {
		return ErrProviderUnavailable
	}

	var extIdValue string
	// 如果提供了 extId 参数，直接使用
	if len(extId) > 0 && extId[0] != "" {
		extIdValue = extId[0]
	} else {
		// 否则从内存映射中查找
		p.mu.RLock()
		var exists bool
		extIdValue, exists = p.phoneToExtId[phoneNumber]
		p.mu.RUnlock()

		if !exists {
			zap.S().Infof("[BigBus666] 释放手机号失败: 手机号 %s 未找到对应的extId（内存映射中不存在）", phoneNumber)
			zap.S().Infof("[BigBus666] 提示: 如果 extId 存储在数据库中，请从数据库读取并作为参数传入")
			return NewProviderError("INVALID_PHONE", fmt.Sprintf("手机号 %s 未找到对应的extId", phoneNumber))
		}
	}

	zap.S().Infof("[BigBus666] ========== 释放手机号 - 开始 ==========")
	zap.S().Infof("[BigBus666] 释放参数: phoneNumber=%s, extId=%s", phoneNumber, extIdValue)

	// 根据 BigBus666 API 文档，释放接口可能是 /d/{customerOutNumber} (delete)
	// 请求参数: {"extId": "xxx"}
	requestData := map[string]string{
		"extId": extIdValue,
	}

	// 转换为 map[string]interface{}
	requestDataInterface := make(map[string]interface{})
	for k, v := range requestData {
		requestDataInterface[k] = v
	}

	// 调用释放接口，通常释放接口的 endpoint 是 "d" (delete)
	// 如果文档中没有明确说明，可能需要根据实际情况调整
	response, err := p.callAPI(ctx, "d", requestDataInterface)
	if err != nil {
		zap.S().Infof("[BigBus666] 释放手机号API调用失败: provider=%s, phone=%s, ext_id=%s, error=%v",
			p.info.ID, phoneNumber, extIdValue, err)
		zap.S().Infof("[BigBus666] ========== 释放手机号 - 失败 ==========")
		return fmt.Errorf("释放手机号失败: %w", err)
	}

	// 解析响应
	var apiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
	}

	if err := json.Unmarshal(response, &apiResponse); err != nil {
		zap.S().Infof("[BigBus666] 解析释放响应失败: provider=%s, error=%v", p.info.ID, err)
		zap.S().Infof("[BigBus666] ========== 释放手机号 - 失败 ==========")
		return fmt.Errorf("解析响应失败: %w", err)
	}

	if apiResponse.Code != 0 || !apiResponse.Success {
		zap.S().Infof("[BigBus666] 释放手机号返回错误: provider=%s, code=%d, message=%s",
			p.info.ID, apiResponse.Code, apiResponse.Message)
		zap.S().Infof("[BigBus666] ========== 释放手机号 - 失败 ==========")
		return NewProviderError("API_ERROR", apiResponse.Message)
	}

	// 清理映射关系
	p.mu.Lock()
	delete(p.phoneToExtId, phoneNumber)
	delete(p.extIdToPhone, extIdValue)
	p.mu.Unlock()

	zap.S().Infof("[BigBus666] 释放手机号成功: provider=%s, phone=%s, ext_id=%s", p.info.ID, phoneNumber, extIdValue)
	zap.S().Infof("[BigBus666] ========== 释放手机号 - 成功 ==========")

	return nil
}

// padKey 填充或截断key到指定长度
func (p *BigBus666Provider) padKey(key string, size int) string {
	keyBytes := []byte(key)
	if len(keyBytes) > size {
		return string(keyBytes[:size])
	}
	if len(keyBytes) < size {
		padding := make([]byte, size-len(keyBytes))
		return string(append(keyBytes, padding...))
	}
	return key
}
