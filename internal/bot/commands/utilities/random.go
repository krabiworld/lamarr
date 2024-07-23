package utilities

import (
	"math/rand"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/types"
	"strconv"
)

type RandomCommand struct{}

func NewRandomCommand() *command.Command {
	return command.New().
		Name("rand").
		Description("Get a random number").
		OptionInt("min", "Minimum number", false).
		OptionInt("max", "Maximum number", false).
		Category(types.CategoryUtilities).
		Handler(RandomCommand{}).
		Build()
}

func (cmd RandomCommand) Handle(ctx *command.Context) error {
	minNum, _ := ctx.OptionAsInt("min", 1)
	maxNum, _ := ctx.OptionAsInt("max", 100)

	if minNum > maxNum {
		return ctx.ReplyError("Min should be less than Max.")
	}

	return ctx.Reply(strconv.FormatInt(rand.Int63n(maxNum-minNum)+minNum, 10))
}
