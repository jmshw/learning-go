package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "192.168.81.13:5025"
	fmt.Printf("Conneting to %s\n", addr)

	conn, e := net.Dial("tcp", addr)
	//conn, e := Dial("udp", "192.168.81.13:80")
	//conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80") // Mandatory brackets
	if e != nil {
		os.Stderr.Write([]byte(e.Error()))
		os.Stderr.Write([]byte("\n"))
		return
	}
	defer conn.Close()

	fmt.Println("Connection is ok.")
	n, err := conn.Write([]byte("*IDN?\n"))
	if err == nil {
		buf := make([]byte, 100)
		fmt.Println("Write finished. Reading...")

		n, err = conn.Read(buf)
		if err == nil {
			fmt.Printf("It returns. n:%d, buf:\"%s\"\n", n, buf)
		}
		/*
			r := bufio.NewReader(conn)
			s, err2 := r.ReadString('\n')
			if err2 == nil {
				fmt.Printf("It returns. buf:\"%s\"\n", s)
			}
		*/
	}
}
