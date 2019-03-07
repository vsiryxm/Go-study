package main

import (
    "fmt"
    "reflect"
)

func main() {
    var numbers3 = [5]int{1, 2, 3, 4, 5}
    fmt.Println(reflect.TypeOf(numbers3))
    fmt.Println(len(numbers3))
    slice3 := numbers3[2 : len(numbers3)] //[3,4,5]
    fmt.Println(cap(slice3))
    fmt.Println(slice3)
    numbers3[2] = 33
    fmt.Println(slice3)


    length := (2)
    capacity := (2)
    fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
}