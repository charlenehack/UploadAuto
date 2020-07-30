/* ========== flag包获取命令行参数 ========== */
func main() {  // 通过flag包获取命令行参数，用法：test.ext -u root -p 123 -P3306
	var user, password string
	var port int
	flag.StringVar(&user, "u", "", "说明：用户名，默认为空")
	flag.StringVar(&password, "p", "", "密码，默认为空")
	flag.IntVar(&port, "P", 3306, "端口号，默认3306"

	flag.Parse()

	for i, v := range os.Args {  // 通过os包获取命令行参数
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
