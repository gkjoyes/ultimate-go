// Sample program to show waitForTask pattern, where the parent goroutine sends a signal to a child goroutine waiting to be told what to do.
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan string)
	go func() {
		defer wg.Done()
		d := <-ch
		fmt.Println("child: received signal:", d)
		time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
	}()

	time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
	task := "task-1"
	ch <- task
	fmt.Println("parent: sent signal:", task)

	wg.Wait()
}
