package main

import "fmt"

func check(words string) bool {
	length := len(words)
	if length == 0 {
		return false
	}

	i := 0
	j := length - 1

	for i <= j {
		if words[i] != words[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s1 := "asdfdsa"
	s2 := "dfggfd"

	ret1 := check(s1)
	ret2 := check(s2)

	fmt.Println(ret1)
	fmt.Println(ret2)
}
