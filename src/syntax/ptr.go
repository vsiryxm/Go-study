/* 指针 */

package main

import (
    "fmt"
)

func main()  {

    fmt.Println("\n------简易指针------")
    var a int8 = 34
    var ptr *int8               //定义指针用*号
    ptr = &a                    //获取变量a的内存地址
    fmt.Println("a的值为：", a)     //34
    fmt.Println("*ptr为：", *ptr) //34
    fmt.Println("ptr为：", ptr)   //0xc0000541c8

    fmt.Println("\n------指针数组------")

    var arr = [3]int64 {1,2,3}
    var ptrarr [3]*int64
    for k,_ := range arr {
        ptrarr[k] = &arr[k]
    }
    fmt.Println("指针指向的内存地址有：", ptrarr)
    fmt.Println("遍历指针数组的值：")
    for k,_ := range ptrarr {
        fmt.Println(*ptrarr[k]);
    }

    fmt.Println("\n------指针向指针的指针------")
    var ptr2 **int8
    ptr2 = &ptr //ptr2为指针向指针的指针
    fmt.Println(&a, ptr, *ptr2) //三者完全等价
    fmt.Println(a, *ptr, **ptr2) //三者完全等价

    fmt.Println("\n------将指针作为函数参数------")
    var aa,bb int64 = 10,20
    fmt.Printf("交换前，aa=%d，bb=%d\n", aa, bb)
    swap(&aa, &bb)
    fmt.Printf("交换后，aa=%d，bb=%d\n", aa, bb)

    fmt.Println("\n------空指针------")
    var ptr3 *int8
    fmt.Printf("ptr3的值为 : %x\n", ptr3)
    if(ptr3 == nil) { //直接这么写0==nil是不成立的，即数字0和nil不相等
        fmt.Println("ptr3是一个空指针")
    }

}

//通过传递引用，交换两个变量的值
func swap(x *int64, y *int64) {
    //var tmp int64
    //tmp = *x
    //*x = *y
    //*y = tmp
    *x,*y = *y,*x //交换函数这样写更加简洁，也是 go 语言的特性，c++ 和 c# 是不能这么干的
}