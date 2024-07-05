package command

import "github.com/bwmarrin/discordgo"

type Context struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Event   *discordgo.InteractionCreate
	GuildID string
}

func (ctx *Context) Reply(embed *discordgo.MessageEmbed) error {
	if ctx.Event == nil {
		_, err := ctx.Session.ChannelMessageSendEmbedReply(
			ctx.Message.ChannelID,
			embed,
			ctx.Message.Reference(),
		)
		return err
	} else {
		return ctx.Session.InteractionRespond(ctx.Event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}
}

func (ctx *Context) Guild() (*discordgo.Guild, error) {
	guild, err := ctx.Session.State.Guild(ctx.GuildID)
	if err != nil {
		guild, err = ctx.Session.Guild(ctx.GuildID)
		if err != nil {
			return nil, err
		}
	}

	return guild, nil
}

func (ctx *Context) MemberByID(id string) (*discordgo.Member, error) {
	member, err := ctx.Session.State.Member(ctx.GuildID, id)
	if err != nil {
		member, err = ctx.Session.GuildMember(ctx.GuildID, id)
		if err != nil {
			return nil, err
		}
	}

	return member, nil
}

func (ctx *Context) GuildOwner(ownerId string) (*discordgo.Member, error) {
	return ctx.MemberByID(ownerId)
}
