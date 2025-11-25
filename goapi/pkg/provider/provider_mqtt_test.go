package provider

import (
	"context"
	"testing"
	"time"
)

// 测试配置 - 使用真实的运营商API
const (
	testMQTTAPIGateway  = "http://szb.jczl70.com:6086"
	testMQTTProviderID  = "6"
	testMQTTProviderKey = "dbc4d84b"
)

// TestNewMQTTProvider 测试创建MQTTProvider
func TestNewMQTTProvider(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		ID:                 "test-mqtt",
		Name:               "Test MQTT Provider",
		APIGateway:         testMQTTAPIGateway,
		ProviderID:         testMQTTProviderID,
		ProviderKey:        testMQTTProviderKey,
		Priority:           100,
		CostPerSMS:         1.0,
		SupportedCountries: []string{"CN"},
		Timeout:            30 * time.Second,
	})

	if provider == nil {
		t.Fatal("Provider为nil")
	}

	info := provider.GetProviderInfo()
	if info == nil {
		t.Fatal("ProviderInfo为nil")
	}

	if info.ID != "test-mqtt" {
		t.Errorf("ID不匹配: 期望 test-mqtt, 得到 %s", info.ID)
	}

	if info.Name != "Test MQTT Provider" {
		t.Errorf("Name不匹配: 期望 Test MQTT Provider, 得到 %s", info.Name)
	}

	if info.Type != "http" {
		t.Errorf("Type不匹配: 期望 http, 得到 %s", info.Type)
	}

	if info.Priority != 100 {
		t.Errorf("Priority不匹配: 期望 100, 得到 %d", info.Priority)
	}

	if info.CostPerSMS != 1.0 {
		t.Errorf("CostPerSMS不匹配: 期望 1.0, 得到 %f", info.CostPerSMS)
	}

	if !provider.IsHealthy(context.Background()) {
		t.Error("期望Provider是健康的")
	}
}

// TestMQTTProvider_GetPhone_Success 测试成功获取手机号（使用真实API）
func TestMQTTProvider_GetPhone_Success(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		ID:          "mqtt-test",
		Name:        "MQTT Test Provider",
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Priority:    100,
		CostPerSMS:  1.0,
		Timeout:     30 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Logf("使用真实API测试获取手机号: id=%s, key=%s", testMQTTProviderID, testMQTTProviderKey)
	result, err := provider.GetPhone(ctx, "qq", "physical")

	if err != nil {
		t.Fatalf("获取手机号失败: %v", err)
	}

	if result == nil {
		t.Fatal("返回结果为nil")
	}

	if result.PhoneNumber == "" {
		t.Error("手机号为空")
	}

	if result.ExtId == "" {
		t.Error("extId为空")
	}

	if result.ProviderID != "mqtt-test" {
		t.Errorf("ProviderID不匹配: 期望 mqtt-test, 得到 %s", result.ProviderID)
	}

	t.Logf("✅ 成功获取手机号: %s, extId: %s", result.PhoneNumber, result.ExtId)
	t.Logf("   成本: %.2f, 有效期至: %s", result.Cost, result.ValidUntil.Format(time.RFC3339))
}

// TestMQTTProvider_GetCode_Success 测试成功获取验证码（使用真实API）
// 注意：这个测试需要先获取一个手机号，然后等待验证码
func TestMQTTProvider_GetCode_Success(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		ID:          "mqtt-test",
		Name:        "MQTT Test Provider",
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     30 * time.Second,
	})

	// 先获取手机号
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Log("步骤1: 获取手机号...")
	phoneResp, err := provider.GetPhone(ctx, "qq", "physical")
	if err != nil {
		t.Fatalf("获取手机号失败: %v", err)
	}

	if phoneResp.ExtId == "" {
		t.Fatal("获取的手机号没有extId，无法测试获取验证码")
	}

	t.Logf("✅ 获取到手机号: %s, extId: %s", phoneResp.PhoneNumber, phoneResp.ExtId)
	t.Log("步骤2: 等待验证码（30秒超时）...")
	t.Log("   提示: 请在30秒内完成注册并接收验证码")

	// 获取验证码（30秒超时）
	codeCtx, codeCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer codeCancel()

	codeResp, err := provider.GetCode(codeCtx, phoneResp.PhoneNumber, 30*time.Second, phoneResp.ExtId)

	if err != nil {
		if err == ErrCodeTimeout {
			t.Logf("⚠️  验证码获取超时（这是正常的，如果在30秒内没有收到验证码）")
		} else {
			t.Errorf("获取验证码失败: %v", err)
		}
		return
	}

	if codeResp == nil {
		t.Fatal("返回结果为nil")
	}

	if codeResp.Code == "" {
		t.Error("验证码为空")
	}

	t.Logf("✅ 成功获取验证码: %s", codeResp.Code)
	t.Logf("   接收时间: %s", codeResp.ReceivedAt.Format(time.RFC3339))
}

// TestMQTTProvider_GetCode_InvalidExtId 测试无效extId
func TestMQTTProvider_GetCode_InvalidExtId(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     10 * time.Second,
	})

	ctx := context.Background()
	result, err := provider.GetCode(ctx, "18888888888", 5*time.Second) // 不提供extId

	if err == nil {
		t.Error("期望返回错误，但没有错误")
	}

	if result != nil {
		t.Error("期望返回nil，但有结果")
	}

	if !IsProviderError(err) {
		t.Error("期望ProviderError类型")
	}
}

// TestMQTTProvider_ReleasePhone_Success 测试成功释放手机号（使用真实API）
func TestMQTTProvider_ReleasePhone_Success(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		ID:          "mqtt-test",
		Name:        "MQTT Test Provider",
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     30 * time.Second,
	})

	// 先获取手机号
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Log("步骤1: 获取手机号...")
	phoneResp, err := provider.GetPhone(ctx, "qq", "physical")
	if err != nil {
		t.Fatalf("获取手机号失败: %v", err)
	}

	if phoneResp.ExtId == "" {
		t.Fatal("获取的手机号没有extId，无法测试释放手机号")
	}

	t.Logf("✅ 获取到手机号: %s, extId: %s", phoneResp.PhoneNumber, phoneResp.ExtId)
	t.Log("步骤2: 释放手机号（status=4: 其它问题）...")

	// 释放手机号
	releaseCtx, releaseCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer releaseCancel()

	err = provider.ReleasePhone(releaseCtx, phoneResp.PhoneNumber, phoneResp.ExtId)

	if err != nil {
		t.Errorf("释放手机号失败: %v", err)
		return
	}

	t.Logf("✅ 成功释放手机号: %s (extId: %s)", phoneResp.PhoneNumber, phoneResp.ExtId)
}

// TestMQTTProvider_ReleasePhone_InvalidExtId 测试无效extId
func TestMQTTProvider_ReleasePhone_InvalidExtId(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     10 * time.Second,
	})

	ctx := context.Background()
	err := provider.ReleasePhone(ctx, "18888888888") // 不提供extId

	if err == nil {
		t.Error("期望返回错误，但没有错误")
	}

	if !IsProviderError(err) {
		t.Error("期望ProviderError类型")
	}
}

// TestMQTTProvider_IsHealthy 测试健康检查
func TestMQTTProvider_IsHealthy(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     10 * time.Second,
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

// TestMQTTProvider_GetPhone_Unhealthy 测试不健康状态下的获取手机号
func TestMQTTProvider_GetPhone_Unhealthy(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     10 * time.Second,
	})

	provider.SetHealthy(false)

	ctx := context.Background()
	result, err := provider.GetPhone(ctx, "qq", "physical")

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

// TestMQTTProvider_GetProviderInfo 测试获取Provider信息
func TestMQTTProvider_GetProviderInfo(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		ID:                 "test-mqtt",
		Name:               "Test MQTT Provider",
		APIGateway:         testMQTTAPIGateway,
		ProviderID:         testMQTTProviderID,
		ProviderKey:        testMQTTProviderKey,
		Priority:           100,
		CostPerSMS:         1.5,
		SupportedCountries: []string{"CN", "US"},
		Timeout:            30 * time.Second,
	})

	info := provider.GetProviderInfo()

	if info == nil {
		t.Fatal("ProviderInfo为nil")
	}

	if info.ID != "test-mqtt" {
		t.Errorf("ID不匹配: 期望 test-mqtt, 得到 %s", info.ID)
	}

	if info.Name != "Test MQTT Provider" {
		t.Errorf("Name不匹配: 期望 Test MQTT Provider, 得到 %s", info.Name)
	}

	if info.Type != "http" {
		t.Errorf("Type不匹配: 期望 http, 得到 %s", info.Type)
	}

	if info.Priority != 100 {
		t.Errorf("Priority不匹配: 期望 100, 得到 %d", info.Priority)
	}

	if info.CostPerSMS != 1.5 {
		t.Errorf("CostPerSMS不匹配: 期望 1.5, 得到 %f", info.CostPerSMS)
	}

	if len(info.SupportedCountries) != 2 {
		t.Errorf("SupportedCountries长度不匹配: 期望 2, 得到 %d", len(info.SupportedCountries))
	}

	if info.SupportedCountries[0] != "CN" {
		t.Errorf("SupportedCountries[0]不匹配: 期望 CN, 得到 %s", info.SupportedCountries[0])
	}

	if info.SupportedCountries[1] != "US" {
		t.Errorf("SupportedCountries[1]不匹配: 期望 US, 得到 %s", info.SupportedCountries[1])
	}
}

// TestMQTTProvider_Real_FullFlow 测试完整流程：获取手机号 -> 获取验证码 -> 释放手机号（使用真实API）
func TestMQTTProvider_Real_FullFlow(t *testing.T) {
	provider := NewMQTTProvider(MQTTConfig{
		ID:          "mqtt-test",
		Name:        "MQTT Test Provider",
		APIGateway:  testMQTTAPIGateway,
		ProviderID:  testMQTTProviderID,
		ProviderKey: testMQTTProviderKey,
		Timeout:     30 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 步骤1: 获取手机号
	t.Log("========== 步骤1: 获取手机号 ==========")
	phoneResp, err := provider.GetPhone(ctx, "qq", "physical")
	if err != nil {
		t.Fatalf("获取手机号失败: %v", err)
	}

	if phoneResp.ExtId == "" {
		t.Fatal("获取的手机号没有extId")
	}

	t.Logf("✅ 手机号: %s", phoneResp.PhoneNumber)
	t.Logf("✅ extId: %s", phoneResp.ExtId)
	t.Logf("✅ 成本: %.2f", phoneResp.Cost)

	// 步骤2: 尝试获取验证码（可选，因为可能需要等待）
	t.Log("\n========== 步骤2: 获取验证码（可选，30秒超时） ==========")
	t.Log("提示: 请在30秒内完成注册并接收验证码，否则会超时")

	codeCtx, codeCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer codeCancel()

	codeResp, err := provider.GetCode(codeCtx, phoneResp.PhoneNumber, 30*time.Second, phoneResp.ExtId)
	if err != nil {
		if err == ErrCodeTimeout {
			t.Log("⚠️  验证码获取超时（这是正常的，如果在30秒内没有收到验证码）")
		} else {
			t.Logf("⚠️  获取验证码失败: %v", err)
		}
	} else {
		t.Logf("✅ 验证码: %s", codeResp.Code)
		t.Logf("✅ 接收时间: %s", codeResp.ReceivedAt.Format(time.RFC3339))
	}

	// 步骤3: 释放手机号
	t.Log("\n========== 步骤3: 释放手机号 ==========")
	releaseCtx, releaseCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer releaseCancel()

	err = provider.ReleasePhone(releaseCtx, phoneResp.PhoneNumber, phoneResp.ExtId)
	if err != nil {
		t.Errorf("❌ 释放手机号失败: %v", err)
		return
	}

	t.Logf("✅ 成功释放手机号: %s (extId: %s)", phoneResp.PhoneNumber, phoneResp.ExtId)
	t.Log("\n========== 完整流程测试完成 ==========")
}
