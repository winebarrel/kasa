package subcmd

import (
	"github.com/winebarrel/kasa"
)

type MvcatCmd struct {
	From string `arg:"" help:"From category."`
	To   string `arg:"" help:"To category."`
}

func (cmd *MvcatCmd) Run(ctx *kasa.Context) error {
	return ctx.Driver.MoveCategory(cmd.From, cmd.To)
}
