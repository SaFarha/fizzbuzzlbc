package handler

import (
	"encoding/json"
	"fizzbuzzlbc/database/models"
	"fizzbuzzlbc/helper"
	"fizzbuzzlbc/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// JSON params of the handler
type postFizzBuzzHandlerParams struct {
	Int1  *uint64 `json:"int1" binding:"required"`
	Int2  *uint64 `json:"int2" binding:"required"`
	Limit *uint64 `json:"limit" binding:"required"`
	Str1  *string `json:"str1" binding:"required"`
	Str2  *string `json:"str2" binding:"required"`
}

// JSON response of the handler
type postFizzBuzzHandlerResponse struct {
	Result []string `json:"result" binding:"required"`
}

// used to mock to test the handler correctly
var (
	ioutilReadAllPostFizzbuzzHandler            = ioutil.ReadAll
	jsonUnmarshallPostFizzbuzzHandler           = json.Unmarshal
	jsonMarshallPostFizzbuzzHandler             = json.Marshal
	helperFizzbuzzPostFizzbuzzHandler           = helper.FizzBuzzHelper
	helperFizzbuzzPostAddRequestStatisticHelper = helper.AddRequestStatisticHelper
)

// PostFizzbuzzHandler is a POST handler that take postFizzBuzzHandlerParams in body and return the result of
// helper.FizzbuzzHelper or a http error.
func (handler *Handlers) PostFizzbuzzHandler(response http.ResponseWriter, request *http.Request) {

	// Close the body
	defer request.Body.Close()
	var wg sync.WaitGroup

	// check body
	dataBody, err := ioutilReadAllPostFizzbuzzHandler(request.Body)
	if nil != err {
		utils.HTTPResponseErrorInJSON(response, http.StatusInternalServerError, err.Error())
		return
	}

	if len(dataBody) == 0 {
		utils.HTTPResponseErrorInJSON(response, http.StatusBadRequest, "Missing body")
		return
	}

	var params postFizzBuzzHandlerParams
	err = jsonUnmarshallPostFizzbuzzHandler(dataBody, &params)
	if nil != err {
		utils.HTTPResponseErrorInJSON(response, http.StatusInternalServerError, "Error in body")
		return
	}

	if nil == params.Limit || nil == params.Int1 || nil == params.Int2 || nil == params.Str2 || nil == params.Str1 {
		utils.HTTPResponseErrorInJSON(response, http.StatusBadRequest, "Missing parameter")
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := helperFizzbuzzPostAddRequestStatisticHelper(handler.ParamRequestRepo, handler.RequestStatRepo, models.FzParamRequest{
			Limit: *params.Limit,
			Int1:  *params.Int1,
			Int2:  *params.Int2,
			Str1:  *params.Str1,
			Str2:  *params.Str2,
		})
		if nil != err {
			fmt.Println("Error create statistics")
		}
	}()

	resp, err := helperFizzbuzzPostFizzbuzzHandler(helper.FizzBuzzHelperParams{
		Limit: *params.Limit,
		Int1:  *params.Int1,
		Int2:  *params.Int2,
		Str1:  *params.Str1,
		Str2:  *params.Str2,
	})
	if nil != err {
		wg.Wait()
		utils.HTTPResponseErrorInJSON(response, http.StatusBadRequest, "Error in parameters: "+err.Error())
		return
	}

	dataResponse, err := jsonMarshallPostFizzbuzzHandler(postFizzBuzzHandlerResponse{
		Result: resp,
	})
	if nil != err {
		wg.Wait()
		utils.HTTPResponseErrorInJSON(response, http.StatusInternalServerError, "Unexpected error")
		return
	}

	wg.Wait()

	// http return
	utils.HTTPResponseInJSON(response, http.StatusOK, dataResponse)
}
