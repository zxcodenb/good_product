package routes

import (
	"item-value/api"

	"github.com/gin-gonic/gin"
)

// SetupRouter 用于初始化所有路由
func SetupRouter(r *gin.Engine) {
	// 创建用户控制器实例
	userController := &api.UserController{}

	// 注册用户相关路由
	r.POST("/user", userController.CreateUser)
	r.POST("/login", userController.Login)

}
