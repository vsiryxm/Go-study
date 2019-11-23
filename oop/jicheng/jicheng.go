package jicheng

import (
	"encoding/json"
	"fmt"
)

/**
 * 使用结构体指针类型匿名字段来做继承
 **/
type Person struct {
	Name string `json:"name"` //名字
	Sex  byte   `json:"sex"`  //性别, 字符类型
	Age  int    `json:"age"`  //年龄
}

type Student struct {
	*Person        //指针类型，这里加`json:"person"`，json输出结果与Student2.Person2中一样的效果
	Score   int    `json:"score"`
	School  string `json:"school"`
}

func Fun1() {

	//第一种调用方式
	s1 := Student{&Person{"海阳之新", 'm', 34}, 100, "华师大"}
	fmt.Println("结构体指针 第一种调用方式：")
	fmt.Println(s1.Name, s1.Sex, s1.Age, s1.Score, s1.School)
	//结果：
	//结构体指针 第一种调用方式
	//海阳之新 109 34 100 华师大

	//第二种调用方式
	var s2 Student          //先定义变量
	s2.Person = new(Person) //再分配空间，妙妙妙，实用
	s2.Name = "Simon"       //可以直接访问了，这样一来，方便提取公共的结构体
	s2.Sex = 'm'
	s2.Age = 33
	s2.Score = 100
	s2.School = "华师大"
	fmt.Println("结构体指针 第二种调用方式：")
	fmt.Println(s2.Name, s2.Sex, s2.Age, s2.Score, s2.School)
	//结果：
	//结构体指针 第二种调用方式：
	//Simon 109 33 100 华师大

	//顺便看看json输出
	jsonByte, _ := json.MarshalIndent(s2, "", "    ")
	fmt.Println(string(jsonByte))
	//结果：
	/*
	   {
	       "name": "Simon",
	       "sex": 109,
	       "age": 33,
	       "score": 100,
	       "school": "华师大"
	   }
	*/
}

//=================================================================
/**
 * 使用结构体类型匿名字段来做继承
 **/
type Person2 struct {
	Name string `json:"name"` //名字
	Sex  byte   `json:"sex"`  //性别, 字符类型
	Age  int    `json:"age"`  //年龄
}

type Student2 struct {
	Person2 `json:"person"` //结构体，不是指针
	Score   int             `json:"score"`
	School  string          `json:"school"`
}

func Fun2() {

	//第一种调用方式
	s1 := Student2{Person2{"海阳之新", 'm', 34}, 100, "华师大"}
	fmt.Println("结构体 第一种调用方式：")
	fmt.Println(s1.Name, s1.Sex, s1.Age, s1.Score, s1.School)
	//结果：
	//结构体 第一种调用方式：
	//海阳之新 109 34 100 华师大

	//第二种调用方式
	var s2 Student2
	var person2 Person2 //先定义变量
	person2 = Person2{
		Name: "Simon",
		Sex:  'm',
		Age:  33,
	} //再分配空间
	s2.Person2 = person2
	s2.Score = 100
	s2.School = "华师大"
	fmt.Println("结构体 第二种调用方式：")
	fmt.Println(s2.Name, s2.Sex, s2.Age, s2.Score, s2.School)
	//结果：
	//结构体 第二种调用方式：
	//Simon 109 33 100 华师大

	//顺便看看json输出
	jsonByte, _ := json.MarshalIndent(s2, "", "    ")
	fmt.Println(string(jsonByte))
	//结果：
	/*
	   {
	       "person": {
	           "name": "Simon",
	           "sex": 109,
	           "age": 33
	       },
	       "score": 100,
	       "school": "华师大"
	   }
	*/
}
