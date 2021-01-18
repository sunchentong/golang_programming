package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func test11(command string) {
	fmt.Printf("-- Exercise 1.1 \n")
	fmt.Println(command)
}

func test12(args []string) {
	fmt.Printf("-- Exercise 1.2 \n")
	for index, value := range args {
		fmt.Printf("index : %d , args[%d] : %s \n", index, index, value)
	}
}

func test13(args []string) {
	fmt.Printf("-- Exercise 1.3 \n")
	var s, sep string

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

func sampleDup() {
	fmt.Printf("-- Sample Dup \n")
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("open %s failed, err : %v \n", arg, err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("filename : %s \n", os.Args[0])
			fmt.Printf("line :%s , n : %d \n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	// exercise 1.1
	test11(os.Args[0])

	// exercise 1.2
	test12(os.Args[1:])

	// exercise 1.3
	test13(os.Args[1:])

	// sample dup1
	sampleDup()
}
