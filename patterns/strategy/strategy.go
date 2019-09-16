package strategy

//策略者模式
//定义：策略模式可以让更换对象的内脏，而装饰者模式可以更换对象的外表。我们可以定义算法，封装它们，动态地切换它们。
//对比：模板模式与策略模式 https://segmentfault.com/a/1190000016199883

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (this *Operation) Operate(l, r int) int {
	return this.Operator.Apply(l, r)
}

//开始动态地往类里添加方法逻辑
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Subition struct{} //我加了个减法方法

func (Subition) Apply(lval, rval int) int { //相当于可以随时重写这个方法
	return lval - rval
}