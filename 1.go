data := url.Values{}
data.Set("name", "哈罗")
data.Set("age", "19")
u, _ := url.ParseRequestURI("http://www.baidu.com")
u.RawQuery = data.Encode()
fmt.Println(u.String())
