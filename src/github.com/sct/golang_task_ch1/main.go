package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// exercise 1.1/1.2
	command := os.Args[0]
	fmt.Println(command)

	for index, value := range os.Args[1:] {
		fmt.Printf("index : %d , args[%d] : %s \n", index, index, value)
	}

	// exercise 1.3
	var s, sep string
	args := os.Args[1:]

	start := time.Now()
	// low performace
	for _, val := range args {
		s += sep + val
		sep = " "
	}
	fmt.Printf("%.4fs elapsed \n", time.Since(start).Seconds())

	start = time.Now()
	// high performance
	strings.Join(args, " ")
	fmt.Printf("%.4fs elapsed \n", time.Since(start).Seconds())
}
