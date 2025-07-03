// Sample program to show how to understand method sets.
package main

import "fmt"

// notifier is an interface that defines notification type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func main() {
	// Create a value of type User and send a notification.
	u := user{name: "x1", email: "x1@gmail.com"}

	// NOTE: Values of type user do not implement the interface.
	// sendNotifiaiton(u)

	sendNotifiaiton(&u)
}

// sendNotification accepts values that implements the notifier interface and send notification.
func sendNotifiaiton(n notifier) {
	n.notify()
}
