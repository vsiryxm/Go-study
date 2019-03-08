/* 变量与常量定义 */
//通过 const 关键字来进行常量的定义。
//通过在函数体外部使用 var 关键字来进行全局变量的声明和赋值。
//通过 type 关键字来进行结构(struct)和接口(interface)的声明。
//通过 func 关键字来进行函数的声明。

package main;

import (
    "fmt"
    "unsafe"
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
    //常量中的数据类型只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。
    const myname string = "海阳之新" //string类型可写可不写，写了是显式定义，不写是隐式定义
    fmt.Println("我是字符串常量myname=", myname)

    const PI float32 = 3.1415926
    fmt.Println("我是浮点型常量PI=", PI)

    const is_publish = true //省略类型bool也可以
    fmt.Println("我是逻辑型常量is_publish=", is_publish)


    fmt.Println("\n常量定义第二种格式：")
    const a,b,c,d = 1,3.1415926,"yxm",true
    fmt.Println(a,b,c,d)

    // 使用 const 块来实现枚举
    fmt.Println("\n常量定义第三种格式：")
    const (
        unknown = 0
        female = 1
        male = 2
    ) //如性别，用1代表男，2代表女，0代表未知
    fmt.Println(unknown, female, male)

    fmt.Println("\n常量的值可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值：")
    const (
        aa = "abc"
        bb = len(aa)
        cc = unsafe.Sizeof(aa) //需要导入unsafe包，计算变量常量在内存占用的字节
    )
    fmt.Println(aa, bb, cc)
    //cc=16 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。

    //iota是Go的一个特殊常量，可以认为是一个可以被编译器修改的常量。
    //位移 <<(乘以2的n次方)、>>(除以2的n次方)
    const (
        aaa = 1<<iota //iota从0开始，编译器会将其值依次自增
        bbb = 3<<iota //iota=1
        ccc //下面没有赋值，被认为沿用上一个表达式，即：3<<iota，但此时iota=2
        ddd
    )
    fmt.Println(aaa, bbb, ccc, ddd) //1 6 12 24

    const (
        aaa1 = 1
        bbb1
        ccc1
    )
    fmt.Println(aaa1, bbb1, ccc1) //1 1 1

    //查看变量所在的内存地址
    //fmt.Println(&num1)

    fmt.Printf("我是全局变量quanju=%d，我可以定义在函数体外面", quanju)


}
