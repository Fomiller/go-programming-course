package main

import (
	"fmt"
)

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for i := 0; i <= 2; i++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}
