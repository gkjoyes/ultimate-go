package task

import (
	"sync"
)

type Worker interface {
	Work()
}

type Task struct {
	ch chan Worker
	wg sync.WaitGroup
}

func New(goroutines int) *Task {
	task := Task{
		ch: make(chan Worker),
	}

	task.wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer task.wg.Done()
			for t := range task.ch {
				t.Work()
			}
		}()
	}

	return &task
}

func (t *Task) Shutdown() {
	close(t.ch)
	t.wg.Wait()
}

func (t *Task) Do(task Worker) {
	t.ch <- task
}
