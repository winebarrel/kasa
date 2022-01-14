package subcmd

import (
	"encoding/json"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/utils"
)

type InfoCmd struct {
	Path string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
}

func (cmd *InfoCmd) Run(ctx *kasa.Context) error {
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

	post.BodyMd = ""
	post.BodyHTML = ""
	rawJson, err := json.MarshalIndent(post, "", "  ")

	if err != nil {
		return err
	}

	ctx.Fmt.Println(string(rawJson))

	return nil
}
