// main.go是项目的入口文件
package main

import (
	"item-value/config" // 导入config包
	"item-value/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main函数是程序的入口点
func main() {
	// 初始化数据库连接
	config.Init()
	log.Println("数据库连接成功")

	// 创建路由
	router := mux.NewRouter()

	// 创建控制器实例并注册路由
	itemController := controller.NewItemController()
	userController := controller.NewUserController()
	itemController.RegisterRoutes(router)
	userController.RegisterRoutes(router)

	// 启动HTTP服务器
	port := ":8080"
	log.Printf("服务器启动，监听端口 %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("服务器启动失败: %s\n", err)
	}
}
