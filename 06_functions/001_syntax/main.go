package main

import "fmt"

func main() {
	foo()
	bar("Forrest")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := zar("Forrest", "Miller")
	fmt.Println(x)
	fmt.Println(y)
}

// function with zero parameters
func foo() {
	fmt.Println("Hello from Foo")
}

// function with an argument of type string
func bar(s string) {
	fmt.Println("Hello,", s)
}

// function returning a string
func woo(str string) string {
	return fmt.Sprint("Hello from woo, ", str)
}

func zar(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `,says "Hello"`)
	b := false
	return a, b

}
