package server

import (
	"context"
	"net/http"
)

// Server Interface is used to contain and use the http server in the most simple way possible.
type Server interface {
	Listen() error
	Shutdown()
}

type server struct {
	server *http.Server
}

// NewServer initialize a new Server interface.
// it takes the port to listen.
func NewServer(port string) Server {

	ret := server{
		server: &http.Server{},
	}

	ret.server.Addr = ":" + port
	ret.setupRouting()

	return &ret
}

// Listen is an encapsulation of the function server.ListenAndServe.
// This function is used to run the server.
func (srv server) Listen() error {

	return srv.server.ListenAndServe()

}

// Shutdown is an encapsulation of the function server.Shutdown.
// This function is used to stop the server.
func (srv server) Shutdown() {

	_ = srv.server.Shutdown(context.Background())

}
