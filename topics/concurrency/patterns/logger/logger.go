// Package logger demonstrates a pattern of using a buffer to maintain log write continuity by
// discarding log data during write latencies.
package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger provides support to throw log lines away if log writes start to timeout due to latency.
type Logger struct {
	write chan string    // Channel to send/receive data to be logged.
	wg    sync.WaitGroup // Help control the shutdown.
}

// New creates a logger value and initializes it for use. The user can pass the size of
// buffer to use for continuity.
func New(w io.Writer, cap int) *Logger {

	l := Logger{
		write: make(chan string, cap), // Buffered channel if cap > 0
	}

	l.wg.Add(1)

	// Create a goroutine that performs the actual write to disk.
	go func() {
		for d := range l.write {
			// Simulate write to disk.
			fmt.Fprintln(w, d)
		}
		// Mark that we are done.
		l.wg.Done()
	}()

	return &l
}

// Shutdown closes the logger and wait for the write goroutine to terminate.
func (l *Logger) Shutdown() {
	close(l.write)
	l.wg.Wait()
}

// Write is used to write the data to the log.
func (l *Logger) Write(data string) {
	select {
	case l.write <- data: // The writing goroutine got it.
	default: // Drop the write.
		fmt.Println("Dropping the write")
	}
}
