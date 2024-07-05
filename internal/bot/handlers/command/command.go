package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/types"
)

type Command struct {
	Name              string
	Description       string
	Category          types.Category
	OwnerCommand      bool
	ModerationCommand bool
	Hidden            bool
	Handler           ICommand
}

func (c *Command) Run(s *discordgo.Session, m *discordgo.MessageCreate, i *discordgo.InteractionCreate) {
	var guildId string

	if i == nil {
		guildId = m.GuildID
	} else {
		guildId = i.GuildID
	}

	ctx := &Context{
		Session: s,
		Message: m,
		Event:   i,
		GuildID: guildId,
	}

	if err := c.Handler.Handle(ctx); err != nil {
		log.Error().Err(err).Msg("Error executing command")
	}
}
