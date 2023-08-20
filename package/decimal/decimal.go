package decimal

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Decimal() {

	fmt.Printf("decimal的默认精度为：%d位小数\n", decimal.DivisionPrecision)

	var price1 float64 = 3.1415926535897932
	var amount1 int = 100

	// var price2 float64 = 3.1415926535897932
	// var amount2 int = 100

	// 乘法
	subtotal, _ := decimal.NewFromFloat(price1).Mul(decimal.NewFromFloat(float64(amount1))).Float64()
	fmt.Println("subtotal =", subtotal) // 314.1592653589793

	// 除法
	var balance float64 = 100.00
	var price2 float64 = 1.1
	result, _ := decimal.NewFromFloat(balance).Div(decimal.NewFromFloat(price2)).Float64()
	fmt.Println("result =", result)

}
