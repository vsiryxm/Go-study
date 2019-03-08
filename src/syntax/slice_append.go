/* 深入理解切片的append函数 */

package main

import "fmt"

func main(){
    s := []int{5}
    fmt.Println("s =", s, "len(s) =", len(s), "cap(s) =", cap(s), "ptr(s) =", &s[0])

    s = append(s,7)
    fmt.Println("s =", s, "len(s) =", len(s), "cap(s) =", cap(s), "ptr(s) =", &s[0])

    s = append(s,9) //每次扩容，内存都重新分配，旧的内存收回
    fmt.Println("s =", s, "len(s) =", len(s), "cap(s) =", cap(s), "ptr(s) =", &s[0])

    x := append(s, 11) //x的内存地址与s的内存地址是一样的
    fmt.Println("x =", x, "len(x) =", len(x), "cap(x) =", cap(x), "ptr(s) =", &s[0], "ptr(x) =", &x[0])

    y := append(s, 12) //y的内存地址与s的内存地址是一样的
    fmt.Println("y =", y, "len(y) =", len(y), "cap(y) =", cap(y), "ptr(s) =", &s[0], "ptr(y) =", &y[0])



}

//输出结果：
//s = [5] len(s) = 1 cap(s) = 1 ptr(s) = 0xc000054080
//s = [5 7] len(s) = 2 cap(s) = 2 ptr(s) = 0xc0000540d0
//s = [5 7 9] len(s) = 3 cap(s) = 4 ptr(s) = 0xc0000520e0
//x = [5 7 9 11] len(x) = 4 cap(x) = 4 ptr(s) = 0xc0000520e0 ptr(x) = 0xc0000520e0
//y = [5 7 9 12] len(y) = 4 cap(y) = 4 ptr(s) = 0xc0000520e0 ptr(y) = 0xc0000520e0
//最后这一行的输出结果有点难以理解
