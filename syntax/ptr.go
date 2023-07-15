/* 指针 */

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("\n------简易指针------")
	var a int8 = 34
	var ptr *int8 //定义指针用*号
	ptr = &a      //获取变量a的内存地址
	//ptr++ //指针不能做加减法运算，编译报错
	fmt.Println("a的值为：", a)     //34
	fmt.Println("*ptr为：", *ptr) //34
	fmt.Println("ptr为：", ptr)   //0xc0000541c8

	// 参考Go语言圣经P61 另一个创建变量的方法是调用用内建的new函数
	// 表达式new(T)将创建一个T类型的匿名变 量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为 *T
	// pp := new(int8) // p, *int 类型, 指向匿名的 int8 变量，等同于下两行的语法
	var abc int8
	pp := &abc
	fmt.Println("pp为：", pp)

	fmt.Println("\n------指针数组------")

	var arr = [3]int64{1, 2, 3}
	var ptrarr [3]*int64
	for k, _ := range arr {
		ptrarr[k] = &arr[k]
	}
	fmt.Println("指针指向的内存地址有：", ptrarr)
	fmt.Println("遍历指针数组的值：")
	for k, _ := range ptrarr {
		fmt.Println(*ptrarr[k])
	}

	fmt.Println("\n------指针向指针的指针------")
	var ptr2 **int8
	ptr2 = &ptr                  //ptr2为指针向指针的指针
	fmt.Println(&a, ptr, *ptr2)  //三者完全等价
	fmt.Println(a, *ptr, **ptr2) //三者完全等价

	fmt.Println("\n------将指针作为函数参数------")
	var aa, bb int64 = 10, 20
	fmt.Printf("交换前，aa=%d，bb=%d\n", aa, bb)
	swap(&aa, &bb)
	fmt.Printf("交换后，aa=%d，bb=%d\n", aa, bb)

	fmt.Println("\n------空指针------")
	var ptr3 *int8
	fmt.Printf("ptr3的值为 : %x\n", ptr3)
	if ptr3 == nil { //直接这么写0==nil是不成立的，即数字0和nil不相等
		fmt.Println("ptr3是一个空指针")
	}

	fmt.Println("\n------指向结构体的指针------")
	//定义结构体指针分四步：
	//1、定义结构体
	//2、定义指针
	//3、实例化结构体
	//4、指针指向结构体实例
	type Person struct {
		name   string
		sex    rune
		height uint8
	}
	var ouyang Person
	ouyang.name = "欧阳"
	ouyang.sex = '男'
	ouyang.height = 175
	//也可以这样赋值：ouyang := Person("欧阳", '男', 175)

	var ptr4 *Person
	ptr4 = &ouyang
	fmt.Println(*ptr4)                                     //输出：{欧阳 30007 175}
	fmt.Println((*ptr4).name, (*ptr4).sex, (*ptr4).height) //按照常规写法，应该用*号去取值，与下面一行的写法等价
	fmt.Println(ptr4.name, ptr4.sex, ptr4.height)          //输出：欧阳 30007 175 Go提供了一种隐式解引用特性，可以直接用`指针名.结构字段`的形式访问值

	//可以在 unsafe.Pointer 和任意类型指针间进⾏转换。
	x := 0x12345678
	p := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(p)      // Pointer -> *[4]byte //没搞明白*[4]byte，不知道是不是数组指针?
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	} //输出：78 56 34 12

}

// 通过传递引用，交换两个变量的值
func swap(x *int64, y *int64) {
	//var tmp int64
	//tmp = *x
	//*x = *y
	//*y = tmp
	*x, *y = *y, *x //交换语句这样写更加简洁，也是 go 语言的特性，c++ 和 c# 是不能这么干的
}

/*
心得笔记：
- 指针实际也是一个变量，只不过它只用于存储内存地址，而内存地址其实是一个正整数，所以指针也可以理解成整型变量。
- 当定义一个空指针，如果是直接输出则不会报错，如果需要访问空指针成员（如struct中的成员），或将他赋值给另一个变量，再打印这个变量就会报空指针异常。
- 在go语言中，指针不能指针移动操作（即加减位置），指针可以做==运算，当指针为nil或指向同一个内存地址时，此时两个指针相等。
- 不管是指针、引用类型，还是其他类型参数，都是值拷贝传递（pass-by-value)。区别无非是拷贝目标对象，还是拷贝指针而已。在函数调用前，会为形参和返回值分配内存空间，并将实参拷贝到形参内存。
- 引用类型是特指slice、map、channel这三种预定义类型，必须使用make函数来创建。
- 如何避免系统报空指针异常？在访问指针成员时，要先判断指针是否为nil，在向指针写入数据时，值必须是一个已经开辟好的内存地址，直接赋值不能算已经开辟好的内存地址，直接赋值后再访问还是会报错。
*/
