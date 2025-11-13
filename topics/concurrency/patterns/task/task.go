// Package task provides a pool of goroutines to perform tasks.
package task

import "sync"

// Worker must be implemented by types that want to use the run pool.
type Worker interface {
	Work()
}

// Task provides a pool of goroutines that can execute any Worker tasks that are submitted.
type Task struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(goroutines int) *Task {
	t := Task{
		// Using an unbuffered channel because we want the guarantee of knowing the work being
		// submitted is actually being worked on after the call to Do() returns.
		work: make(chan Worker),
	}

	// The goroutines are the pool. So we could add code to change the size of the pool later on.
	t.wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer t.wg.Done()
			for w := range t.work {
				w.Work()
			}
		}()
	}

	return &t
}

// Shutdown waits for all the goroutines to shutdown.
func (t *Task) Shutdown() {
	close(t.work)
	t.wg.Wait()
}

// Do submits work to the pool.
func (t *Task) Do(w Worker) {
	t.work <- w
}
