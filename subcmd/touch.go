package subcmd

import (
	"errors"
	"fmt"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/postname"
)

type TouchCmd struct {
	Path   string `arg:"" help:"Post name."`
	Notice bool   `negatable:"" help:"Post with notify."`
}

func (cmd *TouchCmd) Run(ctx *kasa.Context) error {
	post, err := ctx.Driver.Get(cmd.Path)

	if err != nil {
		return err
	}

	if post != nil {
		return fmt.Errorf("post already exists: %s", post.URL)
	}

	newPost := &model.NewPostBody{}
	cat, name := postname.Split(cmd.Path)

	if name == "" {
		return errors.New("post name is empty")
	}

	newPost.Name = name
	newPost.Category = cat
	url, err := ctx.Driver.Post(newPost, 0, cmd.Notice)

	if err != nil {
		return err
	}

	ctx.Fmt.Println(url)

	return nil
}
