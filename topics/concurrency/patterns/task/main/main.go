package main

import (
	"log"
	"sync"
	"time"

	"github.com/gkjoyes/ultimate-go/topics/concurrency/patterns/task"
)

// names provides a set of names to display.
var names = []string{
	"x1",
	"x2",
	"x3",
	"x4",
	"x5",
	"x6",
	"x7",
	"x8",
	"x9",
}

// printer provides special support for printing names.
type printer struct {
	name string
}

// Work implements the Worker interface.
func (p printer) Work() {
	log.Println(p.name)
	time.Sleep(1 * time.Second)
}

func main() {
	const goroutines = 10

	// Create a task pool.
	t := task.New(goroutines)

	var wg sync.WaitGroup
	wg.Add(goroutines * len(names))

	for i := 0; i < goroutines; i++ {
		// Iterate over slice of names.
		for _, name := range names {
			// Create a new printer and provide the specific name.
			p := printer{
				name: name,
			}
			go func() {
				// Submit the task to be worked on. When Do() returns, we know it is being handled.
				t.Do(p)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	// Shutdown the task pool and wait for all existing work to be completed.
	t.Shutdown()
}
