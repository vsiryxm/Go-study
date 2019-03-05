package main

import "fmt"

func main() {
    var number int = 5
    if number := 4; 10 > number {
        number -= 4
        number += 3
        fmt.Print(number)
    } else if 10 < number {
        number -= 2
        fmt.Print(number)
    } else {
        fmt.Print(number)
    }
    number +=4
    fmt.Println(number)
}

//输出39