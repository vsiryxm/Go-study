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

    var (
        num8 int8
        num16 int16
        num32 int32
        num64 int64

        unum8 uint8 //别名类型byte
        unum16 uint16
        unum32 uint32
        unum64 uint64
    )

    //有符号整型可表示范围
    num8 = 127 //-128~127
    num16 = 32767 //-32768~32767
    num32 = 2147483647 //-2147483648~2147483647
    num64 = 9223372036854775807 //-9223372036854775808~9223372036854775807
    fmt.Printf("num8 最大值：%d，num16 最大值：%d，num32 最大值：%d，num64 最大值：%d\n", num8, num16, num32, num64);

    //无符号整型可表示范围
    unum8 = 255 //0~255
    unum16 = 65535 //0~65535
    unum32 = 4294967295 //0~4294967295
    unum64 = 18446744073709551615 //0~18446744073709551615
    fmt.Printf("unum8 最大值：%d，unum16 最大值：%d，unum32 最大值：%d，unum64 最大值：%d\n", unum8, unum16, unum32, unum64);


    // 声明一个整数类型变量并赋值
    var num1 int = -0x1000 //可以用8进制和16进制为所有整型变量赋值，但不能为浮点型变量赋值

    // 这里用到了字符串格式化函数。其中，%X用于以16进制显示整数类型值，%d用于以10进制显示整数类型值。
    fmt.Printf("16进制数 %X 表示的是 %d。\n", num1, ( num1 ))

}