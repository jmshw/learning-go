package main

import (
	"bufio"
	"os"
)

// main
func main() {
	f, _ := os.Open("./text")
	defer f.Close()
	r := bufio.NewReader(f) // make it a bufio to access the ReadString method
	for {
		s, ok := r.ReadString('\n') // Read a line from the input
		if ok == nil {
			os.Stdout.Write([]byte(s))
		} else {
			break
		}

	}
}
