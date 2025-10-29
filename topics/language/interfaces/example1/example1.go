// Sample program that could benefit from polymorphic behavior with interfaces.
package main

import "fmt"

// file defines a system file.
type file struct {
	name string
}

// reads data from the provided file system.
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Ultimate Go</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// reads data from the provided pipe network connection.
func (pipe) read(b []byte) (int, error) {
	s := `{name: "x1", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

func main() {
	// Create two values one of type file and one of type pipe.
	f := file{name: "data.json"}
	p := pipe{name: "data_service"}

	// Call each retrieve function for each concrete type.
	retrieveFile(f)
	retrievePipe(p)
}

// retrieveFile can read from a file and process the data.
func retrieveFile(f file) error {
	data := make([]byte, 100)

	len, err := f.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// retrievePipe can read from a pipe and process the data.
func retrievePipe(p pipe) error {
	data := make([]byte, 100)

	len, err := p.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
