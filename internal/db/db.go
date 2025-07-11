package db

import (
	"errors"
	"github.com/krabiworld/lamarr/internal/config"
	"github.com/krabiworld/lamarr/internal/db/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MustNew() *gorm.DB {
	dial, err := openDialector()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to open dialector")
	}

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Send()
		return nil
	}

	err = db.AutoMigrate(
		&models.Guild{},
		&models.Warn{},
		&models.Stats{},
	)
	if err != nil {
		log.Panic().Err(err).Send()
		return nil
	}

	return db
}

func openDialector() (gorm.Dialector, error) {
	switch config.Get().DatabaseType {
	case "postgres":
		return postgres.Open(config.Get().DatabaseDSN), nil
	case "sqlite":
		return sqlite.Open(config.Get().DatabaseDSN), nil
	default:
		return nil, errors.New("unsupported dialector type")
	}
}
