// Sample program to show how to implement behavior as context.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// client represents a single connection in the room.
type client struct {
	name   string
	reader *bufio.Reader
}

// temporary is declared to test for the existence of the method coming from the net package.
type temporary interface {
	Temporary() bool
}

// BehaviorAsContext shows how to check for the behavior of an interface that can be returned from the net package.
func (c *client) BehaviorAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case temporary:
				if !e.Temporary() {
					log.Println("Not-Temporary: Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}
				log.Println("ERROR:", err)
			}
		}
		fmt.Println("DATA:", line)
	}
}
