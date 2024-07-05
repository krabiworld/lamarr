package command

import (
	"module-go/internal/types"
)

type Argument struct {
	Name        string
	Description string
	Required    bool
	value       string
}

type Command struct {
	Name              string
	Description       string
	Category          types.Category
	OwnerCommand      bool
	ModerationCommand bool
	Hidden            bool
	Arguments         map[string]*Argument
	Handler           ICommand
}
