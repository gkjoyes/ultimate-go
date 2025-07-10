// This is an example that removes the interface pollution by removing the interface and using the concrete type directly.
package main

// Server is our Server implementation.
type Server struct {
	host string
}

// NewServer returns an concrete value of type Server with a server implementation.
func NewServer(host string) *Server {
	return &Server{host}
}

// Start allows the server to begin to accept requests.
func (s *Server) Start() error {
	return nil
}

// Stop shuts the server down.
func (s *Server) Stop() error {
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *Server) Wait() error {
	return nil
}

func main() {
	// Create a new server.
	s := NewServer("localhost")

	s.Start()
	s.Wait()
	s.Stop()
}
