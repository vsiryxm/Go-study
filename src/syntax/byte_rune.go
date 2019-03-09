package main

import "fmt"

func main() {

	var num1 byte = 255 //byte等价于uint8
	fmt.Printf("num1 = %d\n", num1)

	var sex rune = '男'
	fmt.Printf("字符 '%c' 的Unicode代码点是：%s\n", sex, "U+7537")

}
