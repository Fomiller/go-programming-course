package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// create 20 go routines to add numbers to the channel c
	for i := 0; i <= 10; i++ {
		// what runs each loop -- create new go routine send 10 numbers onto channel
		go func() {
			// adding numbers onto channel
			for j := 0; j <= 10; j++ {
				c <- j
			}
		}()
	}
	for k := 0; k <= 100; k++ {
		v := <-c
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}
