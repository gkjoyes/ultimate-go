// Sample program to show how to use a read/write mutex to define critical sections of code that needs synchronous access.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// data is a slice that will be shared.
var data []string

// rwMutex is used to define a critical section of code.
var rwMutex sync.RWMutex

// Number of reads occuring at any given time.
var readCount int64

func main() {
	// wg is used to manage concurency.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a write goroutine.
	go func() {
		for i := 0; i < 10; i++ {
			writer(i)
		}
		wg.Done()
	}()

	// Create eight reader goroutines.
	for i := 0; i < 8; i++ {
		go func(id int) {
			for {
				reader(i)
			}
		}(i)
	}

	// Wait for the write goroutine to finish.
	wg.Wait()
	fmt.Println("Program Complete!")
}

// writer adds a new string to the slice in random intervals.
func writer(i int) {
	// Only allow one goroutine read/write to the slice at a time.
	rwMutex.Lock()
	{
		// Capture the current read count.
		rc := atomic.LoadInt64(&readCount)

		// Perform some work since we have a full lock.
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("Performing Write: RCound[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

// reader wakes up and iterates over the data slice.
func reader(id int) {
	// Any goroutine can read when no write operation is taking place.
	rwMutex.RLock()
	{
		// Increment the read count value by 1.
		rc := atomic.AddInt64(&readCount, 1)

		// Perform some read work and display values.
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d: Performing Read: Length[%d] RCount[%d]\n", id, len(data), rc)

		// Decrement the read count value by 1.
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}
