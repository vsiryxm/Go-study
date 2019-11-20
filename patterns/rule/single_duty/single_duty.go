package single_duty

import "fmt"

//单一职责
//一个类只负责一项职责，尽量做到类的只有一个行为原因引起变化

//接口与类的混合职责示例
type DoSomething interface {
	BringBaby() //带娃
	Coding() //写程序
}

type Person struct {
	Name string
}

func (this Person) BringBaby()  {
	fmt.Println(this.Name + "带娃")
}

func (this Person) Coding()  {
	fmt.Println(this.Name + "编码")
}

//=============================================
//接口与类的单一职责示例

type DoBringBaby interface {
	BringBaby() //带娃
}

type DoCoding interface {
	Coding() //写程序
}

type Man struct {
	Name string
}

type WoMen struct {
	Name string
}

func (this Man) Coding()  {
	fmt.Println(this.Name + "男人编码")
}

func (this WoMen) BringBaby()  {
	fmt.Println(this.Name + "女人带娃")
}

//方法
func Login(mobile, email string)  {
	if mobile != "" {
		fmt.Println("验证手机格式")
		fmt.Println("登录")
	}

	if email != "" {
		fmt.Println("验证邮箱格式")
		fmt.Println("登录")
	}
}

//将验证逻辑剥离出来，登录和验证分开
func CheckData(str string) bool  {
	fmt.Println("验证手机格式")
	fmt.Println("验证邮箱格式")
	return true
}