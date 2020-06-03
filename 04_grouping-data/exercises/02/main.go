package main

import "fmt"

func main() {
	x := []int{22, 45, 33, 44, 21, 232, 123, 56546, 11, 2}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
