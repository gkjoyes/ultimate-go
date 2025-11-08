// Sample program to show cancellation pattern, where the parent goroutine creates a child goroutine to perform some work.
// The parent goroutine is only willing to wait 150 millisecond for that work to be completed.
package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Println("work completed: ", result)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
