package main

import "fmt"

func addSlice1(s []int, x int) {
	// 改变参数 s 长度，形参指针发生变化
	s = append(s, x)
}

func addSlice2(s []int, x int) []int {
	s = append(s, x)
	return s
}

func changeSlice(s []int, i, x int) {
	s[i] = x
}

func main() {
	s1 := []int{}
	for i := 0; i < 5; i++ {
		addSlice1(s1, i)
	}
	fmt.Println("s1: ", s1)

	s2 := []int{}
	for i := 0; i < 5; i++ {
		s2 = addSlice2(s2, i)
	}
	fmt.Println("s2: ", s2)

	s3 := make([]int, 5)
	for i := 0; i < 5; i++ {
		changeSlice(s3, i, i)
	}
	fmt.Println("s3: ", s3)
}
