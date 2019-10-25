package time

import (
	//"fmt"
	"fmt"
	"strconv"
	"time"
)

//参考学习：https://www.jianshu.com/p/9d5636d34f17
func Time()  {
	//2006-01-02 15:04:05 GO的诞生时间

	//注意：time.Unix有两个参数，是将时间戳(或叫int64整数)转换成Time对象；
	//而time.Now().Unix()是没有参数的，直接返回当前时间戳

	var t time.Time
	t = time.Now()
	fmt.Printf("时间: %v, 时区: %v, 时间类型: %T\n", t, t.Location(), t) //实用
	//结果：时间: 2019-10-25 12:51:40.2986356 +0800 CST m=+0.025001401, 时区: Local, 时间类型: time.Time

	ss := t.Unix() //单位s，结果：1571994877
	ms := t.UnixNano()/int64(time.Millisecond) //单位ms，结果：1571994877677
	μs := t.UnixNano()/int64(time.Microsecond) //单位μs，结果：1571994877677535
	ns := t.UnixNano() //单位ns，结果：1571994877677535500

	fmt.Printf("当前时间戳（10位，即s ）为：%d \n", ss)
	fmt.Printf("当前时间戳（13位，即ms）为：%d \n", ms)
	fmt.Printf("当前时间戳（16位，即μs）为：%d \n", μs)
	fmt.Printf("当前时间戳（19位，即ns）为：%d \n", ns)
	//结果：当前时间戳（10位，即s）为：1571990641，当前时间戳（13位，即ms）为：1571990641843 ，当前时间戳（19位，即ns）为：1571990641843259600

	//GMT表示格林威治时间，CST表示四个时区的缩写，也可用来表示中国时区
	fmt.Printf("本地时区: %v\n", t.Location())
	//结果：本地时区: Local

	zoneName, s:= t.Zone()
	fmt.Printf("本地时区: %s，比格林威治时间快：%d秒\n", zoneName, s) //28800秒(s)=8时(h)
	//结果：本地时区: CST，比格林威治时间快：28800秒

	//指定时区
	var local *time.Location
	local, _ = time.LoadLocation("Asia/Shanghai") //还可以指定偏移量（s）
	fmt.Printf("指定时区：%v, 数据类型：%T\n", local, local)
	//结果：指定时区：Asia/Shanghai, 数据类型：*time.Location

	//今天是哪年哪月哪日？（默认日期为英文，可以int(month)转换成数字）
	year, month, day := t.Date()
	fmt.Printf("今天是%d年%d月%d日\n", year, int(month), day)
	//结果：今天是2019年10月25日

	year = t.Year()
	month = t.Month() //默认为英文
	day = t.Day()
	fmt.Printf("今天是%d年%d月%d日\n", year, int(month), day)
	//结果：今天是2019年10月25日

	var mon time.Month
	mon = 4
	fmt.Printf("数字%d月转化为英文：%s\n", mon, mon.String())
	//结果：数字4月转化为英文：April

	//今天是星期几？（默认星期为英文，可以int(week)转换成数字）
	week := t.Weekday()
	fmt.Printf("今天是星期%d\n", int(week))
	//结果：今天是星期5

	var wee time.Weekday
	wee = 5
	fmt.Printf("数字星期%d转化为英文：%s\n", wee, wee.String())
	//结果：数字星期5转化为英文：Friday

	//今天是今年的第几周？
	year, week_int := t.ISOWeek()
	fmt.Printf("今天是%d年的第%d周！\n", year, week_int)
	//结果：今天是2019年的第43周！

	//现在是几点几分几秒几纳秒？
	hour, min, sec := t.Clock()
	fmt.Printf("现在是%d点%d分%d秒\n", hour, min, sec)
	//结果：现在是12点51分40秒

	hour = t.Hour()
	min = t.Minute()
	sec = t.Second()
	nsec := t.Nanosecond()
	fmt.Printf("现在是%d点%d分%d秒%d纳秒\n", hour, min, sec, nsec)
	//结果：现在是12点51分40秒325637200纳秒

	//今天是今年的第几天
	thisDay := t.YearDay()
	fmt.Printf("今天是今年的第%d天\n", thisDay)
	//结果：今天是今年的第298天

	//日期字符串、时间戳字符串都可以转换成Time对象，日期字符串用time.Parse，时间戳字符串先转成int64再用time.Unix
	//1、时间戳或时间戳字符串转换成日期：
	//先将时间戳字符串转换成Time对象：a := time.Unix(int64(时间戳字符串), 0)，然后再用a.Format("2006-01-02 15:04:05")
	//2、日期或日期字符串转时间戳：
	//先将日期字符串用转成Time对象（日期已经是Time对象）：a := time.Parse("日期模版", 日期字符串)，然后a.Unix()即可

	//指定某一天是一年中第几天
	dateStr := "2019-10-25 00:00:00"
	theDay, _ := time.Parse("2006-01-02 15:04:05", dateStr) //先将日期字符串转换成time对象（实用） 模版：2006-01-02 15:04:05
	fmt.Printf("日期字符串%s转换成时间戳为：%d\n", dateStr, theDay.Unix())
	fmt.Printf("%s是%d年的第%d天\n", dateStr, theDay.Year(), theDay.YearDay())
	//结果：2019/10/25是2019年的第298天

	tsStr := "1571986050" //（实用）
	tsInt64, _ := strconv.ParseInt(tsStr, 10, 64)
	theDay = time.Unix(tsInt64, 0) //获取时间戳
	//time.Unix(tmp, 0)
	dateStr = theDay.Format("2006/01/02 15:04:05") //模版：2006-01-02 15:04:05
	fmt.Printf("时间戳%s转换成日期字符串为：%s\n", tsStr, dateStr)

	// 时间序列化
	t_byte, _ := t.MarshalJSON()
	fmt.Println("当前时间序列化:", t_byte)
	//结果：当前时间序列化: [34 50 48 49 57 45 49 48 45 50 53 84 49 50 58 53 49 58 52 48 46 51 50 53 54 51 55 50 43 48 56 58 48 48 34]

	// 时间数据反序列化
	var t_un time.Time
	t_un.UnmarshalJSON(t_byte)
	fmt.Println("当前时间反序列化: ", t_un)
	//结果：当前时间反序列化:  2019-10-25 12:51:40.3256372 +0800 CST



	//Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回
	//d_second := time.Second
	//time.Sleep(d_second)
	//fmt.Printf("阻塞时间为: %d s\n", d_second)
//1571909813
	//
	dst := time.Unix(int64(4128226613), 0) //2100-10-26 17:36:53
	isOk := t.Before(dst)
	fmt.Printf("目标时间%d比现在时间%d大（未来）吗？%v\n", dst.Unix(), t.Unix(), isOk) //结果：true

	dst = time.Unix(int64(1572082613), 0) //2019-10-26 17:36:53
	isOk = t.After(dst)
	fmt.Printf("目标时间%d比现在时间%d小（过去）吗？%v\n", dst.Unix(), t.Unix(), isOk) //结果：true
//Add
	t_new := t.AddDate(0, 1, 1)
	fmt.Println("'t.AddDate': ", t_new)

	t_new2 := t.Sub(dst)
	fmt.Println("'t.AddDate': ", t_new2)
}