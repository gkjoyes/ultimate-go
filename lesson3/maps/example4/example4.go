// Sample program show how iterating over a map is random.
package main

import "fmt"

// user represent someone using our program.
type user struct {
	firstname string
	lastname  string
}

func main() {
	// Declare and initialize the map with values.
	users := map[string]user{
		"u1": {"x1", "y1"},
		"u2": {"x2", "y2"},
		"u3": {"x3", "y3"},
		"u4": {"x4", "y4"},
	}

	display(users)
	display(users) // Notice order is different here.

}

// display iterate over the map printing each key and value.
func display(users map[string]user) {
	for key, val := range users {
		fmt.Println(key, val)
	}
	fmt.Println()
}
