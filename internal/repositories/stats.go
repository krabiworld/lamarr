package repositories

import "module-go/internal/db/models"

type StatsRepository interface {
	FindByID(id uint) (*models.Stats, error)
}
