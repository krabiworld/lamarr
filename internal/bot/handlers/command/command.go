package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/krabiworld/lamarr/internal/types"
)

type Command struct {
	ApplicationCommand *discordgo.ApplicationCommand
	Category           types.Category
	OwnerCommand       bool
	ModerationCommand  bool
	Hidden             bool
	Handler            ICommand
}
