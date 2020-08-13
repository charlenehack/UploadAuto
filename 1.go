resp, err := http.Get("https://www.baidu.com")
doc,e := goquery.NewDocumentFromReader(resp.Body)


doc,e := goquery.NewDocumentFromReader(strings.NewReader(html))
doc,e := goquery.NewDocument(url string)
2、查找内容

ele.Find("#title") //根据id查找
ele.Find(".title") //根据class查找
ele.Find("h2").Find("a") //链式调用
3、获取内容

ele.Html()
ele.Text()
4、获取属性

ele.Attr("href")
ele.AttrOr("href", "")
5、遍历

ele.Find(".item").Each(func(index int, ele *goquery.Selection){
   
})
