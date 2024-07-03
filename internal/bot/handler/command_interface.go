package handler

type ICommand interface {
	Handle(ctx *CommandContext) error
}
