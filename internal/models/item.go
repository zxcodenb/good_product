// models包定义了数据模型
package models

import "github.com/shopspring/decimal"

// Item 物品实体
type Item struct {
	ID          string          `json:"id"`          // 唯一标识符
	Name        string          `json:"name"`        // 名称
	Description string          `json:"description"` // 描述
	Price       decimal.Decimal `json:"price"`       // 价格（使用decimal类型确保精确计算）
	Remark      string          `json:"remark"`      // 备注
}