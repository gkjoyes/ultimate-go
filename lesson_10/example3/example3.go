// Sample program to show waitForTask pattern, where the parent goroutine sends a signal to a child goroutine
// waiting to be told what to do.
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("child: received signal:", d)
	}()

	time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
	ch <- "data"
	fmt.Println("parent: sent signal")

	time.Sleep(time.Second)
}
