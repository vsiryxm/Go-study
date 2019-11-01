package strings

import (
	"fmt"
	"strings"
	)

func Strings()  {

	str := "https://www.baidu.com"
	isOk := strings.HasPrefix(str, "https://")
	fmt.Printf("字符串 %s 的开头是https://吗？%v\n", str, isOk)
	//结果：字符串 https://www.baidu.com 的开头是https://吗？true

	str = "ouyang.jpg"
	isOk = strings.HasSuffix(str, ".jpg")
	fmt.Printf("图片文件后缀 %s 是.jpg格式吗？%v\n", str, isOk)
	//结果：图片文件后缀 ouyang.jpg 是.jpg格式吗？true

	str = "vsiryxm@qq.com"
	subStr := "@"
	pos := strings.Index(str, subStr)
	fmt.Printf("%s 在 %s 中的位置是：%d\n", subStr, str, pos) //位置是从0开始计的
	//结果：@ 在 vsiryxm@qq.com 中的位置是：7
	fmt.Printf("从%s开始截取字符串：%s\n", subStr, str[pos:]) //字符串，在内存里的存储本质上是一个只读的byte数组
	//结果：从@开始截取字符串：@qq.com

	str = "vsiryxm@qq.com"
	subStr = "#"
	pos = strings.Index(str, subStr)
	fmt.Printf("%s 在 %s 中的位置是：%d\n", subStr, str, pos) //不存在时，值为-1
	//结果：# 在 vsiryxm@qq.com 中的位置是：-1

	str = "ssp.artarva.com"
	subStr = "."
	pos = strings.LastIndex(str, subStr) //从字符串尾部开始搜索，搜索到了返回位置，没有则返回-1
	fmt.Printf("%s 在 %s 中的位置是：%d\n", subStr, str, pos)
	//结果：. 在 ssp.artarva.com 中的位置是：11

	subStr = "#"
	pos = strings.LastIndex(str, subStr)
	fmt.Printf("%s 在 %s 中的位置是：%d\n", subStr, str, pos) //不存在时，值为-1
	//结果：# 在 ssp.artarva.com 中的位置是：-1

	str = "ssp.artarva.com"
	subStr = "."
	count := strings.Count(str, subStr)
	fmt.Printf("%s 在 %s 中出现的次数是：%d\n", subStr, str, count)
	//结果：. 在 ssp.artarva.com 中出现的次数是：2

	subStr = "#"
	count = strings.Count(str, subStr)
	fmt.Printf("%s 在 %s 中出现的次数是：%d\n", subStr, str, count)
	//结果：# 在 ssp.artarva.com 中出现的次数是：0

	str = "1"
	str += strings.Repeat("0", 10)
	fmt.Printf("在1后面重复添加10个0是：%s \n", str)
	//结果：在1后面重复添加10个0是：10000000000

	str = "OUYANG"
	str2 := strings.ToLower(str)
	fmt.Printf("%s转换成小写：%s \n", str, str2)
	//结果：OUYANG转换成小写：ouyang

	str = "ouyang"
	str2 = strings.ToUpper(str)
	fmt.Printf("%s转换成大写：%s \n", str, str2)
	//结果：ouyang转换成大写：OUYANG

	str = " ouyang "
	str2 = strings.TrimSpace(str)
	fmt.Printf("%s去掉首尾空格为：%s \n", str, str2)
	//结果： ouyang 去掉首尾空格为：ouyang

	str = " ouyang "
	str2 = strings.Trim(str, " ")
	fmt.Printf("%s去掉首尾空格为：%s \n", str, str2) //等同于上一行输出
	//结果： ouyang 去掉首尾空格为：ouyang

	str = "mongodb://off"
	str2 = strings.TrimLeft(str, "mongodb://")
	fmt.Printf("%s去掉前面的mongodb://为：%s \n", str, str2)
	//结果：mongodb://off去掉前面的mongodb://为：ff
	//坑：以上结果显然是我们不想要，这是一个坑，参见：
	//推荐用strings.TrimPrefix() strings.TrimSuffix() 替换 strings.TrimLeft() strings.TrimRight()

	str = "mongodb://off"
	str2 = strings.TrimPrefix(str, "mongodb://")
	fmt.Printf("%s去掉前面的mongodb://为：%s \n", str, str2)
	//结果：mongodb://off去掉前面的mongodb://为：off

	str = "/upload/ouyang/"
	str2 = strings.TrimRight(str, "ouyang/")
	fmt.Printf("%s去掉后面的ouyang/：%s \n", str, str2)
	//结果：/upload/ouyang/去掉后面的ouyang/为：/upload

	str = "hello world simon"
	arr := strings.Fields(str) //以空格为分隔符，将字符串分割成数组切片
	fmt.Printf("%s以空格分割成数组切片为：%v \n", str, arr)
	//结果：hello world simon以空格分割成数组切片为：[hello world simon]

	str = "a=1 b=2 c=3"
	arr2 := strings.Split(str, " ") //除空格外，可指定其他字符，等同于上一行输出
	fmt.Printf("%s用空格分割成数组切片为：%v \n", str, arr2)
	//结果：a=1 b=2 c=3用空格分割成数组切片为：[a=1 b=2 c=3]

	str = strings.Join(arr2, "&")
	fmt.Printf("数组%v用&连接起来为：%s \n", arr2, str)
	//结果：数组[a=1 b=2 c=3]用&连接起来为：a=1&b=2&c=3

	str = "021-66666666"
	isOk = strings.Contains(str, "-")
	fmt.Printf("%s包含-吗：%v \n", str, isOk)
	//结果：021-66666666包含-吗：true
}