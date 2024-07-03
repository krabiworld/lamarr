package handler

import "github.com/bwmarrin/discordgo"

type CommandContext struct {
	session *discordgo.Session
	event   *discordgo.InteractionCreate
	command *Command
}

func reply(c *CommandContext) {

}
