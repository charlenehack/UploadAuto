func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		panic(err)
	}

	v, err := cfg.GetValue("default", "test")
	v1, err := cfg.Int("default", "test")
	s, err := cfg.GetSection("default")
	fmt.Printf("%T, %v\n", v, v)
	fmt.Printf("%T, %v\n", v1, v1)
	fmt.Printf("%T, %v", s, s)
}
