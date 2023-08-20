package main

import (
	"fmt"
	"reflect"
)

/*
	字段标签（tag）并不是注释，而是用来对字段进行描述的元数据。尽管它不属于数据成员，但却是类型的组成部分。
	在运行期，可用反射获取标签信息。它常被用作格式校验、数据库关系映射等。
	标准库reflect.StructTag提供了分/解析功能。
	————《go语言学习笔记》雨痕 P125
*/

type User struct {
	Name string `姓名`
	Sex  byte   `性别`
}

func main() {

	user1 := User{"xinmin", 1}
	v := reflect.ValueOf(user1)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
	/* 输出：
	姓名: xinmin
	性别: 1
	*/

}
