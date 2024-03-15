package helper

import (
	"errors"
	"fizzbuzzlbc/database/models"
	"fizzbuzzlbc/repositories"
	"fizzbuzzlbc/utils"
	"gorm.io/gorm"
)

// AddRequestStatisticHelper add statistics for fizzbuzz param in database
func AddRequestStatisticHelper(prRepo repositories.FzParamRequestRepository, rsRepo repositories.FzRequestStatsRepository, data models.FzParamRequest) error {

	data.Uid = utils.CreateRequestUid(data)

	_, err := prRepo.GetParamRequestByUid(data.Uid)
	if nil != err && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = prRepo.CreateParamRequest(data)
		if nil != err {
			return err
		}

		return rsRepo.CreateRequestStats(models.FzRequestStat{
			ParamRequestUid: data.Uid,
			Count:           1,
		})

	}

	entityRequestStats, err := rsRepo.GetRequestStatsByUid(data.Uid)
	if nil != err && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rsRepo.CreateRequestStats(models.FzRequestStat{
			ParamRequestUid: data.Uid,
			Count:           1,
		})
	}

	entityRequestStats.Count++

	return rsRepo.UpdateRequestStats(entityRequestStats)
}
