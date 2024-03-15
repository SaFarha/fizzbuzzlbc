package helper

import (
	"fizzbuzzlbc/database/models"
	"fizzbuzzlbc/repositories"
)

// GetMaxFzRequestStat get the most request param used in fizzbuzz route
func GetMaxFzRequestStat(prRepo repositories.FzParamRequestRepository, rsRepo repositories.FzRequestStatsRepository) (*models.FzParamRequest, error) {

	value, err := rsRepo.GetRequestStatsMostCount()
	if err != nil {
		return nil, err
	}

	return prRepo.GetParamRequestByUid(value.ParamRequestUid)
}
