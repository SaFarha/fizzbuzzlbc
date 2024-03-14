package server

import (
	testutils "fizzbuzzlbc/test_utils"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCheckAndSetRequestMiddleware(t *testing.T) {

	fakeHandler := func(w http.ResponseWriter, r *http.Request) {}
	res := testutils.NewMockHTTPResponseWriter()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)

	// Error Content Type not good
	{
		fakeHandler(res, req)
		checkMiddleware := checkAndSetRequestMiddleware(http.MethodPost, fakeHandler)
		checkMiddleware.ServeHTTP(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf(
				"http status not good \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"),
			)
		}

		res.Header().Del("status")
	}

	// Setup for other test
	req.Header.Add("Content-Type", "application/json")

	// Error Method mismatch
	{
		fakeHandler(res, req)
		checkMiddleware := checkAndSetRequestMiddleware(http.MethodGet, fakeHandler)
		checkMiddleware.ServeHTTP(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusMethodNotAllowed) {
			t.Errorf(
				"http status not good \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusMethodNotAllowed),
				res.Header().Get("status"),
			)
		}

		res.Header().Del("status")
	}

	// Setup For other test
	checkMiddleware := checkAndSetRequestMiddleware(http.MethodPost, fakeHandler)

	// Middleware ok
	{
		fakeHandler(res, req)
		checkMiddleware.ServeHTTP(res, req)

		if res.Header().Get("status") != "" {
			t.Errorf(
				"http status not good \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusMethodNotAllowed),
				res.Header().Get("status"),
			)
		}
	}

}
