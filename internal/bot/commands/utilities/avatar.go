package utilities

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
)

type AvatarCommand struct{}

func NewAvatarCommand() *command.Command {
	return command.New().
		Name("avatar").
		Description("User avatar").
		OptionUser("user", "Specific user", false).
		Category(types.CategoryUtilities).
		Handler(AvatarCommand{}).
		Build()
}

func (cmd AvatarCommand) Handle(ctx *command.Context) error {
	user, _ := ctx.OptionAsUser("user", ctx.User())

	e := discord.NewEmbedBuilder().
		SetAuthor(fmt.Sprintf("Avatar of %s", user.Username), "", "")

	if user.AccentColor != nil {
		e.SetColor(*user.AccentColor)
	}

	if user.Avatar != nil {
		e.SetImage(*user.AvatarURL())
	}

	return ctx.ReplyEmbed(e.Build())
}
