package main

import "fmt"

func main() {
	// equals make(chan int, 0), it will block when writing to it.
	//ch := make(chan int)
	ch := make(chan int, 1) //This is ok.

	fmt.Println("Before writing to chan")
	ch <- 1

	fmt.Println("After writing to chan")
	<-ch

	fmt.Println("After Reading from chan")
}
