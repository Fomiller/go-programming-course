package main

import "fmt"

func main() {
	// similar to while statement
	// x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }
	// fmt.Println("done")

	// infinite loop
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")

}
