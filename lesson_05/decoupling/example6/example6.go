// Sample program to demonstrating struct composition.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Data is the structure of the data we are copying.
type Data struct {
	Line string
}

// Puller declares behaviour of pulling data.
type Puller interface {
	Pull(d *Data) error
}

// Storer declares behaviour of storing data.
type Storer interface {
	Store(d *Data) error
}

// X1 is a system we need to pull data from.
type X1 struct {
	Host    string
	Timeout time.Duration
}

// Pull knows how to pull data out of X1.
func (*X1) Pull(d *Data) error {
	switch r.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error reading data from X1")
	default:
		d.Line = "data"
		fmt.Println("IN: ", d.Line)
	}

	return nil
}

// X2 is a system we need to store data into.
type X2 struct {
	Host    string
	Timeout time.Duration
}

// Store knows how to store data into X2.
func (*X2) Store(d *Data) error {
	fmt.Println("OUT:", d.Line)
	return nil
}

// pull knows how to pull bulks of data from X1.
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

// store knows how to store bulks of data into X2.
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

// Copy knows how to pull and store data from the System.
func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(p, data)
		if i > 0 {
			if _, err := store(s, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	x1 := &X1{
		Host:    "localhost:8000",
		Timeout: time.Second,
	}
	x2 := &X2{
		Host:    "localhost:9000",
		Timeout: time.Second,
	}

	if err := Copy(x1, x2, 3); err != io.EOF {
		log.Fatal(err)
	}
}
