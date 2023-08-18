package subcmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/postname"
)

type ImportCmd struct {
	File   string `arg:"" help:"Source file (stdin:'-')."`
	Path   string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
	Notice bool   `negatable:"" default:"true" help:"Post with notify."`
	Wip    bool   `negatable:"" help:"Post as WIP."`
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
		name = filepath.Base(cmd.File)
	}

	newPost := &model.NewPostBody{
		Name:     name,
		Category: cat,
		Wip:      &cmd.Wip,
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
