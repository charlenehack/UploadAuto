/* ========== 函数、自定义数据类型 ========== */
func getSum(n1 int, n2 int) (sum int) {  // 直接为函数返回值命名，return后没必须再接返回值，好处是多个返回值时不用关心返回值位置顺序
	sum = n1 + n2
	return
}

func sum(n1 int, args... int) int {  // 函数可变参数传参，args关键字可自定义
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

type myint int  // 给int类型取别名，此时虽然myint和int都是int类型，但go依然认为是两个类型
type myFuncType func(int, int) int  // 自定义函数数据类型，myFuncType为传入两个int类型，返回一个int类型的函数

func myFun(funvar myFuncType, num1 int, num2 int) int {   // 调用自定义函数类型，funvar为声明函数变量
	return funvar(num1, num2)
}

func init() {  // init函数主要用来完成一些初始化的工作；执行流程：全局变量定义-->init函数-->main函数
	fmt.Println("我是init函数。")
}

func Sum(n1 int, n2 int) int {
	defer fmt.Println("我第3个执行！")   // 当执行到defer时，暂时不执行，先将defer后面的语句压入到独立栈中
	defer fmt.Println("我第2个执行！")   // 当函数执行完毕后，再从栈中按照先入后出的方式出栈，执行

	res := n1 + n2
	fmt.Println("我第1个执行！")
	return res
}
var (add = func(n1 int, n2 int) int {   // 定义一个全局匿名函数并赋值给一个变量
	return n1 + n2
})

func main() {
	res := myFun(getSum, 20, 30)
	println(res)
	res1 := sum(1, 3,5,8,9)
	println(res1)

	res3 := func(n1 int, n2 int) int {   // 匿名函数的使用
		return n1 + n2
	}(10, 20)  // 传参
	println(res3)

	a := func(n1 int, n2 int) int {   // 将匿名函数赋值给一个变量
		return n1 - n2
	}
	res4 := a(30, 20)  // 通过变量名调用匿名函数
	fmt.Printf("a type is %T\n", a)   // 此时a为一个函数数据类型
	println(res4)

	res5 := add(30, 20)   // 调用全局匿名函数
	println(res5)
}
