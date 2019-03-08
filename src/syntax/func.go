package main

import (
    "fmt"
    "mylib/calc"  //包所在的文件夹路径，先去找GOROOT/src/mylib/calc文件夹，再去找GOPATH/src/mylib/calc文件夹
)

func main() {

    fmt.Println("在同一个包下，调用函数，返回多个值且怎样去接收值：")
    a, b, _, strs := multiReturn(1) //注意：获取函数返回的多个值时，如果其中一个值不想接收，可以用下划线表示

    fmt.Println(a, b, strs)

    fmt.Println("在不同的包下，调用函数：")
    c := calc.Add(a, b) //调用格式：包名.方法名（参数1，参数2...），函数首字母必须大写才能被调用（相当于public）
    fmt.Printf("a + b = %d\n", c)

}

func multiReturn(a int64) (int64, int64, int64, string) {
    b, c, d := 2, 3, "this is string"
    return a, int64(b), int64(c), d
}

