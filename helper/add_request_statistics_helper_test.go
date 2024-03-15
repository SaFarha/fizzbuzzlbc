package helper

import (
	"errors"
	"fizzbuzzlbc/database/models"
	"fizzbuzzlbc/mocks"
	"fizzbuzzlbc/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

var (
	fakeFzParam = models.FzParamRequest{
		Int1:  3,
		Int2:  5,
		Limit: 100,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	fakeUid      = utils.CreateRequestUid(fakeFzParam)
	fakeReqStats = models.FzRequestStat{
		ParamRequestUid: fakeUid,
		Count:           1,
	}
)

func TestFailGetPrByUid(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(nil, errors.New("error"))

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestCreatePrFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)
	fakeFzParam.Uid = fakeUid

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(nil, gorm.ErrRecordNotFound)
	mockedPrRepo.On("CreateParamRequest", fakeFzParam).Return(errors.New("error"))

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestCreateRequestStatFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)
	fakeFzParam.Uid = fakeUid

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(nil, gorm.ErrRecordNotFound)
	mockedPrRepo.On("CreateParamRequest", fakeFzParam).Return(nil)
	mockedRsRepo.On("CreateRequestStats", fakeReqStats).Return(errors.New("error"))

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestCreateRequestStatOk(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)
	fakeFzParam.Uid = fakeUid

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(nil, gorm.ErrRecordNotFound)
	mockedPrRepo.On("CreateParamRequest", fakeFzParam).Return(nil)
	mockedRsRepo.On("CreateRequestStats", fakeReqStats).Return(nil)

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.Nil(t, err)
}

func TestGetRequestStatsFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)
	fakeFzParam.Uid = fakeUid

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(&fakeFzParam, nil)
	mockedRsRepo.On("GetRequestStatsByUid", fakeUid).Return(nil, errors.New("error"))

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestGetRequestCreateStatsFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)
	fakeFzParam.Uid = fakeUid

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(&fakeFzParam, nil)
	mockedRsRepo.On("GetRequestStatsByUid", fakeUid).Return(nil, gorm.ErrRecordNotFound)
	mockedRsRepo.On("CreateRequestStats", fakeReqStats).Return(errors.New("error"))

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestGetRequestCreateStatsOk(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)
	fakeFzParam.Uid = fakeUid

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(&fakeFzParam, nil)
	mockedRsRepo.On("GetRequestStatsByUid", fakeUid).Return(nil, gorm.ErrRecordNotFound)
	mockedRsRepo.On("CreateRequestStats", fakeReqStats).Return(nil)

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.Nil(t, err)
}

func TestGetRequestUpdateStatsFailed(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(&fakeFzParam, nil)
	mockedRsRepo.On("GetRequestStatsByUid", fakeUid).Return(&fakeReqStats, nil)
	mockedRsRepo.On("UpdateRequestStats", &fakeReqStats).Return(errors.New("error"))

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestGetRequestUpdateStatsOk(t *testing.T) {
	mockedPrRepo := new(mocks.FzParamRequestRepository)
	mockedRsRepo := new(mocks.FzRequestStatsRepository)

	mockedPrRepo.On("GetParamRequestByUid", fakeUid).Return(&fakeFzParam, nil)
	mockedRsRepo.On("GetRequestStatsByUid", fakeUid).Return(&fakeReqStats, nil)
	mockedRsRepo.On("UpdateRequestStats", &fakeReqStats).Return(nil)

	err := AddRequestStatisticHelper(mockedPrRepo, mockedRsRepo, fakeFzParam)

	assert.Nil(t, err)
}
