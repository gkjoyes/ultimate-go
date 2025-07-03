// Sample program to show how to declare and use function types.
package main

import "fmt"

// event displays global events.
func event(message string) {
	fmt.Println(message)
}

type data struct {
	name string
	age  int
}

// event displays events for this data.
func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

// fireEvent1 uses an anonymous function type.
func fireEvent1(f func(string)) {
	f("anonymous")
}

// handler represent a function for handling events.
type handler func(string)

// fireEvent2 uses a function type.
func fireEvent2(h handler) {
	h("handler")
}

func main() {

	d := data{
		name: "x1",
		age:  30,
	}

	// fireEvent1 accepts any function or method with the right signature.
	fireEvent1(event)
	fireEvent1(d.event)

	// fireEvent2 accepts any function or method of type `handler` or any literal function or method with the right signature.
	fireEvent2(event)
	fireEvent2(d.event)

	// Declare a variable of type handler for the gobal and method based event functions.
	h1 := handler(event)
	h2 := handler(d.event)

	// fireEvent2 accepts values of type `handler`.
	fireEvent2(h1)
	fireEvent2(h2)

	// fireEvent1 accepts any function or method with the right signature.
	fireEvent1(h1)
	fireEvent1(h2)
}
