package testutils

import "github.com/gin-gonic/gin"

// SetupTestRouter sets up a minimal Gin router for testing with a mock customer_id.
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(func(c *gin.Context) {
		// Mock customer_id for authenticated routes
		c.Set("customer_id", int64(1)) // Using int64 for consistency
		c.Next()
	})
	return router
}
