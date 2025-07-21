// Sample program to show what happens when the outer and inner type implement the same interface.
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

// admin represents an admin user with privillages.
type admin struct {
	user
	level string
}

// notify notifies admins of different events.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
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

	// The embedded inner type's implementation of the interface is NOT "promoted" to the outer type.
	sendNotification(&a)

	// We can access the inner type's method directly.
	a.user.notify()

	// The inner type's method is NOT promoted.
	a.notify()
}

// sendNotification accepts values that implement the notifier interface and send notifications.
func sendNotification(n notifier) {
	n.notify()
}
