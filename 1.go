/* ========== 数组、切片 ========== */
// 数组内的元素数据类型必须相同；长度是数组类型的一部分，必须固定不能动态变化；数组创建后如何没有赋值有默认值
func test(arr *[3]int)  {  // 此时arr并不是数组类型而是指针
	(*arr)[0] = 10  // 通过指针获取到数组内存地址，然后进行赋值
}

func main() {
	// 定义数组的几种方式
	var numArr01 [3]int = [3]int{1, 3, 5}
	var numArr02 = [3]int{2, 4, 6}
	var numArr03 = [...]int{7, 8, 9}
	var numArr04 = [...]int{1: 5, 0: 6, 2:7}  // 通过下标指定顺序
	strArr05 := [...]string{"tom", "jerry", "kitty", "jack", "mary"}
	fmt.Println(numArr01, numArr02, numArr03, numArr04, strArr05)
	// 定义切片
	slice1 := strArr05[1:4]  // 定义一个切片，然后让切片去引用一个已经创建好的数组
							// 起始和结束下标可省略，如引用整个数组可简写为 strArr05[:]

	var slice2 []float64 = make([]float64, 5, 10)  // 使用make定义切片，三个参数分别为数据类型、长度、容量；其中容量为可选，但如果指定则容量必须大于长度
	slice2[3] = 10

	b, _ := json.Marshal(numArr02)
	strNumArr02 = string(b)    // 将数组或切片转为字符串

	var slice3 []string = []string{"tom", "jack", "mary"}  // 定义一个切片直接指定具体数组，原理类似make；但没指定容量，此时容量等于长度
	slice3 = append(slice3, "jim", "tim", "jerry")  // 对切片进行动态追加元素
	slice3 = append(slice3, slice1...)   // 切片追加切片
	slice4 := make([]string, 10)
	copy(slice4, slice3)   // 切片拷贝，前面为目标，后面为源；数据类型必须一致，并且都为切片；当源的长度超过目标容量时，丢弃超过部分元素

	fmt.Printf("slice1的长度为：%d， slice1的容量为：%d\n", len(slice1), cap(slice1))
	fmt.Printf("slice2的长度为：%d， slice2的容量为：%d\n", len(slice2), cap(slice2))
	fmt.Printf("slice3的长度为：%d， slice3的容量为：%d\n", len(slice3), cap(slice3))
	fmt.Printf("slice4的长度为：%d， slice4的容量为：%d\n", len(slice4), cap(slice4))

	/* slice1和slice2的区别：
		slice1创建的切片是直接引用数组，这个数组是事先存在的；slice2通过make创建切片，也会创建一个数据由切片在底层进行维护 */
	// 数组、切片遍历
	for index, value := range strArr05 {   // index 下标，不需要可使用下划线占位符忽略；value 对应值
		print(index, value)
	}

	test(&numArr01)   // 通过指针修改数组的值
	fmt.Println(numArr01)

	var arr [4][6]int   // 声明一个二维数组
	arr[1][2] = 3  // 二维数组赋值

	var arr1 = [2][3]int{{1,2,3}, {4,5,6}}  // 声明二维数组并赋值，[2]可用[...]自行推导，[3]必须指定

	for i :=0; i < len(arr); i++ {  // 遍历二维数组1
		for j:= 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println("\n")
	}

	for i, v := range arr1 {  // 遍历二维数组2
		for j, v1 := range v {
			fmt.Print(i, j, v1, "\n")
		}
	}
	fmt.Println(arr1)
}
