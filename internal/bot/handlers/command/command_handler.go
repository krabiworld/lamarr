package command

import (
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
	"module-go/internal/types"
)

type Handler struct {
	CommandsMap  map[string]Command
	CommandsList []Command
	Categories   []types.Category
	guildService services.GuildService
	ownerId      snowflake.ID
}

func NewHandler(commands []Command, categories []types.Category, guildService services.GuildService, ownerId snowflake.ID) *Handler {
	m := make(map[string]Command, len(commands))
	l := make([]Command, len(commands))
	for i, command := range commands {
		m[command.ApplicationCommand.Name] = command
		l[i] = command
	}

	return &Handler{
		CommandsMap:  m,
		CommandsList: l,
		Categories:   categories,
		guildService: guildService,
		ownerId:      ownerId,
	}
}

func (h *Handler) OnInteractionCreate(event *events.ApplicationCommandInteractionCreate) {
	command, ok := h.CommandsMap[event.Data.CommandName()]
	if !ok {
		return
	}

	ctx := &Context{
		e:          event,
		commands:   h.CommandsList,
		categories: h.Categories,
		service:    h.guildService,
		owner:      h.ownerId,
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
