package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/krabiworld/lamarr/internal/services"
	"github.com/krabiworld/lamarr/internal/types"
	"github.com/krabiworld/lamarr/pkg/embed"
	"github.com/rs/zerolog/log"
	"time"
)

type Context struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
	commands    []Command
	categories  []types.Category
	service     services.GuildService
	owner       string
}

func (ctx *Context) Reply(message string, ephemeral ...bool) error {
	data := &discordgo.InteractionResponseData{
		Content: message,
	}

	if len(ephemeral) > 0 && ephemeral[0] {
		data.Flags = discordgo.MessageFlagsEphemeral
	}

	return ctx.session.InteractionRespond(ctx.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
}

func (ctx *Context) ReplyEmbed(embed *discordgo.MessageEmbed, ephemeral ...bool) error {
	data := &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{embed},
	}

	if len(ephemeral) > 0 && ephemeral[0] {
		data.Flags = discordgo.MessageFlagsEphemeral
	}

	return ctx.session.InteractionRespond(ctx.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
}

func (ctx *Context) ReplyError(message string) error {
	return ctx.ReplyEmbed(embed.New().Description(message).Color(types.ColorError).Build(), true)
}

func (ctx *Context) Option(key string) *discordgo.ApplicationCommandInteractionDataOption {
	opts := ctx.interaction.ApplicationCommandData().Options
	for _, opt := range opts {
		if opt.Name == key {
			return opt
		}
	}
	return nil
}

func (ctx *Context) OptionAsUser(key string, defaultUser ...*discordgo.User) (*discordgo.User, bool) {
	opt := ctx.Option(key)
	if opt != nil {
		return opt.UserValue(ctx.session), true
	}

	if len(defaultUser) > 0 {
		return defaultUser[0], true
	}

	return nil, false
}

func (ctx *Context) OptionAsInt(key string, defaultNumber ...int64) (int64, bool) {
	opt := ctx.Option(key)
	if opt != nil {
		return opt.IntValue(), true
	}

	if len(defaultNumber) > 0 {
		return defaultNumber[0], true
	}

	return 0, false
}

func (ctx *Context) OptionAsString(key string, defaultString ...string) (string, bool) {
	opt := ctx.Option(key)
	if opt != nil {
		return opt.StringValue(), true
	}

	if len(defaultString) > 0 {
		return defaultString[0], true
	}

	return "", false
}

func (ctx *Context) Guild() (*discordgo.Guild, error) {
	guild, err := ctx.session.State.Guild(ctx.interaction.GuildID)
	if err != nil {
		guild, err = ctx.session.Guild(ctx.interaction.GuildID)
	}
	return guild, err
}

func (ctx *Context) Channel() (*discordgo.Channel, error) {
	channel, err := ctx.session.State.Channel(ctx.interaction.ChannelID)
	if err != nil {
		channel, err = ctx.session.Channel(ctx.interaction.ChannelID)
	}
	return channel, err
}

func (ctx *Context) Channels() ([]*discordgo.Channel, error) {
	guild, err := ctx.Guild()
	if err != nil {
		return nil, err
	}

	return guild.Channels, nil
}

func (ctx *Context) Presence(id string) (*discordgo.Presence, bool) {
	guild, err := ctx.Guild()
	if err != nil {
		return nil, false
	}

	for _, presence := range guild.Presences {
		if presence.User.ID == id {
			return presence, true
		}
	}

	return nil, false
}

func (ctx *Context) Presences() ([]*discordgo.Presence, error) {
	guild, err := ctx.Guild()
	if err != nil {
		return nil, err
	}

	return guild.Presences, nil
}

func (ctx *Context) Members() ([]*discordgo.Member, error) {
	guild, err := ctx.Guild()
	if err != nil {
		return nil, err
	}

	return guild.Members, nil
}

func (ctx *Context) Member() *discordgo.Member {
	return ctx.interaction.Member
}

func (ctx *Context) MemberByID(id string) (*discordgo.Member, error) {
	member, err := ctx.session.State.Member(ctx.interaction.GuildID, id)
	if err != nil {
		member, err = ctx.session.GuildMember(ctx.interaction.GuildID, id)
	}
	return member, err
}

func (ctx *Context) User() *discordgo.User {
	return ctx.Member().User
}

func (ctx *Context) Owner() bool {
	return ctx.User().ID == ctx.owner
}

func (ctx *Context) Moderator() bool {
	member := ctx.Member()

	if ctx.HasPermission(member.Permissions, discordgo.PermissionAdministrator) || ctx.Owner() {
		return true
	}

	modRole, err := ctx.service.GetModRole(ctx.interaction.GuildID)
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

func (ctx *Context) SelfUser() *discordgo.User {
	return ctx.session.State.User
}

func (ctx *Context) State() *discordgo.State {
	return ctx.session.State
}

func (ctx *Context) Ping() string {
	return ctx.session.HeartbeatLatency().Round(time.Millisecond).String()
}

func (ctx *Context) Commands() []Command {
	return ctx.commands
}

func (ctx *Context) Categories() []types.Category {
	return ctx.categories
}

func (ctx *Context) HasPermission(dst int64, src int) bool {
	return dst&int64(src) != 0
}
