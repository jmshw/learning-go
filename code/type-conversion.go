package main

import "fmt"

func main() {
	var f float32
	f = 12.342
	var n int
	n = (int)(f) //int(f) is ok, but (int)f is failed.
	fmt.Println("n:", n)
}
