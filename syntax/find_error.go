/* 一起来填坑 */
//11处坑

package mylib

import . "fmt"

var cc int8 = 8
dd := 9

func main()
{

    var a,b,c int8 = 1,2,3;
    d := a
    c = a++; b = --c

    var m float32 = 1.1
    n := (int)a //类型强转

    println("Hello World") //println和print都为内置函数
    Println(aaa())
    Printf(aaa())
    Println(cc, dd, m, n)

    var ptr *int8
    ptr = a
    Println(*ptr)

    var sex rune = "男"
    Println(sex)

}

func aaa() int8 {
    return 123
}