/* 变量与常量定义 */
//通过 const 关键字来进行常量的定义。
//通过在函数体外部使用 var 关键字来进行全局变量的声明和赋值。
//通过 type 关键字来进行结构(struct)和接口(interface)的声明。
//通过 func 关键字来进行函数的声明。
//Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用：
//函数名首字母小写即为 private，函数名首字母大写即为 public

package main;

import (
    "fmt"
)

var quanju int = 123 //变量还可以定义在函数体外面，做为全局变量
// quanju2 := 456 //短语句定义不支持在函数体外面


func main()  {
    //定义变量，用关键字var，变量名称为字母数字下划线，但不能以数字开头
    //格式：普通赋值，由关键字var、变量名称、变量类型、=、值
    //若只声明不赋值，则去除=和后面的值
    fmt.Println("变量定义的第一种书写格式：")
    var (
        num1 int
        num2 int
        num3 bool
    )
    num1, num2, num3 = 1, 2, true
    fmt.Println(num1, num2, num3) //输出函数Println

    fmt.Println("变量定义的第二种书写格式：")
    var yyy,xxx,mmm int
    yyy = 123
    xxx = 456
    mmm = 789
    fmt.Println(yyy, xxx, mmm)

    fmt.Println("变量定义的第三种书写格式：")
    var (
        yyy1 int = 123
        xxx1 int = 456
    )
    fmt.Printf("yyy1 = %d, xxx1 = %d\n", yyy1, xxx1) //%d为数字占位符

    yang := "阳阳阳" //短语句定义（这个变量在之前是未定义过的），编译时会自动推导出变量的类型，且这种定义方法仅限在函数体内部声明
    fmt.Printf("短语句定义，yang = %s，编译时会自动推导出变量类型\n", yang)

    fmt.Println("常量定义第一种格式：")
    const myname string = "海阳之新" //string类型可写可不写
    fmt.Println("我是常量myname=",myname)

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

    //查看变量所在的内存地址
    fmt.Println(&num1)

    fmt.Printf("我是全局变量quanju=%d，我可以定义在函数体外面", quanju)


}
