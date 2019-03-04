package main

import "fmt"

func main()  {
    //定义变量，用关键字var
    //格式：普通赋值，由关键字var、变量名称、变量类型、=、值
    //若只声明不赋值，则去除=和后面的值
    var num uint64 = 65535

    // 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
    size := (8)

    fmt.Printf("类型为uint64的整数 %d 需要占用的存储空间为：%d 个字节。\n", num, size)
}