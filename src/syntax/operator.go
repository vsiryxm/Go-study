/* 运算符操作 */

package main

import "fmt"

func main()  {

    base() //算术运算符
    guanxi() //关系运算符




}


func base()  {

    fmt.Println("\n------算术运算符------")

    var a,b,c int = 21,10,0

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

func guanxi()  {

    fmt.Println("\n------关系运算符------")

    var a,b int = 21,10

    if (a == b) {
        fmt.Println("第一行a等于b")
    } else {
        fmt.Println("第一行a不等于b")
    }

    if (a < b) {
        fmt.Println("第二行a小于b")
    } else {
        fmt.Println("第二行a不小于b")
    }

    if (a > b) {
        fmt.Println("第三行a大于b")
    } else {
        fmt.Println("第三行a不大于b")
    }

    if (b <= a) {
        fmt.Println("第四行b小于等于a")
    }

    if (a >= b) {
        fmt.Println("第五行a大于等于b")
    }

}