package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gkjoyes/ultimate-go/task"
)

var names = []string{
	"x1",
	"x2",
	"x3",
	"x4",
	"x5",
	"x6",
	"x7",
	"x8",
}

type Printer struct {
	Name string
}

func (p *Printer) Work() {
	fmt.Println(p.Name)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	goroutines := 10
	t := task.New(goroutines)

	var wg sync.WaitGroup
	wg.Add(goroutines * len(names))

	for i := 0; i < goroutines; i++ {
		for _, name := range names {
			p := &Printer{
				Name: name,
			}
			go func() {
				defer wg.Done()
				t.Do(p)
			}()
		}

	}
	wg.Wait()
	t.Shutdown()
}
