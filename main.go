package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	timeout := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Printf("Completed work: %s\n", result)
	case <-ctx.Done():
		fmt.Println("Timeout!")
	}

}
