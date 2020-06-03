package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[2])

	for i, v := range x {
		fmt.Println(i, v)
	}

}
