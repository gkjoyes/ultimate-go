// Sample programs to test cpu bounded work.
package example3

import (
	"math/rand/v2"
	"sync"
	"sync/atomic"
)

// func main() {
// 	nums := generateList(10000)
// 	fmt.Println(add(nums))
// 	fmt.Println(addConcurrent(nums, runtime.NumCPU()))
// }

func generateList(totalNums int64) []int64 {
	nums := make([]int64, totalNums)
	for i := range totalNums {
		nums[i] = rand.Int64N(totalNums)
	}
	return nums
}

func add(nums []int64) int64 {
	var total int64
	for _, n := range nums {
		total += n
	}
	return total
}

func addConcurrent(nums []int64, grs int) int64 {
	stride := len(nums) / grs

	var wg sync.WaitGroup
	wg.Add(grs)
	var total int64

	for i := 0; i < grs; i++ {
		go func(idx int) {
			defer wg.Done()
			start := idx * stride
			end := start + stride

			var t int64
			for j := start; j < end; j++ {
				t += nums[j]
			}

			atomic.AddInt64(&total, t)
		}(i)
	}
	wg.Wait()
	return total
}
