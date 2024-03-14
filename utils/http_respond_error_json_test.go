package utils

import (
	"errors"
	testutils "fizzbuzzlbc/test_utils"
	"net/http"
	"strconv"
	"testing"
)

func TestHTTPResponseErrorInJsonJSON(t *testing.T) {

	//Setup unit test mocking
	mockHTTPResponseWriter := testutils.NewMockHTTPResponseWriter()
	oldJSONMarshallFunc := jsonMarshalHTTPResponseErrorInJSON

	// TestJsonFailed
	{
		// Setup mocking of failed jsonMarshal
		jsonMarshalHTTPResponseErrorInJSON = func(v any) ([]byte, error) {
			return nil, errors.New("JSON Failed")
		}

		statusCode := 0
		errorString := "error json marshall"
		HTTPResponseErrorInJSON(mockHTTPResponseWriter, statusCode, errorString)

		if mockHTTPResponseWriter.Header().Get("status") != strconv.Itoa(http.StatusInternalServerError) {
			t.Errorf(
				"http status not good \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusInternalServerError),
				mockHTTPResponseWriter.Header().Get("status"),
			)
		}

		mockHTTPResponseWriter.Header().Del("status")
		jsonMarshalHTTPResponseErrorInJSON = oldJSONMarshallFunc
	}

	// Test OK
	{
		statusCode := 500
		errorString := "error test"
		HTTPResponseErrorInJSON(mockHTTPResponseWriter, statusCode, errorString)

		if mockHTTPResponseWriter.Header().Get("status") != strconv.Itoa(statusCode) {
			t.Errorf(
				"http status not good \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(statusCode),
				mockHTTPResponseWriter.Header().Get("status"),
			)
		}

		mockHTTPResponseWriter.Header().Del("status")
	}

}
