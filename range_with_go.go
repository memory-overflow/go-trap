package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		go func() {
			fmt.Println(n)
		}()
	}
	time.Sleep(time.Second)

	for _, n := range nums {
		number := n
		go func() {
			fmt.Println(number)
		}()
	}
	time.Sleep(time.Second)
}