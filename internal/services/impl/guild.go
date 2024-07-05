package impl

import (
	"module-go/internal/db/models"
	"module-go/internal/repositories"
)

type GuildServiceImpl struct {
	r repositories.GuildRepository
}

func NewGuildServiceImpl(r repositories.GuildRepository) *GuildServiceImpl {
	return &GuildServiceImpl{r: r}
}

func (s *GuildServiceImpl) Get(id string) (*models.Guild, error) {
	return s.r.FindByID(id)
}

func (s *GuildServiceImpl) GetPrefix(id string) (string, error) {
	guild, err := s.Get(id)
	if err != nil {
		return "", err
	}

	return guild.Prefix, nil
}

func (s *GuildServiceImpl) Create(guild *models.Guild) error {
	return s.r.Create(guild)
}

func (s *GuildServiceImpl) Update(guild *models.Guild) error {
	return s.r.Update(guild)
}
