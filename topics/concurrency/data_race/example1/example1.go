// Sample program to show how to create race condition in our program. We don't want to do this.

// command: go build -race
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

func main() {
	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				// Capture the value of Counter.
				value := counter

				// Increment our local value of Counter.
				value++

				// Store the value back into Counter.
				counter = value
			}
			wg.Done()
		}()
	}
	// Wait for the goroutine to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
