package utils

import (
	testutils "fizzbuzzlbc/test_utils"
	"strconv"
	"testing"
)

func TestHTTPRespondJSON(t *testing.T) {

	//Setup unit test mocking
	mockHTTPResponseWriter := testutils.NewMockHTTPResponseWriter()

	// Test OK
	{
		statusCode := 200
		data := []byte("test")
		HTTPResponseInJSON(mockHTTPResponseWriter, statusCode, data)

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
