package handler

import "github.com/bwmarrin/discordgo"

type CommandHandler struct {
	OwnerID      string
	ForceGuildID *string
	Commands     map[string]Command
}

func (h *CommandHandler) OnReady(s *discordgo.Session, r *discordgo.Ready) {

}

func (h *CommandHandler) OnInteraction(session *discordgo.Session, msg *discordgo.InteractionCreate) {

}
