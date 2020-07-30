/* ========== 协程、管道 ========== */
/* 如果一个管道只有写，而没有读，当管道容量写满时就会出现阻塞而deadlock；
   写管道和读管道的频率不一致无所谓，只要有在读。
 */
cpuNum := runtime.NumCPU()
fmt.Println("CPU个数：", cpuNum)
// runtime.GOMAXPROCS(cpuNum - 1)   // 设置程序使用CPU个数，go1.8之后默认运行在多核上，不用设置了

var intChan chan int   // 声明一个管道，默认可读写；var intChan chan<- int声明只写管道； var intChan ->chan int声明只读管道（只读管道不可close）
intChan = make(chan int, 10)  // channel使用前必须先make，channel为引用类型；int为可以存放的数据类型，10为管道容量（写入数据不可超过管道容量）
fmt.Printf("intChan的值为%v，intChan本身的地址为%p\n", intChan, &intChan)
intChan <- 10   // 向管道写入数据
intChan <- 20   // 向管道写入数据

fmt.Printf("intChan的长度为%v，intChan的容量为%v\n", len(intChan), cap(intChan))
num := <- intChan  // 从管道取数据，先进先出
<- intChan  // 从管道取数据扔掉不接收，管道没有数据了再取就会报deadlock
fmt.Printf("num = %v\n", num)
close(intChan)  // 关闭管道，一旦关闭只能读不能写

for v := range intChan {  // 在遍历管道时，如果管道没有关闭，遍历结束会报deadlock错误；如果已经关闭则正常遍历完
	fmt.Println(v)
}

for {
	select {   // 循环从管道取数据时，为防止管道没有关闭一直阻塞而报deadlock，可使用select方式读取数据
		case v := <- chan1 :   // 依次从管道读数据
			fmt.Printf("从chan1读取到的数据%v\n", v)
		case v := <- chan2 :  // 上一管道无数据可取自动到下一个case
			fmt.Printf("从chan2读取到的数据%v\n", v)
		default :
			fmt.Printf("管道已无数据可取")
			return
	}
}

// 实例1
func writeData(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("写入数据：%v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool)  {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("获取数据：%v\n", v)
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	<- exitChan  // 阻塞直到能取到数据，因为编译器检测到有协程会往管道里面写数据，所以取不到数据会等待不会报错
}

// 实例2：使用waitGroup优雅退出协程

// 如果定义的不是全局的wait group,则在传值的时候需要传指针类型
func sendData(ch chan string, waitGroup *sync.WaitGroup) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"

	close(ch)  	// 关闭chan,即使使用了 waitGroup也需要关闭channel
	fmt.Printf("send data exited")
	waitGroup.Done()  // 使用 waitGroup给出goroutine的结束信号，调用Done()方法计数器-1
}

//
func getData(ch chan string, waitGroup *sync.WaitGroup) {
	for {
		input, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("getData中的input值: %s\n", input)
	}
	fmt.Printf("get data exited\n")
	waitGroup.Done()
}

func getData2(ch chan string, waitGroup *sync.WaitGroup) {
	for {
		input2, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("getData2中的input值:%s\n", input2)
	}
	fmt.Printf("get data2 exited\n")
	waitGroup.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)  	//一个没有缓冲区的管道
	wg.Add(3)   // 开3个协程，计数器+3
	go sendData(ch, &wg)
	go getData(ch, &wg)
	go getData2(ch, &wg)

	// waitGroup只是用于主线程等待所有的goroutine都执行完成后才关闭主线程
	time.Sleep(time.Second*3)
	wg.Wait()   // 等待计数器=0，说明没有任务了可以优雅退出协程
	fmt.Printf("main goroutine exited\n")
}
