// Sample program to demonstrates how the logger package works.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gkjoyes/ultimate-go/lesson_11/example1/logger"
)

type device struct {
	problem bool
}

// Write implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		// Simulate disk problems.
		time.Sleep(time.Second)
	}
	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	// Number of goroutines that will be writing logs.
	const grs = 10

	var d device
	// Create a logger instance with a buffer sized to the number of goroutines.
	l := logger.New(&d, grs)

	// Create goroutines, each writing to disk.
	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking. Capture interrupt signals to toggle device issues.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		// We have data race here with the Write method. Let's keep things simple to show mechanics.
		d.problem = !d.problem
	}
}
