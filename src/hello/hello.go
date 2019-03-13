package main // 代码包声明语句

// 代码包导入语句
import (
    "fmt" //导入代码包fmt
)

// main函数
func main()  { // 代码块由“{”和“}”包裹
    fmt.Printf("Hello World!\nHello Simon!\n")
    //
    //d,i := [3]int8{1,2,3},0
    //i,d[i] = 1,100
    //println(i,d[0])

    const yyy  = iota
    const bbb  = iota
    fmt.Println(yyy,bbb)

    const (
        i=iota
        j1,j2,j3 = iota,iota,iota
        k = iota
    )
    fmt.Println(i,j1,j2,j3,k)

}


