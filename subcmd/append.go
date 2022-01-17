package subcmd

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/utils"
)

type AppendCmd struct {
	Path   string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
	Body   string `short:"b" required:"" help:"Post body file to append." predictor:"file"`
	Prefix string `help:"Prefix text."`
	Notice bool   `negatable:"" default:"true" help:"Post with notify."`
}

func (cmd *AppendCmd) Run(ctx *kasa.Context) error {
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

	bodys := []string{post.BodyMd}

	if cmd.Prefix != "" {
		bodys = append(bodys, cmd.Prefix)
	}

	var file io.ReadCloser

	if cmd.Body == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.OpenFile(cmd.Body, os.O_RDONLY, 0)

		if err != nil {
			return err
		}

		defer file.Close()
	}

	bodyMd, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	bodys = append(bodys, string(bodyMd))

	newPost := &model.NewPostBody{
		BodyMd: strings.Join(bodys, "\n"),
	}

	url, err := ctx.Driver.Post(newPost, post.Number, cmd.Notice)

	if err != nil {
		return err
	}

	ctx.Fmt.Println(url)

	return nil
}
