package models

import "gorm.io/gorm"

// FzParamRequest is a model in database that represent a unique param request
type FzParamRequest struct {
	gorm.Model
	Uid   string
	Int1  uint64
	Int2  uint64
	Limit uint64
	Str1  string
	Str2  string
}
