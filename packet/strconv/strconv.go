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

	//接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	str = "12345888888" //ff会报错
	//第一个参数指定进制（2到36），如果base为0，则会从字符串前置判断，“0x”是16进制，“0”是8进制，否则是10进制
	//最后一个参数可以是0、8、16、32、64，分别代表int、int8、int16、int32、int64
	//指定最后一个参数，会检查转换后的数字有没有超出类型的范围，超出了会报错
	num2, err2 := strconv.ParseInt(str, 2, 64)
	if err2 != nil  {
		fmt.Printf("错误：%s\n", err2) //转换失败，num=0
		//结果：错误：strconv.ParseInt: parsing "123458888888888888888": value out of range
	} else {
		fmt.Printf("字符串%s（%T）转换成Int型为：%v（%T）\n", str, str, num2, num2)
		//结果：字符串f（string）转换成Bool型为：false（bool）
		//结果：字符串T（string）转换成Bool型为：true（bool）
	}






}