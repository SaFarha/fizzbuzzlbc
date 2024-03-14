package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fizzbuzzlbc/helper"
	testutils "fizzbuzzlbc/test_utils"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPostFizzbuzzHandler(t *testing.T) {

	// Mock function
	oldIoutilReadAll := ioutilReadAllPostFizzbuzzHandler
	oldJSONUnmarshall := jsonUnmarshallPostFizzbuzzHandler
	oldHelperFizzBuzz := helperFizzbuzzPostFizzbuzzHandler
	oldJSONMarshall := jsonMarshallPostFizzbuzzHandler

	// value for test
	res := testutils.NewMockHTTPResponseWriter()
	postBody, _ := json.Marshal(map[string]interface{}{
		"int1":  3,
		"int2":  5,
		"limit": 100,
		"str1":  "fizz",
		"str2":  "buzz",
	})
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
	req.Header.Add("Content-Type", "application/json")

	// Failed ReadAll error
	{
		ioutilReadAllPostFizzbuzzHandler = func(r io.Reader) ([]byte, error) {
			return nil, errors.New("error read all")
		}

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusInternalServerError) {
			t.Errorf("error read all failed, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusInternalServerError),
				res.Header().Get("status"))
		}

		ioutilReadAllPostFizzbuzzHandler = oldIoutilReadAll
		res.Header().Del("status")
	}

	// Failed ReadAll missing body
	{
		ioutilReadAllPostFizzbuzzHandler = func(r io.Reader) ([]byte, error) {
			return nil, nil
		}

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error read all missing body, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		ioutilReadAllPostFizzbuzzHandler = oldIoutilReadAll
		res.Header().Del("status")
	}

	// Setup for other test
	ioutilReadAllPostFizzbuzzHandler = oldIoutilReadAll

	// Json marshall failed
	{
		jsonUnmarshallPostFizzbuzzHandler = func(data []byte, v any) error {
			return errors.New("json unmarshal failed")
		}

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusInternalServerError) {
			t.Errorf("error json unmarshall body, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusInternalServerError),
				res.Header().Get("status"))
		}

		jsonUnmarshallPostFizzbuzzHandler = oldJSONUnmarshall
		res.Header().Del("status")
	}

	// Setup for other test
	jsonUnmarshallPostFizzbuzzHandler = oldJSONUnmarshall

	// params Limit null
	{
		postBody, _ := json.Marshal(map[string]interface{}{
			"int1": 3,
			"int2": 5,
			"str1": "fizz",
			"str2": "buzz",
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
		req.Header.Add("Content-Type", "application/json")

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error params limit is null, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		res.Header().Del("status")
	}

	// params Limit null
	{
		postBody, _ := json.Marshal(map[string]interface{}{
			"int1": 3,
			"int2": 5,
			"str1": "fizz",
			"str2": "buzz",
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
		req.Header.Add("Content-Type", "application/json")

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error params limit is null, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		res.Header().Del("status")
	}

	// params int1 null
	{
		postBody, _ := json.Marshal(map[string]interface{}{
			"limit": 100,
			"int2":  5,
			"str1":  "fizz",
			"str2":  "buzz",
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
		req.Header.Add("Content-Type", "application/json")

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error params int1 is null, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		res.Header().Del("status")
	}

	// params int2 null
	{
		postBody, _ := json.Marshal(map[string]interface{}{
			"int1":  3,
			"limit": 100,
			"str1":  "fizz",
			"str2":  "buzz",
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
		req.Header.Add("Content-Type", "application/json")

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error params int2 is null, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		res.Header().Del("status")
	}

	// params str1 null
	{
		postBody, _ := json.Marshal(map[string]interface{}{
			"int1":  3,
			"int2":  5,
			"limit": 100,
			"str2":  "buzz",
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
		req.Header.Add("Content-Type", "application/json")

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error params str1 is null, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		res.Header().Del("status")
	}

	// params str1 null
	{
		postBody, _ := json.Marshal(map[string]interface{}{
			"int1":  3,
			"int2":  5,
			"limit": 100,
			"str1":  "buzz",
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))
		req.Header.Add("Content-Type", "application/json")

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error params str2 is null, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		res.Header().Del("status")
	}

	// setup for tother test
	req = httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))

	// helper fizzbuzz failed
	{
		helperFizzbuzzPostFizzbuzzHandler = func(params helper.FizzBuzzHelperParams) ([]string, error) {
			return nil, errors.New("helper fizz buzz failed")
		}

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusBadRequest) {
			t.Errorf("error helper failed, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusBadRequest),
				res.Header().Get("status"))
		}

		helperFizzbuzzPostFizzbuzzHandler = oldHelperFizzBuzz
		res.Header().Del("status")
	}

	// setup to other test
	helperFizzbuzzPostFizzbuzzHandler = oldHelperFizzBuzz
	req = httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))

	// failed json marshall
	{
		jsonMarshallPostFizzbuzzHandler = func(v any) ([]byte, error) {
			return nil, errors.New("json marshall failed")
		}

		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusInternalServerError) {
			t.Errorf("error json marshall failed, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusInternalServerError),
				res.Header().Get("status"))
		}

		jsonMarshallPostFizzbuzzHandler = oldJSONMarshall
		res.Header().Del("status")
	}

	// setup for other test
	jsonMarshallPostFizzbuzzHandler = oldJSONMarshall
	req = httptest.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(postBody))

	// Test OK
	{
		PostFizzbuzzHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusOK) {
			t.Errorf("test ok, status not match \nwe need : %s \nand we get: %s\n",
				strconv.Itoa(http.StatusOK),
				res.Header().Get("status"))
		}
	}
}
