// Sample program to show how maps are not safe for concurrent use by default.
package main

import (
	"fmt"
	"sync"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int)

func main() {
	//  wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	// Create first goroutine to update score for A.
	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	// Create second goroutine to update score for B.
	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	// Wait for goroutines to finish.
	wg.Wait()
	fmt.Println("Final Scores:", scores)

}

// Output: runtime will detect concurrent writes and panic.
