package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
	"strings"
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

func (h *Handler) OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || m.Author.System {
		return
	}

	prefix, err := h.s.GetPrefix(m.GuildID)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// remove prefix from message
	content := strings.TrimSpace(strings.TrimPrefix(m.Content, prefix))
	parts := strings.Split(content, " ")
	if len(parts) < 1 {
		return
	}

	command, ok := h.Commands[parts[0]]
	if !ok {
		return
	}

	ctx := &Context{
		Session: s,
		Message: m,
		Command: command,
	}

	if command.Arguments != nil && len(command.Arguments) > 0 {
		i := 0
		for key := range command.Arguments {
			i++

			if command.Arguments[key].Required && len(parts) < i+1 {
				_ = ctx.ReplyError(fmt.Sprintf("Agrument %s is required", key))
				return
			}

			command.Arguments[key].value = parts[i]
		}
	}

	go func() {
		if err := command.Handler.Handle(ctx); err != nil {
			log.Error().Err(err).Msg("Error executing command")
		}
	}()
}
