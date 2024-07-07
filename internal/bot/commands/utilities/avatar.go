package utilities

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"module-go/pkg/embed"
)

type AvatarCommand struct{}

func NewAvatarCommand() *command.Command {
	return command.New().
		Name("avatar").
		Description("User avatar").
		Option(
			discordgo.ApplicationCommandOptionUser,
			"user",
			"Specific user",
			false,
		).
		Category(types.CategoryUtilities).
		Handler(&AvatarCommand{}).
		Build()
}

func (cmd *AvatarCommand) Handle(ctx *command.Context) error {
	user := ctx.OptionAsUser("user", ctx.User())

	e := embed.New().
		Author(fmt.Sprintf("Avatar of %s", user.Username), "").
		Color(user.AccentColor).
		Image(user.AvatarURL("1024")).
		Build()

	return ctx.Reply(e)
}
