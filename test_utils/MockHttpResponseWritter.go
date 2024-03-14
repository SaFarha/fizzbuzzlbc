package testutils

import (
	"net/http"
	"strconv"
)

type mockHTTPResponseWriter struct {
	header http.Header
}

// NewMockHTTPResponseWriter create a new mock of http.ResponseWriter
func NewMockHTTPResponseWriter() http.ResponseWriter {
	return &mockHTTPResponseWriter{
		header: map[string][]string{},
	}
}

func (mock *mockHTTPResponseWriter) Header() http.Header {
	return mock.header
}

func (mock *mockHTTPResponseWriter) WriteHeader(statusCode int) {
	mock.header.Add("status", strconv.Itoa(statusCode))
}

func (mock *mockHTTPResponseWriter) Write(value []byte) (int, error) {
	return len(value), nil
}
