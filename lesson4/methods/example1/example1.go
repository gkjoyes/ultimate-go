// Sample program to show how to declare methods and how the Go compiler supports them.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail implments a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// Values of type user can be used to call methods declared with both value and pointer receivers.
	x1 := user{name: "x1", email: "x1@gmail.com"}
	x1.changeEmail("x1@hotmail.com")
	x1.notify()

	// Pointers of type user can also be user to call methods declared with both value and pointer receivers.
	x2 := &user{name: "x2", email: "x2@gmail.com"}
	x2.changeEmail("x2@hotmail.com")
	x2.notify()
}
