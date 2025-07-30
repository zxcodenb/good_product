// Package service contains business logic implementations
package service

import (
	"item-value/domain/dao"
	"item-value/domain/dto"
	"item-value/domain/model"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ItemService provides business logic for items
type ItemService struct {
	itemDAO *dao.ItemDAO
}

// NewItemService creates a new ItemService instance
func NewItemService() *ItemService {
	return &ItemService{
		itemDAO: dao.NewItemDAO(),
	}
}

// CreateItem creates a new item
func (s *ItemService) CreateItem(req dto.ItemCreateRequest) (*dto.ItemResponse, error) {
	item := &model.Item{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Remark:      req.Remark,
	}

	if err := s.itemDAO.Create(item); err != nil {
		return nil, err
	}

	return &dto.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Remark:      item.Remark,
		CreatedAt:   item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// GetItem retrieves an item by its ID
func (s *ItemService) GetItem(id string) (*dto.ItemResponse, error) {
	item, err := s.itemDAO.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Remark:      item.Remark,
		CreatedAt:   item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateItem updates an existing item
func (s *ItemService) UpdateItem(id string, req dto.ItemUpdateRequest) (*dto.ItemResponse, error) {
	item, err := s.itemDAO.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.Description != "" {
		item.Description = req.Description
	}
	if !req.Price.IsZero() {
		item.Price = req.Price
	}
	if req.Remark != "" {
		item.Remark = req.Remark
	}

	if err := s.itemDAO.Update(item); err != nil {
		return nil, err
	}

	return &dto.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Remark:      item.Remark,
		CreatedAt:   item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// DeleteItem deletes an item by its ID
func (s *ItemService) DeleteItem(id string) error {
	return s.itemDAO.Delete(id)
}

// ListItems retrieves all items
func (s *ItemService) ListItems() (*dto.ItemListResponse, error) {
	items, err := s.itemDAO.List()
	if err != nil {
		return nil, err
	}

	var itemResponses []dto.ItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, dto.ItemResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Remark:      item.Remark,
			CreatedAt:   item.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &dto.ItemListResponse{
		Items: itemResponses,
	}, nil
}

// CalculateTotalPrice calculates the total price of items
func (s *ItemService) CalculateTotalPrice(itemIds []string) (decimal.Decimal, error) {
	var total decimal.Decimal

	for _, id := range itemIds {
		item, err := s.itemDAO.GetByID(id)
		if err != nil {
			return decimal.Zero, err
		}
		total = total.Add(item.Price)
	}

	return total, nil
}