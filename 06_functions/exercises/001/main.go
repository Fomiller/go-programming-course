package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	a := 42
	return a
}

func bar() (int, string) {
	x := 69
	y := "apple"
	return x, y
}
