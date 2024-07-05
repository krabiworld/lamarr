package handlers

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"module-go/internal/db/models"
	"module-go/internal/services"
)

type GuildEvents struct {
	s services.GuildService
}

func NewGuildEvents(s services.GuildService) *GuildEvents {
	return &GuildEvents{s: s}
}

func (e *GuildEvents) OnGuildCreate(_ *discordgo.Session, g *discordgo.GuildCreate) {
	_, err := e.s.Get(g.ID)
	if err == nil {
		return
	}

	// if record not found - create, else log error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := e.s.Create(&models.Guild{ID: g.ID}); err != nil {
			log.Error().Err(err).Send()
		}
		return
	}

	log.Error().Err(err).Send()
}
