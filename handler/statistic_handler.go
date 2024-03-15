package handler

import (
	"encoding/json"
	"errors"
	"fizzbuzzlbc/helper"
	"fizzbuzzlbc/utils"
	"gorm.io/gorm"
	"net/http"
)

// getStatisticHandlerResponse is the StatisticHandler response model
type getStatisticHandlerResponse struct {
	Int1  uint64 `json:"int1" binding:"required"`
	Int2  uint64 `json:"int2" binding:"required"`
	Limit uint64 `json:"limit" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}

// used to mock to test the handler correctly
var (
	jsonMarshallGetStatisticHandler = json.Marshal
	helperGetMaxFzRequestStat       = helper.GetMaxFzRequestStat
)

// StatisticHandler is a GET handler that return the most frequent request has been
func (handler *Handlers) StatisticHandler(response http.ResponseWriter, request *http.Request) {

	// Close the body
	defer request.Body.Close()

	stat, err := helperGetMaxFzRequestStat(handler.ParamRequestRepo, handler.RequestStatRepo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			println("toto")
			utils.HTTPResponseErrorInJSON(response, http.StatusNotFound, "Stats not found")
			return
		}

		utils.HTTPResponseErrorInJSON(response, http.StatusInternalServerError, "Unexpected error")
		return
	}

	dataResponse, err := jsonMarshallGetStatisticHandler(getStatisticHandlerResponse{
		Int1:  stat.Int1,
		Int2:  stat.Int2,
		Limit: stat.Limit,
		Str1:  stat.Str1,
		Str2:  stat.Str2,
	})
	if nil != err {
		utils.HTTPResponseErrorInJSON(response, http.StatusInternalServerError, "Unexpected error")
		return
	}

	utils.HTTPResponseInJSON(response, http.StatusOK, dataResponse)
}
