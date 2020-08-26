layout :=  "2006-01-02 15:04:05"
times, _ := time.Parse(timeLayout, "2014-06-15 08:37:18")  
timeUnix := times.Unix()  
fmt.Printf("times is %+v \n, timeUnix is %+v", times, timeUnix)
