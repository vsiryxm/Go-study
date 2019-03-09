package main

import "fmt"

func main() {

	//浮点数范围：
	// -100000000000000004337729697461918607329029332495193931179177378933611681288968111094132375552.000000 即-1.000000E+92
	// 100000000000000004337729697461918607329029332495193931179177378933611681288968111094132375552.000000  即1.000000E+92
	var num2 float32 = 5.89E-4
	// 可以在变量声明并赋值的语句中，省略变量的类型部分。
	// 不过别担心，Go语言可以推导出该变量的类型。
	var num3 = -99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.010000

	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num2, (num2))
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num3, (num3))

}
