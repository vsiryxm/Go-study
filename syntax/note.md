
《Go语言圣经》读书笔记

官网：https://go.dev/
自动垃圾回收
它没有隐式的数值转换，没有构造函数和析构函数，没有运算符重载，没有默认参数，也没有继承，没有异常，没有宏，没有函数修饰，更没有线程局部存储.
没有类，仅仅通过组合（而不是继承）简单的对象来构建复杂的对象。
。Go语言的动态栈使得轻量级线程goroutine的初始栈可以很小，因此，创建一个 goroutine的代价很小，创建百万级的goroutine完全是可行的。


https://go.dev/play/ 的 Go Playground 提供
可以生成对应的ur，只能导入标准库

https://tour.go-zh.org/welcome/1

下载并安装go语言
https://go.dev/dl/

阅读标准库中任意函数和类型的实现代码：
https://pkg.go.dev/std

Go语言原生支持Unicode，它可以处理全世界任何语言的文本

Go语言的代码通过包（package）组织，包类似于其它语言里的库 （libraries）或者模块（modules）

gofmt自动格式化代码
goimports自动添加或删除import声明

Go言里也采用左闭右开形式, 即，区间包括第一个索引元素，不包括最后一个
比如a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]，不包含最后一个元素

变量会在声明时直接初始化。如果变量没有显 式初始化，则被隐式地赋予其类型的零值（zero value），数值类型是0，字符串类型是空字符串
比如：
var name string // 隐式初始化，值为空串""
var age int32 // 隐式初始化，值为0

符号 := 是短变量声明的一部分

自增语句 i++ 给 i 加1，++和--都只能放在变量名后面


Go语言只有for循环这一种循环语句。for循环有多种形式：
for initialization; condition; post { 
    // zero or more statements 
}
initialization语句是可选的，在循环开始前执行。initalization如果存在，必须是一条简单语句，，即，短变量声明、自增语句、赋值语句或函数调用。
for循环的这三个部分每个都可以省略，如果省略 initialization 和 post ，分号也可以省 略：
for condition { 
    // ... 相当于while循环
}
如果连 condition 也省略了，像下面这样：
for { 
    // ... 死循环 可以通过break或return终止循环
}

for 循环的另一种形式, 在某种数据类型的区间（range）上遍历，如字符串或切 片。


go有一个特殊的变量，即空标识符_（下划线）
空标识符可用于任何语法需要变量名但程序逻辑不需要的时候, 例如, 在循环里，丢弃不需要 的循环索引, 保留元素值。

声明变量，下面这些都是等价的（推荐用前两种）：
s := ""  // 只能用在函数内部，而不能用于包变量
var s string // 默认初始化零值机制，被初始化为""
var s = ""  // 很少用
var s string = "" // 类型冗余
初始值重要的话就显式地指定变量的类型，否则使用隐式初始化。

str += "字符串拼接，每次拼接，都会丢弃str原来的内容，如果连接涉及的数据量很大，这种方式代价高昂"
一种简单且高效的解决方案是使 用 strings 包的 Join 函数
 strings.Join
练习 1.3： 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异
11.4节展示了如何写标准测试程序，以得到系统性的性能评测

--- 
读完P26

 map的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化。这种设计是有意为之的，因为能防止程序依赖特定 遍历顺序，而这是无法保证的。

 rand.Seed(time.Now().UTC().UnixNano()) //使用纳秒作为随机种子来保证随机性

 go的内置函数收集：
    append          -- 用来追加元素到数组、slice中,返回修改后的数组、slice
    close           -- 主要用来关闭channel
    delete            -- 从map中删除key对应的value
    panic            -- 停止常规的goroutine  （panic和recover：用来做错误处理）
    recover         -- 允许程序定义goroutine的panic动作
    imag            -- 返回complex的实部   （complex、real imag：用于创建和操作复数）
    real            -- 返回complex的虚部
    make            -- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
    new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
    cap                -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
    copy            -- 用于复制和连接slice，返回复制的数目
    len                -- 来求长度，比如string、array、slice、map、channel ，返回长度
    print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包

另一个在线教程：
    https://www.topgoer.com/go%E5%9F%BA%E7%A1%80/Golang%E5%86%85%E7%BD%AE%E7%B1%BB%E5%9E%8B%E5%92%8C%E5%87%BD%E6%95%B0.html