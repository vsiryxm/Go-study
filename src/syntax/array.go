package main

/* 数组定义 */
import (
	"fmt"
	"reflect"
)

func main() {

	//第一种定义方法：
	var numbers1 = [3]int{1, 2, 3}
	fmt.Println(len(numbers1))

	//第二种定义方法：
	var numbers2 [5]int //数组这样定义
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3 //数组元素的值可以这样直接赋值做修改
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := 11

	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
	//以上分别对应的值为 2 + 5 + 0 + -1 + 5，其中numbers2[2]未赋值，默认值为0

	fmt.Println(reflect.TypeOf(sum)) //输出变量类型

	fmt.Println("数组的内存地址列表：")
	for k,_ := range numbers2 {
		fmt.Print(&numbers2[k], " ")
	} //0xc000078030 0xc000078038 0xc000078040 0xc000078048 0xc000078050 说明数组在内存中分配不一定是连续的
	fmt.Println()

	//第三种定义方法：
	var numbers3 = [...]int{1, 2, 3} //定义不明确元素个数的数组时，可以用省略号...
	fmt.Println(len(numbers3))

	var numbers4 [3]int
	fmt.Println(numbers4[0], numbers4[1], numbers4[2]) //只定义而没有赋值，所以输出都为默认值0
	fmt.Printf("是：%d\n", len(numbers4))
	fmt.Print(numbers4) //说明：Print、Printf、Println三个函数都支持直接输出数组

	//var numbers_cp [5]int
	//copy(numbers_cp, numbers2) //报错arguments to copy must be slices，不支持数组复制

	//定义二维数组
	fmt.Println("\n\n定义二维数组：")
	var arr = [3][3]int64 {
		{1,2,3},
		{4,5,6},
		{7,8,9}, //如果最后一个元素不以逗号结尾，数组的右花括号不能另起一行
		}
	fmt.Println("\n遍历二维数组方法一：")
	for i:=0; i<3; i++ {
		for j:=0; j<len(arr[i]); j++ {
			fmt.Printf("a[%d][%d]=%d ", i, j, arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\n遍历二维数组方法二：")
	for k,v := range arr  {
		for k1,v1 := range v {
			fmt.Printf("arr[%d][%d]=%d ", k, k1, v1)
		}
		fmt.Println()
	}


}
