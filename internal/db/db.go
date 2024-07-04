package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"module-go/internal/cfg"
	"module-go/internal/db/models"
)

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Get().DBHostname,
		cfg.Get().DBUsername,
		cfg.Get().DBPassword,
		cfg.Get().DBDatabase,
		cfg.Get().DBPort,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = conn.AutoMigrate(
		&models.Guild{},
		&models.Warn{},
		&models.Stats{},
	)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	db = conn
}

func Get() *gorm.DB {
	return db
}
