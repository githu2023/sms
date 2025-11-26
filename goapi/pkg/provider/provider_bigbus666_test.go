package provider

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// 测试配置
const (
	testAPIGateway        = "http://mhm1111z.bigbus666.top:2086/s/m"
	testCustomerOutNumber = "11296564"
	testEncryptKey        = "Z8UrC8H2cNgXe2Jw"
	testProjectName       = "hema"
)

// TestBigBus666Provider_EncryptDecrypt 测试加密解密功能
func TestBigBus666Provider_EncryptDecrypt(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		ID:                "test",
		Name:              "Test Provider",
		APIGateway:        testAPIGateway,
		CustomerOutNumber: testCustomerOutNumber,
		EncryptKey:        testEncryptKey,
		Priority:          100,
		CostPerSMS:        1.0,
		Timeout:           30 * time.Second,
	})

	// 测试数据
	plaintext := `{"projectName":"hema"}`

	// 加密
	encrypted, err := provider.encryptAES(plaintext)
	if err != nil {
		t.Fatalf("加密失败: %v", err)
	}

	if len(encrypted) == 0 {
		t.Fatal("加密结果为空")
	}

	// Base64编码
	base64Str := base64.StdEncoding.EncodeToString(encrypted)
	if base64Str == "" {
		t.Fatal("Base64编码结果为空")
	}

	// Base64解码
	decryptedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		t.Fatalf("Base64解码失败: %v", err)
	}

	// 解密
	decrypted, err := provider.decryptAES(decryptedBytes)
	if err != nil {
		t.Fatalf("解密失败: %v", err)
	}

	// 验证结果
	if string(decrypted) != plaintext {
		t.Errorf("解密结果不匹配: 期望 %s, 得到 %s", plaintext, string(decrypted))
	}
}

// TestBigBus666Provider_PKCS7Padding 测试PKCS7填充
func TestBigBus666Provider_PKCS7Padding(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	tests := []struct {
		name      string
		data      []byte
		blockSize int
	}{
		{"正常填充", []byte("hello"), 16},
		{"正好16字节", []byte("1234567890123456"), 16},
		{"超过16字节", []byte("12345678901234567"), 16},
		{"空数据", []byte(""), 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			padded := provider.pkcs7Padding(tt.data, tt.blockSize)

			// 验证长度是blockSize的倍数
			if len(padded)%tt.blockSize != 0 {
				t.Errorf("填充后长度 %d 不是 %d 的倍数", len(padded), tt.blockSize)
			}

			// 验证可以正确去除填充
			unpadded := provider.pkcs7Unpadding(padded)
			if string(unpadded) != string(tt.data) {
				t.Errorf("去除填充后不匹配: 期望 %s, 得到 %s", string(tt.data), string(unpadded))
			}
		})
	}
}

// TestBigBus666Provider_PadKey 测试密钥填充
func TestBigBus666Provider_PadKey(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	tests := []struct {
		name     string
		key      string
		size     int
		expected int
	}{
		{"正常密钥", "Z8UrC8H2cNgXe2Jw", 16, 16},
		{"短密钥", "short", 16, 16},
		{"长密钥", "this_is_a_very_long_key_that_should_be_truncated", 16, 16},
		{"空密钥", "", 16, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			padded := provider.padKey(tt.key, tt.size)
			if len(padded) != tt.expected {
				t.Errorf("密钥长度不正确: 期望 %d, 得到 %d", tt.expected, len(padded))
			}
		})
	}
}

// TestBigBus666Provider_GetPhone_Success 测试成功获取手机号（使用真实API）
func TestBigBus666Provider_GetPhone_Success(t *testing.T) {
	// 使用本地 Provider 生成一个真实可用的手机号，避免去调用外部 BigBus666 API
	localProvider := NewLocalProvider("local", "Local Provider", 1, nil)
	ctx := context.Background()
	localPhone, err := localProvider.GetPhone(ctx, "wechat", "physical")
	require.NoError(t, err)
	require.NotNil(t, localPhone)

	// 准备一个 BigBus666 Provider 用于加密响应
	encryptProvider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	extID := "ext-" + localPhone.PhoneNumber

	// 启动本地 HTTP Server 模拟 BigBus666 的 API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, fmt.Sprintf("/n/%s", testCustomerOutNumber), r.URL.Path)

		resp := map[string]interface{}{
			"code":    0,
			"message": "success",
			"success": true,
			"data": map[string]interface{}{
				"extId":  extID,
				"mobile": localPhone.PhoneNumber,
			},
		}

		respJSON, _ := json.Marshal(resp)
		encrypted, err := encryptProvider.encryptAES(string(respJSON))
		require.NoError(t, err)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(base64.StdEncoding.EncodeToString(encrypted)))
	}))
	defer server.Close()

	provider := NewBigBus666Provider(BigBus666Config{
		ID:                "test",
		Name:              "Test Provider",
		APIGateway:        server.URL,
		CustomerOutNumber: testCustomerOutNumber,
		EncryptKey:        testEncryptKey,
		Priority:          100,
		CostPerSMS:        1.0,
		Timeout:           5 * time.Second,
		ProjectName:       testProjectName,
	})

	result, err := provider.GetPhone(ctx, testProjectName, "physical")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, localPhone.PhoneNumber, result.PhoneNumber)
	require.Equal(t, extID, result.ExtId)
	require.Equal(t, "test", result.ProviderID)

	// 验证内存映射
	provider.mu.RLock()
	mappedExtID, exists := provider.phoneToExtId[result.PhoneNumber]
	provider.mu.RUnlock()
	require.True(t, exists)
	require.Equal(t, extID, mappedExtID)
}

// TestBigBus666Provider_GetPhone_APIError 测试API错误（使用真实API，测试错误处理）
// 注意：这个测试可能会因为真实API返回错误而失败，这是正常的
func TestBigBus666Provider_GetPhone_APIError(t *testing.T) {
	encryptProvider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"code":    500,
			"message": "参数错误",
			"success": false,
		}
		respJSON, _ := json.Marshal(resp)
		encrypted, err := encryptProvider.encryptAES(string(respJSON))
		require.NoError(t, err)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(base64.StdEncoding.EncodeToString(encrypted)))
	}))
	defer server.Close()

	provider := NewBigBus666Provider(BigBus666Config{
		ID:                "test",
		Name:              "Test Provider",
		APIGateway:        server.URL,
		CustomerOutNumber: testCustomerOutNumber,
		EncryptKey:        testEncryptKey,
		Priority:          100,
		CostPerSMS:        1.0,
		Timeout:           5 * time.Second,
	})

	ctx := context.Background()
	result, err := provider.GetPhone(ctx, testProjectName, "physical")

	require.Error(t, err)
	require.Nil(t, result)
	require.True(t, IsProviderError(err))
}

// TestBigBus666Provider_GetCode_Success 测试成功获取验证码
func TestBigBus666Provider_GetCode_Success(t *testing.T) {
	// 创建Provider用于加密响应
	testProvider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	callCount := 0
	// 创建模拟服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		var response map[string]interface{}
		if callCount == 1 {
			// 第一次调用，返回等待中
			response = map[string]interface{}{
				"code":    0,
				"message": "成功",
				"success": true,
				"data": map[string]interface{}{
					"receiveStatus": 0,
					"message":       "",
				},
			}
		} else {
			// 第二次调用，返回成功
			response = map[string]interface{}{
				"code":    0,
				"message": "成功",
				"success": true,
				"data": map[string]interface{}{
					"receiveStatus": 1,
					"message":       "123456",
				},
			}
		}

		responseJSON, _ := json.Marshal(response)

		// 加密响应
		encrypted, err := testProvider.encryptAES(string(responseJSON))
		if err != nil {
			t.Errorf("加密响应失败: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		base64Response := base64.StdEncoding.EncodeToString(encrypted)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(base64Response))
	}))
	defer server.Close()

	provider := NewBigBus666Provider(BigBus666Config{
		APIGateway:        server.URL + "/s/m",
		CustomerOutNumber: testCustomerOutNumber,
		EncryptKey:        testEncryptKey,
		Timeout:           10 * time.Second,
	})

	// 先设置映射关系
	provider.mu.Lock()
	provider.phoneToExtId["13800138000"] = "12345"
	provider.extIdToPhone["12345"] = "13800138000"
	provider.mu.Unlock()

	ctx := context.Background()
	result, err := provider.GetCode(ctx, "13800138000", 5*time.Second)

	if err != nil {
		t.Fatalf("获取验证码失败: %v", err)
	}

	if result == nil {
		t.Fatal("返回结果为nil")
	}

	if result.Code != "123456" {
		t.Errorf("验证码不匹配: 期望 123456, 得到 %s", result.Code)
	}

	if callCount < 2 {
		t.Errorf("期望至少调用2次API，实际调用 %d 次", callCount)
	}
}

// TestBigBus666Provider_GetCode_Timeout 测试验证码获取超时
func TestBigBus666Provider_GetCode_Timeout(t *testing.T) {
	// 创建Provider用于加密响应
	testProvider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	// 创建模拟服务器，始终返回等待中
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"code":    0,
			"message": "成功",
			"success": true,
			"data": map[string]interface{}{
				"receiveStatus": 0,
				"message":       "",
			},
		}

		responseJSON, _ := json.Marshal(response)

		// 加密响应
		encrypted, err := testProvider.encryptAES(string(responseJSON))
		if err != nil {
			t.Errorf("加密响应失败: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		base64Response := base64.StdEncoding.EncodeToString(encrypted)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(base64Response))
	}))
	defer server.Close()

	provider := NewBigBus666Provider(BigBus666Config{
		APIGateway:        server.URL + "/s/m",
		CustomerOutNumber: testCustomerOutNumber,
		EncryptKey:        testEncryptKey,
		Timeout:           2 * time.Second,
	})

	// 设置映射关系
	provider.mu.Lock()
	provider.phoneToExtId["13800138000"] = "12345"
	provider.extIdToPhone["12345"] = "13800138000"
	provider.mu.Unlock()

	ctx := context.Background()
	result, err := provider.GetCode(ctx, "13800138000", 1*time.Second)

	if err == nil {
		t.Fatal("期望返回超时错误，但没有错误")
	}

	if result != nil {
		t.Error("期望返回nil，但有结果")
	}

	if err != ErrCodeTimeout {
		t.Errorf("期望ErrCodeTimeout，得到 %v", err)
	}
}

// TestBigBus666Provider_GetCode_InvalidPhone 测试无效手机号
func TestBigBus666Provider_GetCode_InvalidPhone(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
		Timeout:    30 * time.Second,
	})

	ctx := context.Background()
	result, err := provider.GetCode(ctx, "13800138000", 5*time.Second)

	if err == nil {
		t.Fatal("期望返回错误，但没有错误")
	}

	if result != nil {
		t.Error("期望返回nil，但有结果")
	}

	if !IsProviderError(err) {
		t.Error("期望ProviderError类型")
	}
}

// TestBigBus666Provider_IsHealthy 测试健康检查
func TestBigBus666Provider_IsHealthy(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
	})

	ctx := context.Background()

	// 默认应该是健康的
	if !provider.IsHealthy(ctx) {
		t.Error("期望Provider是健康的")
	}

	// 设置为不健康
	provider.SetHealthy(false)
	if provider.IsHealthy(ctx) {
		t.Error("期望Provider是不健康的")
	}

	// 重新设置为健康
	provider.SetHealthy(true)
	if !provider.IsHealthy(ctx) {
		t.Error("期望Provider是健康的")
	}
}

// TestBigBus666Provider_GetProviderInfo 测试获取Provider信息
func TestBigBus666Provider_GetProviderInfo(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		ID:                 "test",
		Name:               "Test Provider",
		APIGateway:         testAPIGateway,
		CustomerOutNumber:  testCustomerOutNumber,
		EncryptKey:         testEncryptKey,
		Priority:           100,
		CostPerSMS:         1.5,
		SupportedCountries: []string{"CN", "US"},
		Timeout:            30 * time.Second,
	})

	info := provider.GetProviderInfo()

	if info == nil {
		t.Fatal("ProviderInfo为nil")
	}

	if info.ID != "test" {
		t.Errorf("ID不匹配: 期望 test, 得到 %s", info.ID)
	}

	if info.Name != "Test Provider" {
		t.Errorf("Name不匹配: 期望 Test Provider, 得到 %s", info.Name)
	}

	if info.Type != "http" {
		t.Errorf("Type不匹配: 期望 http, 得到 %s", info.Type)
	}

	if info.CostPerSMS != 1.5 {
		t.Errorf("CostPerSMS不匹配: 期望 1.5, 得到 %f", info.CostPerSMS)
	}

	if len(info.SupportedCountries) != 2 {
		t.Errorf("SupportedCountries长度不匹配: 期望 2, 得到 %d", len(info.SupportedCountries))
	}
}

// TestBigBus666Provider_GetPhone_Unhealthy 测试不健康状态下的获取手机号
func TestBigBus666Provider_GetPhone_Unhealthy(t *testing.T) {
	provider := NewBigBus666Provider(BigBus666Config{
		EncryptKey: testEncryptKey,
		Timeout:    30 * time.Second,
	})

	provider.SetHealthy(false)

	ctx := context.Background()
	result, err := provider.GetPhone(ctx, testProjectName, "physical")

	if err == nil {
		t.Fatal("期望返回错误，但没有错误")
	}

	if result != nil {
		t.Error("期望返回nil，但有结果")
	}

	if err != ErrProviderUnavailable {
		t.Errorf("期望ErrProviderUnavailable，得到 %v", err)
	}
}
