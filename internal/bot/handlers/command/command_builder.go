package command

import (
	"github.com/disgoorg/disgo/discord"
	"module-go/internal/types"
)

type Builder struct {
	cmd *Command
}

func New() *Builder {
	return &Builder{cmd: &Command{
		ApplicationCommand: discord.SlashCommandCreate{},
	}}
}

func (b *Builder) Name(name string) *Builder {
	b.cmd.ApplicationCommand.Name = name
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.cmd.ApplicationCommand.Description = description
	return b
}

func (b *Builder) OptionUser(name, description string, required bool) *Builder {
	b.cmd.ApplicationCommand.Options = append(b.cmd.ApplicationCommand.Options, discord.ApplicationCommandOptionUser{
		Name:        name,
		Description: description,
		Required:    required,
	})
	return b
}

func (b *Builder) OptionInt(name, description string, required bool) *Builder {
	b.cmd.ApplicationCommand.Options = append(b.cmd.ApplicationCommand.Options, discord.ApplicationCommandOptionInt{
		Name:        name,
		Description: description,
		Required:    required,
	})
	return b
}

func (b *Builder) Category(category types.Category) *Builder {
	b.cmd.Category = category
	return b
}

func (b *Builder) Handler(handler ICommand) *Builder {
	b.cmd.Handler = handler
	return b
}

func (b *Builder) Build() *Command {
	return b.cmd
}
