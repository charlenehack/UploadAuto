/* ========== map类型 ========== */
func main() {
	// map声明三种方式
	var a map[string]string   // 声明一个map类型
	a = make(map[string]string, 10)  // 使用map前需先make，给map分配数据空间；[]内为key的类型，后面为value的类型；10为最大存放多少键值对，省略则默认分配一个为起始大小
	a["no1"] = "aaa"   // 如果key不存在，则增加；如果存在则修改
	v, ok := a["no1"]   // 如果map存在指定key，则v为指定key的值，ok返回true；若不存在则v为空，ok返回false
	print(v, ok)
	fmt.Println(len(a))  // 返回map有多少键值对

	var b = make(map[int]string)   // 声明一个map直接make
	b[1] = "bbb"
	delete(b, 1)  // 如果key存在则删除该key-value；如果不存在，不操作不报错
	b = make(map[int]string)  // 通过make一个新空间的方式清空map所有key，也可通过遍历逐一删除

	var c map[string]int = map[string]int{"age": 18}  // 声明一个map并直接赋值
	// c := map[string]int{"age": 18}   // 简写

	for k, v := range a {   // map遍历
		fmt.Print(k, v)
	}

	studentMap := make(map[string]map[string]string)  // map嵌套
	studentMap["stu01"] = make(map[string]string)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["gender"] = "男"

	for _, v1 := range studentMap {  // 嵌套map遍历
		for k, v := range v1 {
			print(k, v)
		}
	}

	fmt.Println(a, b, c, studentMap)

	// map切片
	var d []map[string]string
	d = make([]map[string]string, 2)
	d[0] = make(map[string]string, 2)
	d[0]["name"] = "tom"
	d[0]["age"] = "18"
	d[1] = make(map[string]string, 2)
	d[1]["name"] = "mary"
	d[1]["age"] = "16"

	e := map[string]string{
		"name" : "xiaoming",
		"age" : "20",
	}
	d = append(d, e)   // 切片扩容
	fmt.Println(d)

	// map排序
	map1 := map[int]int{10:10, 20:20, 30:30}  // 定义一个map
	var keys []int  // 定义一个切片
	for k, _ := range map1 {
		keys = append(keys, k)  // 循环map的key加入切片
	}
	sort.Ints(keys)   // key排序

	for _, k := range keys {  // 遍历切片，根据key来输出对应的value
		fmt.Printf("map1[%v]=%v\n", k, map1[k])
	}
}
