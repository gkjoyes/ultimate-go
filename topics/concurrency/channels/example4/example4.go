// Sample program to show pooling pattern, the parent goroutine signals 100 pieces of work to a pool of child goroutines waiting for work to perform.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	grs := runtime.GOMAXPROCS(0)
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(grs)

	for g := range grs {
		go func(id int) {
			defer wg.Done()
			for w := range ch {
				fmt.Printf("child %d: received signal: %s\n", id, w)
			}
			fmt.Printf("child %d: received shutdown signal\n", id)
		}(g)
	}

	const works = 100
	for w := range works {
		work := fmt.Sprintf("work-%d", w)
		ch <- work
		fmt.Println("parent: sent signal:", work)
	}
	close(ch)
	fmt.Println("parent: sent shutdown signal")

	wg.Wait()
}
