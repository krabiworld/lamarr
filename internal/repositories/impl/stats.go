package impl

import (
	"gorm.io/gorm"
	"module-go/internal/db/models"
)

type StatsRepositoryImpl struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) *StatsRepositoryImpl {
	return &StatsRepositoryImpl{db: db}
}

func (r *StatsRepositoryImpl) FindByID(id uint) (*models.Stats, error) {
	var stats *models.Stats
	if err := r.db.First(&stats, id).Error; err != nil {
		return nil, err
	}

	return stats, nil
}
