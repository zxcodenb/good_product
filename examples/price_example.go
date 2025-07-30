// examples包提供使用示例
package main

import (
	"fmt"
	"item-value/internal/models"
	"item-value/utils"

	"github.com/shopspring/decimal"
)

func main() {
	// 创建一个物品示例
	item := models.Item{
		ID:          "1",
		Name:        "示例商品",
		Description: "这是一个示例商品",
		Price:       decimal.NewFromFloat(99.99), // 使用decimal存储价格
		Remark:      "示例备注",
	}

	fmt.Printf("商品信息:\n")
	fmt.Printf("ID: %s\n", item.ID)
	fmt.Printf("名称: %s\n", item.Name)
	fmt.Printf("描述: %s\n", item.Description)
	fmt.Printf("价格: %s元\n", utils.FormatPrice(item.Price))
	fmt.Printf("备注: %s\n", item.Remark)

	// 演示价格计算
	price1 := decimal.NewFromFloat(100.50)
	price2 := decimal.NewFromFloat(50.25)

	fmt.Printf("\n价格计算示例:\n")
	fmt.Printf("价格1: %s元\n", utils.FormatPrice(price1))
	fmt.Printf("价格2: %s元\n", utils.FormatPrice(price2))
	fmt.Printf("价格相加: %s元\n", utils.FormatPrice(utils.AddPrices(price1, price2)))
	fmt.Printf("价格相减: %s元\n", utils.FormatPrice(utils.SubtractPrices(price1, price2)))
	fmt.Printf("价格相乘: %s元\n", utils.FormatPrice(utils.MultiplyPrice(price1, decimal.NewFromFloat(2))))
	fmt.Printf("价格相除: %s元\n", utils.FormatPrice(utils.DividePrice(price1, decimal.NewFromFloat(2))))

	// 演示元和分的转换
	fmt.Printf("\n货币单位转换示例:\n")
	yuan := decimal.NewFromFloat(99.99)
	cent := utils.YuanToCent(yuan)
	fmt.Printf("%s元 = %d分\n", utils.FormatPrice(yuan), cent)
	
	convertedYuan := utils.CentToYuan(cent)
	fmt.Printf("%d分 = %s元\n", cent, utils.FormatPrice(convertedYuan))
}