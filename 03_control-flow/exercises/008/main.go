package main

import "fmt"

func main() {
	switch {
	case 2 == 3:
		fmt.Println("statement one is correct")
	case 2 == 2:
		fmt.Println("statement two is correct")
	}
}
