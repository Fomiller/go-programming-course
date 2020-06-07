package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	{
		y := 42
		fmt.Println(y)
	}
}

func foo() {
	fmt.Println("foo")
}
