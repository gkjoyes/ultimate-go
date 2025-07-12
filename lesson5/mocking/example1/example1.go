// Sample program to show how you can personally mock concrete types when you need to for your own packages or tests.
package main

import "github.com/gkjoyes/ultimate-go/lesson5/mocking/example1/pubsub"

// publisher is an interface to allow this package to mock the pubsub package support.
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {
	return nil
}

func main() {
	// Create a slice of publisher interface values. Assign the address of pubsub.PubSub value and the address of a mock value.
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}

	// Range over the interface values to see how the publisher interface provides the level of decoupling the user needes.
	// The pubsub package did not need to provide the interface type.
	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
