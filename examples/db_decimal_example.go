// examples包提供数据库decimal类型使用示例
package main

import (
	"fmt"
	"item-value/internal/models"
	"item-value/utils"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// createItem 创建一个物品并保存到数据库
func createItem(db *gorm.DB, name string, price float64) (*models.DBItem, error) {
	item := &models.DBItem{
		Name:        name,
		Description: fmt.Sprintf("%s的描述信息", name),
		Price:       models.Decimal{Decimal: decimal.NewFromFloat(price)},
		Remark:      "示例备注",
	}

	result := db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

// getItemByID 根据ID获取物品
func getItemByID(db *gorm.DB, id string) (*models.DBItem, error) {
	var item models.DBItem
	result := db.First(&item, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

// updateItemPrice 更新物品价格
func updateItemPrice(db *gorm.DB, id string, newPrice decimal.Decimal) error {
	result := db.Model(&models.DBItem{}).Where("id = ?", id).Update("price", models.Decimal{Decimal: newPrice})
	return result.Error
}

// listItems 获取所有物品
func listItems(db *gorm.DB) ([]models.DBItem, error) {
	var items []models.DBItem
	result := db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// demonstrateDecimalOperations 演示decimal运算后存储到数据库
func demonstrateDecimalOperations(db *gorm.DB) {
	fmt.Println("=== Decimal运算和数据库存储演示 ===")
	
	// 创建物品
	item1, err := createItem(db, "商品1", 99.99)
	if err != nil {
		fmt.Printf("创建商品失败: %v\n", err)
		return
	}
	fmt.Printf("创建商品: %s, 价格: %s\n", item1.Name, item1.Price.StringFixed(2))

	// 查询物品
	foundItem, err := getItemByID(db, item1.ID)
	if err != nil {
		fmt.Printf("查询商品失败: %v\n", err)
		return
	}
	fmt.Printf("查询到商品: %s, 价格: %s\n", foundItem.Name, foundItem.Price.StringFixed(2))

	// 演示decimal运算并更新价格
	newPrice := foundItem.Price.Add(decimal.NewFromFloat(10.50))
	err = updateItemPrice(db, foundItem.ID, newPrice)
	if err != nil {
		fmt.Printf("更新价格失败: %v\n", err)
		return
	}
	fmt.Printf("价格从 %s 更新为 %s\n", foundItem.Price.StringFixed(2), newPrice.StringFixed(2))

	// 查询更新后的物品
	updatedItem, err := getItemByID(db, foundItem.ID)
	if err != nil {
		fmt.Printf("查询更新后的商品失败: %v\n", err)
		return
	}
	fmt.Printf("更新后商品价格: %s\n", updatedItem.Price.StringFixed(2))

	// 列出所有物品
	items, err := listItems(db)
	if err != nil {
		fmt.Printf("列出商品失败: %v\n", err)
		return
	}
	fmt.Println("\n所有商品:")
	for _, item := range items {
		fmt.Printf("- %s: %s元\n", item.Name, item.Price.StringFixed(2))
	}
}

func ExampleMain() {
	// 初始化数据库连接
	if err := utils.InitDB(); err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		return
	}
	
	// 获取数据库连接实例
	db := utils.GetDB()
	
	// 自动迁移数据库表结构
	err := db.AutoMigrate(&models.DBItem{})
	if err != nil {
		fmt.Printf("自动迁移失败: %v\n", err)
		return
	}
	
	// 演示decimal操作和数据库存储
	demonstrateDecimalOperations(db)
	
	// Output:
	// === Decimal运算和数据库存储演示 ===
	// 创建商品: 商品1, 价格: 99.99
	// 查询到商品: 商品1, 价格: 99.99
	// 价格从 99.99 更新为 110.49
	// 更新后商品价格: 110.49
	// 
	// 所有商品:
	// - 商品1: 110.49元
}