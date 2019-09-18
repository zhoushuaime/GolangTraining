package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
@Time : 2018/1/12 14:25
@Author : joshua
@File : httpServerByGin
@Software: GoLand
*/

var data []byte
var err error

// ReadeSwaggerFile ...
func ReadeSwaggerFile() {
	//data, err = ioutil.ReadFile("./openapi.json")
	data, err = ioutil.ReadFile("./openapi.yaml")
	if err != nil {
		log.Println("read file error:", err)
		return
	}
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Logger())
	server := http.Server{
		Addr:    ":8090",
		Handler: app,
	}
	r := app.Group("/v1/api/swagger.json", ErrorHandler)
	{
		r.GET("/", nil)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGKILL,
		os.Kill,
		syscall.SIGTERM,
		os.Interrupt,
	)
	go func() {
		select {
		case <-ch:
			println("shutdown...")
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("HTTP Server Shutdown error:", err)
			}
			log.Println("HTTP Server has exited")
		}
	}()
	log.Println("HTTP Server Has Started...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP Server start failed: %v\n", err)
	}

}

// ErrorHandler gin
func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.Writer.Status() == http.StatusOK {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(ctx.Writer.Status(), err)
		}
	}()
	ctx.Next()
}
