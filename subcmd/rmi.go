package subcmd

import (
	"github.com/winebarrel/kasa"
)

type RmiCmd struct {
	PostNum int `arg:"" help:"Post number to delete."`
}

func (cmd *RmiCmd) Run(ctx *kasa.Context) error {
	return ctx.Driver.Delete(cmd.PostNum)
}
