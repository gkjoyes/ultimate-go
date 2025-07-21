// Sample program to show why data oriented desgin matters. How data layouts matter more to performance than algorithm efficency.
package example1

import "fmt"

const (
	rows = 4 * 1024
	cols = 4 * 1024
)

// Create a square matrix of 16,777,216 bytes.
var matrix [rows][cols]byte

// node represents a data node for our linked list.
type node struct {
	data byte
	next *node
}

// list points to the head of the linked list.
var list *node

func init() {

	// Create a linked list with same number of elements.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Create a new node and link it backwards.
			list = &node{next: list}

			// Add a value to all even rows.
			if row%2 == 0 {
				list.data = 0xFF
				matrix[row][col] = 0xFF
			}
		}
	}

	// Count the number of elements in the linked list.
	node := list
	count := 0

	for node != nil {
		count++
		node = node.next
	}

	fmt.Printf("Elements in the linked list:%v\n", count)
	fmt.Printf("Elements in the linked matrix:%v\n", rows*cols)
}

func LinkedListTraverse() int {
	count := 0
	node := list

	for node != nil {
		if node.data == 0xFF {
			count++
		}
		node = node.next
	}
	return count
}

func ColumnTraverse() int {
	count := 0

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			if matrix[r][c] == 0xFF {
				count++
			}
		}
	}
	return count
}

func RowTraverse() int {
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0xFF {
				count++
			}
		}
	}
	return count
}
