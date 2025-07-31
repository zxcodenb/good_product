// Package dto contains data transfer objects
package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// PaginationRequest 定义了分页查询的基础请求参数
type PaginationRequest struct {
	Page     int `form:"page"`      // 页码
	PageSize int `form:"page_size"` // 每页数量
}

// ItemListRequest defines the request parameters for listing items with pagination and search
type ItemListRequest struct {
	PaginationRequest
	Name string `form:"name"` // Item name for fuzzy search
}

// ItemCreateRequest represents the request body for creating an item
type ItemCreateRequest struct {
	Name        string          `json:"name" binding:"required"` // Item name
	Description string          `json:"description"`             // Item description
	BuyTime     time.Time       `json:"buy_time"`
	Price       decimal.Decimal `json:"price" binding:"required"` // Item price
	Remark      string          `json:"remark"`                   // Remark
}

// ItemUpdateRequest represents the request body for updating an item
type ItemUpdateRequest struct {
	Name        string          `json:"name"`        // Item name
	Description string          `json:"description"` // Item description
	Price       decimal.Decimal `json:"price"`       // Item price
	Remark      string          `json:"remark"`      // Remark
}

// ItemResponse represents the response body for an item
type ItemResponse struct {
	ItemName     string          `json:"item_name"` // Item name
	AvatarUrl    string          `json:"avatar_url"`
	Description  string          `json:"description"` // Item description
	Price        decimal.Decimal `json:"price"`       // Item price
	AveragePrice decimal.Decimal `json:"average_price"`
	BuyTime      time.Time       `json:"buy_time"` // Buy time
	Days         int             `json:"days"`
	CreateTime   time.Time       `json:"create_time"`
	UpdateTime   time.Time       `json:"update_time"`
}

// ItemListResponse represents the response body for a list of items
type ItemListResponse struct {
	Items []ItemResponse `json:"items"` // List of items
}
