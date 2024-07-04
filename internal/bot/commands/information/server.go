package information

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/colors"
	"strings"
	"time"
)

type ServerCommand struct{}

func (cmd *ServerCommand) Handle(ctx *command.Context) error {
	guild, err := ctx.Guild()
	if err != nil {
		return err
	}

	owner, err := ctx.GuildOwner(guild.OwnerID)
	if err != nil {
		return err
	}

	createdAt, err := discordgo.SnowflakeTimestamp(guild.ID)
	if err != nil {
		return err
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Information about %s", guild.Name),
		Color: colors.DEFAULT,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL("128"),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: guild.BannerURL("512"),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("ID: %s", guild.ID),
		},
		Fields: []*discordgo.MessageEmbedField{
			cmd.MembersField(guild),
			cmd.ChannelsField(guild),
			cmd.StatusField(guild),
			cmd.OwnerField(owner),
			cmd.VerificationLevelField(guild),
			cmd.CreatedAtField(createdAt),
		},
	}

	return ctx.ReplyEmbed(embed)
}

func (cmd *ServerCommand) MembersField(guild *discordgo.Guild) *discordgo.MessageEmbedField {
	botCount, memberCount := 0, 0

	for _, member := range guild.Members {
		if member.User.Bot {
			botCount++
		} else {
			memberCount++
		}
	}

	return &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("Members (%d)", guild.MemberCount),
		Value:  fmt.Sprintf("Members: **%d**\nBots: **%d**", memberCount, botCount),
		Inline: true,
	}
}

func (cmd *ServerCommand) ChannelsField(guild *discordgo.Guild) *discordgo.MessageEmbedField {
	total, textChannels, voiceChannels, stageChannels := 0, 0, 0, 0

	for _, channel := range guild.Channels {
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

	return &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("Channels (%d)", total),
		Value:  strings.TrimSpace(builder.String()),
		Inline: true,
	}
}

func (cmd *ServerCommand) StatusField(guild *discordgo.Guild) *discordgo.MessageEmbedField {
	online, idle, dnd, offline := 0, 0, 0, 0

	for _, presence := range guild.Presences {
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
		builder.WriteString(fmt.Sprintf("Online: **%d**\n", online))
	}

	if idle > 0 {
		builder.WriteString(fmt.Sprintf("Idle: **%d**\n", idle))
	}

	if dnd > 0 {
		builder.WriteString(fmt.Sprintf("Do Not Disturb: **%d**\n", dnd))
	}

	if offline > 0 {
		builder.WriteString(fmt.Sprintf("Offline: **%d**\n", offline))
	}

	return &discordgo.MessageEmbedField{
		Name:   "By Status",
		Value:  strings.TrimSpace(builder.String()),
		Inline: true,
	}
}

func (cmd *ServerCommand) OwnerField(owner *discordgo.Member) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   "Owner",
		Value:  owner.Mention(),
		Inline: true,
	}
}

func (cmd *ServerCommand) VerificationLevelField(guild *discordgo.Guild) *discordgo.MessageEmbedField {
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

	return &discordgo.MessageEmbedField{
		Name:   "Verification Level",
		Value:  verificationLevel,
		Inline: true,
	}
}

func (cmd *ServerCommand) CreatedAtField(createdAt time.Time) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   "Created At",
		Value:  fmt.Sprintf("<t:%[1]d:D> (<t:%[1]d:R>)", createdAt.Unix()),
		Inline: true,
	}
}
