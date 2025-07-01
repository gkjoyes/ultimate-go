// Sample program to show how maps are reference types.
package main

import "fmt"

func main() {

	// Initialize a map with values.
	scores := map[string]int{
		"p1": 20,
		"p2": 10,
	}

	// Pass the map to a function to perform some mutation.
	double(scores, "p2")

	// See the change is visible in our map.
	fmt.Println("Score: ", scores["p2"])
}

// double multiplies specific player score by 2.
func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
