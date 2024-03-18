package main

import (
	"amdCourier/exec/userServer/router"
	"os"
	"os/signal"
	"syscall"
)

// @title amdCourier-userServer
// @version 1.0
// @description amdCourier-userServer
// @contact.name 陶然
func main() {
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}
