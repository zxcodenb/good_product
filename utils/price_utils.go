package utils

import (
	"github.com/shopspring/decimal"
)

// PriceUtils 价格处理工具类
type PriceUtils struct{}

// NewFromFloat 从float64创建decimal.Decimal值
func NewFromFloat(value float64) decimal.Decimal {
	return decimal.NewFromFloat(value)
}

// NewFromInt 从int64创建decimal.Decimal值
func NewFromInt(value int64) decimal.Decimal {
	return decimal.NewFromInt(value)
}

// YuanToCent 将元转换为分（以整数形式存储价格）
// 输入为元，输出为对应的分数值
func YuanToCent(yuan decimal.Decimal) int64 {
	cent := yuan.Mul(decimal.NewFromInt(100))
	return cent.IntPart()
}

// CentToYuan 将分转换为元
func CentToYuan(cent int64) decimal.Decimal {
	return decimal.NewFromInt(cent).Div(decimal.NewFromInt(100))
}

// FormatPrice 格式化价格显示，保留两位小数
func FormatPrice(price decimal.Decimal) string {
	return price.StringFixed(2)
}

// AddPrices 精确加法运算
func AddPrices(price1, price2 decimal.Decimal) decimal.Decimal {
	return price1.Add(price2)
}

// MultiplyPrice 价格乘法运算
func MultiplyPrice(price decimal.Decimal, multiplier decimal.Decimal) decimal.Decimal {
	return price.Mul(multiplier)
}

// SubtractPrices 精确减法运算
func SubtractPrices(price1, price2 decimal.Decimal) decimal.Decimal {
	return price1.Sub(price2)
}

// DividePrice 价格除法运算
func DividePrice(price, divisor decimal.Decimal) decimal.Decimal {
	return price.Div(divisor)
}
