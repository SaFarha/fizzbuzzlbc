package helper

import (
	"errors"
	"fizzbuzzlbc/database/models"
	"fizzbuzzlbc/mocks"
	"fizzbuzzlbc/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	fakeFzParam2 = models.FzParamRequest{
		Int1:  3,
		Int2:  5,
		Limit: 100,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	fakeUid2 = utils.CreateRequestUid(fakeFzParam)

	fakeReqStats2 = models.FzRequestStat{
		ParamRequestUid: fakeUid,
		Count:           1,
	}
)

func TestGetRequestStatsMostCountFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)

	mockedRsRepo.On("GetRequestStatsMostCount").Return(nil, errors.New("error"))

	_, err := GetMaxFzRequestStat(mockedPrRepo, mockedRsRepo)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestMaxFzRequestGetPrByUidFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)

	mockedRsRepo.On("GetRequestStatsMostCount").Return(&fakeReqStats2, nil)
	mockedPrRepo.On("GetParamRequestByUid", fakeReqStats2.ParamRequestUid).Return(nil, errors.New("error"))

	_, err := GetMaxFzRequestStat(mockedPrRepo, mockedRsRepo)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestMaxFzRequestOk(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)

	mockedRsRepo.On("GetRequestStatsMostCount").Return(&fakeReqStats2, nil)
	mockedPrRepo.On("GetParamRequestByUid", fakeReqStats2.ParamRequestUid).Return(&fakeFzParam2, nil)

	_, err := GetMaxFzRequestStat(mockedPrRepo, mockedRsRepo)

	assert.Nil(t, err)
}
