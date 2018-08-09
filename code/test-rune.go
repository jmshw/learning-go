package main

import "fmt"

func main() {
	str := "abc中文"
	fmt.Printf("len(str):%d\n", len(str))
	r := []rune(str)
	r = r[0:4]
	fmt.Printf("len([]rune(str)):%d\n", len([]rune(str)))
	fmt.Printf("r:%v\n", r)
	str2 := string(r)
	fmt.Printf("str2:%v\n", str2)
}
