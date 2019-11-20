package liskov_replace

import "fmt"

type Person interface {
	EatFood() //吃东西
}

type Person struct {
	Name string
}

type Man struct {

}

func (this Person) Coding()  {
	
}

func main()  {
	var person1 Person
	person1.Name = "张三"


}

//http://www.findme.wang/blog/detail/id/453.html
//https://www.cnblogs.com/xuxu8511/p/3296546.html
//https://www.cnblogs.com/wangxusummer/p/5250798.html