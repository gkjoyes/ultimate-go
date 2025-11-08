// Sample program to show fanOutSem pattern, where a semaphore is added to the fan out pattern to restrict the number of child goroutines that can be schedule to run.
package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"time"
)

func main() {
	children := 2000
	ch := make(chan string, children)

	grs := runtime.GOMAXPROCS(0)
	sem := make(chan bool, grs)

	for c := range children {
		go func(child int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
				result := fmt.Sprintf("result: %d", child)
				ch <- result
				fmt.Println("child: sent signal:", result)
			}
			<-sem
		}(c)
	}

	for children > 0 {
		result := <-ch
		children--
		fmt.Printf("parent: received signal: %s\n", result)
	}
}
