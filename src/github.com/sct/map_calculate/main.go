package main

import (
	"fmt"
	"strings"
)

func calculate(words string) (result map[string]int) {
	s1 := strings.Split(words, " ")
	result = make(map[string]int, 10)
	for _, item := range s1 {
		_, ok := result[item]
		if ok {
			result[item]++
		} else {
			result[item] = 1
		}
	}
	return result
}

// 统计句子中的单词个数
// eg : how do you do
// 返回: [do:2 how:1 you:1]
func main() {
	words := "how do you do"
	m1 := calculate(words)
	fmt.Println(m1)
}
