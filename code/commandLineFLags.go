package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dnssec := flag.Bool("dnssec", false, "Request DNSSEC records") // 0
	port := flag.String("port", "53", "Set the query port")        // 1
	flag.Usage = func() {                                          // 2
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults() // 3
	}

	flag.Parse() // 4

	// 0:Define a bool flag, -dnssec. The variable must be a pointer otherwise the
	//   package can not set its value;
	// 1:Idem, but for a port option;
	// 2:Slightly redefine the Usage function, to be a litter more verbose;
	// 3:For every flag given, PrintDefaults will output the help string;
	// 4:Parse the flags and fill the variables.

	// After the flags have been parsed you can use then:
	if *dnssec {
		fmt.Println("dnssec is:", *dnssec)
		fmt.Println("port is:", *port)
	}
}
