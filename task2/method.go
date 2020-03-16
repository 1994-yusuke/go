package main

import "fmt"

type Int int

func (n *Int) Inc() {
	i := *n + 2
	*n = i
}

func (n *Int) Plus() int {
	return int(*n * 2)
}

func main() {
	var a Int
	fmt.Println(a)
	a.Plus()
	fmt.Println(a)
	a.Inc()
	fmt.Println(a)
}
