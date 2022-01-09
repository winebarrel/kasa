package subcmd

import (
	"encoding/json"
	"fmt"

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

	fmt.Println(string(rawJson))

	return nil
}
