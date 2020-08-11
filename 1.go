func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp.Body failed, err: %v\n", err)
		return
	}
	fmt.Println(string(body))
}
