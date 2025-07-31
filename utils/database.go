package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func (config *DatabaseConfig) Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database, config.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitDB 初始化数据库连接并返回 *gorm.DB 实例
func InitDB() (*gorm.DB, error) {
	config := NewDatabaseConfig() // 使用硬编码的配置
	dsn := config.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return db, nil
}

// GetDSN 获取数据库连接字符串
func (dc *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dc.Username, dc.Password, dc.Host, dc.Port, dc.Database, dc.Charset)
}

// CloseDB 关闭数据库连接
func CloseDB(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("failed to get underlying sql.DB")
		}
		if err := sqlDB.Close(); err != nil {
			fmt.Println("failed to close database connection")
		}
	}
}
