/* 切片 */
//切片不是动态数组，切片是提供了一个访问数组内存片段的合法手段，利用切片功能，我们可以很自由地访问数组的任何一个片段
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//第一种定义方法：
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //完整的写法：var s []int = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}，但go讲究简洁
	printSlice("s", s)

	//第二种定义方法：
	var s2 = make([]int, 3, 5) //make(类型，长度，容量)，如果省略容量参数，则容量与长度一致
	printSlice("s2", s2)

	//var s3 []int = {1,2,3} //错误的写法
	var s3 []int //只定义不赋值，则len为0，cap为0，也叫空切片
	printSlice("s3", s3)
	if s3 == nil { //if后面如果只有一个表达式，加上括号是不报错的
		fmt.Println("切片s3是空的")
	}

	fmt.Print(s) //说明：Print、Printf、Println三个函数都支持直接输出切片

	//切片截取操作
	fmt.Println("\n---切片截取操作---")

	fmt.Println("s[1:4] ==", s[1:4]) //从下标（或称索引）为1的元素开始截取，截取到4-1个元素
	fmt.Println("s[:3] ==", s[:3])   //省略上限，则从0开始
	fmt.Println("s[4:] ==", s[4:])   //省略下限，则一直到切片末尾元素

	s4 := make([]int, 0, 5)
	//s4[2] = 234
	printSlice("s4", s4)

	s5 := s[:2]                        //切出一片来给一个新数组，实际是一个引用
	fmt.Println("s[0]的内存地址：", &s[0])   //输出s[0]的内存地址
	fmt.Println("s5[0]的内存地址：", &s5[0]) //输出s5[0]的内存地址，与s[0]是一样的
	printSlice("s5", s5)
	s[1] = 111 //改变s[1]的值后，s5的值也发生了变化，说明s5的值是s的一个引用地址，也可以说是一个数组指针，但实际上还是一个切片
	printSlice("s5", s5)
	fmt.Printf("s5是一个 %v 切片类型", reflect.TypeOf(s5)) //返回[]int

	//切片扩容操作
	//容量是根据切片上一次的长度的倍数来扩展的
	fmt.Println("\n---切片扩容操作---")
	s = append(s, 10)                //上一次s的长度为9，容量为9，现在加入1个元素后，容量不够用了，扩展到len*2，长度+1
	fmt.Println("s扩容前的内存地址：", &s[0]) //输出s的内存地址
	printSlice("s", s)
	s = append(s, 11, 12, 13, 14, 15, 16, 17, 18, 19) //上一次s的长度为10，容量为18，现在加入9个元素后，容量不够用了，扩展到上一次容量*2，长度+1
	fmt.Println("s扩容后的内存地址：", &s[0])
	//fmt.Printf("s扩容后的内存地址：%p\n", s) //在扩容后，再次输出s的内存地址，两次不一样，说明开辟了新的内存空间，旧的内存废弃回收
	printSlice("s", s)

	//拼接两个切片
	yy := []int{1, 2, 3}
	xx := []int{4, 5, 6}
	mm := append(yy, xx...)
	fmt.Printf("yy的内存地址：%p\n", yy)
	fmt.Printf("mm的内存地址：%p\n", mm)
	printSlice("mm", mm)

	fmt.Println("\n---切片复制操作---")
	s_cp := make([]int, len(s), cap(s))
	copy(s_cp, s)
	printSlice("s_cp", s_cp)
	s[1] = 111111 //再次改变s[1]的值，发现s_cp[1]不再跟着变化，说明在内存中复制并开辟了一块新的存储空间
	printSlice("s_cp", s_cp)

	s_cp2 := make([]int, len(s)*2, cap(s)*2) //此时len扩充2倍的话，后面新增的元素都被赋值为0
	copy(s_cp2, s)
	printSlice("s_cp2", s_cp2)
}

func printSlice(name string, x []int) { //注意形参写法：x []int
	fmt.Printf("我是切片%s，长度len为：%d，容量cap为：%d，值为：%v\n", name, len(x), cap(x), x)
}
