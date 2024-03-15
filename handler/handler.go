package handler

import (
	"fizzbuzzlbc/database"
	"fizzbuzzlbc/repositories"
)

// Handlers contains tools uses by handler func
type Handlers struct {
	ParamRequestRepo repositories.FzParamRequestRepository

	RequestStatRepo repositories.FzRequestStatsRepository
}

// NewHandlers creates a new Handlers object
func NewHandlers(db *database.Database) Handlers {

	return Handlers{
		ParamRequestRepo: repositories.NewFzParamRequestRepository(db.Db),
		RequestStatRepo:  repositories.NewRequestStatsRepository(db.Db),
	}

}
