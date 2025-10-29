// Sample program to show how maps behave when you read an absent key.
package main

import "fmt"

func main() {
	// Create a map to track scores for players in a game.
	scores := make(map[string]int)

	// Read the element at key "x1". It is absent so we get the zero value for this map's value type.
	score := scores["x1"]
	fmt.Println("Score: ", score)

	// If we need to check for the presence of a key we use a 2 variable assignment.
	score, ok := scores["x1"]
	fmt.Println("Score: ", score, "Present: ", ok)

	// We can leverage the zero-value behavior to write convenient code this this:
	scores["x1"]++

	// Without this behavior we would have to code in a defensive way like this:
	if s, ok := scores["x1"]; ok {
		scores["x1"] = s + 1
	} else {
		scores["x1"] = 1
	}

	score, ok = scores["x1"]
	fmt.Println("Score: ", score, "Present: ", ok)
}
