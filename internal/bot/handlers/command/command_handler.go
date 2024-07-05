package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
	"strings"
)

type Handler struct {
	Commands map[string]*Command
	s        services.GuildService
}

func NewCommandHandler(commands []*Command, s services.GuildService) *Handler {
	mapCommands := make(map[string]*Command, len(commands))

	for _, command := range commands {
		mapCommands[command.Name] = command
	}

	return &Handler{
		Commands: mapCommands,
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

	go command.Run(s, m, nil)
}

func (h *Handler) OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command, ok := h.Commands[i.ApplicationCommandData().Name]
	if !ok {
		return
	}

	go command.Run(s, nil, i)
}
