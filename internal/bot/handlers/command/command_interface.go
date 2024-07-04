package command

type ICommand interface {
	Handle(ctx *Context) error
}
