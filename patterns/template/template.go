package template

//定义：模版方法设计模式允许把对象不同的部分抽象，在同一段代码中执行相同的逻辑，增加可拓展性。在Go语言中，实现由底层对象实现，而行为由顶层方法控制。
//理解：有点像策略模式更换核心的做法，区别时，模板设计模式已经预先知道有哪些
//对比：模板模式与策略模式 https://segmentfault.com/a/1190000016199883

import (
	"fmt"
)

type Tmpl interface {
	DoSomething()
	DoAnything()
}

type TmplA struct {
}

func (this *TmplA) DoSomething() {
	fmt.Println("TmplA.DoSomething")
}

func (this *TmplA) DoAnything() {
	fmt.Println("TmplA.DoAnything")
}

type TmplB struct {
}

func (this *TmplB) DoSomething() {
	fmt.Println("TmplB.DoSomething")
}

func (this *TmplB) DoAnything() {
	fmt.Println("TmplB.DoAnything")
}

func Operate(tmpl Tmpl) {
	tmpl.DoSomething()
	tmpl.DoAnything()
}
