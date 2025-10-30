// This is an example that creates interface pollution by improperly using an interface when one is not needed.
package main

// Server defines a contract for tcp servers.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}

// server is our Server implementation.
type server struct {
	host string
}

// NewServer returns an interface value of type Server with a server implementation.
func NewServer(host string) Server {
	// SMELL - Storing an unexported type pointer in the interface.
	return &server{host}
}

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
	return nil
}

// Stop shuts the server down.
func (s *server) Stop() error {
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
	return nil
}

func main() {
	// Create a new server.
	s := NewServer("localhost")

	s.Start()
	s.Wait()
	s.Stop()
}
