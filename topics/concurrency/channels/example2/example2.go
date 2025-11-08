// Sample program to show fanout pattern, where the parent goroutine creates 2000 child goroutines and waits for them to signal the results.
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	children := 2000
	ch := make(chan string, children)

	for c := range children {
		go func(child int) {
			time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
			result := fmt.Sprintf("result-%d", child)
			ch <- result
			fmt.Println("child: sent signal:", result)
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println("parent: received signal:", d)
	}
}
