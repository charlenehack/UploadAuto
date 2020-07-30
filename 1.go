/* ========== 指针 ========== */
func main() {
	i := 10
	fmt.Printf("i的值为%d, i的内存地址为：%v\n", i, &i)

	var ptr *int = &i  // 声明一个指针变量，类型为*int，此时ptr的值为&i，指针的数据类型必须与对应值的数据类型一致；指针变量存放的是一个地址，这个地址指向的空间存放的才是值
	fmt.Println("ptr的值为：", *ptr)   // 指针变量加上*获取内存空间对应的值

	*ptr = 11   // 通过内存地址来修改变量的值
	fmt.Printf("i的值为：%d, i的内存地址为：%v", i, &i)

	a := 10
	b := 20
	a, b = b, a
	fmt.Printf("a = %d, b = %d", a, b)
}

