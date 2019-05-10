/* 变量的作用域 */
package main

import "fmt"

var config string = "我是全局变量"
var a int64 = 1
var b int64 = 2

func main()  {

    for a:=10; a<=15; a++ { //在for循环内的作用域
        fmt.Println(a) //局部变量，不影响全局a的值，输出：10到15
    }

    fmt.Println(a) //输出：1

    if config:="欧阳";config == "欧阳" { //在if条件内的作用域
        fmt.Println(config)
    }

    fmt.Println(config) //输出：我是全局变量
    config = "张三"
    fmt.Println(config) //输出：张三

    fmt.Println(add(a, &b))
    fmt.Println(a, b) //输出：1 3  全局a的值仍然没有改变，b的值是通过引用传递才做了修改

    for a=0; a<=5; a++ {
        fmt.Println(a) //这才改变了全局变量的值
    } //此时a=6，始终都会多+1，因为for循环里最后执行的是a++
    a++
    fmt.Println(a) //7

}

func add(a int64, b *int64) int64  { //在函数内的作用域
    a += 1
    *b += 1
    return a + *b
}