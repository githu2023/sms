package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func main() {
	client := &http.Client{Timeout: 15 * time.Second}

	token := login(client)
	log.Printf("Token: %s", token)

	initialBalance := getBalance(client, token)
	fmt.Printf("初始余额: %.2f\n", initialBalance)

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(idx int) {
			defer wg.Done()
			if err := getPhone(client, token); err != nil {
				fmt.Printf("[%d] get_phone failed: %v\n", idx, err)
			}
		}(i)
	}
	wg.Wait()

	finalBalance := getBalance(client, token)
	fmt.Printf("完成后余额: %.2f\n", finalBalance)
	fmt.Printf("扣减: %.2f\n", initialBalance-finalBalance)
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

func getBalance(client *http.Client, token string) float64 {
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
		Currency string  `json:"currency"`
	}
	if err := json.Unmarshal(wrapper.Data, &data); err != nil {
		log.Fatalf("balance decode data: %v", err)
	}
	return data.Balance
}

func getPhone(client *http.Client, token string) error {
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
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}
