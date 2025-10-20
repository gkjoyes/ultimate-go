// Sample program to show memory for an array is contiguous.
package main

import "fmt"

func main() {
	users := [4]string{"x1", "x2", "x3", "x4"}

	for i, u := range users {
		fmt.Printf("Value: [%v]\tAddress: [%p]\tIndexAddress: [%p]\n", u, &u, &users[i])
	}
}
