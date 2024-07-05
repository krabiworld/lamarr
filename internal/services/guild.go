package services

import "module-go/internal/db/models"

type GuildService interface {
	Get(id string) (*models.Guild, error)
	GetPrefix(id string) (string, error)
	Create(guild *models.Guild) error
	Update(guild *models.Guild) error
}
