package command

import (
	"github.com/disgoorg/disgo/discord"
	"module-go/internal/types"
)

type Command struct {
	ApplicationCommand discord.SlashCommandCreate
	Category           types.Category
	OwnerCommand       bool
	ModerationCommand  bool
	Hidden             bool
	Handler            ICommand
}
