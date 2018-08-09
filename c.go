package main

// #include <stdlib.h>
// #include <string.h>
// #include <stdio.h>
import "C"
import "fmt"

func arc4random() uint32 {
	return uint32(C.arc4random())
}

func strlen(str *C.char) uint32 {
	return uint32(C.strlen(str))
}

func main() {
	value := arc4random()
	fmt.Printf("value:%d\n", value)

	//str := []byte{'a', 'b', 'c'}
	//	cStr := (*C.char)(str)
	//	nLen := strlen(str)
	//	fmt.Printf("nLen:%d\n", nLen)
}
