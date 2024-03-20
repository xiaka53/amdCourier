package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// 白名单过滤
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range strings.Split(os.Getenv("GOLANG_SERVER_SERVER_IPAUTH"), ",") {
			if c.ClientIP() == host || host == "*" {
				isMatched = true
			}
		}
		if !isMatched {
			ResponseError(c, InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
