// Package model contains the domain models
package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Item represents an item entity
type Item struct {
	ID          string          `json:"id" gorm:"column:id;type:varchar(255);primaryKey"`        // Unique identifier
	ItemName    string          `json:"item_name" gorm:"column:name;type:varchar(100);not null"` // Item name
	AvatarUrl   string          `json:"avatar_url" gorm:"column:avatar;type:varchar(255)"`
	Description string          `json:"description" gorm:"column:description;type:text"`       // Item description
	Price       decimal.Decimal `json:"price" gorm:"column:price;type:decimal(10,2);not null"` // Item price
	BuyTime     time.Time       `json:"buy_time" gorm:"column:buy_time;type:datetime"`         // Buy time
	CreateTime  time.Time       `json:"create_time" gorm:"column:create_time;type:datetime"`   // Item creation time
	UpdateTime  time.Time       `json:"update_time" gorm:"column:update_time;type:datetime"`   // Item update time
}

// TableName sets the table name for the Item model
func (Item) TableName() string {
	return "items"
}
