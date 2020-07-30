
/* ========== switch..case ========== */
func main() {
	var key, c byte
	fmt.Println("请输入一个字符 a, b, c, d")
	fmt.Scanf("%c", &key)

	switch key {   // switch后也可不带表达式，当不带表达式时用case后的表达作为条件进行判断，类似if else分支来使用
					// 变量声明也可以直接写在switch后面，用分号结束；如 switch key := 10 {
		case 'a':    // case后各个表达式的数据类型必须和switch后表达式数据类型一致，多个表达式逗号分隔
			fmt.Println("A")
		case 'b':
			fmt.Println("B")
			fallthrough  // switch穿透，带此关键字不跳出继续执行下一个case
		case c + 1:
			fmt.Println("D")
		default:  // default语句非必须
			fmt.Println("NO")
	}
}
