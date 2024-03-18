package router

import (
	"amdCourier/server"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

// 启动服务
func HttpServerRun() {
	switch os.Getenv("GOLANG_SERVER_USERSERVER_DEBUG") {
	case "test":
		gin.SetMode(gin.TestMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           os.Getenv("GOLANG_SERVER_USERSERVER_ADDR"),
		Handler:        r,
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
		MaxHeaderBytes: 1 << uint(20),
	}

	if HttpSrvHandler.Addr == "" {
		HttpSrvHandler.Addr = "8552"
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun%s\n", HttpSrvHandler.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun%s err:%v\n", HttpSrvHandler.Addr, err)
		}
	}()
}

// 关闭服务
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.ServerDatabaseClose()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stoppend\n")
}
