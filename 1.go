/* ========== 单元测试 ========== */
// 要测试的函数必须与测试用例处于同一目录内，测试用例文件名必须以_test.go结尾
func TestAddUpper(t *testing.T)  {   // 函数名必须为TestXxx格式
	res := addUpper(10)   // 要测试的函数
	if res != 55 {
		t.Fatalf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	t.Logf("AddUpper(10)执行正确。")
}

# go test -v  // 在目录内执行命令，调用测试用例，所有以Test开头的函数都被执行；-v 运行正确或错误都输出日志，不加-v只输出错误日志
# go test -v cal_test.go cal.go  // 测试单个文件时，带上被测试的原文件
# go test -v -test.run TestAddUpper  // 测试单个方法
