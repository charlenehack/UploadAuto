/* ========== 循环 ========== */
func main() {
	i := 0
	for ; i <= 10; i++ {
		fmt.Println("hello", i)
	}

	j := 1
	for j <= 10 {
		fmt.Println("hello", j)
		j++
	}

	for { // 等价于 for ; ; {
		fmt.Println("我是死循环，请用break帮我退出")
		break
	}

	var str string = "hello, world!"
	//	str1 = []rune(str)  // 如果字符串中存在中文会乱码，需先转为切片类型
	for i := 0; i < len(str); i++ { // 字符串遍历一
		fmt.Printf("%c\n", str[i])
	}
	for index, val := range str { // 字符串遍历二，中文不会乱码，一个汉字占3字节
		fmt.Printf("index = %d, val = %c", index, val)
	}

	var count int = 0
	rand.Seed(time.Now().UnixNano()) // 生成随机数需要给rand设置一个种子，UnixNano表示刷新频率为纳秒级，可根据需求设置频率
	for {
		a := rand.Intn(100) // 生成随机数
		count++
		fmt.Printf("生成的随机数为：%d, 生成次数：%d\n", a, count)
		if a == 50 {
			break // break默认中止最近的循环，可给循环打上标签，通过break指定标签中止标签对应的循环
		}
	}

	println("我是label1")
	if count > 0 {
		goto label3 // 跳到指定标签处执行
	}
	println("我是label2")
	label3:
	println("我是label3")
}
