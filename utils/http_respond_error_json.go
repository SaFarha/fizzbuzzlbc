package utils

import (
	"encoding/json"
	"net/http"
)

// JSON Error response
type errorResponse struct {
	Error string `json:"error"`
}

// used to mock to test the handler correctly
var (
	jsonMarshalHTTPResponseErrorInJSON = json.Marshal
)

// HTTPResponseErrorInJSON is an abstraction to respond an error in json format.
func HTTPResponseErrorInJSON(response http.ResponseWriter, statusCode int, error string) {

	errorStruct := errorResponse{
		Error: error,
	}

	responseJSON, err := jsonMarshalHTTPResponseErrorInJSON(errorStruct)
	if nil != err {
		response.WriteHeader(http.StatusInternalServerError)
		_, _ = response.Write(nil)
		return
	}

	response.WriteHeader(statusCode)
	_, _ = response.Write(responseJSON)
}
