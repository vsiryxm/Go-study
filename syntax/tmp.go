package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name string
	Sex  int32
}

func main() {

	var v *int32
	a := int32(123)
	v = &a
	saleNum := int32(100)
	shelveNum := int32(666)
	key := make([]string, 3)
	key[0] = strconv.Itoa(int(*v))
	key[1] = strconv.Itoa(int(saleNum))
	key[2] = strconv.Itoa(int(shelveNum))
	member := strings.Join(key, ":")
	fmt.Println(member)
	// student := new(Student)
	// fmt.Println(student.Name) // 输出为空字符串，不会报空指针异常

	// 空指针异常 即指针没有指向任何内存地址，当访问这个指针时，就会报空指针
	// var ptr *int
	// fmt.Println(*ptr) // panic: runtime error: invalid memory address or nil pointer dereference

	//注意， range 会复制对象。
	//a := [3]int{0, 1, 2}
	//for i, v := range a { // index、 value 都是从复制品中取出。
	//    if i == 0 { // 在修改前，我们先修改原数组。
	//        a[1], a[2] = 999, 999
	//        fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
	//    }
	//    a[i] = v + 100 // 使⽤复制品中取出的 value 修改原数组。
	//}
	//fmt.Println(a)

	//建议改⽤引⽤类型，其底层数据不会被复制。
	//s := []int{1, 2, 3, 4, 5}
	//for i, v := range s { // 复制 struct slice { pointer, len, cap }。
	//    if i == 0 {
	//        s = s[:3] // 对 slice 的修改，不会影响 range。
	//        s[2] = 100 // 对底层数据的修改。
	//    }
	//    println(i, v)
	//}
	//输出：
	//0 1
	//1 2
	//2 100
	//3 4
	//4 5
	//另外两种引⽤类型 map、 channel 是指针包装，⽽不像 slice 是 struct。

	//data := [...]int{0, 1, 2, 3, 4, 5, 6}
	//slice := data[1:4:5]
	//fmt.Println(len(slice), cap(slice))

	// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 := s[2:5]    // [2 3 4]
	// s2 := s1[2:6:7] // [4 5 6 7 0]  5
	// s3 := s2[3:6]
	// fmt.Println(s, s1, cap(s1), s2, cap(s2), s3)

}

//func add(x,y int64) (int64) {
//    z = x + y
//    return z
//}
