package provider

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLocalProvider(t *testing.T) {
	provider := NewLocalProvider("test-001", "Test Provider", 100, nil)

	assert.NotNil(t, provider)
	assert.Equal(t, "test-001", provider.GetProviderInfo().ID)
	assert.Equal(t, "Test Provider", provider.GetProviderInfo().Name)
	assert.Equal(t, "local", provider.GetProviderInfo().Type)
	assert.Equal(t, 100, provider.GetProviderInfo().Priority)
	assert.Equal(t, 0.01, provider.GetProviderInfo().CostPerSMS)
	assert.True(t, provider.IsHealthy(context.Background()))
}

func TestLocalProvider_GetPhone_Success(t *testing.T) {
	provider := NewLocalProviderWithConfig("phone-test", "Phone Test", 100, LocalProviderConfig{
		SuccessRate: 100, // 100% 成功率确保测试通过
		MinDelayMs:  0,   // 无延迟
		MaxDelayMs:  0,
	})
	ctx := context.Background()

	response, err := provider.GetPhone(ctx, "qq", "physical")
	require.NoError(t, err)
	assert.NotEmpty(t, response.PhoneNumber)
	assert.NotEmpty(t, response.CountryCode)
	assert.Greater(t, response.Cost, 0.0)
	assert.True(t, time.Now().Before(response.ValidUntil))
}

func TestLocalProvider_GetPhone_UnsupportedBusinessType(t *testing.T) {
	provider := NewLocalProvider("unsupported-test", "Unsupported Test", 100, nil)
	ctx := context.Background()

	_, err := provider.GetPhone(ctx, "unsupported-type", "physical")
	assert.Error(t, err)
	assert.True(t, IsProviderError(err))
}

func TestLocalProvider_GetCode_Success(t *testing.T) {
	provider := NewLocalProviderWithConfig("code-test", "Code Test", 100, LocalProviderConfig{
		SuccessRate: 100,
		MinDelayMs:  0,
		MaxDelayMs:  10, // 很短的延迟
	})
	ctx := context.Background()

	// 先获取手机号
	phoneResp, err := provider.GetPhone(ctx, "qq", "physical")
	require.NoError(t, err)

	// 获取验证码
	codeResp, err := provider.GetCode(ctx, phoneResp.PhoneNumber, 5*time.Second)
	require.NoError(t, err)
	assert.NotEmpty(t, codeResp.Code)
	assert.NotEmpty(t, codeResp.Message)
	assert.Equal(t, provider.GetProviderInfo().ID, codeResp.ProviderID)
}

func TestLocalProvider_GetCode_InvalidPhone(t *testing.T) {
	// 注意：由于 LocalProvider.GetCode 现在可以处理不在内存中的手机号（使用默认业务类型）
	// 这个测试需要调整，或者测试一个真正无效的场景
	provider := NewLocalProviderWithConfig("invalid-phone", "Invalid Phone Test", 100, LocalProviderConfig{
		SuccessRate:   100,
		MinDelayMs:    0,
		MaxDelayMs:    0,
		BusinessTypes: nil, // 使用默认业务类型
	})
	ctx := context.Background()

	// 由于现在可以处理不在内存中的手机号，这个测试改为验证可以生成验证码
	codeResp, err := provider.GetCode(ctx, "invalid-phone-number", 5*time.Second)
	// 现在应该成功，因为会使用默认业务类型
	require.NoError(t, err)
	assert.NotEmpty(t, codeResp.Code)
}

func TestLocalProvider_HealthStatus(t *testing.T) {
	provider := NewLocalProvider("health-test", "Health Test", 100, nil)
	ctx := context.Background()

	assert.True(t, provider.IsHealthy(ctx))

	provider.SetHealthy(false)
	assert.False(t, provider.IsHealthy(ctx))

	provider.SetHealthy(true)
	assert.True(t, provider.IsHealthy(ctx))
}

func TestLocalProvider_GetStats(t *testing.T) {
	provider := NewLocalProvider("stats-test", "Stats Test", 100, nil)

	stats := provider.GetStats()
	assert.NotNil(t, stats)
	assert.Equal(t, "stats-test", stats["provider_id"])
	assert.Equal(t, "Stats Test", stats["provider_name"])
}

func TestLocalProvider_SupportedBusinessTypes(t *testing.T) {
	provider := NewLocalProvider("business-test", "Business Test", 100, nil)

	business := provider.GetSupportedBusiness()
	assert.NotNil(t, business)
	assert.Greater(t, len(business), 0)

	// 检查默认支持的业务类型
	assert.Contains(t, business, "qq")
	assert.Contains(t, business, "wechat")
}

func TestLocalProvider_ContextTimeout(t *testing.T) {
	// 创建一个延迟很高的provider来测试超时
	provider := NewLocalProviderWithConfig("timeout-test", "Timeout Test", 100, LocalProviderConfig{
		SuccessRate: 100,
		MinDelayMs:  500, // 500毫秒延迟
		MaxDelayMs:  600, // 600毫秒延迟
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := provider.GetPhone(ctx, "qq", "physical")
	assert.Error(t, err)
	assert.Equal(t, context.DeadlineExceeded, err)
}

func TestLocalProvider_ContextCancellation(t *testing.T) {
	provider := NewLocalProviderWithConfig("cancel-test", "Cancel Test", 100, LocalProviderConfig{
		SuccessRate: 100,
		MinDelayMs:  1000, // 1秒延迟
		MaxDelayMs:  1500,
	})

	ctx, cancel := context.WithCancel(context.Background())

	// 在另一个goroutine中取消
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := provider.GetPhone(ctx, "qq", "physical")
	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}

func TestLocalProvider_DifferentCountries(t *testing.T) {
	// 使用100%成功率确保测试通过
	provider := NewLocalProviderWithConfig("country-test", "Country Test", 100, LocalProviderConfig{
		SuccessRate: 100, // 100% 成功率
		MinDelayMs:  0,
		MaxDelayMs:  0,
	})
	ctx := context.Background()

	// 测试不同业务类型返回不同的国家代码
	countries := make(map[string]bool)
	for i := 0; i < 10; i++ {
		resp, err := provider.GetPhone(ctx, "qq", "physical")
		require.NoError(t, err)
		countries[resp.CountryCode] = true
	}

	// 应该至少有一个国家代码
	assert.Greater(t, len(countries), 0)
}

func TestLocalProvider_CardTypes(t *testing.T) {
	provider := NewLocalProvider("test-local", "http://localhost:8080", 80, nil)
	ctx := context.Background()

	// 测试物理卡
	physicalResp, err := provider.GetPhone(ctx, "qq", "physical")
	require.NoError(t, err)
	assert.NotEmpty(t, physicalResp.PhoneNumber)

	// 测试虚拟卡
	virtualResp, err := provider.GetPhone(ctx, "qq", "virtual")
	require.NoError(t, err)
	assert.NotEmpty(t, virtualResp.PhoneNumber)
}

func TestLocalProvider_WithDatabaseBusinessTypes(t *testing.T) {
	// 测试从数据库读取的业务类型配置
	businessTypes := []BusinessTypeConfig{
		{
			BusinessCode: "wx",
			BusinessName: "微信",
			Price:        1.0,
		},
		{
			BusinessCode: "qq",
			BusinessName: "QQ",
			Price:        0.5,
		},
	}

	provider := NewLocalProvider("db-test", "DB Test", 100, businessTypes)
	ctx := context.Background()

	// 测试 wx 业务类型
	wxResp, err := provider.GetPhone(ctx, "wx", "physical")
	require.NoError(t, err)
	assert.NotEmpty(t, wxResp.PhoneNumber)

	// 测试 qq 业务类型
	qqResp, err := provider.GetPhone(ctx, "qq", "physical")
	require.NoError(t, err)
	assert.NotEmpty(t, qqResp.PhoneNumber)

	// 测试不支持的业务类型
	_, err = provider.GetPhone(ctx, "unsupported", "physical")
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "不支持的业务类型"))
}
