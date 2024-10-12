package information

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"module-go/pkg/embed"
	"strings"
	"time"
)

type ServerCommand struct{}

func NewServerCommand() command.Command {
	return command.New().
		Name("server").
		Description("Information about server").
		Category(types.CategoryInformation).
		Handler(ServerCommand{}).
		Build()
}

func (cmd ServerCommand) Handle(ctx *command.Context) error {
	guild, err := ctx.Guild()
	if err != nil {
		return err
	}

	members, err := ctx.Members()
	if err != nil {
		return err
	}

	channels, err := ctx.Channels()
	if err != nil {
		return err
	}

	presences, err := ctx.Presences()
	if err != nil {
		return err
	}

	owner, err := ctx.MemberByID(guild.OwnerID)
	if err != nil {
		return err
	}

	createdAt, err := discordgo.SnowflakeTimestamp(guild.ID)
	if err != nil {
		return err
	}

	e := embed.New().
		Title(fmt.Sprintf("Information about %s", guild.Name)).
		Color(types.ColorDefault).
		Footer(fmt.Sprintf("ID: %s", guild.ID)).
		Field(cmd.MembersField(members)).
		Field(cmd.ChannelsField(channels)).
		Field(cmd.StatusField(presences)).
		Field(cmd.OwnerField(owner)).
		Field(cmd.VerificationLevelField(guild)).
		Field(cmd.CreatedAtField(createdAt))

	if guild.Icon != "" {
		e.Thumbnail(guild.IconURL("512"))
	}

	if guild.Banner != "" {
		e.Image(guild.BannerURL("1024"))
	}

	return ctx.ReplyEmbed(e.Build())
}

func (cmd ServerCommand) MembersField(members []*discordgo.Member) (string, string, bool) {
	botCount, memberCount := 0, 0

	for _, member := range members {
		if member.User.Bot {
			botCount++
		} else {
			memberCount++
		}
	}

	name := fmt.Sprintf("Members (%d)", len(members))
	value := fmt.Sprintf("Members: **%d**\nBots: **%d**", memberCount, botCount)
	return name, value, true
}

func (cmd ServerCommand) ChannelsField(channels []*discordgo.Channel) (string, string, bool) {
	total, textChannels, voiceChannels, stageChannels := 0, 0, 0, 0

	for _, channel := range channels {
		switch channel.Type {
		case discordgo.ChannelTypeGuildText:
			textChannels++
		case discordgo.ChannelTypeGuildVoice:
			voiceChannels++
		case discordgo.ChannelTypeGuildStageVoice:
			stageChannels++
		default:
			continue
		}

		total++
	}

	var builder strings.Builder

	if textChannels > 0 {
		builder.WriteString(fmt.Sprintf("Text: **%d**\n", textChannels))
	}

	if voiceChannels > 0 {
		builder.WriteString(fmt.Sprintf("Voice: **%d**\n", voiceChannels))
	}

	if stageChannels > 0 {
		builder.WriteString(fmt.Sprintf("Stage: **%d**\n", stageChannels))
	}

	name := fmt.Sprintf("Channels (%d)", total)
	value := strings.TrimSpace(builder.String())
	return name, value, true
}

func (cmd ServerCommand) StatusField(presences []*discordgo.Presence) (string, string, bool) {
	online, idle, dnd, offline := 0, 0, 0, 0

	for _, presence := range presences {
		switch presence.Status {
		case discordgo.StatusOnline:
			online++
		case discordgo.StatusIdle:
			idle++
		case discordgo.StatusDoNotDisturb:
			dnd++
		default:
			offline++
		}
	}

	var builder strings.Builder

	if online > 0 {
		builder.WriteString(fmt.Sprintf("%sOnline: **%d**\n", types.EmojiOnline, online))
	}

	if idle > 0 {
		builder.WriteString(fmt.Sprintf("%sIdle: **%d**\n", types.EmojiIdle, idle))
	}

	if dnd > 0 {
		builder.WriteString(fmt.Sprintf("%sDo Not Disturb: **%d**\n", types.EmojiDnd, dnd))
	}

	if offline > 0 {
		builder.WriteString(fmt.Sprintf("%sOffline: **%d**\n", types.EmojiOffline, offline))
	}

	name := "By Status"
	value := strings.TrimSpace(builder.String())
	return name, value, true
}

func (cmd ServerCommand) OwnerField(owner *discordgo.Member) (string, string, bool) {
	return "Owner", owner.Mention(), true
}

func (cmd ServerCommand) VerificationLevelField(guild *discordgo.Guild) (string, string, bool) {
	var verificationLevel string

	switch guild.VerificationLevel {
	case discordgo.VerificationLevelNone:
		verificationLevel = "None"
	case discordgo.VerificationLevelLow:
		verificationLevel = "Low"
	case discordgo.VerificationLevelMedium:
		verificationLevel = "Medium"
	case discordgo.VerificationLevelHigh:
		verificationLevel = "High"
	case discordgo.VerificationLevelVeryHigh:
		verificationLevel = "Very High"
	}

	return "Verification Level", verificationLevel, true
}

func (cmd ServerCommand) CreatedAtField(createdAt time.Time) (string, string, bool) {
	name := "Created At"
	value := fmt.Sprintf("<t:%[1]d:D> (<t:%[1]d:R>)", createdAt.Unix())
	return name, value, true
}
