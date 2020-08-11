import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Header)
	// 响应字符串
	fmt.Fprintln(w, "test")
	w.Write([]byte("OK"))
	// 响应json格式
	w.Header().Set("Content-Type", "application/json")  // 设置响应头中内容的类型
	m := map[string]string{"username": "tom", "password": "123456"}
	j, _ := json.Marshal(m)
	w.Write(j)
	// 让客户端重定向
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func Template(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("index.html")   // 解析模块文件
	// t := template.Must(template.ParseFiles("index.html"))  // Must方法自动处理异常，不需在接收
	t.Execute(w, "hello")   // 返回的内容在html中用{{.}}接收
	t1, _ := template.ParseFiles("1.html", "2.html")
	t1.ExecuteTemplate(w, "2.html", "我要在2.html中显示。")  // 如果传了多个模块文件，执行时需要指定模块文件，默认执行第1个
}

func main() {
	http.HandleFunc("/test", handler)
	http.HandleFunc("/template", Template)
	http.ListenAndServe(":8080", nil)
}
