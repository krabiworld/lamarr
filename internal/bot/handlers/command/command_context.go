package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
	"module-go/internal/types"
	"module-go/pkg/embed"
)

type Context struct {
	session      *discordgo.Session
	event        *discordgo.InteractionCreate
	command      *Command
	guildService services.GuildService
	ownerId      string
}

func (ctx *Context) Reply(embed *discordgo.MessageEmbed) error {
	return ctx.session.InteractionRespond(ctx.event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func (ctx *Context) ReplyError(message string) error {
	return ctx.Reply(embed.New().Description(message).Color(types.ERROR.Int()).Build())
}

func (ctx *Context) Option(key string) *discordgo.ApplicationCommandInteractionDataOption {
	opts := ctx.event.ApplicationCommandData().Options
	for _, opt := range opts {
		if opt.Name == key {
			return opt
		}
	}
	return nil
}

func (ctx *Context) OptionAsUser(key string, defaultUser ...*discordgo.User) *discordgo.User {
	opt := ctx.Option(key)
	if opt != nil {
		return opt.UserValue(ctx.session)
	}

	if len(defaultUser) > 0 {
		return defaultUser[0]
	}

	return nil
}

func (ctx *Context) Guild() (*discordgo.Guild, error) {
	guild, err := ctx.session.State.Guild(ctx.event.GuildID)
	if err != nil {
		guild, err = ctx.session.Guild(ctx.event.GuildID)
	}
	return guild, err
}

func (ctx *Context) Member() *discordgo.Member {
	return ctx.event.Member
}

func (ctx *Context) User() *discordgo.User {
	return ctx.event.Member.User
}

func (ctx *Context) MemberByID(id string) (*discordgo.Member, error) {
	member, err := ctx.session.State.Member(ctx.event.GuildID, id)
	if err != nil {
		member, err = ctx.session.GuildMember(ctx.event.GuildID, id)
	}
	return member, err
}

func (ctx *Context) GuildOwner(ownerId string) (*discordgo.Member, error) {
	return ctx.MemberByID(ownerId)
}

func (ctx *Context) Owner() bool {
	return ctx.User().ID == ctx.ownerId
}

func (ctx *Context) Moderator() bool {
	member := ctx.Member()

	if ctx.HasPermission(member.Permissions, discordgo.PermissionAdministrator) || ctx.Owner() {
		return true
	}

	modRole, err := ctx.guildService.GetModRole(ctx.event.GuildID)
	if err != nil {
		log.Error().Err(err).Send()
		return false
	}

	if modRole != nil {
		for _, role := range member.Roles {
			if role == *modRole {
				return true
			}
		}
	}

	return false
}

func (ctx *Context) HasPermission(dst int64, src int) bool {
	return dst&int64(src) != 0
}
