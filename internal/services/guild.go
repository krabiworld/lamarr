package services

import (
	"github.com/disgoorg/snowflake/v2"
	"module-go/internal/db/models"
)

type GuildService interface {
	Get(id snowflake.ID) (*models.Guild, error)
	GetModRole(id snowflake.ID) (*string, error)
	Create(guild *models.Guild) error
	Update(guild *models.Guild) error
}
