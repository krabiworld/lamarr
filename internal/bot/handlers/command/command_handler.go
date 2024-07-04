package command

import (
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	Commands map[string]*Command
}

func NewCommandHandler(commands []*Command) *Handler {
	mapCommands := make(map[string]*Command, len(commands))

	for _, command := range commands {
		mapCommands[command.Command.Name] = command
	}

	return &Handler{
		Commands: mapCommands,
	}
}

func (h *Handler) OnInteraction(session *discordgo.Session, i *discordgo.InteractionCreate) {
	command, ok := h.Commands[i.ApplicationCommandData().Name]
	if !ok {
		return
	}

	go command.Run(session, i)
}
