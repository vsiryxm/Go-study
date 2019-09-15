package proxy

import "testing"

func TestProxy(test *testing.T) {
	proxy := new(ProObject)
	proxy.ObjDo("Well")
	proxy2 := new(ProObject) //第二次实例化时，直接访问内存中的对象，而不做new(Object)操作
	proxy2.ObjDo("编程")
}
