package main

/* 数组定义 */
import (
    "fmt"
    "reflect"
)

func main()  {

    //第一种定义方法：
    var numbers1 = [3]int{1, 2, 3}
    fmt.Println(len(numbers1))

    //第二种定义方法：
    var numbers2 [5]int  //数组这样定义
    numbers2[0] = 2
    numbers2[3] = numbers2[0] - 3 //数组元素的值可以这样直接赋值做修改
    numbers2[1] = numbers2[2] + 5
    numbers2[4] = len(numbers2)
    sum := 11

    fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
    //以上分别对应的值为 2 + 5 + 0 + -1 + 5，其中numbers2[2]未赋值，默认值为0

    fmt.Println(reflect.TypeOf(sum)) //输出变量类型

    //第三种定义方法：
    var numbers3 = [...]int{1, 2, 3} //定义不明确元素个数的数组时，可以用省略号...
    fmt.Println(len(numbers3))

    var numbers4 [3]int
    fmt.Println(numbers4[0], numbers4[1], numbers4[2]) //只定义而没有赋值，所以输出都为默认值0
    fmt.Printf("是：%d\n",len(numbers4))
    fmt.Print(numbers4) //说明：Print、Printf、Println三个函数都支持直接输出数组

    //var numbers_cp [5]int
    //copy(numbers_cp, numbers2) //报错arguments to copy must be slices，不支持数组复制
}