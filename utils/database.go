package utils

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局的数据库连接实例
var (
	db     *gorm.DB
	dbOnec sync.Once
)

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
}

// NewDatabaseConfig 创建一个包含硬编码配置的新实例
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "zx197546",
		Database: "good_product",
		Charset:  "utf8mb4",
	}
}

// Init 初始化数据库连接
func Init() {
	var err error
	config := NewDatabaseConfig() // 使用硬编码的配置
	dsn := config.GetDSN()

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 如果数据库连接失败，程序直接 panic
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
}

// GetDSN 获取数据库连接字符串
func (dc *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dc.Username, dc.Password, dc.Host, dc.Port, dc.Database, dc.Charset)
}

func GetDB() *gorm.DB {
	dbOnec.Do(Init)
	return db
}
