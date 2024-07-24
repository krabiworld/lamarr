package command

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
	"github.com/rs/zerolog/log"
	"module-go/internal/services"
	"module-go/internal/types"
)

type Context struct {
	e          *events.ApplicationCommandInteractionCreate
	commands   []Command
	categories []types.Category
	service    services.GuildService
	owner      snowflake.ID
}

func (ctx *Context) Reply(message string, ephemeral ...bool) error {
	return ctx.e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent(message).
		SetEphemeral(len(ephemeral) > 0 && ephemeral[0]).
		Build())
}

func (ctx *Context) ReplyEmbed(embed discord.Embed, ephemeral ...bool) error {
	return ctx.e.CreateMessage(discord.NewMessageCreateBuilder().
		SetEmbeds(embed).
		SetEphemeral(len(ephemeral) > 0 && ephemeral[0]).
		Build())
}

func (ctx *Context) ReplyError(message string) error {
	embed := discord.NewEmbedBuilder().SetDescription(message).SetColor(types.ColorError).Build()
	return ctx.ReplyEmbed(embed, true)
}

func (ctx *Context) Data() discord.SlashCommandInteractionData {
	return ctx.e.SlashCommandInteractionData()
}

func (ctx *Context) Option(key string) (discord.SlashCommandOption, bool) {
	opts := ctx.e.SlashCommandInteractionData().Options
	for _, opt := range opts {
		if opt.Name == key {
			return opt, true
		}
	}
	return discord.SlashCommandOption{}, false
}

func (ctx *Context) OptionAsUser(key string, defaultUser ...discord.User) (discord.User, bool) {
	user, ok := ctx.Data().OptUser(key)
	if ok {
		return user, true
	}

	if len(defaultUser) > 0 {
		return defaultUser[0], true
	}

	return discord.User{}, false
}

func (ctx *Context) OptionAsInt(key string, defaultNumber ...int64) (int64, bool) {
	integer, ok := ctx.Data().OptInt(key)
	if ok {
		return int64(integer), true
	}

	if len(defaultNumber) > 0 {
		return defaultNumber[0], true
	}

	return 0, false
}

func (ctx *Context) OptionAsString(key string, defaultString ...string) (string, bool) {
	str, ok := ctx.Data().OptString(key)
	if ok {
		return str, true
	}

	if len(defaultString) > 0 {
		return defaultString[0], true
	}

	return "", false
}

func (ctx *Context) Guild() (discord.Guild, error) {
	guild, ok := ctx.e.Guild()
	if !ok {
		restGuild, err := ctx.e.Client().Rest().GetGuild(*ctx.e.GuildID(), false)
		if err != nil {
			return discord.Guild{}, err
		}
		guild = restGuild.Guild
	}
	return guild, nil
}

func (ctx *Context) Channels() ([]discord.GuildChannel, error) {
	return ctx.e.Client().Rest().GetGuildChannels(*ctx.e.GuildID())
}

func (ctx *Context) Presence(id snowflake.ID) (discord.Presence, bool) {
	return ctx.e.Client().Caches().Presence(*ctx.e.GuildID(), id)
}

func (ctx *Context) Presences() []discord.Presence {
	presences := make([]discord.Presence, 0)
	ctx.e.Client().Caches().PresenceForEach(*ctx.e.GuildID(), func(presence discord.Presence) {
		presences = append(presences, presence)
	})
	return presences
}

func (ctx *Context) Members() ([]discord.Member, error) {
	return ctx.e.Client().MemberChunkingManager().RequestAllMembers(*ctx.e.GuildID())
}

func (ctx *Context) Member() discord.ResolvedMember {
	return *ctx.e.Member()
}

func (ctx *Context) MemberByID(id snowflake.ID) (discord.Member, error) {
	member, ok := ctx.e.Client().Caches().Member(*ctx.e.GuildID(), id)
	if !ok {
		restMember, err := ctx.e.Client().Rest().GetMember(*ctx.e.GuildID(), id)
		if err != nil {
			return discord.Member{}, err
		}

		member = *restMember
	}
	return member, nil
}

func (ctx *Context) User() discord.User {
	return ctx.e.User()
}

func (ctx *Context) Owner() bool {
	return ctx.User().ID == ctx.owner
}

func (ctx *Context) Moderator() bool {
	member := ctx.Member()

	if member.Permissions.Has(discord.PermissionAdministrator) || ctx.Owner() {
		return true
	}

	modRole, err := ctx.service.GetModRole(*ctx.e.GuildID())
	if err != nil {
		log.Error().Err(err).Send()
		return false
	}

	if modRole != nil {
		for _, role := range member.RoleIDs {
			if role.String() == *modRole {
				return true
			}
		}
	}

	return false
}

func (ctx *Context) Commands() []Command {
	return ctx.commands
}

func (ctx *Context) Categories() []types.Category {
	return ctx.categories
}
