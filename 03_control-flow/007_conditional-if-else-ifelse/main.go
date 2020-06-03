package main

import "fmt"

func main() {

	x := 43
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 42 {
		fmt.Println("our value was 42")
	} else {
		fmt.Println("our value was not 40 or 42")
	}
}
