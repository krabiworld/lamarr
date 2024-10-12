package repositories

import "github.com/krabiworld/lamarr/internal/db/models"

type StatsRepository interface {
	FindByID(id uint) (*models.Stats, error)
}
