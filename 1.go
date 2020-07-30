/* ========== 结构体类型 ========== */
type books struct {   // 声明一个结构体
	title string `json:"title"`   // `json/"title"`为结构体tag，一般用于返回json时首字母大写的问题
	author string
	subject string
	book_id int
}

func NewBooks(t, a, s string, b int) *books {   // 工厂模式，当结构体首字母为小写无法被其它包调用时
	return &books{
		title: t,
		author: a,
		subject: s,
		book_id: b,
	}
}

func main() {
	var book1 books

	book1.title = "aaa"
	book1.author = "liu.a"
	book1.subject = "bbb"
	book1.book_id = 3476

	fmt.Println(books{"eee", "liu", "ggg", 28998})  //创建一个新的结构体，值与声明字段一一对应
	fmt.Println(books{title: "hhh", author: "liu", subject: "iii", book_id: 2878})  // 使用key:value格式创建，此方式若声明字段没有赋值则为0或空
	fmt.Printf("book1 title: %s\n", book1.title)
	fmt.Printf("book1 author: %s\n", book1.author)
	fmt.Printf("book1 subject: %s\n", book1.subject)
	fmt.Printf("book1 book_id: %d\n", book1.book_id)

	book2 := new(books)   // 通过指针方式给字段赋值，也可以写为 book2 := &books{} （可在{}内直接赋值，{"lll", "liu:, "mmm", 5454}）
	(*book2).title = "jjj"
	(*book2).author = "kkk"
	book2.book_id = 2878  // 简写

	var book3 = &books{"lll", "mmm", "nnn", 7828}
	book4 := &books{title: "lll", author: "mmm", subject: "nnn", book_id: 1298}

	book5 := NewBooks("aaa", "bbb", "ccc", 9992)   // 调用工厂模式，用指针取值

	fmt.Println(book2.book_id, book3, book4, *book5)

	printBook(book1)  // 调用函数，使用指针变量访问结构体时传入&books

	book1.test()   // 调用结构体方法
	book1.test1()  // 此时book1为(&book1)的简写
	fmt.Println(&book1)  // 默认调用book1的String()方法，相当于book1.String()的简写
}

func printBook(book books)  {   // 将结构体作为参数传给函数， 使用指针变量访问结构体 *books
	fmt.Printf("book1 title: %s\n", book.title)
	fmt.Printf("book1 author: %s\n", book.author)
	fmt.Printf("book1 subject: %s\n", book.subject)
	fmt.Printf("book1 book_id: %d\n", book.book_id)
}

func (b books) test()  {   // 给结构体绑定一个方法
	fmt.Println(b.title)
}

func (c *books) test1()  {  // 结构体的指针和方法绑定，传内存地址提高效率
	fmt.Println(c.title)  // 此时的c为(*c)的简写
}

func (d *books) String()  string{  // 如果一个类型实现了String()这个方法，那么fmt.Println默认就会调用这个变量的String()进行输出
	str := fmt.Sprintf("title=[%v], subject=[%v]", d.title, d.subject)
	return str
}

/* 方法和函数的区别：
	对于普通函数，接收者为值类型时，不能将指针类型数据直接传递；接收者为指针类型时，不能直接传值类型；
	对于方法，接收者为值类型时，可以直接传指针类型数据；接收者为指针类型时，也可以传值类型；实际为值类型还是指针类型看绑定方法的接收者决定。
 */
