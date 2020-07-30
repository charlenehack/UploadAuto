/* ======== 时间和日期相关函数 ========== */
func main() {
	now := time.Now()  // 获取当前时间
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())  // int(now.Month())可将月转为数字形式
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())
	// 格式化日期时间赋值给变量
	dateStr := fmt.Sprintf("当前时间：%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)
	fmt.Printf(now.Format("2006-01-02 15:04:05"))  // 格式化日期时间的另一种方式，分隔符可自定义，各个数字是固定的不能改
	fmt.Printf(now.Format("2006-01-02"))  // 可以根据数字按需取出想要的字段，如"01"只取出月份
	fmt.Printf(now.Format("15:04:05"))
	time.Sleep(time.Second)   // 休眠一秒
	time.Sleep(time.Millisecond * 100)  // 休眠一毫秒
	fmt.Printf("Unix时间戳：%v，Unixnano时间戳：%v", now.Unix(), now.UnixNano())
	// 时间加减
	m, _ := time.ParseDuration("-1m")  	// 10分钟前
	m1 := now.Add(m)

	h, _ := time.ParseDuration("-1h")  	// 8个小时前
	h1 := now.Add(8 * h)

	d, _ := time.ParseDuration("-24h")   // 一天前
	d1 := now.Add(d)


	mm, _ := time.ParseDuration("1m")    // 10分钟后
	mm1 := now.Add(mm)

	hh, _ := time.ParseDuration("1h")    // 8小时后
	hh1 := now.Add(hh)

	dd, _ := time.ParseDuration("24h")  	// 一天后
	dd1 := now.Add(dd)

	subM := now.Sub(m1)  	// Sub 计算两个时间差
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)
}
