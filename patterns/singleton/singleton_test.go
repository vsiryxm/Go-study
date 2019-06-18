package singleton

import (
    "fmt"
    "testing"
)
func Test1(test *testing.T){
    s := New()
    s["this"] = "that"
    s2 := New() //再次创建时，会先检查内存里是否已经存在这个实例了
    fmt.Println("This is ", s2["this"]) //输出that，说明并没有重新分配内存，即没有重新创建实例
}