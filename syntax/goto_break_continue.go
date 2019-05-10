package main

import "fmt"

func main()  {

    fmt.Println("---goto使用---");
    var i int
    for {
        println(i)
        i++
        if i>2 { goto OUYANG }
    }
    OUYANG:
        println("我跳到了OUYANG")
    //OUYANG2:
    //    println("我跳到了OUYANG2") //编译报错：label OUYANG2 defined and not used 提示定义而没有使用


    //break 可⽤于 for、 switch、 select，⽽ continue 仅能⽤于 for 循环。
    fmt.Println("---continue与break的使用---");
    YY1:
    for x := 0; x < 3; x++ {
    YY2:
        for y := 0; y < 5; y++ {
            if y > 2 { continue YY2 }
            if x > 1 { break YY1 }
            print(x, ":", y, " ")
        }
        println()
    }

    x := 100
    switch {
    case x >= 0:
        if x == 0 { break }
        println(x)
    }

}
