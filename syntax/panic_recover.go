//错误处理
//没有结构化异常，使⽤用 panic 抛出错误， recover 捕获错误。

package main

import "fmt"

//import "fmt"

func main()  {

    test() //panic error！
    test2() //defer panic

    //捕获函数 recover 只有在延迟调⽤内直接调⽤才会终⽌止错误，否则总是返回 nil。任何未捕获的错误都会沿调⽤用堆栈向外传递。
    fmt.Println("---recover的正确位置---")
    test4()
    fmt.Println("---recover的错误位置---")
    test3()

}

func test()  {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err.(string)) //将interface{}转型为具体类型
        }
    }()
    panic("panic error！")
}

func test2()  {

    defer func() {
        fmt.Println(recover()) //捕获最后一次异常
    }()

    defer func() {
        panic("defer panic") //第二次抛出异常
    }()

    panic("test panic") //第一次抛出异常

}

func test3() {
    defer recover() //无⽆效！
    defer fmt.Println(recover()) //⽆无⽆效！
    defer func() {
        func() {
            println("defer inner")
            recover() //⽆无效！
        }()
    }()
    panic("test panic")
}

//正确的写法：
func except() {
    recover()
}
func test4()  {
    defer except()
    panic("test4 panic")
}