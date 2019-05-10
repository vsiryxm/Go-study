package main //在main.go文件中调用CrosAccess函数，执行go build main.go func.go可以完成函数调用

import "fmt"

func CrosAccess()  {
    fmt.Println("哈哈哈哈，我定义在func.go文件里，能被同一个包、不同文件调用构建！")
}
