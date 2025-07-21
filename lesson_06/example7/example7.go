// Sample program to show how wrapping errors work with the stdlib.
package main

import (
	"errors"
	"fmt"
)

// AppError represents a custom error type.
type AppError struct {
	State int
}

// Error implements the error interface.
func (c *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", c.State)
}

// Cause iterates through all the wrapped errors until the root error value is reached.
func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}

func main() {
	// Make the function call and validate the error.
	if err := firstCall(10); err != nil {

		// How to use the As function.
		var ap *AppError
		if errors.As(err, &ap) {
			fmt.Println("As says it is an AppError")
		}

		// Use type as context to determine cause.
		switch v := Cause(err).(type) {
		case *AppError:
			fmt.Println("Custom App Error:", v.State)
		default:
			fmt.Println("Default Error")
		}

		fmt.Printf("%v\n", err)
	}
}

// firstCall makes a call to a second function and wraps any error.
func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return fmt.Errorf("firstCall -> secondCall(%d) : %w", i, err)
	}
	return nil
}

// secondCall makes a call to a third function and wraps any error.
func secondCall(i int) error {
	if err := thirdCall(); err != nil {
		return fmt.Errorf("secondCall -> thirdCall() : %w", err)
	}
	return nil
}

// thridCall create an error value we will validate.
func thirdCall() error {
	return &AppError{99}
}
