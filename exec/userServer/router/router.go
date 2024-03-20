package router

import (
	_ "amdCourier/exec/userServer/docs"
	"amdCourier/exec/userServer/dto"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 路由初始化
func InitRouter(middlewares ...gin.HandlerFunc) (router *gin.Engine) {
	dto.InitValidate()
	router = gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares...)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
