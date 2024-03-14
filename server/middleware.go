package server

import (
	"fizzbuzzlbc/utils"
	"net/http"
)

// This is the middleware used to the route in the Server.
func checkAndSetRequestMiddleware(method string, next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		response.Header().Add("Content-Type", "application/json")

		if contentType := request.Header.Get("Content-Type"); contentType != "application/json" {
			utils.HTTPResponseErrorInJSON(response, http.StatusBadRequest, "content-type not defined correctly")
			return
		}

		// Check the request type
		if request.Method != method {
			utils.HTTPResponseErrorInJSON(response, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		next.ServeHTTP(response, request)
	})

}
