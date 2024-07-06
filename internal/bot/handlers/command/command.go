package command

import (
	"github.com/bwmarrin/discordgo"
	"module-go/internal/types"
)

type Command struct {
	ApplicationCommand *discordgo.ApplicationCommand
	Category           types.Category
	OwnerCommand       bool
	ModerationCommand  bool
	Hidden             bool
	Handler            ICommand
}
