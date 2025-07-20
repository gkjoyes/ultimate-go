// Sample program to show a more complicated race condition using an interface value.
package main

import (
	"fmt"
	"os"
	"sync"
)

// Speaker allows for speaking behavior.
type Speaker interface {
	Speak() bool
}

// Ben is a person who can speak.
type Ben struct {
	name string
}

// Speak allows Ben to say hello. It returns false if the method is called through the interface value after a partial write.
func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
		return false
	}
	return true
}

// Jerry is a person who can speak.
type Jerry struct {
	name string
}

// Speak allows Jerry to say hello. It returns false if the method is called through the interface value after a partial write.
func (j *Jerry) Speak() bool {
	if j.name != "Jerry" {
		fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
		return false
	}
	return true
}

func main() {
	// Create values of type Ben and Jerry.
	ben := Ben{"Ben"}
	jerry := Jerry{"Jerry"}

	var person Speaker

	// Have a goroutine constantly assign the pointer of the Ben value to the interface and then Speak.
	go func() {
		for {
			person = &ben
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	// Have a goroutine constantly assign the pointer of the Jerry value to the interface and then Speak.
	go func() {
		for {
			person = &jerry
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	// Just hold main from returing. The data race will cause the program to exit.
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
