/* 变量与常量定义 */

package main;

import (
    "fmt"
)

func main()  {
    //定义变量，用关键字var
    //格式：普通赋值，由关键字var、变量名称、变量类型、=、值
    //若只声明不赋值，则去除=和后面的值
    fmt.Println("变量定义的第一种书写格式：")
    var (
        num1 int
        num2 int
        num3 int
    )
    num1, num2, num3 = 1, 2, 3
    fmt.Println(num1, num2, num3) //输出函数Println

    fmt.Println("变量定义的第二种书写格式：")
    var yyy,xxx int
    yyy = 123
    xxx = 456
    fmt.Println(yyy, xxx)

    fmt.Println("变量定义的第三种书写格式：")
    var (
        yyy1 int = 123
        xxx1 int = 456
    )
    fmt.Printf("yyy1 = %d, xxx1 = %d\n", yyy1, xxx1) //%d为数字占位符

    yang := "阳阳阳"
    fmt.Printf("推导定义yang = %s\n", yang)

    fmt.Println("常量定义第一种格式：")
    const myname string = "海阳之新" //string类型可写可不写
    fmt.Println(myname)

    fmt.Println("常量定义第二种格式：")
    const a,b = 1,2
    fmt.Println(a,b)

    // 使用 const 块来实现枚举
    fmt.Println("常量定义第三种格式：")
    const (
        c    = 0
        golang  = 1
        java = 2
    )
    fmt.Println(c, golang, java)
}
