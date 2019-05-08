/*
 在php中，json_decode函数是将json字符串转换成数组。
 在go中，我们通过json.Unmarshal来完成类似功能，将json转换成struct。
 */

package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Name interface{} `json:"name"` //如果成员以小写开头，则访问不到，值为nil
    Height interface{} `json:"height"`
    IsLike interface{} `json:"is_like"` //如果json中的key与这里的变量名不同，可以使用`json:key`的形式明确声明
}

func main()  {
    var err error
    person := Person{}
    var data string = "{\"name\":\"simon\", \"height\":\"175\", \"is_like\":\"true\"}"
    arr := []byte(data)
    fmt.Println(arr)
    err = json.Unmarshal(arr, &person)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(person)

}