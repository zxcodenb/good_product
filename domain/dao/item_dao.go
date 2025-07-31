package dao

import (
	"item-value/domain/dto"
	"item-value/domain/model"

	"gorm.io/gorm"
)

// ItemDAO 结构体用于封装与物品相关的数据库操作
type ItemDAO struct {
	db *gorm.DB
}

// NewItemDAO 是一个构造函数，用于创建一个新的 ItemDAO 实例
func NewItemDAO(db *gorm.DB) *ItemDAO {
	return &ItemDAO{
		db: db,
	}
}

// GetByItemID 根据ID从数据库中查找物品
func (d *ItemDAO) GetByItemID(id string) (*model.Item, error) {
	var item model.Item
	result := d.db.Where("id = ?", id).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

// Create 创建一个新的物品
func (d *ItemDAO) Create(item *model.Item) error {
	result := d.db.Create(item)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update 更新物品信息
func (d *ItemDAO) Update(item *model.Item) error {
	result := d.db.Save(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete 删除物品
func (d *ItemDAO) Delete(id string) error {
	result := d.db.Delete(&model.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 分页查询物品
func (d *ItemDAO) List(req *dto.ItemListRequest) ([]*model.Item, error) {
	var items []*model.Item
	query := d.db.Model(&model.Item{})

	// Add search condition for name if provided
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	offset := (req.Page - 1) * req.PageSize
	result := query.Offset(offset).Limit(req.PageSize).Order("create_time DESC").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
