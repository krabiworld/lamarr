package command

import (
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
)

type Handler struct {
	Commands     map[string]*Command
	guildService services.GuildService
	ownerId      snowflake.ID
}

func NewHandler(commands []*Command, guildService services.GuildService, ownerId snowflake.ID) *Handler {
	m := make(map[string]*Command)
	for _, command := range commands {
		m[command.ApplicationCommand.Name] = command
	}

	return &Handler{
		Commands:     m,
		ownerId:      ownerId,
		guildService: guildService,
	}
}

func (h *Handler) OnInteractionCreate(event *events.ApplicationCommandInteractionCreate) {
	command, ok := h.Commands[event.Data.CommandName()]
	if !ok {
		return
	}

	ctx := &Context{
		e:       event,
		service: h.guildService,
		owner:   h.ownerId,
	}

	if command.OwnerCommand && !ctx.Owner() {
		return
	}

	if command.ModerationCommand && !ctx.Moderator() {
		return
	}

	go func() {
		if err := command.Handler.Handle(ctx); err != nil {
			log.Error().Err(err).Msg("Error executing command")
		}
	}()
}
