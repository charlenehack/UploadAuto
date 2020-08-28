urli := url.URL{}
urlproxy, _ := urli.Parse("https://127.0.0.1:8080")
client := http.Client{
    Transport: &http.Transport{
        Proxy: http.ProxyURL(urlproxy),  // 进行类型转换
        Dial: (&net.Dialer{Timeout: 10 * time.Second,}).Dial,   // 不配置超时时间无需
    },
}
resp, err := client.Get("http://myip.top")

req, err := http.NewRequest("GET", "http://myip.top", nil)
req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
resp, err := client.Do(req)
if err != nil {
	fmt.Printf("get failed, err: %v\n", err)
	return
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
