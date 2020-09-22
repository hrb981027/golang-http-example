package main

import (
	"api-user-login/config"
	"api-user-login/route"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	conf := config.LoadConf()

	gin.SetMode(conf.AppModel)

	engine := gin.New()
	route.SetupRouter(engine)

	server := &http.Server{
		Addr: ":" + conf.AppPort,
		Handler: engine,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|              go-gin               |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|      Port:" + conf.AppPort + "      Pid:" + fmt.Sprintf("%d", os.Getpid()) + "      |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
