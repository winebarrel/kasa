package subcmd

import (
	"encoding/json"

	"github.com/winebarrel/kasa"
)

type InfoCmd struct {
	PostNum int `arg:"" help:"Post number."`
}

func (cmd *InfoCmd) Run(ctx *kasa.Context) error {
	post, err := ctx.Driver.GetFromPageNum(cmd.PostNum)

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
