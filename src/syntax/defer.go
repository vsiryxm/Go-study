//defer ⽤于注册延迟调⽤用，这些调用⽤直到 return 前才被执⾏，通常⽤于释放资源或错误处理。
//滥⽤用defer可能会导致性能问题，尤其是在⼀个 "⼤循环" ⾥。

package main

import "fmt"

func main()  {

    //多个 defer 注册，按 FILO（First In Last Out 先进后出）次序执行⾏。哪怕函数或某个延迟调⽤发⽣错误，这些调⽤用依旧会被执⾏。
    test(10) //如果是传非0值，如10，输出c 10 b a
    // 输出：
    // c
    // b
    // a
    // panic: runtime error: integer divide by zero

    //延迟调⽤参数在注册时求值或复制，可⽤指针或闭包 "延迟" 读取。
    test2()
    //x= 20 y= 120
    //defer: 10 120

}

func test(x int)  {
    defer println("a")
    defer println("b")
    defer func() {
        println(100/x)
    }()
    defer  fmt.Println("c")
}

func test2()  {
    x,y := 10,20
    defer func(i int) {
        println("defer:", i, y) //y被闭包引用
    }(x) //x在运行下面的程序前就被复制，即在注册时求值或复制

    x += 10
    y += 100
    println("x=", x, "y=", y)
}