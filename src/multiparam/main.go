package main

import "fmt"

func main()  {
    if err := is_odd_number(1,2,3,4,5); err != nil {
        fmt.Println(err)
    }
}

//是否为奇数
func is_odd_number(n ...int) error {
    for _, i := range n {
        if i % 2 == 0 {
            return fmt.Errorf("不是奇数", n)
        }
        fmt.Println(n, "是奇数")
    }
    return nil
}
