// Sample program to show boundedWorkPooling pattern, where a pool of child goroutines is created to service a fixed amount of work.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	grs := runtime.GOMAXPROCS(0)
	ch := make(chan string, grs)

	var wg sync.WaitGroup
	wg.Add(grs)

	for child := range grs {
		go func(c int) {
			defer wg.Done()
			for w := range ch {
				fmt.Printf("child: %d received signal: %s\n", c, w)
			}
			fmt.Printf("child: %d received shutdown signal\n", c)
		}(child)
	}

	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}
	for _, w := range work {
		ch <- w
	}
	close(ch)
	wg.Wait()
}
