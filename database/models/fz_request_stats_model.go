package models

import "gorm.io/gorm"

// FzRequestStat contains statistics of a ParamRequest
type FzRequestStat struct {
	gorm.Model
	ParamRequestUid string
	Count           uint64
}
