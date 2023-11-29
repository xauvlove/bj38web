package main

import (
	"bj38web/web/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化路由
	router := gin.Default()

	// 路由匹配
	router.Static("/home", "view")

	router.GET("/api/v1.0/session", controller.GetSession)

	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd)

	// 启动
	router.Run(":8080")
}
