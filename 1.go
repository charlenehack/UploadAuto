/* ========== 反射 ========== */
// 完成基本数据类型、reflect.value类型、interface{}类型之间的转换
func reflectTest1(b interface{})  {
	rType := reflect.TypeOf(b)   // 获取到传入变量的type
	fmt.Printf("rType = %v\n", rType)
	rVal := reflect.ValueOf(b)   // 获取到传入变量的值，类型为reflect.value
	a := rVal.Int()  // 通过调用reflect.value的方法去得到值
	k1 := rVal.Kind()  // 获取传入变量的类别，	另一种方式 k2 := rType.Kind()
	fmt.Printf("rVal = %v 类型为%T，a = %v 类型为%T\n", rVal, rVal, a, a)
	iV := rVal.Interface()  // 将reflect.value类型转为接口类型
	c := iV.(int)  // 通过类型断言将接口类型转为需要的类型
	fmt.Println("b = ", c, k1)
}

// 结构体的反射
func reflectTest2(b interface{})  {
	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iV的值为%v，iV类型为%T\n", iV, iV)
	stu, ok := iV.(Student)  // 如果有多个结构体可使用switch..case来断言多个结构体类型
	if ok {
		fmt.Println(stu.Name)
	}
}

// 通过反射修改传入变量的值
func reflect01(b interface{})  {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(200)  // Elem()返回指针指向地址具体的值，SetXxx设置值
}

type Student struct {
	Name string
	Age int
}

func main() {
	num := 100
	reflectTest1(num)
	reflect01(&num)  // 传入地址修改值
	fmt.Println(num)

	str := Student{"Tom", 19}
	reflectTest2(str)
}
