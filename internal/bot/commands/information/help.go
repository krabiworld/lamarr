package information

import (
	"fmt"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"module-go/pkg/embed"
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

	e := embed.New().Color(types.ColorDefault)

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

					e.Field(cmd.ApplicationCommand.Name, cmd.ApplicationCommand.Description, false)
				}

				e.Title(fmt.Sprintf("Commands of category %s", category.String()))
				return ctx.ReplyEmbed(e.Build())
			}
		}

		for _, cmd := range commands {
			if !cmd.Hidden && strings.Contains(strings.ToLower(cmd.ApplicationCommand.Name), strings.ToLower(query)) {
				e.Field(cmd.ApplicationCommand.Name, cmd.ApplicationCommand.Description, false)
				e.Title(fmt.Sprintf("Information of command %s", cmd.ApplicationCommand.Name))
				return ctx.ReplyEmbed(e.Build())
			}
		}

		return ctx.ReplyError(fmt.Sprintf("Command or category **%s** not found.", query))
	} else {
		e.Title("Available commands").
			Description("For additional information enter `help category` to get information about category or `help command` to get information about command.")

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

			e.Field(fmt.Sprintf("%[1]s (help %[1]s)", category.String()), builder.String(), false)
		}

		return ctx.ReplyEmbed(e.Build())
	}
}
