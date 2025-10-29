// Sample program to show how to embed a type into another type and the relationship between the inner and outer type.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify notifies users of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	user  // Embedded Type
	level string
}

func main() {
	// Create an admin user.
	a := admin{
		user: user{
			name:  "x1",
			email: "x1@gmail.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	a.user.notify()

	// The inner type's method is promoted.
	a.notify()
}
