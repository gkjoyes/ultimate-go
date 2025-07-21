// Sample program to show how what we are doing is NOT embedding a type but just using a type as a field.
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

// admin represents an admin user with privillages.
type admin struct {
	person user // NOT Embedding
	level  string
}

func main() {
	// Create an admin user.
	a := admin{
		person: user{
			name:  "x1",
			email: "x1@gmail.com",
		},
		level: "super",
	}
	// We can access fields methods.
	a.person.notify()
}
