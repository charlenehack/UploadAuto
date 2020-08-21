A := []string{"111", "222", "333", "444", "555", "666"}
	
rand.Seed(time.Now().UnixNano())
fmt.Println(A[rand.Intn(len(A))])
