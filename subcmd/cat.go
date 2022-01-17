package subcmd

import (
	"errors"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/utils"
)

type CatCmd struct {
	Path string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
}

func (cmd *CatCmd) Run(ctx *kasa.Context) error {
	num, err := utils.GetPostNum(ctx.Team, cmd.Path)

	if err != nil {
		return err
	}

	var post *model.Post

	if num > 0 {
		post, err = ctx.Driver.GetFromPageNum(num)
	} else {
		post, err = ctx.Driver.Get(cmd.Path)
	}

	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post not found")
	}

	ctx.Fmt.Println(post.BodyMd)

	return nil
}
