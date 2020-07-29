/* ========== 基本数据类型例子 ========== */

func main() {
	var a int32 = 87
	var b float32 = float32(a)   //  数据类型之间转换
	var str, str1, str2, str3, str4 string
	var c bool = true
	var d bool
	var e int64
	const f int = 10 // 常量声明必须赋值，数据类型可省略会自动类型推导，常量不能修改
	const (
		g = iota   // 表示给g赋0
		h  // 上一个的基础上+1
		i  // i,j = iota, iota则i,j都为2，+1以行为基础
	)

	str = fmt.Sprintf("%d", a)  // 整型转字符串
	str1 = strconv.FormatInt(int64(a), 10)  // 使用strconv包进行数据类型转换
	str2 = strconv.Itoa(int(a))  	// 数字转字符串另一个函数，参数只接受int型，非int型时需转换
	str3 = strconv.FormatBool(c)
	str4 = strconv.FormatFloat(float64(b), 'f', 5, 32)  // 接收参数须为float64，f 格式，10 精度表示保留几位小数点，32 表示f的数据来源（32/64）
	d, _ = strconv.ParseBool(str3)   // 字符串转布尔型，返回两个值，不关心第二个值使用下划线忽略
	e, _ = strconv.ParseInt(str, 10, 64)  // 字符串转数字，返回为int64；当给的字符串无法转为有效的数字时会直接将其转为0

	fmt.Printf("a = %d, a type is %T\n", e, d)
	fmt.Printf("b = %f, b type is %T\n", b, b)
	fmt.Printf("str = %q, str type is %T\n", str, str)  // %s 输出字符串，%q 输出字符串带上双引号
	fmt.Printf("str1 = %s, str type is %T\n", str4, str3)
	fmt.Printf("str2 = %q, str type is %T\n", str2, str1)
}
