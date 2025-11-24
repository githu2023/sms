package dto

// BusinessDTO 业务类型相关的数据传输对象

// BusinessTypeResponse 业务类型响应
type BusinessTypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// CustomerBusinessTypeResponse 客户业务类型响应
type CustomerBusinessTypeResponse struct {
	ID           int64  `json:"id"`
	BusinessCode string `json:"business_code"`
	BusinessName string `json:"business_name"`
	Weight       int    `json:"weight"`
}
