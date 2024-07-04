package repositories

import (
	"module-go/internal/db/models"
)

type WarnRepository interface {
	FindByID(id uint) (*models.Warn, error)
	FindAllByGuildAndMember(guildID string, memberID uint) ([]*models.Warn, error)
}
