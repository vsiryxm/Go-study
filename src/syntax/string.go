package main

import (
"fmt"
)

func main() {
    // 声明一个string类型变量并赋值
    var str1 string = "\\\"" //如果有转义字符，用双引号即可，如果没有，用反引号`
    var str4 string = "abcd"
    var str3 string = "海阳之新"

    // 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
    fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s \n", str1+str3, `\"`)

    var str2 string = `哈哈哈哈，我没有转义字符，可以直接用反引号`
    fmt.Printf("反引号（所见即所得）str2 = %s \n", str2)

    fmt.Printf("str3的长度为：%d\n", len(str3)); //在UniCode中，1个汉字占3个字节，len为Go的内置函数，不需要引用包
    fmt.Printf("str1的长度为：%d\n", len(str1)); //这里只算转义结果的长度
    fmt.Printf("str4的长度为：%d\n", len(str4));

}