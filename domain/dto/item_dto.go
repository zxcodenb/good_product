// Package dto contains data transfer objects
package dto

import "github.com/shopspring/decimal"

// ItemCreateRequest represents the request body for creating an item
type ItemCreateRequest struct {
	Name        string          `json:"name" binding:"required"`        // Item name
	Description string          `json:"description"`                    // Item description
	Price       decimal.Decimal `json:"price" binding:"required"`       // Item price
	Remark      string          `json:"remark"`                         // Remark
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
	ID          string          `json:"id"`          // Unique identifier
	Name        string          `json:"name"`        // Item name
	Description string          `json:"description"` // Item description
	Price       decimal.Decimal `json:"price"`       // Item price
	Remark      string          `json:"remark"`      // Remark
	CreatedAt   string          `json:"created_at"`  // Creation time
	UpdatedAt   string          `json:"updated_at"`  // Update time
}

// ItemListResponse represents the response body for a list of items
type ItemListResponse struct {
	Items []ItemResponse `json:"items"` // List of items
}