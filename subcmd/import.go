package subcmd

import (
	"errors"
	"io"
	"os"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/postname"
)

type ImportCmd struct {
	File   string `arg:"" help:"Source file (stdin:'-')."`
	Path   string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
	Notice bool   `negatable:"" default:"true" help:"Post with notify."`
}

func (cmd *ImportCmd) Run(ctx *kasa.Context) error {
	var file io.Reader

	if cmd.File == "-" {
		file = os.Stdin
	} else {
		f, err := os.OpenFile(cmd.File, os.O_RDONLY, 0)

		if err != nil {
			return err
		}

		defer f.Close()
		file = f
	}

	cat, name := postname.Split(cmd.Path)

	if name == "" {
		return errors.New("post name is empty")
	}

	newPost := &model.NewPostBody{
		Name:     name,
		Category: cat,
	}

	bodyMd, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	newPost.BodyMd = string(bodyMd)
	url, err := ctx.Driver.Post(newPost, 0, cmd.Notice)

	if err != nil {
		return err
	}

	ctx.Fmt.Println(url)

	return nil
}
