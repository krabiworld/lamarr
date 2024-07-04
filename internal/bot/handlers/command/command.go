package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/types"
)

type Command struct {
	Command           *discordgo.ApplicationCommand
	Category          types.Category
	OwnerCommand      bool
	ModerationCommand bool
	Hidden            bool
	Handler           ICommand
}

func (c *Command) Run(session *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx := &Context{
		Session: session,
		Event:   i,
		command: c,
	}

	if err := c.Handler.Handle(ctx); err != nil {
		log.Error().Err(err).Msg("Error executing command")
	}
}
