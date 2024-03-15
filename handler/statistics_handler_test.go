package handler

import (
	"errors"
	"fizzbuzzlbc/database/models"
	"fizzbuzzlbc/mocks"
	"fizzbuzzlbc/repositories"
	testutils "fizzbuzzlbc/test_utils"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestStatisticsHandler(t *testing.T) {
	handlers := &Handlers{
		ParamRequestRepo: &mocks.FzParamRequestRepository{},
		RequestStatRepo:  &mocks.FzRequestStatsRepository{},
	}

	var (
		//oldGetstatsMarshall = jsonMarshallGetStatisticHandler
		oldGetStatgetMax = helperGetMaxFzRequestStat
	)

	// value for test
	res := testutils.NewMockHTTPResponseWriter()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	req.Header.Add("Content-Type", "application/json")

	// failed GetMaxFzRequest error
	{
		helperGetMaxFzRequestStat = func(prRepo repositories.FzParamRequestRepository, rsRepo repositories.FzRequestStatsRepository) (*models.FzParamRequest, error) {
			return nil, errors.New("error")
		}

		handlers.StatisticHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusInternalServerError) {
			t.Errorf("error not found status \nwe need: %s \nand we get: %s \n",
				strconv.Itoa(http.StatusInternalServerError),
				res.Header().Get("status"))
		}
	}

	helperGetMaxFzRequestStat = oldGetStatgetMax
	res = testutils.NewMockHTTPResponseWriter()
	req = httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	{
		helperGetMaxFzRequestStat = func(prRepo repositories.FzParamRequestRepository, rsRepo repositories.FzRequestStatsRepository) (*models.FzParamRequest, error) {
			return &models.FzParamRequest{
				Uid:   "toto",
				Int1:  3,
				Int2:  4,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			}, nil
		}

		handlers.StatisticHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusOK) {
			t.Errorf("error not found status \nwe need: %s \nand we get: %s \n",
				strconv.Itoa(http.StatusOK),
				res.Header().Get("status"))
		}
	}

	helperGetMaxFzRequestStat = oldGetStatgetMax
	res = testutils.NewMockHTTPResponseWriter()
	req = httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// failed GetMaxFzRequest error
	{
		helperGetMaxFzRequestStat = func(prRepo repositories.FzParamRequestRepository, rsRepo repositories.FzRequestStatsRepository) (*models.FzParamRequest, error) {
			return nil, gorm.ErrRecordNotFound
		}

		handlers.StatisticHandler(res, req)

		if res.Header().Get("status") != strconv.Itoa(http.StatusNotFound) {
			t.Errorf("error not found status \nwe need: %s \nand we get: %s \n",
				strconv.Itoa(http.StatusNotFound),
				res.Header().Get("status"))
		}

		helperGetMaxFzRequestStat = oldGetStatgetMax
	}
}
