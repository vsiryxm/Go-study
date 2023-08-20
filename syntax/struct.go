/* 定义结构体 */
//定义结构体指针分四步：
//1、定义结构体
//2、定义指针
//3、实例化结构体
//4、指针指向结构体实例

package main

import (
	"fmt"
	"os"
	"unsafe"
)

// 格式：type 自定义结构体名称 struct
type Person struct {
	Name   string
	Sex    rune //rune等同于int32
	Height uint8
}

func main() {
	var ouyang Person
	var daming Person

	ouyang.Name = "欧阳" //赋值方法一
	ouyang.Sex = '男'
	ouyang.Height = 175

	daming.Name = "大明"
	daming.Sex = '男'
	daming.Height = 177

	ella := Person{"艾拉", '女', 165} //赋值方法二，使用短语句定义

	fmt.Println(ouyang.Name) //Sex显示到网页，男可以用&#30007; 女可以用&#22899;
	fmt.Println(daming)
	fmt.Println(ella)

	fmt.Println("---结构体Person指针---")
	var ptr1 *Person //定义结构体指针，格式：var 指针名称 *结构体名
	ptr1 = &ouyang

	fmt.Printf("指针指向的内存地址为：%p\n", ptr1)                    //需要用格式输出，否则显示的是&{欧阳 30007 175}
	fmt.Println(*ptr1)                                     //输出内容，但这里的内容是一个结构体{欧阳 30007 175}
	fmt.Println((*ptr1).Name, (*ptr1).Sex, (*ptr1).Height) //按照常规写法，应该用*号去取值，与下面一行的写法等价
	fmt.Println(ptr1.Name, ptr1.Sex, ptr1.Height)          //Go提供了一种隐式解引用特性，可以直接用`指针名.结构字段`的形式访问值

	fmt.Println("---将结构体Person指针作为函数参数---")
	//达到与上面同样效果
	printPerson(&ouyang)

	a := [10]struct{}{}
	b := a[:]
	c := [0]int{}
	// unsafe是一个内置的标准库包，提供了一些访问底层内存和数据结构的工具。这个包的使用需要特别谨慎，
	// 因为它可以绕过 Go 语言的类型系统和内存安全机制，可能会导致不安全的行为。
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c)) // a和b的长度都为0

	// 这类长度为0的对象通常都指向runtime.zerobase变量 ————《go语言学习笔记》雨痕 P121
	// 空结构可作为通道元素类型，用于事件通知
	fmt.Printf("%p, %p, %p\n", &a[0], &b[0], &c) // 0x102367098, 0x102367098, 0x102367098

	// 匿名字段，又称嵌入字段或嵌入类型
	// 匿名字段的类型，可以是结构体、指针、接口等
	// 嵌入类型不是继承，无法实现多态
	type Scholar struct { // 学士
		Person           // 人的基本信息 嵌入类型
		Education string // 学历
	}
	scholar := Scholar{
		Person: Person{ // 匿名字段在赋值时，要用类型名作为键名来赋值
			Name:   "xinmin",
			Sex:    '男',
			Height: 175,
		},
		Education: "本科学士学位",
	}
	fmt.Printf("%#v\n", scholar) // main.Scholar{Person:main.Person{Name:"xinmin", Sex:30007, Height:0xaf}, Education:"本科学士学位"}
	fmt.Printf("%v\n", scholar.Name)

	type IsOK map[int]string
	type FileData struct {
		os.File
		IsOK
		Remark string
	}
	file1 := FileData{
		File:   os.File{}, // 如果嵌入类型为其他包中的类型，则隐式字段名不包括包名 ————《go语言学习笔记》雨痕 P122
		IsOK:   map[int]string{1: "xinmin"},
		Remark: "文件备注信息",
	}
	fmt.Printf("%#v\n", file1)
	fmt.Printf("%v\n", file1)

	type Student struct {
		Name string
	}
	type Scholar2 struct { // 学士
		Person // 多个匿名字段中，有同名字段时，比如这里两个类型都有Name字段
		Student
	}
	scholar2 := Scholar2{}
	// 在读写时就要显式访问了
	scholar2.Person.Name = "yyy"
	scholar2.Student.Name = "xxx"
	fmt.Println(scholar2)

	type Scholar3 struct { // 学士
		Student
		Name string // 与匿名字段中的Name重名
	}
	scholar3 := Scholar3{}
	scholar3.Name = "yyy"
	scholar3.Student.Name = "yyy" // 在访问Student下的Name时，就要显式指定了
	fmt.Println(scholar3)
}

func printPerson(ptr1 *Person) {
	fmt.Println(ptr1.Name, ptr1.Sex, ptr1.Height)
}
