package application

import (
	"fizzbuzzlbc/configuration"
	"fizzbuzzlbc/server"
)

// Application Interface is used to contain all things needed by the project.
type Application interface {
	Run() error

	ShutDown()
}

type application struct {
	srv server.Server
}

// NewApplication initializes a new Application interface.
func NewApplication() (Application, error) {

	conf := configuration.LoadConfig()

	srv, err := server.NewServer(conf)
	if nil != err {
		return nil, err
	}

	return &application{
		srv: srv,
	}, nil

}

// Run is used to run the application.
func (app application) Run() error {

	return app.srv.Listen()

}

// ShutDown is used to stop the application.
func (app application) ShutDown() {

	app.srv.Shutdown()

}
