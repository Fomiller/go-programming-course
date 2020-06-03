package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)
	x = append(x, 55, 555, 44, 31234)
	fmt.Println(x)
	y := []int{234, 123, 455, 666}
	x = append(x, y...)
	fmt.Println(x)
}
