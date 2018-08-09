package main

import "fmt"

type T1 int
type T2 string
type T3 float32

type P struct {
	T3
}

type T struct {
	T1
	*T2
	x, y int
}

func main() {
	t := new(T)
	fmt.Printf("%v, %v, %v\n", t.T1, t.T2, t.x)
}
