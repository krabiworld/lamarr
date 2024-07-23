package handlers

import (
	"errors"
	"github.com/disgoorg/disgo/events"
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

func (g *GuildEvents) OnGuildCreate(e events.GuildJoin) {
	_, err := g.s.Get(e.GuildID)
	if err == nil {
		return
	}

	// if record not found - create, else log error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := g.s.Create(&models.Guild{ID: e.GuildID.String()}); err != nil {
			log.Error().Err(err).Send()
		}
		return
	}

	log.Error().Err(err).Send()
}
