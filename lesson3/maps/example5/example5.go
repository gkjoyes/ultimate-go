// Sample program to show how to walk through a map by alphabetical key order.
package main

import (
	"fmt"
	"sort"
)

// user represent someone using our program.
type user struct {
	firstname string
	lastname  string
}

func main() {
	// Declare and initialize the map with values.
	users := map[string]user{
		"u1": {"x1", "y1"},
		"u3": {"x3", "y3"},
		"u2": {"x2", "y2"},
		"u4": {"x4", "y4"},
	}

	// Pull the keys from the map.
	keys := make([]string, 0, len(users))
	for key := range users {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically.
	sort.Strings(keys)

	// Walk throught the keys and pull each value from the map.
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}
