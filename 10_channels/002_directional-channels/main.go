package main

import "fmt"

func main() {
	// recieve channel
	// c := make(<-chan int, 2)
	// send channel
	// c := make(chan<- int, 2)
	// bi directional channel
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------------")
	fmt.Printf("%T\n", c)
}
