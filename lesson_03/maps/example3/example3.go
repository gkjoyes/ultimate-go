// Sample program to show how only types that can have equality defined on them can be a map key.
package main

// user represent someone using our program.
type user struct {
	firstName string
	lastName  string
}

type users []user

func main() {
	/*
	   // Declare and make a map that uses a slice as the key.
	   userMap := make(map[users]int)

	   // Iterate over the map.

	   	for key, val := range userMap {
	   		fmt.Println(key, val)
	   	}
	*/
}
