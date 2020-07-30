/* ========== 闭包 ========== */
// 闭包是一个函数和与其相关的引用环境组合的一个整体
func AddUpper() func(int) int {  // 定义一个AddUpper函数，返回的数据类型是func(int) int
	n := 10
	return func(x int) int {  // 返回一个匿名函数，但这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {  // 闭包使用实例，传入一个文件名，如果是指定后缀则返回文件名，否则加上指定后缀
	return func(name string) string {   // 返回的匿名函数与外部变量suffix形成一个闭包，因为返回的函数引用到suffix这个变量
		if !strings.HasSuffix(name, suffix) {   // HasSuffix方法判断字符串是否以..结尾
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(2))
	fmt.Println(f(5))  // 当反复调用f函数时，因为n是初始化一次，因此每调用一次值就会改变而不是重新初始化
	fmt.Println(f(7))
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("aaa"))
}
