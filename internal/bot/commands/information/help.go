package information

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"strings"
)

type HelpCommand struct{}

func NewHelpCommand() command.Command {
	return command.New().
		Name("help").
		Description("List of all commands").
		OptionString("query", "Command or category", false).
		Category(types.CategoryInformation).
		Handler(HelpCommand{}).
		Build()
}

func (c HelpCommand) Handle(ctx *command.Context) error {
	query, ok := ctx.OptionAsString("query")
	commands := ctx.Commands()
	categories := ctx.Categories()

	embed := discord.NewEmbedBuilder().
		SetColor(types.ColorDefault)

	if ok {
		for _, category := range categories {
			if category == types.CategoryUnspecified {
				continue
			}

			if strings.Contains(strings.ToLower(category.String()), strings.ToLower(query)) {
				for _, cmd := range commands {
					if cmd.Hidden || category != cmd.Category {
						continue
					}

					embed.AddField(cmd.ApplicationCommand.Name, cmd.ApplicationCommand.Description, false)
				}

				embed.SetTitle(fmt.Sprintf("Commands of category %s", category.String()))
				return ctx.ReplyEmbed(embed.Build())
			}
		}

		for _, cmd := range commands {
			if !cmd.Hidden && strings.Contains(strings.ToLower(cmd.ApplicationCommand.Name), strings.ToLower(query)) {
				embed.AddField(cmd.ApplicationCommand.Name, cmd.ApplicationCommand.Description, false)
				embed.SetTitle(fmt.Sprintf("Information of command %s", cmd.ApplicationCommand.Name))
				return ctx.ReplyEmbed(embed.Build())
			}
		}

		return ctx.ReplyError(fmt.Sprintf("Command or category **%s** not found.", query))
	} else {
		embed.SetTitle("Available commands").
			SetDescription("For additional information enter `help category` to get information about category or `help command` to get information about command.")

		for _, category := range categories {
			if category == types.CategoryUnspecified {
				continue
			}

			builder := strings.Builder{}

			for _, cmd := range commands {
				if cmd.Hidden || category != cmd.Category {
					continue
				}

				builder.WriteString(fmt.Sprintf("`%s` ", cmd.ApplicationCommand.Name))
			}

			embed.AddField(fmt.Sprintf("%[1]s (help %[1]s)", category.String()), builder.String(), false)
		}

		return ctx.ReplyEmbed(embed.Build())
	}
}
