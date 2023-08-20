/* string类型 */
//string在go语言中是不可变值类型，在内存中的存储结构为：指针指向UTF-8字节数组
//默认值为空字符串""
//可以用索引号访问某字节

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 声明一个string类型变量并赋值
	var str1 string = "\\\"" //如果有转义字符，用双引号即可，如果没有，用反引号`
	var str4 string = "abcd"
	var str3 string = "海阳之新"
	var str5 string = "哈哈，" + //支持跨行，但是加号必须在行尾
		"我可以换行，但是加号必须在上一行末尾"

	// 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s \n", str1+str3, `\"`)

	var str2 string = `哈哈哈哈，我没有转义字符，可以直接用反引号`
	fmt.Printf("反引号（所见即所得）str2 = %s \n", str2)

	fmt.Printf("str3的长度为：%d\n", len(str3)) //在UniCode中，1个汉字占3个字节，len为Go的内置函数，不需要引用包
	fmt.Printf("str1的长度为：%d\n", len(str1)) //这里只算转义结果的长度
	fmt.Printf("str4的长度为：%d\n", len(str4))
	fmt.Println(str5)

	//可以用索引号访问某字节
	fmt.Println(str4[0]) //97  为a的ascii码

	//将一个整数通过string来转换，会当成Ascii码来转换成字符
	fmt.Println(string(65)) // A

	fmt.Println(int('A')) // 65

	//原生字符串，即用反引号括起来的字符串，在原生字符串中转义字符不会转义输出，而是原样输出
	//原生字符串正文内容不能包含`，但可以换成8进制或16进制使用+号拼接
	//应用场景：原生字符串面值用于编写正则表达式会很方便，因为正则表达式往往会包含很多反斜杠。
	//原生字符串面值同时被广泛应用于HTML模板、JSON面值、命令行提示信息以及那些需要扩展到多行的场景。
	fmt.Println(`\n我在反引号里就不转义了哦
	但换行与空格会保留输出
	`)

	fmt.Println(`
	{
		"code": 0,
		"msg": "OK",
		"data": {
			"areacode": 86,
			"mobile": "19918956336"
		}
	}
	`)

	//截取字符串，通过切片的形式对rune
	fmt.Println("---截取字符串---")
	ss := "Hello, World!"
	ss1 := ss[:5] //字符串本身是一个[]rune数组，在数组的基础上切片，获取不同的值
	ss2 := ss[7:]
	fmt.Println(ss)
	fmt.Println(ss1) //Hello
	fmt.Println(ss2) //World!

	//对字符串进行加工
	//要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。无⽆论哪种转换，都会重新分配内存，并复制字节数组。
	fmt.Println("---对字符串进行修改---")
	bs := []byte(ss)
	bs[0] = 'h' //修改首字母为小写h
	ss = string(bs)
	fmt.Println(ss)
	bs2 := []rune(ss)
	bs2[0] = '阳' //修改首字母为阳
	ss = string(bs2)
	fmt.Println(ss)

	fmt.Println("---对字符串进行遍历---")
	//⽤for 循环遍历字符串时，也有 byte 和 rune 两种⽅式。
	ss3 := "ouyang"
	for i := 0; i < len(ss3); i++ { //byte
		fmt.Printf("%c,", ss3[i])
	}
	fmt.Println()
	for _, r := range ss3 { //rune
		fmt.Printf("%c,", r)
	}
	fmt.Println()

	// 在一个字符串中有中英文时，使用range遍历可以隐式地解码，也可以显式地调用utf8.DecodeRuneInString解码
	ss4 := "Hello,世界"
	for i, r := range ss4 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	/*
		返回结果：
		0       'H'     72
		1       'e'     101
		2       'l'     108
		3       'l'     108
		4       'o'     111
		5       ','     44
		6       '世'    19990
		9       '界'    30028
	*/

	// 使用range统计中英文字符串的字符个数
	str := "Hello,世界"
	fmt.Println("length1=", len(str)) // len()函数并不是指字符个数，而是字节长度，这里返回的是12个字节长度，在utf8编码中，汉字占3个字节
	length := 0
	for range str { // for range 还可以省略k，v
		length++
	}
	fmt.Println("length2=", length)                      // 返回8个字符长度
	fmt.Println("length3=", utf8.RuneCountInString(ss4)) // 8

}
