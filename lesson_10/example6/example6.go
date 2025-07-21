// Sample program to show boundedWorkPooling pattern, where a pool of child goroutines is created to service a fixed
// amount of work.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}

	grs := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(grs)

	ch := make(chan string, grs)
	for c := range grs {
		go func(child int) {
			defer wg.Done()
			for w := range ch {
				fmt.Printf("child: %d recived signal: %s\n", child, w)
			}
			fmt.Printf("child: %d received shutdown signal\n", child)
		}(c)
	}

	for _, w := range work {
		ch <- w
	}
	close(ch)
	wg.Wait()
}
