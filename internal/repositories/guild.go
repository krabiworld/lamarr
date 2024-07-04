package repositories

import "module-go/internal/db/models"

type GuildRepository interface {
	FindByID(id string) (*models.Guild, error)
}
