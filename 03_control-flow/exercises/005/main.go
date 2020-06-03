package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("when %v is divided by 4, the remander is %v\n", x, x%4)
	}
}
