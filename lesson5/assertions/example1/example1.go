// Sample program demonstrating when implicit interface conversions are provided by the compiler.
package main

import "fmt"

// Mover provides support for moving things.
type Mover interface {
	Move()
}

// Locker provides support for locking and unlocking things.
type Locker interface {
	Lock()
	Unlock()
}

// MoveLocker provides support for moving and locking things.
type MoveLocker interface {
	Mover
	Locker
}

// bike represents a concrete type for the example.
type bike struct {
	name string
}

// Move can change the position of a bike.
func (bike) Move() {
	fmt.Println("Moving the bike")
}

// Lock prevents a bike from moving.
func (bike) Lock() {
	fmt.Println("Locking the bike")
}

// Unlock allows a bike to be moved.
func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	// Declare variables of the MoveLocker and Mover interfaces set to their zero value.
	var ml MoveLocker
	var m Mover

	// Create a value of type bike and assign the value to the MoveLocker interface value.
	ml = bike{}

	// An interface value of type MoveLocker can be implicitly converted into a value of type Mover.
	// They both declare a method named move.
	m = ml

	// Interface type Mover does not declare methods named Lock and Unlock. Therefore, the compiler can't perform
	// an implicit conversion to assign a value of interface type Mover to an interface value of type MoveLocker.
	// ml = m

	// We can perform a type assertion at runtime to support the assignment.
	b := m.(bike)
	ml = b
}
