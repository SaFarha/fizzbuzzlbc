package repositories

import (
	"fizzbuzzlbc/database/models"
	"gorm.io/gorm"
)

// FzRequestStatsRepository is an interface to manipulate request stats entity
type FzRequestStatsRepository interface {
	CreateRequestStats(models.FzRequestStat) error
	GetRequestStatsByUid(uid string) (*models.FzRequestStat, error)
	UpdateRequestStats(*models.FzRequestStat) error
	GetRequestStatsMostCount() (*models.FzRequestStat, error)
}

type fzRequestStatsRepository struct {
	db *gorm.DB
}

func NewRequestStatsRepository(db *gorm.DB) FzRequestStatsRepository {
	return &fzRequestStatsRepository{
		db: db,
	}
}

func (r *fzRequestStatsRepository) CreateRequestStats(data models.FzRequestStat) error {

	resDb := r.db.Create(&data)

	return resDb.Error
}

func (r *fzRequestStatsRepository) GetRequestStatsByUid(uid string) (*models.FzRequestStat, error) {
	entityRequestStats := &models.FzRequestStat{}
	resDb := r.db.Take(entityRequestStats, "param_request_uid = ?", uid)

	return entityRequestStats, resDb.Error
}

func (r *fzRequestStatsRepository) UpdateRequestStats(data *models.FzRequestStat) error {
	return r.db.Save(data).Error
}

func (r *fzRequestStatsRepository) GetRequestStatsMostCount() (*models.FzRequestStat, error) {
	entityRequestStats := &models.FzRequestStat{}
	res := r.db.Order("count DESC").Take(entityRequestStats)

	return entityRequestStats, res.Error
}
