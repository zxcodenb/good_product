// Package model contains the domain models
package model

import (
	"github.com/shopspring/decimal"
)

// Item represents an item entity
type Item struct {
	ID          string          `json:"id" gorm:"column:id;type:varchar(36);primaryKey"`       // Unique identifier
	Name        string          `json:"name" gorm:"column:name;type:varchar(100);not null"`    // Item name
	Description string          `json:"description" gorm:"column:description;type:text"`       // Item description
	Price       decimal.Decimal `json:"price" gorm:"column:price;type:decimal(10,2);not null"` // Item price
	Remark      string          `json:"remark" gorm:"column:remark;type:varchar(255)"`         // Remark
}

// TableName sets the table name for the Item model
func (Item) TableName() string {
	return "items"
}
