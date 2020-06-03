package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not priont")
	case (2 == 4):
		fmt.Println("should not print")
	case (3 == 3):
		fmt.Println("print")
		fallthrough
	case (4 == 4):
		fmt.Println("print as well")
	default:
		fmt.Println("default")
	}
}
