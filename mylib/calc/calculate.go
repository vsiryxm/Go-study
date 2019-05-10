package calc //包名：在调用时以包名为准，包名可以跟文件同名，也可以与文件夹同名，惯例是与所在文件夹同名

//自定义函数 func 函数名(参数1，参数2...) (返回参数类型1，返回参数类型2...)
func Add(_a int64, _b int64) (int64) { //如果函数为小写，则相当于设置了private
    var _c int64
    _c = _a + _b
    return _c
}

func Sub(_a int64, _b int64) (int64) {
    var _c int64
    _c = _a - _b
    return _c
}

func Mul(_a int64, _b int64) (int64) {
    var _c int64
    _c = _a * _b
    return _c
}

func Div(_a int64, _b int64) (int64) {
    var _c int64
    _c = _a / _b
    return _c
}


