package impl

import (
	"github.com/krabiworld/lamarr/internal/db/models"
	"gorm.io/gorm"
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

func (r *GuildRepositoryImpl) Create(guild *models.Guild) error {
	return r.db.Create(guild).Error
}

func (r *GuildRepositoryImpl) Update(guild *models.Guild) error {
	return r.db.Save(guild).Error
}
