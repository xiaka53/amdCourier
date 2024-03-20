package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 捕获所有的panic，并返回错误信息
func RecoverMiddleware(model string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if model != "debug" {
					ResponseError(c, 500, errors.New("内部错误"))
					return
				} else {
					ResponseError(c, 500, errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
		c.Next()
	}
}
