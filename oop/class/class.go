package class

type Person struct {
	Name string `json:"name"` //名字
	Sex  byte   `json:"sex"`  //性别, 字符类型
	Age  int    `json:"age"`  //年龄
}
