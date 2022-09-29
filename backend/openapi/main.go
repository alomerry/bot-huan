package main

import (
	"bot-huan/openapi/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/tylerb/graceful.v1"
	"time"
)

func main() {
	var (
		r = gin.Default()
	)

	r.GET("/ping", controller.Ping)
	r.POST("/webhook", controller.Webhook())
	r.POST("/report", controller.Report)

	graceful.Run(fmt.Sprintf(":%s", "4376"), 5*time.Second, r)
}
