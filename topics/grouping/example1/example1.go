// This is an example of using type hierarchies with a OOP pattern.
// This is not something we want to do in Go. Go does not have the concept of sub-typing.
// All types are their own and the concept of base and derived types do not exist in Go.
// This pattern does not provide a good design principle in a Go program.
package main

import "fmt"

// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak provides generic behavior for all animals and how they speak.
func (a *Animal) Speak() {
	fmt.Printf("UGH! My name is %s, it is %t I am a mammal\n", a.Name, a.IsMammal)
}

// Dog contains everything an Animal is but specific attributes that only a Dog has.
type Dog struct {
	Animal
	PackFactor int
}

// Speak knows how to speak like a dog.
func (d *Dog) Speak() {
	fmt.Printf("Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n", d.Name, d.IsMammal, d.PackFactor)
}

// Cat contains everything an Animal is but specific attributes that only a Cat has.
type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf("Meow! My name is %s, is it %t I am a mammal with a climb factor of %d.\n", c.Name, c.IsMammal, c.ClimbFactor)
}

func main() {
	/*

		// Create a list of Animals that know how to speak.
		animals := []Animal{
			Dog{
				Animal: Animal{
					Name:     "x1",
					IsMammal: true,
				},
				PackFactor: 5,
			},
			Cat{
				Animal: Animal{
					Name:     "x2",
					IsMammal: true,
				},
				ClimbFactor: 4,
			},
		}

		// Have the Animals Speaks.
		for _, animal := range animals {
			animal.Speak()
		}
	*/
}
