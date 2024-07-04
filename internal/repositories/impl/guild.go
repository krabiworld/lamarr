package impl

import (
	"gorm.io/gorm"
	"module-go/internal/db/models"
)

type GuildRepositoryImpl struct {
	db *gorm.DB
}

func NewGuildRepository(db *gorm.DB) *GuildRepositoryImpl {
	return &GuildRepositoryImpl{db: db}
}

func (r *GuildRepositoryImpl) FindByID(id string) (*models.Guild, error) {
	var guild *models.Guild
	if err := r.db.First(&guild, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return guild, nil
}
