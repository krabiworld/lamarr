package services

import (
	"github.com/krabiworld/lamarr/internal/db/models"
)

type GuildService interface {
	Get(id string) (*models.Guild, error)
	GetModRole(id string) (*string, error)
	Create(guild *models.Guild) error
	Update(guild *models.Guild) error
}
