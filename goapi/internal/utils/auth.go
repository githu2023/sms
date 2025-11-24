package utils

import (
	"errors"
	"sms-platform/goapi/internal/common"

	"github.com/gin-gonic/gin"
)

// ErrUnauthorized 表示未授权错误
var ErrUnauthorized = errors.New("unauthorized")

// GetCustomerIDFromContext 从Gin上下文中获取customerID
// 现在统一处理int64类型，因为中间件已经确保设置为int64
func GetCustomerIDFromContext(c *gin.Context) (int64, error) {
	customerID, exists := c.Get("customer_id")
	if !exists {
		return 0, ErrUnauthorized
	}

	// 中间件应该已经设置为int64类型
	if id, ok := customerID.(int64); ok {
		if id <= 0 {
			return 0, ErrUnauthorized
		}
		return id, nil
	}

	// 如果不是int64，说明中间件设置有问题
	return 0, ErrUnauthorized
}

// RequireCustomerID 是GetCustomerIDFromContext的便捷封装
// 如果获取失败，直接响应错误并返回false，成功则返回true和customerID
func RequireCustomerID(c *gin.Context) (int64, bool) {
	customerID, err := GetCustomerIDFromContext(c)
	if err != nil {
		common.RespondError(c, common.CodeUnauthorized)
		return 0, false
	}
	return customerID, true
}
