// Sample program to show drop pattern, where the parent goroutine signals 2000 pieces of work to a single child goroutine that can't handle all the work.
// If child is not ready/don't have enough capacity, that work is discarded and dropped.
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	const cap = 100
	ch := make(chan string, cap)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for work := range ch {
			fmt.Println("child: received signal:", work)
			time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
		}
	}()

	const works = 2000
	for w := range works {
		work := fmt.Sprintf("work-%d", w)
		select {
		case ch <- work:
			fmt.Println("parent: sent signal: ", work)
		default:
			fmt.Println("parent: dropped data:", work)
		}
	}
	close(ch)
	fmt.Println("parent: sent shutdown signal")
	wg.Wait()
}
