package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"module-go/internal/cfg"
	"module-go/internal/db/models"
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
		log.Fatal().Err(err).Send()
	}

	err = db.AutoMigrate(
		&models.Guild{},
		&models.Warn{},
		&models.Stats{},
	)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return db
}
