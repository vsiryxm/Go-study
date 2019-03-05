package main

/* 数组定义 */
import (
    "fmt"
    "reflect"
)

func main()  {
    //第一种定义方法：
    var numbers [5]int  //数组这样定义
    numbers[0] = 2
    numbers[3] = numbers[0] -3 //数组元素的值可以这样直接赋值做修改
    numbers[1] = numbers[2] + 5
    numbers[4] = len(numbers)
    sum := 11

    fmt.Println(reflect.TypeOf(sum)) //输出变量类型
    fmt.Printf("%v\n", (sum == numbers[0]+numbers[1]+numbers[2]+numbers[3]+numbers[4]))
    //以上分别对应的值为 2 + 5 + 0 + -1 + 5，其中numbers[2]未赋值，默认值为0
    //var a,b int8 = 1,3
    //fmt.Printf("a与b是否相等，结果是：%v\n", a==b)

    //第二种定义方法：
    var numbers2 = [3]int{1, 2, 3}
    fmt.Println(len(numbers2))

    //第三种定义方法：
    var numbers3 [3]int
    fmt.Println(numbers3[0], numbers3[1], numbers3[2]) //只定义而没有赋值，所以输出都为默认值0

    //第四种定义方法：
    var numbers4 = [...]int{1, 2, 3} //定义不明确元素个数的数组时，可以用省略号...
    fmt.Println(len(numbers4))

}