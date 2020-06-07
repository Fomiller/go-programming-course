package main

import "fmt"

func main() {
	defer foo()
	defer zam()
	bar()
}

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
func zam() {
	fmt.Println("zam")
}
