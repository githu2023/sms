package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 配置
const (
	BaseURL         = "http://localhost:6060"
	TestUser        = "testuser_1763969884"
	TestPassword    = "TestPassword123!"
	APISecret       = "363dbbc5a6519dcf604579f57f2ef87f6b405f40aceef7db01f1ede27f18f7fd"
	MerchantNo      = "532401"
	ClientAPIPrefix = "/client/v1"
	APIPrefix       = "/api/v1"
)

// 响应结构体
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

type APITokenResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
	UserID  int64   `json:"user_id,omitempty"`
}

type UserProfileResponse struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
}

type PhoneResponse struct {
	PhoneNumber string  `json:"phone_number"`
	CountryCode string  `json:"country_code"`
	Cost        float64 `json:"cost"`
	ValidUntil  string  `json:"valid_until"`
	ProviderID  string  `json:"provider_id"`
	Balance     float64 `json:"balance"`
}

type CodeResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	PhoneInfo []struct {
		PhoneNumber string `json:"phone_number"`
		Code        string `json:"code"`
		Status      string `json:"status"`
		Message     string `json:"message"`
	} `json:"phone_info"`
}

type GetPhoneRequest struct {
	BusinessType string `json:"business_type"`
	CardType     string `json:"card_type"`
	Count        int    `json:"count"`
}

type GetCodeRequest struct {
	PhoneNumbers []string `json:"phone_numbers"`
}

var client = &http.Client{
	Timeout: 30 * time.Second,
}

var clientToken, apiToken string

// 发送HTTP请求
func makeRequest(method, url string, headers map[string]string, body interface{}) (*APIResponse, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body error: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v, body: %s", err, string(respBody))
	}

	return &apiResp, nil
}

func parseResponseData(resp *APIResponse, target interface{}) error {
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return fmt.Errorf("marshal data error: %v", err)
	}
	return json.Unmarshal(dataBytes, target)
}

// 测试用例函数
func testUserLogin() error {
	fmt.Println("\n[1] 测试用户登录...")
	loginData := map[string]string{
		"username": TestUser,
		"password": TestPassword,
	}

	resp, err := makeRequest("POST", BaseURL+ClientAPIPrefix+"/login", nil, loginData)
	if err != nil {
		return fmt.Errorf("登录失败: %v", err)
	}

	if resp.Code != 200 {
		return fmt.Errorf("登录失败: code=%d, message=%s", resp.Code, resp.Message)
	}

	var loginResp LoginResponse
	if err := parseResponseData(resp, &loginResp); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	clientToken = loginResp.Token
	fmt.Printf("✅ 用户登录成功，Token: %s...\n", clientToken[:min(50, len(clientToken))])
	return nil
}

func testGetAPIToken() error {
	fmt.Println("\n[2] 测试API Token获取...")
	tokenData := map[string]string{
		"merchant_no": MerchantNo,
		"secret":      APISecret,
	}

	resp, err := makeRequest("POST", BaseURL+APIPrefix+"/get_token", nil, tokenData)
	if err != nil {
		return fmt.Errorf("获取API Token失败: %v", err)
	}

	if resp.Code != 200 {
		return fmt.Errorf("获取API Token失败: code=%d, message=%s", resp.Code, resp.Message)
	}

	var tokenResp APITokenResponse
	if err := parseResponseData(resp, &tokenResp); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	apiToken = tokenResp.Token
	fmt.Printf("✅ API Token获取成功，Token: %s...\n", apiToken[:min(50, len(apiToken))])
	return nil
}

func testGetUserProfile() error {
	fmt.Println("\n[3] 测试获取用户信息...")
	headers := map[string]string{"Authorization": "Bearer " + clientToken}
	resp, err := makeRequest("GET", BaseURL+ClientAPIPrefix+"/profile", headers, nil)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %v", err)
	}

	if resp.Code != 200 {
		return fmt.Errorf("获取用户信息失败: code=%d, message=%s", resp.Code, resp.Message)
	}

	var profile UserProfileResponse
	if err := parseResponseData(resp, &profile); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	fmt.Printf("✅ 用户信息获取成功: ID=%d, 用户名=%s, 邮箱=%s, 余额=%.2f\n",
		profile.ID, profile.Username, profile.Email, profile.Balance)
	return nil
}

func testGetBalance(apiPrefix, token, apiType string) error {
	fmt.Printf("\n[4] 测试%s余额查询...\n", apiType)
	headers := map[string]string{"Authorization": "Bearer " + token}
	resp, err := makeRequest("GET", BaseURL+apiPrefix+"/balance", headers, nil)
	if err != nil {
		return fmt.Errorf("%s余额查询失败: %v", apiType, err)
	}

	if resp.Code != 200 {
		return fmt.Errorf("%s余额查询失败: code=%d, message=%s", apiType, resp.Code, resp.Message)
	}

	var balance BalanceResponse
	if err := parseResponseData(resp, &balance); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	fmt.Printf("✅ %s余额查询成功，余额: %.2f\n", apiType, balance.Balance)
	return nil
}

func testGetPhone(apiPrefix, token, apiType string) error {
	fmt.Printf("\n[5] 测试%s获取手机号...\n", apiType)
	headers := map[string]string{"Authorization": "Bearer " + token}
	phoneData := GetPhoneRequest{
		BusinessType: "wx",
		CardType:     "physical",
		Count:        1,
	}

	resp, err := makeRequest("POST", BaseURL+apiPrefix+"/get_phone", headers, phoneData)
	if err != nil {
		return fmt.Errorf("%s获取手机号失败: %v", apiType, err)
	}

	if resp.Code != 200 {
		return fmt.Errorf("%s获取手机号失败: code=%d, message=%s", apiType, resp.Code, resp.Message)
	}

	// 解析响应
	dataBytes, _ := json.Marshal(resp.Data)
	var phoneDataResp struct {
		Phones []PhoneResponse `json:"phones"`
	}
	if err := json.Unmarshal(dataBytes, &phoneDataResp); err == nil && len(phoneDataResp.Phones) > 0 {
		phone := phoneDataResp.Phones[0]
		fmt.Printf("✅ %s获取手机号成功: %s, 成本: %.2f, 余额: %.2f\n",
			apiType, phone.PhoneNumber, phone.Cost, phone.Balance)
		return nil
	}

	fmt.Printf("✅ %s获取手机号成功\n", apiType)
	return nil
}

func testGetCode(apiPrefix, token, apiType string, phoneNumbers []string) error {
	if len(phoneNumbers) == 0 {
		return fmt.Errorf("没有手机号可测试")
	}

	fmt.Printf("\n[6] 测试%s获取验证码...\n", apiType)
	headers := map[string]string{"Authorization": "Bearer " + token}
	codeData := GetCodeRequest{
		PhoneNumbers: phoneNumbers,
	}

	resp, err := makeRequest("POST", BaseURL+apiPrefix+"/get_code", headers, codeData)
	if err != nil {
		return fmt.Errorf("%s获取验证码失败: %v", apiType, err)
	}

	if resp.Code != 200 {
		fmt.Printf("⚠️  %s获取验证码返回: code=%d, message=%s (可能是等待中)\n", apiType, resp.Code, resp.Message)
		return nil
	}

	// 解析响应
	dataBytes, _ := json.Marshal(resp.Data)
	var codeResp CodeResponse
	if err := json.Unmarshal(dataBytes, &codeResp); err == nil {
		if len(codeResp.PhoneInfo) > 0 {
			info := codeResp.PhoneInfo[0]
			if info.Code != "" {
				fmt.Printf("✅ %s获取验证码成功: %s -> %s\n", apiType, info.PhoneNumber, info.Code)
			} else {
				fmt.Printf("⚠️  %s验证码等待中: %s (状态: %s)\n", apiType, info.PhoneNumber, info.Status)
			}
		}
	}

	return nil
}

func testUnauthorizedAccess() {
	fmt.Println("\n[7] 测试无Token访问...")
	headers := map[string]string{}
	resp, err := makeRequest("GET", BaseURL+ClientAPIPrefix+"/balance", headers, nil)
	if err != nil {
		fmt.Printf("✅ 无Token访问正确失败: %v\n", err)
		return
	}

	// 401 或 40101 都表示未授权，应该被拒绝
	if resp.Code == 401 || resp.Code == 40101 {
		fmt.Printf("✅ 无Token访问正确被拒绝: code=%d, message=%s\n", resp.Code, resp.Message)
	} else {
		fmt.Printf("❌ 无Token访问应该被拒绝但却成功了: code=%d\n", resp.Code)
	}
}

// 测试过期退款功能
func testExpiredRefund() error {
	fmt.Println("\n[8] 测试过期退款功能...")
	fmt.Println("注意: 此测试需要手动设置分配记录为过期状态，或等待定时器处理")
	fmt.Println("测试步骤:")
	fmt.Println("1. 获取一个手机号")
	fmt.Println("2. 手动修改数据库中的分配记录创建时间为过期时间")
	fmt.Println("3. 等待定时器运行（约5秒）")
	fmt.Println("4. 检查余额是否已退款")
	fmt.Println("5. 检查交易记录中是否有退款记录")
	return nil
}

func main() {
	// 支持命令行参数
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "balance":
			// 只测试余额
			if err := testUserLogin(); err != nil {
				fmt.Printf("❌ %v\n", err)
				return
			}
			if err := testGetAPIToken(); err != nil {
				fmt.Printf("❌ %v\n", err)
				return
			}
			if clientToken != "" {
				testGetBalance(ClientAPIPrefix, clientToken, "客户端")
			}
			if apiToken != "" {
				testGetBalance(APIPrefix, apiToken, "编程API")
			}
			return
		case "phone":
			// 只测试获取手机号
			if err := testUserLogin(); err != nil {
				fmt.Printf("❌ %v\n", err)
				return
			}
			if err := testGetAPIToken(); err != nil {
				fmt.Printf("❌ %v\n", err)
				return
			}
			if apiToken != "" {
				testGetPhone(APIPrefix, apiToken, "编程API")
			}
			return
		case "code":
			// 只测试获取验证码（需要先有手机号）
			if len(os.Args) < 3 {
				fmt.Println("用法: go run api_tester.go code <phone_number>")
				return
			}
			if err := testGetAPIToken(); err != nil {
				fmt.Printf("❌ %v\n", err)
				return
			}
			if apiToken != "" {
				testGetCode(APIPrefix, apiToken, "编程API", []string{os.Args[2]})
			}
			return
		case "refund":
			// 测试过期退款
			testExpiredRefund()
			return
		}
	}

	// 完整测试流程
	fmt.Println("=== SMS平台完整API测试 ===")
	fmt.Printf("测试用户: %s\n", TestUser)
	fmt.Printf("商户号: %s\n", MerchantNo)
	fmt.Printf("API密钥: %s...\n", APISecret[:min(20, len(APISecret))])
	fmt.Println("========================================")

	// 1. 用户认证
	if err := testUserLogin(); err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	// 2. API Token获取
	if err := testGetAPIToken(); err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	// 3. 获取用户信息
	if err := testGetUserProfile(); err != nil {
		fmt.Printf("❌ %v\n", err)
	}

	// 4. 测试客户端API余额
	if clientToken != "" {
		testGetBalance(ClientAPIPrefix, clientToken, "客户端")
	}

	// 5. 测试编程API余额
	if apiToken != "" {
		testGetBalance(APIPrefix, apiToken, "编程API")
	}

	// 6. 测试获取手机号
	if apiToken != "" {
		if err := testGetPhone(APIPrefix, apiToken, "编程API"); err == nil {
			// 如果获取手机号成功，可以测试获取验证码
			// 注意：这里需要从响应中获取手机号，简化处理
			fmt.Println("\n提示: 获取验证码测试需要手机号，请使用: go run api_tester.go code <phone_number>")
		}
	}

	// 7. 测试无Token访问
	testUnauthorizedAccess()

	// 8. 测试过期退款说明
	testExpiredRefund()

	fmt.Println("\n========================================")
	fmt.Println("✅ API测试完成")
	fmt.Println("\n可用命令:")
	fmt.Println("  go run api_tester.go balance  - 只测试余额")
	fmt.Println("  go run api_tester.go phone    - 只测试获取手机号")
	fmt.Println("  go run api_tester.go code <phone> - 测试获取验证码")
	fmt.Println("  go run api_tester.go refund   - 测试过期退款说明")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
