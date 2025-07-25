package example3

import (
	"runtime"
	"testing"
)

var nums []int64

func init() {
	nums = generateList(10000)
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(nums)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(nums, runtime.NumCPU())
	}
}
