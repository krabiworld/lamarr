package information

import (
	"fmt"
	"github.com/krabiworld/lamarr/internal/bot/handlers/command"
	"github.com/krabiworld/lamarr/internal/types"
	"github.com/krabiworld/lamarr/internal/uptime"
	"github.com/krabiworld/lamarr/pkg/embed"
)

type StatsCommand struct{}

func NewStatsCommand() command.Command {
	return command.New().
		Name("stats").
		Description("Bot statistics").
		Category(types.CategoryInformation).
		Handler(StatsCommand{}).
		Build()
}

func (cmd StatsCommand) Handle(ctx *command.Context) error {
	guilds := ctx.State().Guilds

	channels := 0
	members := 0
	for _, server := range guilds {
		channels += len(server.Channels)
		members += server.MemberCount
	}

	main := fmt.Sprintf("**Servers:** %d\n**Users:** %d\n**Channels:** %d", len(guilds), members, channels)
	platform := fmt.Sprintf("**Ping:** %s\n**Uptime:** <t:%d:R>", ctx.Ping(), uptime.Get())

	selfUser := ctx.SelfUser()

	e := embed.New().
		Title("Bot statistics").
		Thumbnail(selfUser.AvatarURL("512")).
		Field("Main", main, true).
		Field("Platform", platform, true).
		Build()

	return ctx.ReplyEmbed(e)
}
