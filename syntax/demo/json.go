package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Name interface{} `json:"name"`
    Sex  interface{} `json:"sex2"` //如果json字符串和结构体中的变量名称不一样，可以通过`json:key`形式明确对应关系
}

func main()  {
    var err error
    person := Person{}
    data := "{\"name\":\"simon\", \"sex2\":\"man\"}"
    str := []byte(data)
    err = json.Unmarshal(str, &person) //将json转换成结构体
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(person)


}
