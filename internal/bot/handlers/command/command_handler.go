package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
)

type Handler struct {
	Commands     map[string]*Command
	guildService services.GuildService
	ownerId      string
}

func NewCommandHandler(commands map[string]*Command, guildService services.GuildService, ownerId string) *Handler {
	return &Handler{
		Commands:     commands,
		ownerId:      ownerId,
		guildService: guildService,
	}
}

func (h *Handler) OnInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command, ok := h.Commands[i.ApplicationCommandData().Name]
	if !ok {
		return
	}

	ctx := &Context{
		session:      s,
		event:        i,
		command:      command,
		guildService: h.guildService,
		ownerId:      h.ownerId,
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
