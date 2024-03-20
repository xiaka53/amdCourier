package middleware

import (
	"github.com/gin-gonic/gin"
)

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("origin")) // 允许来源域
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}
