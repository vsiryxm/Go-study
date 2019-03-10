/* 用函数做为参数传递 */

package main

import (
	"fmt"
	"math"
)

type cb func(int64) int64 //1、先声明一个函数类型

func main() {

	fmt.Println("\n通过函数变量的形式来定义函数：")
	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Printf("求%v的平方根：%v\n", 9, getSqrt(9))

	fmt.Println("\n通过函数接口的形式来定义函数：")
	testCallBack(1, callBack) //2、定义触发函数
	testCallBack(2, func(x int64) int64 {
		fmt.Printf("我是回调2，x:%d\n", x)
		return x
	}) //2、或者定义一个函数闭包，将这个闭包作为参数传递

}

func testCallBack(x int64, f cb) {
	f(x)
}

//回调函数
func callBack(x int64) int64  {
	fmt.Printf("我是回调1，x:%d\n", x)
	return x
}
