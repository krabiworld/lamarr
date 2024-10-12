package information

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/krabiworld/lamarr/internal/bot/handlers/command"
	"github.com/krabiworld/lamarr/internal/types"
	"github.com/krabiworld/lamarr/pkg/embed"
	"strings"
)

type UserCommand struct{}

func NewUserCommand() command.Command {
	return command.New().
		Name("user").
		Description("Information about user").
		OptionUser("user", "Specific user", false).
		Category(types.CategoryInformation).
		Handler(UserCommand{}).
		Build()
}

func (cmd UserCommand) Handle(ctx *command.Context) error {
	user, _ := ctx.OptionAsUser("user", ctx.User())
	member, err := ctx.MemberByID(user.ID)
	if err != nil {
		return err
	}

	userPresence, ok := ctx.Presence(user.ID)
	if !ok {
		return errors.New("user presence not found")
	}

	description := fmt.Sprintf(
		"**Username:** %s\n%s\n%s\n%s\n%s",
		user.Mention(),
		cmd.Status(userPresence.Status),
		cmd.Activities(userPresence.Activities),
		cmd.JoinedAt(member),
		cmd.CreatedAt(user),
	)

	e := embed.New().
		Author(user.Username, "").
		Color(user.AccentColor).
		Description(description).
		Footer("ID: " + user.ID)

	if user.Avatar != "" {
		e.Author(user.Username, user.AvatarURL("512")).Thumbnail(user.AvatarURL("512"))
	}

	if len(member.Roles) > 0 {
		e.Field("Roles", cmd.Roles(member.Roles), false)
	}

	if user.Banner != "" {
		e.Image(user.BannerURL("1024"))
	}

	return ctx.ReplyEmbed(e.Build())
}

func (cmd UserCommand) Status(userStatus discordgo.Status) string {
	var status string

	switch userStatus {
	case discordgo.StatusOnline:
		status = types.EmojiOnline + "Online"
	case discordgo.StatusIdle:
		status = types.EmojiIdle + "Idle"
	case discordgo.StatusOffline:
		status = types.EmojiOffline + "Do Not Disturb"
	default:
		status = types.EmojiOffline + "Offline"
	}

	return fmt.Sprintf("**Status:** %s", status)
}

func (cmd UserCommand) Activities(userActivities []*discordgo.Activity) string {
	var builder strings.Builder

	for _, activity := range userActivities {
		switch activity.Type {
		case discordgo.ActivityTypeGame:
			builder.WriteString("**Playing:** ")
		case discordgo.ActivityTypeStreaming:
			builder.WriteString("**Streaming:** ")
		case discordgo.ActivityTypeListening:
			builder.WriteString("**Listening to:** ")
		case discordgo.ActivityTypeWatching:
			builder.WriteString("**Watching:** ")
		case discordgo.ActivityTypeCustom:
			builder.WriteString("**Custom:** ")
		case discordgo.ActivityTypeCompeting:
			builder.WriteString("**Competing to:** ")
		}
		builder.WriteString(activity.Name)
		builder.WriteString("\n")
	}

	return strings.TrimSuffix(builder.String(), "\n")
}

func (cmd UserCommand) JoinedAt(member *discordgo.Member) string {
	return fmt.Sprintf("**Joined At:** <t:%[1]d:D> (<t:%[1]d:R>)", member.JoinedAt.Unix())
}

func (cmd UserCommand) CreatedAt(user *discordgo.User) string {
	createdAt, err := discordgo.SnowflakeTimestamp(user.ID)
	if err != nil {
		return "invalid id"
	}

	return fmt.Sprintf("**Created At:** <t:%[1]d:D> (<t:%[1]d:R>)", createdAt.Unix())
}

func (cmd UserCommand) Roles(roles []string) string {
	var builder strings.Builder
	for _, role := range roles {
		builder.WriteString(fmt.Sprintf("<@&%s> ", role))
	}
	return strings.TrimSuffix(builder.String(), " ")
}
