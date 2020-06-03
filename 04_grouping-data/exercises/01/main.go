package main

import "fmt"

func main() {
	x := [5]int{2, 3, 4, 5, 56}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
