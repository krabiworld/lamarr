package information

import "module-go/internal/bot/handlers/command"

type UserCommand struct{}

func (cmd *UserCommand) Handle(ctx *command.Context) error {
	member, err := ctx.ArgAsMember("user")
	if err != nil {
		member = ctx.Message.Member
		member.User = ctx.Message.Author
	}

	return ctx.ReplyError("Member name " + member.DisplayName())
}
