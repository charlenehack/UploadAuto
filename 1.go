/* ========== 接口 ========== */
// interface体现了程序的多态和高内聚低偶合的思想
type Usb interface {
	Start() int  // 定义一组方法，但不需要实现；可以有返回值
	Stop()  // 不能包含任何变量
}

type A interface {
	test1()
}

type B interface {  // 一个接口继承多个接口时，需要将各个接口的方法都实现；继承的多个接口不能有相同的方法
	Usb
	A
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

func (p Phone) Start() int{
	fmt.Println("手机开始工作。")
	return 8
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作。")
}

func (p Phone) Call()  {
	fmt.Printf("%v手机正在打电话。", p.Name)
}

func (c Camera) Start() int {
	fmt.Println("相机开始工作。")
	return 88
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。")
}

func (c Camera) test1() {
	fmt.Println("相机正在测试。")
}
type Computer struct {

}

func (c Computer) Working(usb Usb)  {  // 接收一个接口类型变量，只要实现了接口，就实现了接口里面声明的所有方法
	usb.Start()   // 通过接口变量调用方法
	usb.Stop()
	phone, ok := usb.(Phone)  // 使用类型断言调用Phone结构体的特有方法Call()
	if ok {
		phone.Call()
	}
}

func main() {
	computer := Computer{}  // 创建结构体变量
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)  // 将结构体变量也接口变量传入方法
	computer.Working(camera)

	var phone1 Phone // Phone实现了Usb接口
	var p1 Usb = phone1  // 因此可将phone1赋值给Usb接口类型
	p1.Start()

	var camera1 Camera
	var c1 B = camera1
	c1.test1()

	var t interface{} = 8   // 可以把任何一个变量赋给空接口；空接口没有任何方法，即所有类型都实现了空接口
	fmt.Println(t)

	var usbArr [3]Usb  // 定义一个Usb接口数组，存放Phone和Camera的结构体变量，体现出多态数组
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Camera{"佳能"}
	for _, v := range usbArr {   // 遍历数组将结构体变量传入接口类型
		computer.Working(v)
	}
}
