package subcmd

import (
	"github.com/winebarrel/kasa"
)

type RmCmd struct {
	PostNum int `arg:"" help:"Post number to delete."`
}

func (cmd *RmCmd) Run(ctx *kasa.Context) error {
	return ctx.Driver.Delete(cmd.PostNum)
}
