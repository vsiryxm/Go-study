/* 一起来填坑 */

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

    println("Hello World")
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