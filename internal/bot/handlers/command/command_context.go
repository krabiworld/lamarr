package command

import "github.com/bwmarrin/discordgo"

type Context struct {
	Session *discordgo.Session
	Event   *discordgo.InteractionCreate
	command *Command
}

func (ctx *Context) Reply(message string) error {
	return ctx.Session.InteractionRespond(ctx.Event.Interaction, &discordgo.InteractionResponse{
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func (ctx *Context) ReplyEmbed(embed *discordgo.MessageEmbed) error {
	return ctx.Session.InteractionRespond(ctx.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func (ctx *Context) Guild() (*discordgo.Guild, error) {
	guild, err := ctx.Session.State.Guild(ctx.Event.GuildID)
	if err != nil {
		guild, err = ctx.Session.Guild(ctx.Event.GuildID)
		if err != nil {
			return nil, err
		}
	}

	return guild, nil
}

func (ctx *Context) MemberByID(id string) (*discordgo.Member, error) {
	member, err := ctx.Session.State.Member(ctx.Event.GuildID, id)
	if err != nil {
		member, err = ctx.Session.GuildMember(ctx.Event.GuildID, id)
		if err != nil {
			return nil, err
		}
	}

	return member, nil
}

func (ctx *Context) GuildOwner(ownerId string) (*discordgo.Member, error) {
	return ctx.MemberByID(ownerId)
}
