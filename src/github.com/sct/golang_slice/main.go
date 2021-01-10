package main

import (
	"fmt"
	"sort"
)

func main() {

	// slice的赋值拷贝
	fmt.Printf("====================== create slice ...\r\n")
	s1 := make([]int, 3) // make empty slice , contain 3 elements [0,0,0]
	s2 := s1
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	// slice的遍历，采用for range的方式
	fmt.Printf("====================== loop for slice \r\n")
	for idx, item := range s1 {
		fmt.Println(idx, item)
	}

	// slice的遍历，采用for循环方式
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}

	// slice append 方法使用
	fmt.Printf("====================== slice append usage \r\n")
	var s3 []int
	s3 = append(s3, 1, 2, 3)
	s4 := []int{4, 5, 6}
	s3 = append(s3, s4...)
	fmt.Println(s3)

	// slice copy 方法
	// copy(dest_slice, src_slice, []T)
	fmt.Printf("====================== slice copy usage \r\n")
	a := []int{1, 2, 3, 4, 5}
	s5 := make([]int, 5)
	copy(s5, a)
	fmt.Println(a)
	fmt.Println(s5)

	s5[0] = 100
	fmt.Println(a)
	fmt.Println(s5)

	// delete element from slice
	// eg: delete element b[2] from b, then b = b[:2] + b[3:]...
	fmt.Printf("====================== delete element[2] from slice \r\n")
	b := []int{30, 31, 32, 33, 34, 35, 36, 37}
	b = append(b[:2], b[3:]...)
	fmt.Println(b)

	// exercise 1
	// output : [  0,1,2,3,4,5,6,7,8,9]
	fmt.Printf("====================== exercise 1 \r\n")
	c := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		c = append(c, fmt.Sprintf("%v", i))
	}
	fmt.Println(c)

	// exercise 2
	// sort slice
	fmt.Printf("====================== exercise 2 \r\n")
	var d = [...]int{3, 7, 8, 9, 1}
	sort.Ints(d[:]) // 数组转切片，调用sor.Ints进行升序排序
	fmt.Println(d)

	//降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(d[:])))
	fmt.Println(d)
}
