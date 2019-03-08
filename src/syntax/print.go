/* 格式化输出函数 */

package main

import "fmt"

func main()  {
    var arr = [3]int {1, 2, 3}
    fmt.Print("使用Print输出：字符串 ", 123, true, arr, ' ',aaa(), "\n") //当打印函数时，后面必须带括号，不然就是打印函数所在的内存地址
    fmt.Println("使用Println输出：字符串 ", 123, true, arr, ' ',aaa(), "\n")
    fmt.Printf("使用Printf输出：%s，整型：%d，逻辑型：%v，数组：%v，空串：%s，函数返回：%d\n", "字符串", 123, true, arr, " ",aaa())
    //fmt.Printf(aaa()); //会报错，不能这样直接输出，必须使用占位符配合输出

}

func aaa() (int32) {
    return 456;
}
