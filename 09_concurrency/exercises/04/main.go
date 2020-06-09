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
	var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()

			v := counter
			runtime.Gosched()
			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count: ", counter)
}
