package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	{
		ID: "202303151200_create_param_request_table",
		Migrate: func(db *gorm.DB) error {
			type FzParamRequest struct {
				gorm.Model

				Uid   string `gorm:"index"`
				Int1  uint64
				Int2  uint64
				Limit uint64
				Str1  string
				Str2  string
			}

			return db.AutoMigrate(&FzParamRequest{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("fz_param_requests")
		},
	},
	{
		ID: "202303151205_create_request_stats_table",
		Migrate: func(db *gorm.DB) error {
			type FzRequestStat struct {
				gorm.Model
				ParamRequestUid string `gorm:"index"`
				Count           uint64
			}

			return db.AutoMigrate(&FzRequestStat{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("fz_request_stats")
		},
	},
}
