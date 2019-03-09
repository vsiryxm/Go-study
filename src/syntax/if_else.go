package main

import "fmt"

func main() {
	var number int = 5
	if number := 4; 10 > number {
		number -= 4
		number += 3
		fmt.Print(number)
	} else if 10 < number { //if后面如果只有一个表达式，加上括号是不报错的
		number -= 2
		fmt.Print(number)
	} else {
		fmt.Print(number)
	}
	number += 4
	fmt.Println(number)

}

//输出39
