package proxy

//代理模式
//定义：代理模式使得一个对象可以给另一个对象提供访问控制。截取所有访问。
//应用场景：代理模式可以使用在很多地方，例如网络连接，内存中大的对象，一个文件，或者其他消耗大的对象，或者是不可能被复制的对象。
//理解：等于在一个类之外加了一层，通过这一层做一些访问控制或其它操作。

import (
	"fmt"
	"sync"
)

type IObject interface {
	ObjDo(action string)
}

type Object struct {
	action string
}

func (this *Object) ObjDo(action string) {
	fmt.Println("I can do:" + action)
}

type ProObject struct {
	obj *Object
}

var one = new(sync.Once) //仅实例化一次，相当于单例模式，创建唯一代理

func (this *ProObject) ObjDo(action string) {
	one.Do(func() {
		fmt.Println("我唯一代理了Object，其他人谁也别跟我抢")
		if this.obj == nil {
			this.obj = new(Object)
		}
	})
	this.obj.ObjDo(action)
}
