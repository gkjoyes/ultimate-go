// Sample program to show the syntax of the type asserations.
package main

import (
	"fmt"
	"log"
)

// user defines a user in our application.
type user struct {
	id   int
	name string
}

// finder represents the ablility to find users.
type finder interface {
	find(id int) (*user, error)
}

// userSVC is a service for dealing with users.
type userSVC struct {
	host string
}

// find implements the finder interface using pointer samantics.
func (*userSVC) find(id int) (*user, error) {
	return &user{id: id, name: "x1"}, nil
}

func main() {
	svc := userSVC{host: "128.162.3.4:3456"}

	if err := run(&svc); err != nil {
		log.Fatal(err)
	}
}

// run performs the find operation against the concrete data that is passed into the call.
func run(f finder) error {
	u, err := f.find(1234)
	if err != nil {
		return err
	}
	fmt.Printf("Found user %+v\n", u)

	// You can use a type assertion to get a copy of the userSVC pointer that is stored inside the interface.
	// if ok is false, not such value is present in the interface.
	svc, ok := f.(*userSVC)
	if ok {
		fmt.Println("Host: ", svc.host)
	}
	return nil
}
