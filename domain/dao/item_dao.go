package dao

import (
	"item-value/config"
	"item-value/domain/model"
)

func GetByItemID(id string) (*model.Item, error) {

	var item model.Item

	result := config.DB.Where("id = ?", id).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}
