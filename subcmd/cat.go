package subcmd

import (
	"fmt"

	"github.com/winebarrel/kasa"
)

type CatCmd struct {
	Path string `arg:"" optional:"" help:"Post name."`
}

func (cmd *CatCmd) Run(ctx *kasa.Context) error {
	post, err := ctx.Driver.Get(cmd.Path)

	if err != nil {
		return err
	}

	if post != nil {
		fmt.Println(post.BodyMd)
	} else {
		fmt.Printf("Page not found: %s\n", cmd.Path)
	}

	return nil
}
