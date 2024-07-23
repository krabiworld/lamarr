package information

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"strings"
	"time"
)

type ServerCommand struct{}

func NewServerCommand() *command.Command {
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

	presences := ctx.Presences()

	owner, err := ctx.MemberByID(guild.OwnerID)
	if err != nil {
		return err
	}

	createdAt := guild.ID.Time()

	e := discord.NewEmbedBuilder().
		SetTitle(fmt.Sprintf("Information about %s", guild.Name)).
		SetColor(types.ColorDefault.Int()).
		SetFooter(fmt.Sprintf("ID: %s", guild.ID), "").
		AddField(cmd.MembersField(members)).
		AddField(cmd.ChannelsField(channels)).
		AddField(cmd.StatusField(presences)).
		AddField(cmd.OwnerField(owner)).
		AddField(cmd.VerificationLevelField(guild)).
		AddField(cmd.CreatedAtField(createdAt))

	if guild.Icon != nil {
		e.SetThumbnail(*guild.IconURL())
	}

	if guild.Banner != nil {
		e.SetImage(*guild.BannerURL())
	}

	return ctx.ReplyEmbed(e.Build())
}

func (cmd ServerCommand) MembersField(members []discord.Member) (string, string, bool) {
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

func (cmd ServerCommand) ChannelsField(channels []discord.GuildChannel) (string, string, bool) {
	total, textChannels, voiceChannels, stageChannels := 0, 0, 0, 0

	for _, channel := range channels {
		switch channel.Type() {
		case discord.ChannelTypeGuildText:
			textChannels++
		case discord.ChannelTypeGuildVoice:
			voiceChannels++
		case discord.ChannelTypeGuildStageVoice:
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

func (cmd ServerCommand) StatusField(presences []discord.Presence) (string, string, bool) {
	online, idle, dnd, offline := 0, 0, 0, 0

	for _, presence := range presences {
		switch presence.Status {
		case discord.OnlineStatusOnline:
			online++
		case discord.OnlineStatusIdle:
			idle++
		case discord.OnlineStatusDND:
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

	name := "By Status"
	value := strings.TrimSpace(builder.String())
	return name, value, true
}

func (cmd ServerCommand) OwnerField(owner discord.Member) (string, string, bool) {
	return "Owner", owner.Mention(), true
}

func (cmd ServerCommand) VerificationLevelField(guild discord.Guild) (string, string, bool) {
	var verificationLevel string

	switch guild.VerificationLevel {
	case discord.VerificationLevelNone:
		verificationLevel = "None"
	case discord.VerificationLevelLow:
		verificationLevel = "Low"
	case discord.VerificationLevelMedium:
		verificationLevel = "Medium"
	case discord.VerificationLevelHigh:
		verificationLevel = "High"
	case discord.VerificationLevelVeryHigh:
		verificationLevel = "Very High"
	}

	return "Verification Level", verificationLevel, true
}

func (cmd ServerCommand) CreatedAtField(createdAt time.Time) (string, string, bool) {
	name := "Created At"
	value := fmt.Sprintf("<t:%[1]d:D> (<t:%[1]d:R>)", createdAt.Unix())
	return name, value, true
}
