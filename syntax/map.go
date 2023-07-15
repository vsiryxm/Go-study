package main

import "fmt"

func main() {

	fmt.Println("---第一种定义方法---")
	var city map[string]string
	city = make(map[string]string)
	//完整写法： var city map[string]string = make(map[string]string)

	city["shanghai"] = "上海"
	city["loudi"] = "娄底"
	city["changsha"] = "长沙"
	city["suzhou"] = "苏州"

	// 遍历map
	// map的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化。
	// 这种设计是有意为之的，因为能防止程序依赖特定 遍历顺序，而这是无法保证的。
	for k, v := range city {
		fmt.Printf("%s的中文是：%s \n", k, v)
	}

	fmt.Println("\n---第二种定义方法---")
	city2 := map[string]string{"shanghai": "上海", "loudi": "娄底", "changsha": "长沙", "suzhou": "苏州"}
	fmt.Println(city2)

	fmt.Println("\n---查看元素是否存在---")
	value, isExsit := city["shanghai"] //第一个变量为值，第二个变量为逻辑true或false
	if isExsit {
		fmt.Println("shanghai元素存在且值为：", value)
	}

	value2, isExsit2 := city["xinhua"]
	if isExsit2 {
		fmt.Println("xinhua元素存在且值为：", value2)
	} else {
		fmt.Println("xinhua元素不存在")
	}

	fmt.Println("\n---查找元素是否存在于map中---")
	hasKey(city, "suzhou", "hunan", "beijing")

	delete(city, "shanghai")
	fmt.Println("shanghai被删除：", city)

	// 利用map来查询重复数据
	counts := make(map[string]int)
	stageNames := []string{"海阳之新", "东胖", "xinmin", "东胖", "vsiryxm", "海阳之新", "海阳之新"}
	for _, v := range stageNames {
		counts[v]++
	}
	for m, n := range counts {
		if n > 1 {
			fmt.Printf("%s重复了，有%d个\n", m, n)
		}
	}

}

// 查找元素是否存在于map中
func hasKey(city map[string]string, c ...string) {
	var value string
	var has bool
	for _, v := range c {
		value, has = city[v]
		if has {
			fmt.Printf("%s元素存在且值为：%s\n", v, value)
		} else {
			fmt.Printf("%s元素不存在\n", v)
		}
	}
}
