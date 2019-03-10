package main

import (
    "fmt"
)

func main() {

    fmt.Println("\n第一种使用方法：")
    //格式：for 初始化表达式; 条件是否成立; 自增项
    for i := 0; i < 10; i++ {
        fmt.Printf("第%d次循环\n", i)
    }

    fmt.Println("\n第二种使用方法：")
    //格式：for 条件是否成立
    var j int8 = 0
    for j < 10 { //类似while
        fmt.Printf("第%d次循环\n", j)
        j++
    }

    fmt.Println("\n第三种使用方法（用于遍历数组、切片、map）：")
    //格式：for 条件是否成立
    var arr = [10]int8{10,20,30,40,50,60,70,80,90,100}
    for k,v := range arr { //类似PHP中的foreach
        fmt.Printf("第%d个元素的值为：%d\n", k, v)
    }

    fmt.Println("\n遍历二维数组方法一：")
    var arr2 = [3][3]int64 {
        {1,2,3},
        {4,5,6},
        {7,8,9}, //如果最后一个元素不以逗号结尾，数组的右花括号不能另起一行
    }
    for i:=0; i<3; i++ {
        for j:=0; j<len(arr2[i]); j++ {
            fmt.Printf("arr2[%d][%d]=%d ", i, j, arr2[i][j])
        }
        fmt.Println()
    }

    fmt.Println("\n遍历二维数组方法二：")
    for k,v := range arr2  {
        for k1,v1 := range v {
            fmt.Printf("arr2[%d][%d]=%d ", k, k1, v1)
        }
        fmt.Println()
    }

    fmt.Println("\n循环嵌套输出九九乘法表：")
    for m:=1; m<=9; m++ {
        for n:=1; n<=m; n++ {
            fmt.Printf("%dx%d=%d ", n, m, m*n)
        }
        fmt.Println("")
    }

    fmt.Println("\n求1~100以内的素数：")
    var m,n int8
    for m=2; m<=100; m++ {

        for n=2; n<=(m/n); n++ {
            if(m%n == 0) {
                break //中断上一级for循环
            }
        }
        if (n > (m/n)) {
            fmt.Printf("%d是素数\n", m)
        }

    }

}
