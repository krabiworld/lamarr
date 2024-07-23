package impl

import (
	"github.com/disgoorg/snowflake/v2"
	"module-go/internal/db/models"
	"module-go/internal/repositories"
)

type GuildServiceImpl struct {
	r repositories.GuildRepository
}

func NewGuildServiceImpl(r repositories.GuildRepository) *GuildServiceImpl {
	return &GuildServiceImpl{r: r}
}

func (s *GuildServiceImpl) Get(id snowflake.ID) (*models.Guild, error) {
	return s.r.FindByID(id.String())
}

func (s *GuildServiceImpl) GetModRole(id snowflake.ID) (*string, error) {
	guild, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	return guild.Mod, nil
}

func (s *GuildServiceImpl) Create(guild *models.Guild) error {
	return s.r.Create(guild)
}

func (s *GuildServiceImpl) Update(guild *models.Guild) error {
	return s.r.Update(guild)
}
