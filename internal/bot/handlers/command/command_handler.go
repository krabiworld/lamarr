package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
)

type Handler struct {
	Commands map[string]*Command
	s        services.GuildService
}

func NewCommandHandler(commands map[string]*Command, s services.GuildService) *Handler {
	return &Handler{
		Commands: commands,
		s:        s,
	}
}

func (h *Handler) OnMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command, ok := h.Commands[i.ApplicationCommandData().Name]
	if !ok {
		return
	}

	ctx := &Context{
		session: s,
		event:   i,
		command: command,
	}

	go func() {
		if err := command.Handler.Handle(ctx); err != nil {
			log.Error().Err(err).Msg("Error executing command")
		}
	}()
}
