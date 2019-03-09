/* 运算符操作 */

package main

import (
	"fmt"
)

func main() {
	var a int = 650
	b := string(a)
	fmt.Println(b)

	base()   //算术运算符
	guanxi() //关系运算符
	luoji()  //逻辑运算符
	wei()    //位运算符
	fuzhi()  //赋值运算符
	other()  //其他运算符：指针* 和 &地址

}

//算术运算符
func base() {

	fmt.Println("\n------算术运算符------")

	var a, b, c int = 21, 10, 0

	c = a + b
	fmt.Printf("第一行c=%d\n", c)

	c = a - b
	fmt.Printf("第二行c=%d\n", c)

	c = a * b
	fmt.Printf("第三行c=%d\n", c)

	c = a / b
	fmt.Printf("第四行c=%d\n", c)

	c = a % b //求余数
	fmt.Printf("第五行c=%d\n", c)

	a++
	fmt.Printf("第六行a=%d\n", a)

	a = 21
	a--
	fmt.Printf("第七行a=%d\n", a)

}

//关系运算符
func guanxi() {

	fmt.Println("\n------关系运算符------")

	var a, b int = 21, 10

	if a == b {
		fmt.Println("第一行a等于b")
	} else {
		fmt.Println("第一行a不等于b")
	}

	if a < b {
		fmt.Println("第二行a小于b")
	} else {
		fmt.Println("第二行a不小于b")
	}

	if a > b {
		fmt.Println("第三行a大于b")
	} else {
		fmt.Println("第三行a不大于b")
	}

	if b <= a {
		fmt.Println("第四行b小于等于a")
	}

	if a >= b {
		fmt.Println("第五行a大于等于b")
	}

}

//逻辑运算符
func luoji() {
	fmt.Println("\n------关系运算符------")

	var a, b bool = true, false

	if a && b {
		fmt.Println("第一行条件为真")
	} else {
		fmt.Println("第一行条件为假")
	}

	if a || b {
		fmt.Println("第二行条件为真")
	} else {
		fmt.Println("第二行条件为假")
	}

	if !(a && b) { //逻辑非
		fmt.Println("第三行条件为真")
	} else {
		fmt.Println("第三行条件为假")
	}

}

//位运算符
//支持 & | ^ << >>
func wei() {
	fmt.Println("\n------位运算符------")
	var a, b int16 = 34, 60
	//a = 00100010
	//b = 00111100

	fmt.Println("a & b = ", a&b) //00100000 => 32
	fmt.Println("a | b = ", a|b) //00111110 => 62
	fmt.Println("a ^ b = ", a^b) //00011110 => 30 异或：两个二进制位不同时为1，其它都为0

	fmt.Println("a = ", a)
	fmt.Println("a << 3 = ", a<<3) //34*2的3次方 => 272 如果a换成int8型，会溢出但不报错误，结果为16，说明做位移运算时要注意做溢出检查
	fmt.Println("a >> 3 = ", a>>3) //34/2的3次方 => 4

}

//赋值运算符
func fuzhi() {
	fmt.Println("\n------赋值运算符------")
	var a, b, c, d int64 = 2, 7, 0, 0

	c += a
	fmt.Println("c += a ： ", c) //2
	d -= a
	fmt.Println("d -= a ： ", d) //-2
	c *= a
	fmt.Println("c *= a ： ", c) //4
	c /= 2
	fmt.Println("c /= 2 ： ", c) //2
	c %= b
	fmt.Println("c %= b ： ", c) //2
	a <<= 3
	fmt.Println("a <<= 3 ：", a) //16
	b >>= 3
	fmt.Println("b >>= 3 ： ", b) //0
	a &= 3
	fmt.Println("a &= 3 ： ", a) //0
	b |= 3
	fmt.Println("b |= 3 ： ", b) //3
	a ^= 3
	fmt.Println("a ^= 3 ： ", a) //3 异或：两个二进制位不同时为1，其它都为0

}

//其他运算符：指针* 和 &地址
func other() {
	fmt.Println("\n------其他运算符：指针* 和 &地址------")
	var a int8 = 34
	var ptr *int8               //定义指针用*号
	ptr = &a                    //获取变量a的内存地址
	fmt.Println("a的值为：", a)     //34
	fmt.Println("*ptr为：", *ptr) //34
	fmt.Println("ptr为：", ptr)   //0xc0000541c8

}
