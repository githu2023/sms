package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"

// RequestID adds a unique request ID to each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(RequestIDHeader, requestID)
		c.Writer.Header().Set(RequestIDHeader, requestID)
		c.Next()
	}
}
