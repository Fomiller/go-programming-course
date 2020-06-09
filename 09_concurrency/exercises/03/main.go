package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count: ", counter)
}
