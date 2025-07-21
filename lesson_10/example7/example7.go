// Sample program to show drop pattern, where the parent goroutine signals 2000 pieces of work to a single child goroutine
// that can't handle all the work. If child is not ready/don't have enough capacity, that work is discarded and dropped.
package main

import (
	"fmt"
	"time"
)

func main() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child: recived signal:", p)
		}
	}()

	const work = 2000
	for w := range work {
		data := fmt.Sprintf("data: %d", w)
		select {
		case ch <- data:
			fmt.Println("parent: sent signal: ", data)
		default:
			fmt.Println("parent: dropped data:", data)
		}
	}
	close(ch)
	fmt.Println("parent: sent shutdown signal")
	time.Sleep(time.Second)
}
