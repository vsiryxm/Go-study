
package main

import (
    "fmt"
)

func main()  {
    arr := [10]int32{'a','b','c','d','e','f','g','h','i','j'}
    s1 := arr[1:5]
    fmt.Println(s1) //b c d e  len=4 cap=9
    s2 := s1[2:4]
    fmt.Println(s2) //d e len=2 cap=7
    s1[2] = 'z'
    fmt.Println(s2[0]) //z
    fmt.Printf("s1的内存地址为：%p，s2的内存地址为：%p", s1[0], s2[0]); //同一个内存地址
    s2 = append(s2, 'y','x','m'); //等于或超过容量，重新开辟内存容量
    fmt.Printf("s1的内存地址为：%p，s2的内存地址为：%p", s1[0], s2[0]); //不同内存地址

}

