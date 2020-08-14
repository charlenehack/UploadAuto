func main() {
	match, _ = regexp.MatchString("H(.*)d!", "Hello, world!")  // 返回true或false
	match, _ = regexp.Match("H(.*)d!", []byte("Hello, world!"))    // 返回true或false
	
	r, _ := regexp.Compile("H(.*)d!")
	reg := regexp.MustCompile("H(.*)d!")  // 与Compile区别在于少一个返回值err


	match = reg.MatchString("Hello, world!")  // 返回true或false
	fmt.Println(reg.FindString("Hello, world! haha"))  // 返回第一个匹配到的子串， Hello, world！
	fmt.Println(reg.FindAllString("Hello, world! haha", -1))  // 返回所有匹配到的子串，n表示匹配数量，[Hello, world!]
	fmt.Println(reg.FindStringIndex("Hello, world! haha"))  // 返回匹配到子串的起止索引， [0 13]
	fmt.Println(reg.FindStringSubmatch("Hello, world! haha"))  // 返回第一个匹配到的子串"H(.*)d!"和局部子串"(.*)"， [Hello, world! ello, worl]
	fmt.Println(reg.FindAllStringSubmatch("Hello, world! haha", -1)) // 返回所有匹配到的子串和局部子串， [[Hello, world! ello, worl]]
	fmt.Println(reg.ReplaceAllString("Hello world! haha", "html"))  // 将匹配到的结果替换，html haha
}
