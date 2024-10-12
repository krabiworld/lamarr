package utilities

import (
	"fmt"
	"github.com/krabiworld/lamarr/internal/bot/handlers/command"
	"github.com/krabiworld/lamarr/internal/types"
	"github.com/krabiworld/lamarr/pkg/embed"
)

type AvatarCommand struct{}

func NewAvatarCommand() command.Command {
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

	e := embed.New().
		Author(fmt.Sprintf("Avatar of %s", user.Username), "").
		Color(user.AccentColor).
		Image(user.AvatarURL("1024")).
		Build()

	return ctx.ReplyEmbed(e)
}
