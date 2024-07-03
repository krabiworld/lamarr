package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type Command struct {
	Command           *discordgo.ApplicationCommand
	Category          Category
	OwnerCommand      bool
	ModerationCommand bool
	Hidden            bool
	Handler           ICommand
}

func (c *Command) Run(session *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx := &CommandContext{
		Session: session,
		Event:   i,
		command: c,
	}

	if err := c.Handler.Handle(ctx); err != nil {
		log.Error().Err(err).Msg("Error executing command")
	}
}
