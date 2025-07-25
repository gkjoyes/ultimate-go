// Sample program to show you need to validate your benchmark results.
// go test -run none -bench . -benchtime 3s -benchmem
package example2

import "testing"

// n contains the data to sort.
var n []int

// Generate the numbers to sort.
func init() {
	for i := 0; i < 1_000_000; i++ {
		n = append(n, i)
	}
}

func BenchmarkSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		single(n)
	}
}

func BenchmarkUnlimited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unlimited(n)
	}
}
func BenchmarkNumCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numCPU(n, 0)
	}
}
