package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// declare counter as an int64 with a zero value of '0'
	var counter int64
	var gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			// access the memory address of the counter variable and iterate it by '1'
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count: ", counter)
}
