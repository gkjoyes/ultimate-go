// Sample program to show waitForResult pattern, where the parent goroutine waits for the child goroutine to finish some work to signal the result.
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
		result := "result-1"
		ch <- result
		fmt.Println("child: sent signal:", result)
	}()

	d := <-ch
	fmt.Println("parent: received signal:", d)
	time.Sleep(time.Second)
}
