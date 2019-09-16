package strategy

import (
	"fmt"
	"testing"
)

//策略者模式
//定义：策略模式可以让更换对象的内脏，而装饰者模式可以更换对象的外表。我们可以定义算法，封装它们，动态地切换它们。
//对比：模板模式与策略模式 https://segmentfault.com/a/1190000016199883

func TestStrategy(test *testing.T) {
	fmt.Println("装载加法方法：")
	operation := Operation{Addition{}} //往里塞个加方法，后面可以定义减方法或其他方法，也是这么个塞法，灵活机动
	a := 1985
	b := 1024
	res := operation.Operate(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, res)
	fmt.Println("装载减法方法：")
	operation = Operation{Subition{}}
	res = operation.Operate(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, res)
}
