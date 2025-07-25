// Sample program to show how merge sort works.
package example2

import (
	"runtime"
	"sync"
)

// func main() {
// res := single([]int{5, 3, 4, 6, 1, 2}, 0)
// res := unlimited([]int{5, 3, 4, 6, 1, 2}, 0)
// res := numCPU([]int{5, 3, 4, 6, 1, 2}, 0)
// fmt.Println(res)
// }

// single uses a single goroutine to perform the merge sort.
func single(n []int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Sort the left side.
	l := single(n[:i])

	// Sort the right side.
	r := single(n[i:])

	// Merge ordered list.
	return merge(l, r)
}

// unlimited uses a goroutine for every split to perform the merge sort.
func unlimited(n []int) []int {
	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Maintain the ordered left and right side lists.
	var l, r []int

	// For each split we will have 2 goroutines.
	var wg sync.WaitGroup
	wg.Add(2)

	// Sort the left side concurrently.
	go func() {
		l = unlimited(n[:i])
		wg.Done()
	}()

	// Sort the right side concurrently.
	go func() {
		r = unlimited(n[i:])
		wg.Done()
	}()

	// Wait for the splitting to end.
	wg.Wait()

	// Place things in order and merge ordered lists.
	return merge(l, r)
}

// numCPU uses number of gorouttines equal to number of cores.
func numCPU(n []int, level int) []int {
	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Maintain the ordered left and right side lists.
	var l, r []int

	maxLevel := runtime.NumCPU()

	if level <= maxLevel {
		level++

		// For each split we will have 2 goroutines.
		var wg sync.WaitGroup
		wg.Add(2)

		// Sort the left side concurrently.
		go func() {
			l = numCPU(n[:i], level)
			wg.Done()
		}()

		// Sort the right side concurrently.
		go func() {
			r = numCPU(n[i:], level)
			wg.Done()
		}()

		// Wait for the splitting to end.
		wg.Wait()

		// Place things in order and merge ordered lists.
		return merge(l, r)
	}

	// Sort the left and right side on this goroutine.
	l = numCPU(n[:i], level)
	r = numCPU(n[i:], level)

	// Place things in order and merge ordered lists.
	return merge(l, r)
}

// merge performs the merging of the two lists in proper order.
func merge(l, r []int) []int {
	res := make([]int, 0, len(l)+len(r))

	for {
		switch {
		case len(l) == 0:
			return append(res, r...)
		case len(r) == 0:
			return append(res, l...)
		case l[0] < r[0]:
			res = append(res, l[0])
			l = l[1:]
		case r[0] < l[0]:
			res = append(res, r[0])
			r = r[1:]
		}
	}
}
