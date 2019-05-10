package main

import (
    "fmt"
    //"math"
    "runtime"
    "sync"
)

func main()  {
    runtime.GOMAXPROCS(2)
    var a float32 = 1.2345678998888888888
    fmt.Println(a)
    //go func() {
    //    fmt.Println("Hello world！")
    //}()

    wg := new(sync.WaitGroup)
    wg.Add(2)
    
}
