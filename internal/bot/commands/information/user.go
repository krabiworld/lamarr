package information

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"module-go/pkg/embed"
	"strings"
)

type UserCommand struct{}

func NewUserCommand() *command.Command {
	return command.New().
		Name("user").
		Description("Information about user").
		Option(
			discordgo.ApplicationCommandOptionUser,
			"user",
			"Specific user",
			false,
		).
		Category(types.CategoryInformation).
		Handler(&UserCommand{}).
		Build()
}

func (cmd *UserCommand) Handle(ctx *command.Context) error {
	guild, err := ctx.Guild()
	if err != nil {
		return err
	}

	user := ctx.OptionAsUser("user", ctx.User())
	member, err := ctx.MemberByID(user.ID)
	if err != nil {
		return err
	}

	var userPresence *discordgo.Presence
	for _, presence := range guild.Presences {
		if presence.User.ID == user.ID {
			userPresence = presence
			break
		}
	}

	if userPresence == nil {
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
		Author(user.Username, user.AvatarURL("64")).
		Color(user.AccentColor).
		Description(description).
		Thumbnail(user.AvatarURL("512")).
		Footer("ID: " + user.ID)

	if len(member.Roles) > 0 {
		e.Field("Roles", cmd.Roles(member.Roles), false)
	}

	if user.Banner != "" {
		e.Image(user.BannerURL("512"))
	}

	return ctx.Reply(e.Build())
}

func (cmd *UserCommand) Status(userStatus discordgo.Status) string {
	var status string

	switch userStatus {
	case discordgo.StatusOnline:
		status = "Online"
	case discordgo.StatusIdle:
		status = "Idle"
	case discordgo.StatusDoNotDisturb:
		status = "Do Not Disturb"
	default:
		status = "Offline"
	}

	return fmt.Sprintf("**Status:** %s", status)
}

func (cmd *UserCommand) Activities(userActivities []*discordgo.Activity) string {
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

func (cmd *UserCommand) JoinedAt(member *discordgo.Member) string {
	return fmt.Sprintf("**Joined At:** <t:%[1]d:D> (<t:%[1]d:R>)", member.JoinedAt.Unix())
}

func (cmd *UserCommand) CreatedAt(user *discordgo.User) string {
	createdAt, err := discordgo.SnowflakeTimestamp(user.ID)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("**Created At:** <t:%[1]d:D> (<t:%[1]d:R>)", createdAt.Unix())
}

func (cmd *UserCommand) Roles(roles []string) string {
	var builder strings.Builder
	for _, role := range roles {
		builder.WriteString(fmt.Sprintf("<@&%s> ", role))
	}
	return strings.TrimSuffix(builder.String(), " ")
}
