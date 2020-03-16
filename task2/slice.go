package main

import "fmt"

func main() {
	n1 := []int{1, 11, 111, 1111}
	sum := 0
	for _, v := range n1 {
		sum += v
	}
	fmt.Println(sum)
}
