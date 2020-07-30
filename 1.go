/* ========== 字符串处理常用函数 ========== */
func main() {
	str := "hello, 我爱北京天安门"
	fmt.Printf("字符串长度为：%d\n", len(str))   // 中文占3个字节
	str1 := []rune(str)   // 将字符串转为切片类型
	fmt.Printf("字符串类型：%T\n", str1)
	n, err := strconv.Atoi("123")  // 字符串转整型，转换不成功时err接收错误
	str = strconv.Itoa(12345)  // 整数转字符串
	bytes := []byte("hello")   // 字符串转[]byte
	str = string([]byte{97, 98, 99})  // byte转字符串
	str = strconv.FormatInt(123, 2)  // 十进制转2，8，16进制，第一个参数十进制，第二个参数要转成的进制
	strings.Contains("hello,world", "llo")  // 查找子串是否在指定的字符串中，存在返回true，不存在返回false
	strings.Count("hello,world", "l")  // 统计字符串中有几个指定的子串
	strings.EqualFold("abc", "ABC")  // 不区分大小写比较字符串，返回true / false；（ ==是区分大小写的）
	strings.Index("comepro.com", "com")  // 返回子串在字符串第一次出现的index值，没有没有返回-1；也可用来查找字符串中是否存在子串
	strings.LastIndex("comepro.com", "com") // 返回子串在字符串中最后一次出现的位置，没有返回-1
	strings.Replace("hello,world", "h", "H", 1)  // 将字符串中指定子串替换成另一子串，n 指定替换几个，-1全部替换
	strings.Split("hello,world", ",")  // 按指定字符将字符串拆分成字符串数组
	strings.ToLower("GO")  // 大写转小写
	strings.ToUpper("go")  // 小写转大写
	strings.TrimSpace("  ABC  ")  // 去除字符串左右两边的空格，TrimSpaceLeft 去左 / TrimSpaceRight 去右
	strings.Trim("hello!" , "!")  // 去除字符串左右两边的指定字符
	strings.HasPrefix("Hello", "H")  // 判断字符串是否以指定字符串开头
	strings.HasSuffix("hello", "o")  // 判断字符串是否以指定字符串结尾
	fmt.Println(n, err, str, bytes)
	// 字符串拼接效率最高方式
	var build strings.Builder
	build.WriteString("aaaaaaa")
	build.WriteString("bbbbbbb")
	ab := build.String()
	fmt.Println(ab)
}
