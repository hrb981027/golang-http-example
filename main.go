package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-http-example/global"
	"golang-http-example/initialize"
	"golang-http-example/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	initialize.Viper()
	initialize.Gorm()
	initialize.Redis()

	Run()
}

func Run() {
	var err error

	gin.SetMode("release")

	engine := gin.New()
	route.SetupRouter(engine)

	server := &http.Server{
		Addr:    ":" + global.CONFIG.System.Port,
		Handler: engine,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|              go-gin               |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|      Port:" + global.CONFIG.System.Port + "      Pid:" + fmt.Sprintf("%d", os.Getpid()) + "     |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
