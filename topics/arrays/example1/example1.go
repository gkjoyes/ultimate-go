// Sample program to show why data oriented design matters. How data layouts matter more to performance than algorithm efficiency.
package main

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
var head = &node{}

func init() {
	list := head

	// Create a linked list with same number of elements.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Create a new node and link it forward.
			list.next = &node{}
			list = list.next

			// Add a value to all even rows and corresponding nodes.
			if row%2 == 0 {
				list.data = 0xFF
				matrix[row][col] = 0xFF
			}
		}
	}

	// Count the number of elements in the linked list.
	node := head.next
	count := 0

	for node != nil {
		count++
		node = node.next
	}

	fmt.Printf("Elements in the linked list: \t%v\n", count)
	fmt.Printf("Elements in the matrix: \t%v\n", rows*cols)
}

func LinkedListTraverse() int {
	node := head.next
	count := 0

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

func main() {
	fmt.Println("List Traversal:\t", LinkedListTraverse())
	fmt.Println("Row Traversal:\t", RowTraverse())
	fmt.Println("Column Traversal:\t", ColumnTraverse())
}
