/* ========== 文件操作 ========== */
func main() {
	file, err := os.Open("/Users/charlene_lau/Desktop/11.py")  // 返回一个文件指针
	if err != nil {
		fmt.Println("打开文件失败，", err)
	}

	fmt.Printf("%v", file)

	defer file.Close()   // 函数退出时关闭文件句柄

	reader := bufio.NewReader(file)  // 带缓冲区读取，默认缓冲区为4096
	for {
		str, err := reader.ReadString('\n')  // 逐行读取，每读到一个换行就结束
		if err == io.EOF {   // io.EOF表示文件的末尾
			break   // 读到末尾退出循环
		}
		fmt.Print(str)
	}
}

func main() {   // 创建文件并写入内容
	filePath := "/Users/charlene_lau/Desktop/1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	/*
	    O_RDONLY  // 只读模式打开文件
	    O_WRONLY  // 只写模式打开文件
	    O_RDWR    // 读写模式打开文件
	    O_APPEND  // 写操作时将数据附加到文件尾部
	    O_CREATE  // 如果不存在将创建一个新文件
	    O_EXCL    // 和O_CREATE配合使用，文件必须不存在
	    O_SYNC    // 打开文件用于同步I/O
	    O_TRUNC   // 如果可能，打开时清空文件
	 */
	if err != nil {
		fmt.Println("打开文件错误，", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)  // 带缓存写入
	writer.WriteString("hello, text")  // 内容先入写缓存
	writer.Flush()  // 将缓存内容写入到文件
}

func main() {   //  读写文件的另一种方式，适合小文件
	content, err := ioutil.ReadFile("/Users/charlene_lau/Desktop/11.py")
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	fmt.Println(string(content))

	err = ioutil.WriteFile("/Users/charlene_lau/Desktop/11.txt", content, 0666)
	if err != nil {
		fmt.Println("写入文件错误", err)
	}
}

func PathExists(path string) (bool, error)  {   // 判断文件或目录是否存在的函数
	_, err := os.Stat(path)
	if err == nil {    // 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {   // 文件拷贝函数
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误，", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	writer.Flush()

	return io.Copy(writer, reader)
}
