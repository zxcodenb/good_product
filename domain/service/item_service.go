package service

import (
	"item-value/domain/dao"
	"item-value/domain/dto"
	"item-value/domain/model"
	"item-value/utils"

	"gorm.io/gorm"
)

type ItemService struct {
	itemDAO *dao.ItemDAO
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{
		itemDAO: dao.NewItemDAO(db),
	}
}

// 新增物品
func (s *ItemService) CreateItem(req dto.ItemCreateRequest) (bool, error) {
	item := model.Item{
		ID:          utils.GenerateSnowflakeID(),
		ItemName:    req.Name,
		Price:       utils.NewFromFloat(req.Price.InexactFloat64()),
		BuyTime:     req.BuyTime,
		Description: req.Description,
	}
	err := s.itemDAO.Create(&item)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 修改物品
func (s *ItemService) UpdateItem(id string, req dto.ItemUpdateRequest) (bool, error) {
	item, err := s.itemDAO.GetByItemID(id)
	if err != nil {
		return false, err
	}
	item.ItemName = req.Name
	item.Price = utils.NewFromFloat(req.Price.InexactFloat64())
	item.Description = req.Description
	err = s.itemDAO.Update(item)
	if err != nil {
		return false, err
	}
	return true, nil

}

// 获取物品详情
func (s *ItemService) GetItem(id string) (*dto.ItemResponse, error) {

	item, err := s.itemDAO.GetByItemID(id)
	if err != nil {
		return nil, err
	}
	return &dto.ItemResponse{
		ItemName:     item.ItemName,
		AvatarUrl:    item.AvatarUrl,
		Price:        item.Price,
		AveragePrice: utils.AveragePrice(item.Price, item.BuyTime),
		BuyTime:      item.BuyTime,
		Days:         utils.GetDays(item.BuyTime),
		Description:  item.Description,
		CreateTime:   item.CreateTime,
		UpdateTime:   item.UpdateTime,
	}, nil
}

// 获取物品列表
func (s *ItemService) ListItems(req *dto.ItemListRequest) ([]*dto.ItemResponse, error) {
	items, err := s.itemDAO.List(req)

	if err != nil {
		return nil, err
	}

	responses := make([]*dto.ItemResponse, len(items))
	for i, item := range items {

		responses[i] = &dto.ItemResponse{
			ItemName:     item.ItemName,
			Price:        item.Price,
			AveragePrice: utils.AveragePrice(item.Price, item.BuyTime),
			AvatarUrl:    item.AvatarUrl,
			BuyTime:      item.BuyTime,
			Days:         utils.GetDays(item.BuyTime),
		}
	}
	return responses, nil
}
