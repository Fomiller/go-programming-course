package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()

}

func foo() {
	fmt.Println("Message from foo")
	wg.Done()
}
func bar() {
	fmt.Println("Message from bar")
	wg.Done()
}
