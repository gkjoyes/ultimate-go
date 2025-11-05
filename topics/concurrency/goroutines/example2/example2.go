// Sample program to show how the goroutine scheduler will time slice goroutines on a single thread.
package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	// Create the first goroutine and manage its lifecycle here.
	go func() {
		defer wg.Done()
		printHashes("A")
	}()

	// Create the second goroutine and manage its lifecycle here.
	go func() {
		defer wg.Done()
		printHashes("B")
	}()

	// Wait for the goroutine to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printHashes calculates the sha1 hash for a range of numbers and prints each in hex encoding.
func printHashes(prefix string) {

	for i := 0; i < 5000; i++ {
		// Convert i to a string.
		num := strconv.Itoa(i)

		// Calculate hash for string num.
		sum := sha1.Sum([]byte(num))

		// Print prefix: 5-digit-number: hex encoded hash
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}
	fmt.Println("Completed", prefix)
}
