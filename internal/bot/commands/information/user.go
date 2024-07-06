package information

import "module-go/internal/bot/handlers/command"

type UserCommand struct{}

func (cmd *UserCommand) Handle(ctx *command.Context) error {
	user := ctx.OptionAsUser("user", ctx.User())

	return ctx.ReplyError("Member name " + user.Mention())
}
