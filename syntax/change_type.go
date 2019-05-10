package main

import (
    "fmt"
    "unsafe"
)

func main()  {

    //1、试图将字符串转成数字
    //a := 1 + "2"
    //fmt.Println(a)
    //编译报错：cannot convert "2" (type untyped string) to type int
    //结论：字符串不能转换成数字

    //2、试图将数字转成字符串
    fmt.Println(string(65)) //A
    s := 65
    fmt.Println(string(s)) //A
    chgString(string(s)) //A
    //结论：将数字常量、整型变量转换成字符串是可以的

    //3、试图将整型转换成浮点型
    var sum int64 = 10
    var count int64 = 3
    var f float64
    var i int64
    i = sum /count //直接取整，不四舍五入
    fmt.Println(i)
    f = float64(sum) / float64(count) //保留16位小数
    fmt.Println(f)
    //整型转换成float型可行，会保留小数


    //4、底层类型的定义不能是递归的，参考：https://studygolang.com/articles/14018
    type A string //相当于底层用了string，可以理解成从string分叉出一个新类型，可以说string是A的底层类型
    type B A
    var b B = "haha"
    var c B = "123"
    var a A = "aaa"

    //fmt.Println(a+c) //类型不匹配，编译报错，mismatched types A and B，因为a的底层类型是A，c的底层类型是B
    fmt.Println(B(a)+c) //将a的类型转换成B类型即可通过编译
    fmt.Println(b+c) //haha123

    //type S string
    //type T map[S]float64
    //var y map[string]float64 = make(T) //定义哈希数组
    //fmt.Println(y)
    //编译报错：cannot use make(T) (type T) as type map[string]float64 in assignment
    //原因：因为T的底层类型是map[S]float64，而不是map[string]float64

    //5、不支持隐式类型转换
    //var b byte = 100
    //var n int = b //不支持隐式类型转换
    //var _ int = int(b)

    //6、不支持非bool类型作bool使用
    //var a2 := 100
    //if a2 {
    //    fmt.Println(a)
    //}

    //7、可以在 unsafe.Pointer 和任意类型指针间进⾏转换。
    x := 0x12345678
    p := unsafe.Pointer(&x) // *int -> Pointer
    n := (*[4]byte)(p) // Pointer -> *[4]byte  //没搞明白*[4]byte，不知道是不是数组指针?
    for i := 0; i < len(n); i++ {
        fmt.Printf("%X ", n[i])
    } //输出：78 56 34 12



}

func chgString(s string)  {
    fmt.Println(s)
}