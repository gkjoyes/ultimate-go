// Sample program to show how to implement behaviour as context.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// client represents a single connection in the room.
type client struct {
	name   string
	reader *bufio.Reader
}

// TypeAsContext shows how to check multiple types of possible custom error types that can be returned from the net package.
func (c *client) TypeAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Println("Not-Temporary: Client leaving chat")
					return
				}
			case *net.AddrError:
				if !e.Temporary() {
					log.Println("Not-Temporary: Client leaving chat")
					return
				}
			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("Not-Temporary: Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}
				log.Println("Error:", err)
			}
		}
		fmt.Println("Data:", line)
	}
}

// temporary is declared to test for the existence of the method coming from the net package.
type temporary interface {
	Temporary() bool
}

// BehaviourAsConext shows how to check for the behaviour of an interface that can be returned from the net package.
func (c *client) BehaviourAsContext() {
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
