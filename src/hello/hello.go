package main // 代码包声明语句

// 代码包导入语句
import (
    "fmt" //导入代码包fmt
    "mylib/calc"  //包所在的文件夹路径，先去找GOROOT/src/mylib/lib文件夹，再去找GOPATH/src/mylib/lib文件夹
)

// main函数
func main()  { // 代码块由“{”和“}”包裹
    fmt.Printf("Hello World!\nHello Simon!\n")

    var a,b,c int64
    a = 123
    b = 1
    c = calc.Add(a, b) //调用格式：包名.方法名（参数1，参数2...）
    fmt.Printf("a + b = %d\n", c)
    //fmt.Printf(aaa()); //会报错，不能这样直接输出，必须使用占位符配合输出
    fmt.Println(aaa());
    fmt.Print(aaa());

}

func aaa() (int32) {
    return 123;
}

