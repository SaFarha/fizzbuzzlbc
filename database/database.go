package database

import (
	"fizzbuzzlbc/configuration"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Database struct {
	Db *gorm.DB
}

func InitDatabase(config configuration.Configuration) (*Database, error) {
	connection := fmt.Sprintf(
		`host=%s port=%d user=%s password='%s' dbname=%s sslmode=%s`,
		config.PostgresqlHost,
		config.PostgresqlPort,
		config.PostgresqlUser,
		config.PostgresqlPassword,
		config.PostgresqlDbName,
		config.PostgresqlSsl)

	db, err := gorm.Open(
		postgres.Open(connection),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Warn,
					Colorful:      true,
				},
			),
		},
	)
	if err != nil {
		return nil, err
	}

	// ping db
	database, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
