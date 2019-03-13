package main

import "fmt"

/* switch特点：
1、不同的 case 之间不使用 break 分隔，默认只会执行一个 case。
2、使用 fallthrough 会强制执行后面的 case 语句（不会去判断下一个 case 的表达式是否为 true）。
   用了fallthrough后，后面必须还要有case或者default，不然报错，也可用 break 终止
3、支持多条件匹配
*/

func main() {

	var a uint8 = 90
	//第一种使用方法：如果条件放在case中，那么switch后面不带任何条件表达式
	switch {
	case a >= 60 && a < 70:
		fmt.Println("及格")
	case a >= 70 && a < 80:
		fmt.Println("一般")
	case a >= 80 && a < 90:
		fmt.Println("良好")
	case a >= 90:
		fmt.Println("优秀")
		fallthrough //相当于在别的语言里不写break，比如js，程序会继续执行下一条 case，且它不会去判断下一个 case 的表达式是否为 true
		//用了fallthrough后，后面必须还要有case或者default，不然报错，go语言官方认为如果是在最后一条case语句中就没必要写这个关键字了
	default: //可写可不写
		fmt.Println("未及格")
	}

	//第二种使用方法：如果条件放在switch中，那么case后面不带任何条件表达式，而是值
	switch a {
	case 70:
		fmt.Println("一般")
	case 80:
		fmt.Println("良好")
	case 90:
		fmt.Println("优秀")
	default: //可写可不写
		fmt.Println("未及格")
	}

	//var myname string = "欧阳"
	switch myname := "欧阳"; { //可以带初始化语句
	case myname=="欧阳":
		if(myname!="欧阳") { break } //也可以手工增加break操作
		fmt.Println("你是欧阳") //默认自动加break
	case myname=="Ella":
		fmt.Println("你是Ella")
	case myname=="大明":
		fmt.Println("你是大明")
	default: //可写可不写
		fmt.Println("辉哥")
	}

	//用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}
	switch is_type := x.(type) { //x.(type)获取一个变量的类型
	case byte, int8: //多个值匹配的用法
		fmt.Printf("myname变量是%T型", is_type)
	case string:
		fmt.Printf("myname变量是%T型", is_type)
	case uint:
		fmt.Printf("myname变量是%T型", is_type)
	case bool:
		fmt.Printf("myname变量是%T型", is_type)
	default: //可写可不写
		fmt.Printf("未知类型：%T", is_type)
	}

}
