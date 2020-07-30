/* ========== 错误处理异常捕获 ========== */
func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error=", err)
		}
	}()
	a, b := 10, 0
	res := a / b
	print(res)
}

func open(name string) (err error)  {
	if name == "1.txt" {
		return nil
	} else {
		return errors.New("打开文件错误。")   // 返回一个error类型的自定义错误
	}
}

func test1() {
	err := open("2.txt")
	if err != nil {
		panic(err)   // panic捕获到错误后立即终止程序，并打印错误信息
	}
}

func main() {
	test()
	test1()
	print("aaaa")
}
