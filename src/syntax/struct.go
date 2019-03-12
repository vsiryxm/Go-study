/* 定义结构体 */
//定义结构体指针分四步：
//1、定义结构体
//2、定义指针
//3、实例化结构体
//4、指针指向结构体实例

package main

import "fmt"

// 格式：type 自定义结构体名称 struct
type Person struct {
    name string
    sex rune //rune等同于int32
    height uint8
}
func main()  {
    var ouyang Person
    var daming Person

    ouyang.name = "欧阳" //赋值方法一
    ouyang.sex = '男'
    ouyang.height = 175

    daming.name = "大明"
    daming.sex = '男'
    daming.height = 177

    ella := Person{"艾拉", '女', 165} //赋值方法二，使用短语句定义

    fmt.Println(ouyang.name) //sex显示到网页，男可以用&#30007; 女可以用&#22899;
    fmt.Println(daming)
    fmt.Println(ella)

    fmt.Println("---结构体Person指针---")
    var ptr1 *Person //定义结构体指针，格式：var 指针名称 *结构体名
    ptr1 = &ouyang

    fmt.Println(ptr1) //输出的不是内存地址，而是&{欧阳 30007 175}
    fmt.Println(*ptr1) //输出的是{欧阳 30007 175}，而不是内存地址
    fmt.Println((*ptr1).name, (*ptr1).sex, (*ptr1).height) //按照常规写法，应该用*号去取值，与下面一行的写法等价
    fmt.Println(ptr1.name, ptr1.sex, ptr1.height) //Go提供了一种隐式解引用特性，可以直接用`指针名.结构字段`的形式访问值

    fmt.Println("---将结构体Person指针作为函数参数---")
    //达到与上面同样效果
    printPerson(&ouyang)

}

func printPerson(ptr1 *Person)  {
    fmt.Println(ptr1.name, ptr1.sex, ptr1.height)
}