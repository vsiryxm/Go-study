package main

import "fmt"

func main() {

	var num1 byte = 255 //byte等价于uint8，也就是uint8的别名
	fmt.Printf("num1 = %d\n", num1)

	var sex rune = '男'
	fmt.Printf("字符 '%c' 的Unicode代码点是：%s\n", sex, "U+7537")

	// Unicode编码表中文查询：http://unicode-table.com/cn/
	var me rune = '\u6211' //我  支持\uFFFF、\U7FFFFFFF、\xFF格式
	fmt.Println(me)

}
