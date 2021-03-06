package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := foo(xi...)
	p := bar(xi)
	fmt.Println(t)
	fmt.Println(p)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
