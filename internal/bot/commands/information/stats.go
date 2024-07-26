package information

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
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
	servers := ctx.Caches().GuildsLen()
	members := ctx.Caches().MembersAllLen()
	channels := ctx.Caches().ChannelsLen()
	selfUser, _ := ctx.Caches().SelfUser()
	ping := ctx.Gateway().Latency().String()

	main := fmt.Sprintf("**Servers:** %d\n**Users:** %d\n**Channels:** %d", servers, members, channels)
	platform := fmt.Sprintf("**Ping:** %s", ping)

	embed := discord.NewEmbedBuilder().
		SetTitle("Bot statistics").
		SetColor(types.ColorDefault).
		AddField("Main", main, true).
		AddField("Platform", platform, true)

	if selfUser.Avatar != nil {
		embed.SetThumbnail(*selfUser.AvatarURL())
	}

	return ctx.ReplyEmbed(embed.Build())
}
