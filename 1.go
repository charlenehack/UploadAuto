/* ========== 继承 ========== */
type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("name=%v，age=%v，score=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student  // 嵌入匿名结构体，以继承结构体属性及方法
	int // 基本数据类型也可作为匿名结构体嵌入
}


func (p *Pupil) testing()  {
	fmt.Println("小学生正在考试。。")
}

type Graduate struct {
	s Student  // 嵌入有名结构体
}

type Middle struct {
	*Student  // 以指针方式嵌入匿名结构体
}

func (p *Graduate) testing()  {
	fmt.Println("大学生正在考试。。")
}

func main()  {
	pupil1 := &Pupil{}  // 指针类型，值类型: var pupil1 Pupil
	pupil1.Student.Name = "tom"
	pupil1.Age = 10  // 嵌入的匿名结构体字段访问可以简化，pupil.Student.Age的简写
	pupil1.testing()
	pupil1.Student.SetScore(100)   // 当结构体和匿名结构体有相同字段或方法时，采用就近访问原则，即字段名或方法名.前面的结构体
	pupil1.ShowInfo() // 当嵌入两个或多个匿名结构体，并且匿名结构体中有相同字段或方法时，必须明确指定匿名结构体名字，不可简写

	pupil2 := Pupil{Student{"jim", 10, 80}}  // 嵌入匿名结构体后在创建结构体变量时直接赋值，多继承时依次写用逗号分隔
	middle1 := Middle{&Student{Name: "jerry", Score: 90, Age: 28}}  // 指针方式匿名结构体赋值
	fmt.Println(pupil2.Name, pupil2.Age, pupil2.Score)
	fmt.Println(*middle1.Student, middle1.Name, middle1.Score)  // 指针方式取值

	var graduate1 Graduate
	graduate1.s.Name = "mary"  // 访问有名结构体字段或方法时，必须带上有名结构体名字，不可简写
}
