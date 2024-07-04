package handler

import "github.com/bwmarrin/discordgo"

type CommandContext struct {
	Session *discordgo.Session
	Event   *discordgo.InteractionCreate
	command *Command
}

func (ctx *CommandContext) Reply(message string) error {
	return ctx.Session.InteractionRespond(ctx.Event.Interaction, &discordgo.InteractionResponse{
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func (ctx *CommandContext) ReplyEmbed(embed *discordgo.MessageEmbed) error {
	return ctx.Session.InteractionRespond(ctx.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func (ctx *CommandContext) Guild() (*discordgo.Guild, error) {
	guild, err := ctx.Session.State.Guild(ctx.Event.GuildID)
	if err != nil {
		guild, err = ctx.Session.Guild(ctx.Event.GuildID)
		if err != nil {
			return nil, err
		}
	}

	return guild, nil
}

func (ctx *CommandContext) MemberByID(id string) (*discordgo.Member, error) {
	member, err := ctx.Session.State.Member(ctx.Event.GuildID, id)
	if err != nil {
		member, err = ctx.Session.GuildMember(ctx.Event.GuildID, id)
		if err != nil {
			return nil, err
		}
	}

	return member, nil
}

func (ctx *CommandContext) GuildOwner(ownerId string) (*discordgo.Member, error) {
	return ctx.MemberByID(ownerId)
}
