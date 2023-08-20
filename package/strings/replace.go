package strings

import (
	"golang.org/x/exp/errors/fmt"
	"strings"
	)

//返回将myStr中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
func Replace()  {

	myStr := "helloworld123helloworld123helloworld123helloworld"
	old := "123"
	new := "simon"
	fmt.Println("原字符串为：", myStr)

	newStr := strings.Replace(myStr, old, new, -1) //n=-1时
	fmt.Println("将字符串中所有123替换成simon：", newStr)
	//结果：helloworldsimonhelloworldsimonhelloworldsimonhelloworld

	newStr = strings.Replace(myStr, old, new, 0) //n=0时
	fmt.Println("n为0时不替换：", newStr)
	//结果：helloworld123helloworld123helloworld123helloworld

	newStr = strings.Replace(myStr, old, new, 1) //n>0时
	fmt.Println("n>0时替换n次：", newStr)
	//结果： helloworldsimonhelloworld123helloworld123helloworld

	newStr = strings.Replace(myStr, old, new, 3) //n>0时
	fmt.Println("n>0且等于old出现次数时替换全部：", newStr)
	//结果：helloworldsimonhelloworldsimonhelloworldsimonhelloworld

	newStr = strings.Replace(myStr, old, new, 4) //n>0时
	fmt.Println("n>0且超出old出现次数时替换全部：", newStr)
	//结果：helloworldsimonhelloworldsimonhelloworldsimonhelloworld

}