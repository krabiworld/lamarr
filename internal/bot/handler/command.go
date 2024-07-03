package handler

type Command struct {
	Name              string
	Description       string
	Category          Category
	OwnerCommand      bool
	ModerationCommand bool
	Hidden            bool
	Children          []Command
	// OptionData options
	// SubcommandData
}

func (c *Command) Run() {

}
