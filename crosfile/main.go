package main



func main()  {
    //跨文件访问函数，但必须声明在一个包名下
    //这样做的好处在于，如果一个文件代码太多，可以分成多个文件，在go build构建的时候把文件名都带上
    CrosAccess()
}
