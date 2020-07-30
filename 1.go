/* ========== json序列化与反序列化 ========== */
func main() {
	str := "{\"name\": \"Tom\", \"age\": 18, \"brithday\": \"2008-08-08\", \"address\": \"beijing\"}"
	str1 := `{"name": "Tom", "age": 18, "brithday": "2008-08-08", "address": "beijing"}`  // 可以用反引号省去转义符
	var info map[string]interface{}   // 反序列化时map可以不用先make，因为make操作已被封装到Unmarshal函数中
	err := json.Unmarshal([]byte(str), &info)
	if err != nil {
		fmt.Println("反序列化失败。", err)
	}

	strJson, err := json.Marshal(info)   //  序列化
	fmt.Println(info, string(strJson))
}
