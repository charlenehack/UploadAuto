func main() {
	url := "http://127.0.0.1:9090/post"
	// json格式数据
	contentType := "application/json"
	data := `{"name": "哈罗", "age": "18"}`
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=哈罗&age=18"
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get response failed, err: %v\n", err)
		return
	}
	fmt.Println(string(body))
}
