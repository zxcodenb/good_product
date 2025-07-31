package routes

import (
	"item-value/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter 用于初始化所有路由
func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// 创建用户控制器实例
	userController := controllers.NewUserController(db)
	itemController := controllers.NewUItemController(db)

	userRoutes := r.Group("/users")
	{
		// 注册用户相关路由
		userRoutes.POST("/user", userController.CreateUser)
		userRoutes.POST("/login", userController.Login)

	}

	// 注册物品相关路由
	itemRoutes := r.Group("/items")
	{
		itemRoutes.POST("", itemController.CreateItem)
		itemRoutes.GET("", itemController.List)
		itemRoutes.GET("/:id", itemController.GetItemById)
	}

}
