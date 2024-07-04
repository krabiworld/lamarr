package services

import "module-go/internal/db/models"

type GuildService interface {
	Find(id uint) (*models.Guild, error)
	Create(guild *models.Guild) error
	Update(guild *models.Guild) error
}
