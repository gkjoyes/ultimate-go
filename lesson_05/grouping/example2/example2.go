// This is an example of using composition and interfaces. This is something we want to do in Go.
// We will group common types by their behaviour and not by their state.
// This pattern does provide a good design principle in Go programming.
package main

import "fmt"

// Speaker provide a common behaviour for all concrete types to follow if they want to be a part of this group.
type Speaker interface {
	Speak()
}

// Dog contains everything a Dog needs.
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete types that know how to speak.
func (d *Dog) Speak() {
	fmt.Printf("Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n", d.Name, d.IsMammal, d.PackFactor)
}

// Cat contains everything a Cat needs.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak knows how to speak lika cat.
// This makes a Cat now part of a group of concrete types that know how to speak.
func (c *Cat) Speak() {
	fmt.Printf("Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n", c.Name, c.IsMammal, c.ClimbFactor)
}

func main() {
	// Create a list of Animals that know how to speak.
	speakers := []Speaker{
		&Dog{
			Name:       "x1",
			IsMammal:   true,
			PackFactor: 5,
		},
		&Cat{
			Name:        "x2",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, s := range speakers {
		s.Speak()
	}
}
