/* 函数声明 */
//通过 func 关键字来进行函数的声明。
//在Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用：
//函数名首字母小写即为 private，函数名首字母大写即为 public

package main

import (
	"fmt"
	"mylib/calc" //包所在的文件夹路径，先去找GOROOT/src/mylib/calc文件夹，再去找GOPATH/src/mylib/calc文件夹
)

func main() {

	fmt.Println("\n在同一个包下，调用函数，返回多个值且怎样去接收值：")
	a, b, _, strs := multiReturn(1) //注意：获取函数返回的多个值时，如果其中一个值不想接收，可以用下划线表示

	fmt.Println(a, b, strs)

	fmt.Println("\n在不同的包下，调用函数：")
	c := calc.Add(a, b) //调用格式：包名.方法名（参数1，参数2...），函数首字母必须大写才能被调用（相当于public）
	fmt.Printf("a + b = %d\n", c)

	fmt.Println("\n不确定的多个参数调用函数：")
	if err := is_odd_number(1,3,5,7,9,10); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n通过传递引用，交换两个变量的值：")
	var aa,bb int64 = 10,20
	fmt.Printf("交换前，aa=%d，bb=%d\n", aa, bb)
	swap(&aa, &bb)
	fmt.Printf("交换后，aa=%d，bb=%d\n", aa, bb)

}

func multiReturn(a int64) (int64, int64, int64, string) { //小写开头的函数相当于private
	b, c, d := 2, 3, "this is string"
	return a, int64(b), int64(c), d //Go中不存在隐式转换，所有类型转换必须显式声明，且转换只能发生在两种相互兼容的类型之间
}

//传递多个不确定的参数的函数定义方法，用...省略号
//是否为奇数
func is_odd_number(n ...int) error {
	for _, i := range n {
		if i % 2 == 0 {
			return fmt.Errorf("不是奇数", n)
		}
		fmt.Print(i, "是奇数\n")
	}
	return nil
}

//通过传递引用，交换两个变量的值
func swap(x *int64, y *int64) {
	var tmp int64
	tmp = *x
	*x = *y
	*y = tmp
}
