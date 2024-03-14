package utils

import "net/http"

// HTTPResponseInJSON function is a simple encapsulation to write http response correctly.
func HTTPResponseInJSON(response http.ResponseWriter, statusCode int, data []byte) {

	response.WriteHeader(statusCode)
	_, _ = response.Write(data)

}
