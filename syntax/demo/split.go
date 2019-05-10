package main

import (
    "fmt"
    //"os"
    //"strings"
)

func main()  {
    //fmt.Println(len(os.Args), os.Args[1]);
    //ip := os.Args[1]; //接收cli参数，命令名称默认为第一个参数
    //tmp := strings.Split(ip, "."); //相当于php中的split或explode函数
    //fmt.Println(len(tmp), tmp[1]);
    k := 123
    a := k+100
    c := k*a
    x := getMyname();
    fmt.Println(x, k, c, a);
}

func getMyname() string {
    y := "yang"
    return y;
}
