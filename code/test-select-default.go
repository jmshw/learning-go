package main

import "fmt"
import "time"

func main() {
	ch := make(chan int)
	ch2 := make(chan bool)

	go func() {
		//ch <- 1
	}()

	time.Sleep(time.Millisecond * 100)

	select {
	default:
		fmt.Println("default")
	case <-ch:
		fmt.Println("<-ch")
	case <-ch2:
		fmt.Println("<-ch2")
	}
}
