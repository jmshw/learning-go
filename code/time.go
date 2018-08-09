package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%d seconds since the Epoc\n", now.Unix())
	fmt.Printf("%d nanoseconds since the Epoc\n", now.UnixNano())

	fmt.Printf("Today is %s\n", now.Format("Monday"))
	fmt.Printf("The time is %s\n", now.Format(time.Kitchen))

	fmt.Println("The raw time is:", now)
	//It outputs:
	//The raw time is: 2018-06-13 08:05:23.517876814 +0800 CST m=+0.000432403
}
