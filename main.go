// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printHash("A")
	}()

	go func() {
		defer wg.Done()
		printHash("B")
	}()
	wg.Wait()
}

func printHash(prefix string) {

	for i := 0; i < 50000; i++ {
		s := strconv.Itoa(i)
		hash := sha1.Sum([]byte(s))
		fmt.Printf("%s: %05d: %x\n", prefix, i, hash)

	}
}
