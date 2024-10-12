package db

import (
	"fmt"
	"github.com/krabiworld/lamarr/internal/cfg"
	"github.com/krabiworld/lamarr/internal/db/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitAndGet() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Get().DBHostname,
		cfg.Get().DBUsername,
		cfg.Get().DBPassword,
		cfg.Get().DBDatabase,
		cfg.Get().DBPort,
	)

	config := &gorm.Config{
		Logger: Logger{},
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Error().Err(err).Send()
		return nil
	}

	err = db.AutoMigrate(
		&models.Guild{},
		&models.Warn{},
		&models.Stats{},
	)
	if err != nil {
		log.Error().Err(err).Send()
		return nil
	}

	return db
}
