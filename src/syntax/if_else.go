/* 条件分支 */
//（1） 不需使用括号将条件包含起来
//（2） 大括号{}必须存在，即使只有一行语句
//（3） 左括号必须在if或else的同一行
//（4） 在if之后，条件语句之前，可以添加变量初始化语句，使用；进行分隔
//（5） 在有返回值的函数中，最终的return不能在条件语句中

package main

import "fmt"

func main() {
	var number bool = true
	if number := 4; 10 > number { //这里定义的number仅在if域里有用，也修改不了上面一行定义的number
		number += 3
		fmt.Println(number) //7
	} else if 10 < number { //if后面如果只有一个表达式，加上括号是不报错的
		number -= 2
		fmt.Println(number)
	} else {
		fmt.Println(number)
	}
	fmt.Println(number) //true


}

