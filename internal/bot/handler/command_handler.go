package handler

import (
	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	OwnerID      string
	ForceGuildID string
	Commands     map[string]*Command
}

func NewCommandHandler(ownerId, forceGuildId string, commands []*Command) *CommandHandler {
	mapCommands := make(map[string]*Command, len(commands))

	for _, command := range commands {
		mapCommands[command.Command.Name] = command
	}

	return &CommandHandler{
		OwnerID:      ownerId,
		ForceGuildID: forceGuildId,
		Commands:     mapCommands,
	}
}

func (h *CommandHandler) OnInteraction(session *discordgo.Session, i *discordgo.InteractionCreate) {
	command, ok := h.Commands[i.ApplicationCommandData().Name]
	if !ok {
		return
	}

	go command.Run(session, i)
}
