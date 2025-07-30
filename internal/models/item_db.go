// models包定义了数据模型
package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/shopspring/decimal"
)

// Decimal 自定义decimal类型，用于数据库交互
type Decimal struct {
	decimal.Decimal
}

// Scan 实现sql.Scanner接口，用于从数据库读取decimal值
func (d *Decimal) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	
	// 将数据库中的值转换为decimal.Decimal
	switch v := value.(type) {
	case float64:
		d.Decimal = decimal.NewFromFloat(v)
	case int64:
		d.Decimal = decimal.NewFromInt(v)
	case string:
		dec, err := decimal.NewFromString(v)
		if err != nil {
			return err
		}
		d.Decimal = dec
	case []byte:
		dec, err := decimal.NewFromString(string(v))
		if err != nil {
			return err
		}
		d.Decimal = dec
	default:
		return fmt.Errorf("无法将 %T 类型转换为 decimal.Decimal", value)
	}
	
	return nil
}

// Value 实现driver.Valuer接口，用于将decimal值写入数据库
func (d Decimal) Value() (driver.Value, error) {
	if d.Decimal.IsZero() {
		return "0", nil
	}
	return d.Decimal.String(), nil
}

// DBItem 数据库中的物品实体
// 该结构体用于与数据库交互，将Decimal映射到MySQL的decimal字段
type DBItem struct {
	ID          string  `json:"id" gorm:"column:id;type:varchar(36);primaryKey"`               // 唯一标识符
	Name        string  `json:"name" gorm:"column:name;type:varchar(100);not null"`            // 名称
	Description string  `json:"description" gorm:"column:description;type:text"`               // 描述
	Price       Decimal `json:"price" gorm:"column:price;type:decimal(10,2);not null"`         // 价格，映射到MySQL的decimal(10,2)类型
	Remark      string  `json:"remark" gorm:"column:remark;type:varchar(255)"`                 // 备注
}

// TableName 设置数据库表名
func (DBItem) TableName() string {
	return "items"
}