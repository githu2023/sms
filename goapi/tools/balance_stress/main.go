package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL      = "http://localhost:6060"
	username     = "testuser_1763969884"
	password     = "TestPassword123!"
	businessType = "wx"
	cardType     = "physical"
	concurrency  = 100
)

type balanceInfo struct {
	Balance float64 `json:"balance"`
	Frozen  float64 `json:"frozen_amount"`
}

type phoneResult struct {
	Cost float64
	Err  error
}

func main() {
	client := &http.Client{Timeout: 15 * time.Second}

	token := login(client)
	log.Printf("Token: %s", token)

	initial := getBalance(client, token)
	fmt.Printf("初始余额: %.2f, 冻结: %.2f\n", initial.Balance, initial.Frozen)

	resultsCh := make(chan phoneResult, concurrency)

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(idx int) {
			defer wg.Done()
			cost, err := getPhone(client, token)
			if err != nil {
				fmt.Printf("[%d] get_phone failed: %v\n", idx, err)
				resultsCh <- phoneResult{Err: err}
				return
			}
			resultsCh <- phoneResult{Cost: cost}
		}(i)
	}
	wg.Wait()
	close(resultsCh)

	var successCount, failCount int
	totalCost := 0.0
	for res := range resultsCh {
		if res.Err != nil {
			failCount++
			continue
		}
		successCount++
		totalCost += res.Cost
	}

	final := getBalance(client, token)
	fmt.Printf("完成后余额: %.2f, 冻结: %.2f\n", final.Balance, final.Frozen)
	fmt.Printf("成功次数: %d, 失败次数: %d, 总扣费: %.2f\n", successCount, failCount, totalCost)

	expectedBalance := initial.Balance - totalCost
	if math.Abs(final.Balance-expectedBalance) <= 0.01 {
		fmt.Println("✅ 余额校验通过（考虑浮点误差）。")
	} else {
		fmt.Printf("⚠️ 余额异常：期望 %.2f, 实际 %.2f\n", expectedBalance, final.Balance)
	}
}

func login(client *http.Client) string {
	reqBody := map[string]string{
		"username": username,
		"password": password,
	}
	b, _ := json.Marshal(reqBody)
	resp, err := client.Post(baseURL+"/client/v1/login", "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("login request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var wrapper struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		log.Fatalf("login decode wrapper: %v", err)
	}
	var data struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(wrapper.Data, &data); err != nil {
		log.Fatalf("login decode data: %v", err)
	}
	if data.Token == "" {
		log.Fatalf("login failed: %s", string(body))
	}
	return data.Token
}

func getBalance(client *http.Client, token string) balanceInfo {
	req, _ := http.NewRequest("GET", baseURL+"/client/v1/balance", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("balance request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var wrapper struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		log.Fatalf("balance decode wrapper: %v", err)
	}
	var data struct {
		Balance  float64 `json:"balance"`
		Frozen   float64 `json:"frozen_amount"`
		Currency string  `json:"currency"`
	}
	if err := json.Unmarshal(wrapper.Data, &data); err != nil {
		log.Fatalf("balance decode data: %v", err)
	}
	return balanceInfo{Balance: data.Balance, Frozen: data.Frozen}
}

func getPhone(client *http.Client, token string) (float64, error) {
	reqBody := map[string]interface{}{
		"business_type": businessType,
		"card_type":     cardType,
		"count":         1,
	}
	b, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", baseURL+"/client/v1/get_phone", bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var wrapper struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return 0, fmt.Errorf("decode wrapper: %w", err)
	}
	if wrapper.Code != 200 {
		return 0, fmt.Errorf("code %d", wrapper.Code)
	}

	var data struct {
		Phones []struct {
			Cost float64 `json:"cost"`
		} `json:"phones"`
	}
	if err := json.Unmarshal(wrapper.Data, &data); err != nil {
		return 0, fmt.Errorf("decode data: %w", err)
	}
	if len(data.Phones) == 0 {
		return 0, fmt.Errorf("no phone returned")
	}
	return data.Phones[0].Cost, nil
}
