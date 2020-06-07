package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 43, 45, 6}
	s := sum(ii...)
	fmt.Println("all numbers:", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers: ", s2)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
