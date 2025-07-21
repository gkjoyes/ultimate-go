// Sample program to show waitForResult pattern, where the parent goroutine waits for the child goroutine to finish
// some work to signal the result.

package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child: sent signal")
	}()

	d := <-ch
	fmt.Println("parent: recived signal:", d)
	time.Sleep(time.Second)
}
