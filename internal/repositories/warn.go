package repositories

import (
	"github.com/krabiworld/lamarr/internal/db/models"
)

type WarnRepository interface {
	FindByID(id uint) (*models.Warn, error)
	FindAllByGuildAndMember(guildID string, memberID uint) ([]*models.Warn, error)
}
