package handlers

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/krabiworld/lamarr/internal/db/models"
	"github.com/krabiworld/lamarr/internal/services"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type GuildEvents struct {
	s services.GuildService
}

func NewGuildEvents(s services.GuildService) *GuildEvents {
	return &GuildEvents{s: s}
}

func (g *GuildEvents) OnGuildCreate(_ *discordgo.Session, e *discordgo.GuildCreate) {
	_, err := g.s.Get(e.ID)
	if err == nil {
		return
	}

	// if record not found - create, else log error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := g.s.Create(&models.Guild{ID: e.ID}); err != nil {
			log.Error().Err(err).Send()
		}
		return
	}

	log.Error().Err(err).Send()
}
