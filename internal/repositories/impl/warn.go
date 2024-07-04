package impl

import (
	"gorm.io/gorm"
	"module-go/internal/db/models"
)

type WarnRepositoryImpl struct {
	db *gorm.DB
}

func NewWarnRepository(db *gorm.DB) *WarnRepositoryImpl {
	return &WarnRepositoryImpl{db: db}
}

func (r *WarnRepositoryImpl) FindByID(id uint) (*models.Warn, error) {
	var warn *models.Warn
	if err := r.db.First(&warn, id).Error; err != nil {
		return nil, err
	}

	return warn, nil
}

func (r *WarnRepositoryImpl) FindAllByGuildAndMember(guildID string, memberID uint) ([]*models.Warn, error) {
	var warns []*models.Warn
	if err := r.db.Find(&warns, "guild_id = ? AND member_id = ?", guildID, memberID).Error; err != nil {
		return nil, err
	}

	return warns, nil
}
