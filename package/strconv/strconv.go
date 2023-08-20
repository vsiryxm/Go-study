package strings

import (
	"fmt"
	"strconv"
	)

func Strconv()  {

	str := "-100" //可以有负号，不能有小数点，如100.00会转换出错
	num, err := strconv.Atoi(str) //字符串转Int
	if err != nil  {
		fmt.Printf("错误：%s\n", err) //转换失败，num=0
		//结果：错误：strconv.Atoi: parsing "100.00": invalid syntax
	} else {
		fmt.Printf("字符串%s（%T）转换成数字为：%d（%T）\n", str, str, num, num)
		//结果：字符串100转换成数字为：100
	}

	num = 200 //必须是int型
	str = strconv.Itoa(num)
	fmt.Printf("数字%d（%T）转换成字符串为：%s（%T）\n", num, num, str, str)

	//接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	str = "True" //ff会报错
	isOk, err := strconv.ParseBool(str)
	if err != nil  {
		fmt.Printf("错误：%s\n", err) //转换失败，num=0
		//结果：错误：strconv.ParseBool: parsing "ff": invalid syntax
	} else {
		fmt.Printf("字符串%s（%T）转换成Bool型为：%v（%T）\n", str, str, isOk, isOk)
		//结果：字符串f（string）转换成Bool型为：false（bool）
		//结果：字符串T（string）转换成Bool型为：true（bool）
	}
	fmt.Println(strconv.ParseBool("TRue")) //不支持
	//结果：false strconv.ParseBool: parsing "TRue": invalid syntax
	fmt.Println(strconv.ParseBool("FALse")) //不支持
	//结果：false strconv.ParseBool: parsing "FALse": invalid syntax

	//====================================================================

	//ParseInt可以将各种进制（2到36）的数据转换成int（0到64）型，功能很强大
	//ParseInt(数字字符串，数字字符串的进制，转换后不能超过的数据范围)
	//第二个参数指定进制（2到36），指明第一个参数是什么进制类型数据，如果base为0，则会从字符串前置判断，“0x”是16进制，“0”是8进制，否则是10进制
	//最后一个参数可以是0、8、16、32、64，分别代表int、int8、int16、int32、int64
	//指定最后一个参数，会检查转换后的数字有没有超出类型的范围，超出了会报错
	str = "123" //十进制数据
	num2, err2 := strconv.ParseInt(str, 10, 64)
	if err2 != nil  {
		fmt.Printf("错误：%s\n", err2) //转换失败，num=0
		//结果：错误：strconv.ParseInt: parsing "1234567890000000000": value out of range
	} else {
		fmt.Printf("字符串%s（十进制）转换成%T型为：%v\n", str, num2, num2)
		//结果：字符串123（十进制）转换成int64型为：123
	}

	str = "1010" //二进制数据
	num3, _ := strconv.ParseInt(str, 2, 64)
	fmt.Printf("字符串%s（二进制）转换成%T型为：%v\n", str, num3, num3)
	//结果：字符串1010（二进制）转换成int64型为：10

	str = "567" //八进制数据
	num4, _ := strconv.ParseInt(str, 8, 64)
	fmt.Printf("字符串%s（八进制）转换成%T型为：%v\n", str, num4, num4)
	//结果：字符串567（八进制）转换成int64型为：375

	str = "0x999" //十六进制数据，base=0时，自动取字符串前缀判断是什么进制
	num5, _ := strconv.ParseInt(str, 0, 64)
	fmt.Printf("字符串%s（十六进制）转换成%T型为：%v\n", str, num5, num5)
	//结果：字符串999（十六进制）转换成int64型为：2457

	str = "999" //十七进制数据，支持2~36之间任意数字进制
	num6, _ := strconv.ParseInt(str, 17, 64)
	fmt.Printf("字符串%s（十七进制）转换成%T型为：%v\n", str, num6, num6)
	//结果：字符串999（十七进制）转换成int64型为：2763

	str = "42"
	num7, err := strconv.ParseUint(str, 10, 64) //转换成uint类型
	fmt.Printf("字符串%s转换成%T为：%d\n", str, num7, num7)
	//结果：字符串42转换成uint64为：42

	str = "3.1415"
	num8, err := strconv.ParseFloat(str, 64)
	fmt.Printf("字符串%s转换成%T为：%f\n", str, num8, num8)
	//结果：字符串3.1415转换成float64为：3.141500

	//====================================================================

	//FormatX函数将给定类型格式化为string类型
	//将bool、float、int、uint类型转换成string
	//FormatBool()、FormatFloat()、FormatInt()、FormatUint()
	str = strconv.FormatBool(true)
	fmt.Printf("true转换成字符串为：%v\n", str)
	//结果：true转换成字符串为：true

	str = strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("3.1415转换成字符串为：%v\n", str)
	//结果：3.1415转换成字符串为：3.1415E+00

	//十进制int数字转换各种进制数，第二个参数表示要转换成什么进制，有效值为2<=base<=36
	str = strconv.FormatInt(-12, 2) //表示将十进制12（int）转换为2进制数
	fmt.Printf("-12转换成字符串为：%v\n", str)
	//结果：-12转换成字符串为：-1100

	//十进制uint数字转换各种进制数
	str = strconv.FormatUint(12, 2)
	fmt.Printf("12转换成字符串为：%v\n", str)
	//结果：12转换成字符串为：1100

	//====================================================================

	//AppendX类函数用于将X转换成字符串后append到一个slice中
	//AppendBool()、AppendFloat()、AppendInt()、AppendUint()
	arr := []byte("我是一个切片，在切片后面加上各种进制数字：")
	//将转换为2进制的string，追加到slice中
	arr = strconv.AppendInt(arr, -12, 2)
	fmt.Println(string(arr))
	//结果：我是一个切片，在切片后面加上各种进制数字：-1100


}
//参考：https://www.cnblogs.com/f-ck-need-u/p/9863915.html