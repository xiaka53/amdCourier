package router

import (
	_ "amdCourier/exec/userServer/docs"
	"amdCourier/exec/userServer/dto"
	"amdCourier/exec/userServer/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 路由初始化
func InitRouter(middlewares ...gin.HandlerFunc) (router *gin.Engine) {
	dto.InitValidate()
	router = gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares...)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.RecoverMiddleware(), middleware.IPAuthMiddleware())
	router.Use(middleware.AccessMiddleware()) // 跨域

	return
}
