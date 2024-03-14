package server

import (
	"fizzbuzzlbc/handler"
	"net/http"
)

// this function is used like a router to initialize all the route of the Server.
func (srv *server) setupRouting() {

	mux := http.NewServeMux()

	mux.Handle("/fizzbuzz", checkAndSetRequestMiddleware(http.MethodPost, handler.PostFizzbuzzHandler))

	srv.server.Handler = mux
}
