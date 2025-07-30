// storage包提供内存存储实现
package storage

import (
	"fmt"
	"item-value/internal/models"
	"sync"

	"github.com/google/uuid"
)

// MemoryStorage 是Store接口的内存实现
type MemoryStorage struct {
	mu    sync.RWMutex        // 读写锁，保证并发安全
	items map[string]models.Item // 项目存储映射表，键为ID，值为项目
}

// NewMemoryStorage 创建并返回一个新的MemoryStorage实例
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		items: make(map[string]models.Item),
	}
}

// CreateItem 向存储中添加一个新项目
func (s *MemoryStorage) CreateItem(item models.Item) (models.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查项目名称是否为空
	if item.Name == "" {
		return models.Item{}, fmt.Errorf("项目名称不能为空")
	}

	// 生成唯一ID并存储项目
	item.ID = uuid.New().String()
	s.items[item.ID] = item
	return item, nil
}

// GetItem 根据ID从存储中检索项目
func (s *MemoryStorage) GetItem(id string) (models.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 查找指定ID的项目
	item, exists := s.items[id]
	if !exists {
		return models.Item{}, fmt.Errorf("未找到ID为'%s'的项目", id)
	}
	return item, nil
}

// UpdateItem 更新存储中的现有项目
func (s *MemoryStorage) UpdateItem(id string, item models.Item) (models.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查项目是否存在
	_, exists := s.items[id]
	if !exists {
		return models.Item{}, fmt.Errorf("未找到ID为'%s'的项目", id)
	}

	item.ID = id // 确保ID保持不变
	s.items[id] = item
	return item, nil
}

// DeleteItem 根据ID从存储中删除项目
func (s *MemoryStorage) DeleteItem(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查项目是否存在
	_, exists := s.items[id]
	if !exists {
		return fmt.Errorf("未找到ID为'%s'的项目", id)
	}

	// 删除项目
	delete(s.items, id)
	return nil
}

// ListItems 返回存储中的所有项目
func (s *MemoryStorage) ListItems() ([]models.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 遍历所有项目并添加到切片中
	var items []models.Item
	for _, item := range s.items {
		items = append(items, item)
	}
	return items, nil
}