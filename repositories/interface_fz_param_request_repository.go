package repositories

import (
	"fizzbuzzlbc/database/models"
	"gorm.io/gorm"
)

// FzParamRequestRepository is an interface to manipulate request stats entity
type FzParamRequestRepository interface {
	GetParamRequestByUid(uid string) (*models.FzParamRequest, error)
	CreateParamRequest(models.FzParamRequest) error
}

type fzParamRequestRepository struct {
	db *gorm.DB
}

func NewFzParamRequestRepository(db *gorm.DB) FzParamRequestRepository {
	return &fzParamRequestRepository{
		db: db,
	}
}

func (r *fzParamRequestRepository) GetParamRequestByUid(uid string) (*models.FzParamRequest, error) {

	entityRequestParam := &models.FzParamRequest{}
	resDb := r.db.Take(entityRequestParam, "uid = ?", uid)

	return entityRequestParam, resDb.Error
}

func (r *fzParamRequestRepository) CreateParamRequest(data models.FzParamRequest) error {

	resDb := r.db.Create(&data)

	return resDb.Error
}
