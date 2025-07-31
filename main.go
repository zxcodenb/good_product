package main

import (
	"item-value/routes"
	"item-value/utils"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// 初始化数据库连接
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	// 程序退出时关闭数据库连接
	defer closeDB(db)

	// 创建 Gin 实例
	r := gin.Default()

	// 注册路由
	routes.SetupRouter(r, db)

	// 启动服务器
	r.Run(":8080")
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to get sql.DB: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("failed to close database: %v", err)
	}
}
