package main

import (
	"item-value/routes"

	"github.com/gin-gonic/gin"
)

// @title Item Value API
// @version 1.0
// @description This is a sample server for item value management.
// @host localhost:8080
// @BasePath /

func main() {
	// 创建 Gin 实例
	r := gin.Default()

	// 注册路由
	routes.SetupRouter(r)

	// 启动服务器
	r.Run(":8080")
}