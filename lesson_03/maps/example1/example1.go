// Sample program to show how to initialize a map, write to it, then read and delete from it.
package main

import "fmt"

// user represent someone using our program.
type user struct {
	forename string
	surname  string
}

func main() {
	// Declare a map that store values of type user with a key of type string.
	users := make(map[string]user)

	// Add key/value pairs to the map.
	users["u1"] = user{"x1", "y1"}
	users["u2"] = user{"x2", "y2"}
	users["u3"] = user{"x3", "y3"}
	users["u4"] = user{"x4", "y4"}

	// Read the value at a specific key.
	u1 := users["u1"]
	fmt.Printf("%+v\n", u1)

	// Replace the value at the x1 key.
	users["u1"] = user{"m1", "n1"}

	// Read the u1 again.
	fmt.Printf("%+v\n", users["u1"])

	// Delete the value at a specfic key.
	delete(users, "u2")

	// Check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "u100")
}
