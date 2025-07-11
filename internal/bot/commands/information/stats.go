package information

import (
	"fmt"
	"github.com/krabiworld/lamarr/internal/bot/handlers/command"
	"github.com/krabiworld/lamarr/internal/types"
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
	//servers := ctx.Caches().GuildsLen()
	//members := ctx.Caches().MembersAllLen()
	//channels := ctx.Caches().ChannelsLen()
	selfUser := ctx.SelfUser()

	//main := fmt.Sprintf("**Servers:** %d\n**Users:** %d\n**Channels:** %d", servers, members, channels)
	platform := fmt.Sprintf("**Ping:** %s", ctx.Ping())

	e := embed.New().
		Title("Bot statistics").
		Thumbnail(selfUser.AvatarURL("512")).
		Field("Main", "main", true).
		Field("Platform", platform, true).
		Build()

	return ctx.ReplyEmbed(e)
}
