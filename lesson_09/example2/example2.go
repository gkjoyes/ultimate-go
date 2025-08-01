// Sample program to show how to use the atomic package to provide safe access to numberic types.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all groutines.
var counter int64

func main() {
	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two groutines
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
