package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l") // create a *cmd
	//err := cmd.Run()                     // run it

	// capture the standard output from the command:
	buf, err := cmd.Output() // buf is a []byte

	fmt.Println("err:", err)
	fmt.Printf("buf:%s\n", buf)
}
