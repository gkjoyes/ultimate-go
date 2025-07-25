// GOGC=off GOMAXPROCS=1 go test -bench . -benchtime 3s
// GOGC=off go test -bench . -benchtime 3s
package example4

import (
	"runtime"
	"testing"
)

var docs []string

func init() {
	docs = generateList(1e4)
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find(docs, "Go")
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(docs, "Go", runtime.NumCPU())
	}
}
