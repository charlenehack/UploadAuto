// 断言实例：编写一个函数，可以判断输入的参数是什么类型
func Type(items... interface{})  {
	for i, x := range items {
		switch x.(type) {
			case bool:
				fmt.Printf("第%v个参数是bool类型，值是%v\n", i, x)
			case float32:
				fmt.Printf("第%v个参数是float32类型，值是%v\n", i, x)
			case float64:
				fmt.Printf("第%v个参数是float64类型，值是%v\n", i, x)
			case int, int32, int64:
				fmt.Printf("第%v个参数是整型，值是%v\n", i, x)
			case string:
				fmt.Printf("第%v个参数是字符串类型，值是%v\n", i, x)
			default:
				fmt.Printf("第%v个参数类型不确定，值是%v\n", i, x)
		}
	}
}

func main() {
	n1, n2, n3, n4, n5, n6 := 1.1, 2.2, 45, true, "tom", 300
	Type(n1, n2, n3, n4, n5, n6)
}
