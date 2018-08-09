package main

import "fmt"

func main() {
	var i int
	var inter interface{}
	inter = i
	switch t := inter.(type) {
	default:
		fmt.Println("Unknown:", t)
	case int:
		fmt.Println("int:", t)
	}

	n := 3
	switch n {
	default:
		fmt.Println("default")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}
