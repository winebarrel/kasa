package subcmd

import (
	"errors"

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

	if post == nil {
		return errors.New("post not found")
	}

	ctx.Fmt.Println(post.BodyMd)

	return nil
}
