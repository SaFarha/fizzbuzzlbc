package server

import (
	"fizzbuzzlbc/handler"
	"net/http"
)

// this function is used like a router to initialize all the route of the Server.
func (srv *server) setupRouting(handlers handler.Handlers) {

	mux := http.NewServeMux()

	mux.Handle("/fizzbuzz", checkAndSetRequestMiddleware(http.MethodPost, handlers.PostFizzbuzzHandler))

	mux.Handle("/statistics", checkAndSetRequestMiddleware(http.MethodGet, handlers.StatisticHandler))

	srv.server.Handler = mux
}
