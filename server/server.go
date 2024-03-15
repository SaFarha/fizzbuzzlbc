package server

import (
	"context"
	"fizzbuzzlbc/configuration"
	"fizzbuzzlbc/database"
	"fizzbuzzlbc/handler"
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
func NewServer(config configuration.Configuration) (Server, error) {

	ret := server{
		server: &http.Server{},
	}

	// new database
	db, err := database.InitDatabase(config)
	if nil != err {
		return nil, err
	}

	handlers := handler.NewHandlers(db)

	ret.server.Addr = ":" + config.Port

	ret.setupRouting(handlers)

	return &ret, nil
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
