// Sample program to show pooling pattern, the parent goroutine signals 100 pieces of work to a pool of child goroutines
// waiting for work to perform.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan string)

	grs := runtime.GOMAXPROCS(0)
	for c := range grs {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d: recived signal: %s\n", child, d)
			}
			fmt.Printf("child %d: recived shutdown signal\n", child)
		}(c)
	}

	const work = 100
	for w := range work {
		ch <- fmt.Sprintf("data: %d", w)
		fmt.Println("parent: sent signal:", w)
	}
	close(ch)
	fmt.Println("parent: sent shutdown signal")
	time.Sleep(time.Second)
}
