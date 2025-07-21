// Sample program to show how to create goroutines and how the goroutine scheduler behaves with two contexts.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Allocate two logical processors for the schduler to use.
	runtime.GOMAXPROCS(2)
}

func main() {
	// wg is used to wait for the program to finish.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Decalre an anonymous function and create a groutine.
	go func() {

		// Display the lowercase alphabets three times.
		for i := 0; i < 3; i++ {
			for r := 'a'; r <= 'z'; r++ {
				fmt.Printf("%c", r)
			}
		}
		wg.Done()
	}()

	// Declare an anonymous function and create a groutine.
	go func() {
		for i := 0; i < 3; i++ {
			for r := 'A'; r < 'Z'; r++ {
				fmt.Printf("%c", r)
			}
		}
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
