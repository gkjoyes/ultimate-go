// Sample program to show that you cannot take the address of an element in a map.
package main

// player represent someone playing our game.
type player struct {
	name  string
	score int
}

func main() {
	players := map[string]player{
		"p1": {"x1", 40},
		"p2": {"x2", 20},
	}

	/*
		// Trying to take the address of a map element fails.
		p1 := &players["p1"]
		p1.score++
	*/

	// Instead take the element, modify it, and put it back.
	p1 := players["p1"]
	p1.score++
	players["p1"] = p1
}
