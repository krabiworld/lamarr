package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
	"module-go/internal/types"
)

type Handler struct {
	CommandsMap  map[string]Command
	CommandsList []Command
	Categories   []types.Category
	guildService services.GuildService
	ownerId      string
}

func NewHandler(commands []Command, categories []types.Category, guildService services.GuildService, ownerId string) *Handler {
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

func (h *Handler) OnInteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	command, ok := h.CommandsMap[interaction.ApplicationCommandData().Name]
	if !ok {
		return
	}

	ctx := &Context{
		session:     session,
		interaction: interaction,
		commands:    h.CommandsList,
		categories:  h.Categories,
		service:     h.guildService,
		owner:       h.ownerId,
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
