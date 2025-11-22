package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse 统一API响应结构
type APIResponse struct {
	Code ErrorCode   `json:"code"`           // 错误码
	Msg  string      `json:"msg"`            // 响应消息
	Data interface{} `json:"data,omitempty"` // 响应数据，可能为nil
}

// SuccessResponse 成功响应
func SuccessResponse(data interface{}) *APIResponse {
	return &APIResponse{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Message(),
		Data: data,
	}
}

// ErrorResponse 错误响应
func ErrorResponse(code ErrorCode) *APIResponse {
	return &APIResponse{
		Code: code,
		Msg:  code.Message(),
		Data: nil,
	}
}

// ErrorResponseWithMsg 自定义消息的错误响应
func ErrorResponseWithMsg(code ErrorCode, msg string) *APIResponse {
	return &APIResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// RespondSuccess 发送成功响应
func RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse(data))
}

// RespondError 发送错误响应
func RespondError(c *gin.Context, code ErrorCode) {
	c.JSON(code.HTTPStatus(), ErrorResponse(code))
}

// RespondErrorWithMsg 发送自定义消息的错误响应
func RespondErrorWithMsg(c *gin.Context, code ErrorCode, msg string) {
	c.JSON(code.HTTPStatus(), ErrorResponseWithMsg(code, msg))
}

// Pagination 分页信息结构
type Pagination struct {
	Total int `json:"total"` // 总记录数
	Page  int `json:"page"`  // 当前页码
	Limit int `json:"limit"` // 每页数量
	Pages int `json:"pages"` // 总页数
}

// PagedResponse 分页响应结构
type PagedResponse struct {
	Items      interface{} `json:"items"`      // 数据列表
	Pagination Pagination  `json:"pagination"` // 分页信息
}

// NewPagedResponse 创建分页响应
func NewPagedResponse(items interface{}, total, page, limit int) *PagedResponse {
	pages := (total + limit - 1) / limit // 计算总页数
	return &PagedResponse{
		Items: items,
		Pagination: Pagination{
			Total: total,
			Page:  page,
			Limit: limit,
			Pages: pages,
		},
	}
}

// RespondPagedSuccess 发送分页成功响应
func RespondPagedSuccess(c *gin.Context, items interface{}, total, page, limit int) {
	data := NewPagedResponse(items, total, page, limit)
	RespondSuccess(c, data)
}
