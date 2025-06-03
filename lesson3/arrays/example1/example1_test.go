// Tests to show data oriented design matters.

// go test -run none -bench . -benchtime 3s

package example1

import "testing"

var res int

func BenchmarkLinkedListTraverse(b *testing.B) {
	var count int

	for i := 0; i < b.N; i++ {
		count = LinkedListTraverse()
	}
	res = count
}

func BenchmarkColumnTraverse(b *testing.B) {
	var count int

	for i := 0; i < b.N; i++ {
		count = ColumnTraverse()
	}
	res = count
}

func BenchmarkRowTraverse(b *testing.B) {
	var count int

	for i := 0; i < b.N; i++ {
		count = RowTraverse()
	}
	res = count
}
