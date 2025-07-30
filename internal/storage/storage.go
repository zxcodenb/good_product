// storage包定义了存储接口和实现
package storage

import "item-value/internal/models"

// Store 定义了项目存储的接口
type Store interface {
	// CreateItem 创建一个新的项目
	CreateItem(item models.Item) (models.Item, error)
	
	// GetItem 根据ID获取项目
	GetItem(id string) (models.Item, error)
	
	// UpdateItem 根据ID更新项目
	UpdateItem(id string, item models.Item) (models.Item, error)
	
	// DeleteItem 根据ID删除项目
	DeleteItem(id string) error
	
	// ListItems 获取所有项目列表
	ListItems() ([]models.Item, error)
}