package information

import (
	"errors"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"strings"
)

type UserCommand struct{}

func NewUserCommand() *command.Command {
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

	e := discord.NewEmbedBuilder().
		SetAuthor(user.Username, "", "").
		SetDescription(description).
		SetFooter("ID: "+user.ID.String(), "")

	if user.AccentColor != nil {
		e.SetColor(*user.AccentColor)
	}

	if user.Avatar != nil {
		e.SetAuthor(user.Username, "", *user.AvatarURL())
		e.SetThumbnail(*user.AvatarURL())
	}

	if len(member.RoleIDs) > 0 {
		e.AddField("Roles", cmd.Roles(member.RoleIDs), false)
	}

	if user.Banner != nil {
		e.SetImage(*user.BannerURL())
	}

	return ctx.ReplyEmbed(e.Build())
}

func (cmd UserCommand) Status(userStatus discord.OnlineStatus) string {
	var status string

	switch userStatus {
	case discord.OnlineStatusOnline:
		status = "Online"
	case discord.OnlineStatusIdle:
		status = "Idle"
	case discord.OnlineStatusOffline:
		status = "Do Not Disturb"
	default:
		status = "Offline"
	}

	return fmt.Sprintf("**Status:** %s", status)
}

func (cmd UserCommand) Activities(userActivities []discord.Activity) string {
	var builder strings.Builder

	for _, activity := range userActivities {
		switch activity.Type {
		case discord.ActivityTypeGame:
			builder.WriteString("**Playing:** ")
		case discord.ActivityTypeStreaming:
			builder.WriteString("**Streaming:** ")
		case discord.ActivityTypeListening:
			builder.WriteString("**Listening to:** ")
		case discord.ActivityTypeWatching:
			builder.WriteString("**Watching:** ")
		case discord.ActivityTypeCustom:
			builder.WriteString("**Custom:** ")
		case discord.ActivityTypeCompeting:
			builder.WriteString("**Competing to:** ")
		}
		builder.WriteString(activity.Name)
		builder.WriteString("\n")
	}

	return strings.TrimSuffix(builder.String(), "\n")
}

func (cmd UserCommand) JoinedAt(member discord.Member) string {
	return fmt.Sprintf("**Joined At:** <t:%[1]d:D> (<t:%[1]d:R>)", member.JoinedAt.Unix())
}

func (cmd UserCommand) CreatedAt(user discord.User) string {
	return fmt.Sprintf("**Created At:** <t:%[1]d:D> (<t:%[1]d:R>)", user.CreatedAt().Unix())
}

func (cmd UserCommand) Roles(roles []snowflake.ID) string {
	var builder strings.Builder
	for _, role := range roles {
		builder.WriteString(fmt.Sprintf("<@&%s> ", role.String()))
	}
	return strings.TrimSuffix(builder.String(), " ")
}
