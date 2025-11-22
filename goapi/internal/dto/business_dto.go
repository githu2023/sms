package dto

// BusinessDTO 业务类型相关的数据传输对象

// BusinessTypeResponse 业务类型响应
type BusinessTypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
