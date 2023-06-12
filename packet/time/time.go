package time

import (
	//"fmt"
	"fmt"
	"strconv"
	"time"
)

func Time() {
	//2006-01-02 15:04:05 GO的诞生时间

	//注意：time.Unix有两个参数，是将时间戳(或叫int64整数)转换成Time对象；
	//而time.Now().Unix()是没有参数的，直接返回当前时间戳

	var t time.Time
	t = time.Now()
	fmt.Printf("时间: %v, 时区: %v, 时间类型: %T\n", t, t.Location(), t) //实用
	//结果：时间: 2019-10-26 16:17:54.9133283 +0800 CST m=+0.004000201, 时区: Local, 时间类型: time.Time

	ss := t.Unix()                               //单位s
	ms := t.UnixNano() / int64(time.Millisecond) //单位ms
	us := t.UnixNano() / int64(time.Microsecond) //单位us
	ns := t.UnixNano()                           //单位ns

	fmt.Printf("当前时间戳（10位，即s ）为：%d \n", ss)
	//结果：当前时间戳（10位，即s ）为：1572077874
	fmt.Printf("当前时间戳（13位，即ms）为：%d \n", ms)
	//结果：当前时间戳（13位，即ms）为：1572077874913
	fmt.Printf("当前时间戳（16位，即μs）为：%d \n", us)
	//结果：当前时间戳（16位，即μs）为：1572077874913328
	fmt.Printf("当前时间戳（19位，即ns）为：%d \n", ns)
	//结果：当前时间戳（19位，即ns）为：1572077874913328300

	//GMT表示格林威治时间，CST表示四个时区的缩写，也可用来表示中国时区
	fmt.Printf("本地时区: %v\n", t.Location())
	//结果：本地时区: Local

	zoneName, s := t.Zone()
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
	//结果：今天是2019年10月26日

	year = t.Year()
	month = t.Month() //默认为英文
	day = t.Day()
	fmt.Printf("今天是%d年%d月%d日\n", year, int(month), day)
	//结果：今天是2019年10月26日

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
	//结果：现在是16点17分54秒

	hour = t.Hour()
	min = t.Minute()
	sec = t.Second()
	nsec := t.Nanosecond()
	fmt.Printf("现在是%d点%d分%d秒%d纳秒\n", hour, min, sec, nsec)
	//结果：现在是16点17分54秒913328300纳秒

	//今天是今年的第几天
	thisDay := t.YearDay()
	fmt.Printf("今天是今年的第%d天\n", thisDay)
	//结果：今天是今年的第299天

	//日期字符串、时间戳字符串都可以转换成Time对象，日期字符串用time.Parse，时间戳字符串先转成int64再用time.Unix
	//1、时间戳或时间戳字符串转换成日期：
	//先将时间戳字符串转换成Time对象：a := time.Unix(int64(时间戳字符串), 0)，然后再用a.Format("2006-01-02 15:04:05")
	//2、日期或日期字符串转时间戳：
	//先将日期字符串用转成Time对象（日期已经是Time对象）：a := time.Parse("日期模版", 日期字符串)，然后a.Unix()即可

	//指定某一天是一年中第几天
	dateStr := "2019-10-26 00:00:00"
	theDay, _ := time.Parse("2006-01-02 15:04:05", dateStr) //先将日期字符串转换成time对象（实用） 模版：2006-01-02 15:04:05
	fmt.Printf("日期字符串%s转换成时间戳为：%d\n", dateStr, theDay.Unix())
	//结果：日期字符串2019-10-26 00:00:00转换成时间戳为：157204800
	fmt.Printf("%s是%d年的第%d天\n", dateStr, theDay.Year(), theDay.YearDay())
	//结果：2019/10/26是2019年的第299天

	tsStr := "1572079137" //（实用）2019-10-26 16:38:57
	tsInt64, _ := strconv.ParseInt(tsStr, 10, 64)
	theDay = time.Unix(tsInt64, 0) //时间戳转time.Time格式
	//time.Unix(tmp, 0)
	dateStr = theDay.Format("2006/01/02 15:04:05") //模版：2006-01-02 15:04:05
	fmt.Printf("时间戳%s转换成日期字符串为：%s\n", tsStr, dateStr)
	//结果：时间戳1572079137转换成日期字符串为：2019/10/26 16:38:57

	dateStr = theDay.Format("2006/01/02") //模版：2006-01-02
	fmt.Printf("时间戳%s转换成日期字符串为：%s\n", tsStr, dateStr)
	//结果：时间戳1572079137转换成日期字符串为：2019/10/26

	dst := time.Unix(int64(4128226613), 0) //2100-10-26 17:36:53
	isOk := t.Before(dst)                  //相当于if t < dst {}
	fmt.Printf("当前时间%d比目标时间%d要小吗？%v\n", t.Unix(), dst.Unix(), isOk)
	//结果：当前时间1572072745比目标时间4128226613要小吗？true

	dst = time.Unix(int64(1571986325), 0) //2019-10-25 14:52:05
	isOk = t.After(dst)                   //相当于if t > dst {}
	fmt.Printf("当前时间%d比目标时间%d要大吗？%v\n", t.Unix(), dst.Unix(), isOk)
	//结果：当前时间1572072745比目标时间1571986325要大吗？true

	dst = time.Unix(int64(1571986325), 0) //2019-10-25 14:52:05
	isOk = dst.Equal(dst)                 //相当于if dst == dst {}
	fmt.Printf("%d与%d相等吗？%v\n", dst.Unix(), dst.Unix(), isOk)
	//结果：1571986325与1571986325相等吗？true

	//时间运算，支持ns, us (µs), ms, s, m, h, d.
	//减运算
	dst = time.Unix(int64(1571986325), 0) //2019-10-25 0:0:0
	dura, _ := time.ParseDuration("-1ms")
	res := dst.Add(dura)
	fmt.Printf("%d的1毫秒前是：%d\n", dst.Unix(), res.Unix()) //如果时间戳是秒为单位，减掉1ms，小数点将直接舍去，相当于减去了1秒，所以毫秒运算，时间戳得用13位
	//结果：1571986325的1毫秒前是：1571986324
	dura, _ = time.ParseDuration("-1s")
	res = dst.Add(dura)
	fmt.Printf("%d的1秒钟前是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1秒钟前是：1571986324
	dura, _ = time.ParseDuration("-1m")
	res = dst.Add(dura)
	fmt.Printf("%d的1分钟前是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1分钟前是：1571986265
	dura, _ = time.ParseDuration("-1h")
	res = dst.Add(dura)
	fmt.Printf("%d的1小时前是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1小时前是：1571982725
	dura, _ = time.ParseDuration("-24h") //不能用y，d
	res = dst.Add(dura)
	fmt.Printf("%d的1天前是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1天前是：1571899925

	//可以综合使用，如减去1小时1分1秒，上限支持到小时，下限支持到纳秒（ns），支持小数点，支持同时3个单位（如-1h1m1s）
	//不支持年（y）、天（d），如果到天和年这一级别，可以用AddDate方法来运算
	dura, _ = time.ParseDuration("-1h1m1s")
	res = dst.Add(dura)
	fmt.Printf("-1h1m1s换算成秒为：%f\n", dura.Seconds())
	//结果：-1h1m1s换算成秒为：-3661.000000
	fmt.Printf("%d的1小时1分钟1秒钟前是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1小时1分钟1秒钟前是：1571982664

	dura, _ = time.ParseDuration("-1.1h")
	fmt.Printf("-1.1h换算成秒为：%f\n", dura.Seconds())
	//结果：-1.1h换算成秒为：-3960.000000
	res = dst.Add(dura)
	fmt.Printf("%d的1小时1分钟前是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1小时1分钟前是：1571982365

	//加运算
	dura, _ = time.ParseDuration("1h1m1s")
	fmt.Printf("1h1m1s换算成秒为：%f\n", dura.Seconds())
	//结果：1h1m1s换算成秒为：3661.000000
	res = dst.Add(dura)
	fmt.Printf("%d的1小时1分钟1秒钟后是：%d\n", dst.Unix(), res.Unix())
	//结果：1571986325的1小时1分钟1秒钟后是：1571989986

	t_new := dst.AddDate(1, 1, 1)
	fmt.Printf("%s + 1年1月1天 = %s\n", dst.Format("2006-01-02 15:04:05"), t_new.Format("2006-01-02 15:04:05"))
	//结果：2019-10-25 14:52:05 + 1年1月1天 = 2020-11-26 14:52:05

	// 时间比较：指定日期是否在今天范围内
	nowTime := time.Now().Local()
	beginTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)          // 获取一个今天 00:00:00:00的日期
	endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 999999999, time.Local) // 获取一个今天 23:59:59:999999999的日期
	createdAt := time.Date(2023, 6, 7, 12, 0, 0, 0, time.Local)
	isToday := createdAt.After(beginTime) && createdAt.Before(endTime) // 即 > && <
	fmt.Println("isToday=", isToday)
	// 结果：true

	// 时间序列化
	t_byte, _ := t.MarshalJSON()
	fmt.Println("当前时间序列化:", t_byte)
	//结果：当前时间序列化: [34 50 48 49 57 45 49 48 45 50 54 84 49 54 58 51 53 58 49 52 46 57 57 49 56 49 55 52 43 48 56 58 48 48 34]

	// 时间数据反序列化
	var t_un time.Time
	t_un.UnmarshalJSON(t_byte)
	fmt.Println("当前时间反序列化: ", t_un)
	//结果：当前时间反序列化:  2019-10-26 16:39:27.2602463 +0800 CST

	//Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回
	d_second := time.Second
	time.Sleep(d_second)
	fmt.Printf("阻塞时间为: %d ns\n", d_second)

}

//参考学习：https://www.jianshu.com/p/9d5636d34f17
//https://studygolang.com/static/pkgdoc/pkg/time.htm
