// Sample program to show how the default error type is implemented.
package main

import "fmt"

// https://pkg.go.dev/builtin#error
type error interface {
	Error() string
}

// https://go.dev/src/errors/errors.go
type errorString struct {
	s string
}

// https://go.dev/src/errors/errors.go
func (e *errorString) Error() string {
	return e.s
}

// https://go.dev/src/errors/errors.go
func New(text string) error {
	return &errorString{text}
}

func main() {
	if err := webCall(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Life is Good")
}

// webCall performs a web operation.
func webCall() error {
	return New("Bad Request")
}
