// Sample program to show how embedded types work with interfaces.
package main

import "fmt"

// notifier defines notification behavior.
type notifier interface {
	notify()
}

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
	user
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

	// The embedded inner type's implementation of the interface is "promoted" to the outer type.
	sendNotification(&a)
}

// sendNotification accepts values that implement the notifier interface and send notifications.
func sendNotification(n notifier) {
	n.notify()
}
