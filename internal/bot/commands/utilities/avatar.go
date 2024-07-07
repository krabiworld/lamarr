package utilities

import (
	"fmt"
	"module-go/internal/bot/handlers/command"
	"module-go/pkg/embed"
)

type AvatarCommand struct{}

func (cmd *AvatarCommand) Handle(ctx *command.Context) error {
	user := ctx.OptionAsUser("user", ctx.User())

	e := embed.New().
		Author(fmt.Sprintf("Avatar of %s", user.Username), "").
		Color(user.AccentColor).
		Image(user.AvatarURL("1024")).
		Build()

	return ctx.Reply(e)
}
