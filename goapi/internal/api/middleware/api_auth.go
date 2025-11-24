package middleware

import (
	"fmt"
	"sms-platform/goapi/internal/common"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// APITokenAuthMiddleware validates API JWT token from Authorization header.
func APITokenAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			common.RespondError(c, common.CodeUnauthorized)
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			common.RespondError(c, common.CodeUnauthorized)
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			common.RespondErrorWithMsg(c, common.CodeUnauthorized, fmt.Sprintf("Invalid token: %v", err))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check expiration
			if exp, ok := claims["exp"].(float64); ok {
				if int64(exp) < time.Now().Unix() {
					common.RespondError(c, common.CodeUnauthorized)
					c.Abort()
					return
				}
			}

			// Check audience claim for "api"
			if aud, ok := claims["aud"].(string); !ok || aud != "api" {
				common.RespondError(c, common.CodeForbidden)
				c.Abort()
				return
			}

			userID := claims["user_id"]
			// 转换为int64类型
			var customerID int64
			switch id := userID.(type) {
			case float64:
				customerID = int64(id)
			case int64:
				customerID = id
			case int:
				customerID = int64(id)
			default:
				common.RespondError(c, common.CodeUnauthorized)
				c.Abort()
				return
			}
			c.Set("customer_id", customerID) // 设置为customer_id且确保为int64
			c.Next()
		} else {
			common.RespondError(c, common.CodeUnauthorized)
			c.Abort()
		}
	}
}
