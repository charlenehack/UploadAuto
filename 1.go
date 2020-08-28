dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:8080",
    &proxy.Auth{User:"username", Password:"password"},
    &net.Dialer {
        Timeout: 30 * time.Second,
        KeepAlive: 30 * time.Second,
    },
)
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1：8080", nil, proxy.Direct)

transport := &http.Transport{
    Proxy: nil,
    Dial: dialer.Dial,
    TLSHandshakeTimeout: 10 * time.Second,
}
client := &http.Client { Transport: transport }
response, err := client.Get("http://mengqi.info")

req, err := http.NewRequest("GET", "http://myip.top", nil)
req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
resp, err := client.Do(req)
