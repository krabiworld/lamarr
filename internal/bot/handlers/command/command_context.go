package command

import (
	"github.com/bwmarrin/discordgo"
	"module-go/internal/types"
)

type Context struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Command *Command
}

func (ctx *Context) Reply(embed *discordgo.MessageEmbed) error {
	_, err := ctx.Session.ChannelMessageSendEmbedReply(ctx.Message.ChannelID, embed, ctx.Message.Reference())
	return err
}

func (ctx *Context) ReplyError(message string) error {
	embed := &discordgo.MessageEmbed{Description: message, Color: types.ERROR.Int()}
	_, err := ctx.Session.ChannelMessageSendEmbedReply(ctx.Message.ChannelID, embed, ctx.Message.Reference())
	return err
}

func (ctx *Context) Arg(key string) string {
	return ctx.Command.Arguments[key].value
}

func (ctx *Context) Guild() (*discordgo.Guild, error) {
	guild, err := ctx.Session.State.Guild(ctx.Message.GuildID)
	if err != nil {
		guild, err = ctx.Session.Guild(ctx.Message.GuildID)
		if err != nil {
			return nil, err
		}
	}

	return guild, nil
}

func (ctx *Context) MemberByID(id string) (*discordgo.Member, error) {
	member, err := ctx.Session.State.Member(ctx.Message.GuildID, id)
	if err != nil {
		member, err = ctx.Session.GuildMember(ctx.Message.GuildID, id)
		if err != nil {
			return nil, err
		}
	}

	return member, nil
}

func (ctx *Context) GuildOwner(ownerId string) (*discordgo.Member, error) {
	return ctx.MemberByID(ownerId)
}
