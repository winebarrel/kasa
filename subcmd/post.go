package subcmd

import (
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/utils"
)

type PostCmd struct {
	Name     string   `short:"n" help:"Post title."`
	Body     string   `short:"b" help:"Post body file." predictor:"file"`
	Path     string   `arg:"" optional:"" help:"Post number to update."`
	Tags     []string `short:"t" help:"Post tags."`
	Category string   `short:"c" help:"Post category."`
	Wip      bool     `negatable:"" help:"Post as WIP."`
	Message  string   `short:"m" help:"Post message."`
	Notice   bool     `negatable:"" help:"Post with notify."`
}

func (cmd *PostCmd) Run(ctx *kasa.Context) error {
	postNum := 0

	if cmd.Path == "" {
		if cmd.Name == "" {
			return errors.New("missing flags: --name=STRING")
		}

		if cmd.Body == "" {
			return errors.New("missing flags: --body=STRING")
		}
	} else {
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

		postNum = post.Number
	}

	var bodyMd []byte

	if cmd.Body != "" {
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

		var err error
		bodyMd, err = ioutil.ReadAll(file)

		if err != nil {
			return err
		}
	} else {
		bodyMd = []byte{}
	}

	newPost := &model.NewPostBody{
		Name:     cmd.Name,
		BodyMd:   string(bodyMd),
		Tags:     cmd.Tags,
		Category: cmd.Category,
		Wip:      esa.Bool(cmd.Wip),
		Message:  cmd.Message,
	}

	url, err := ctx.Driver.Post(newPost, postNum, cmd.Notice)

	if err != nil {
		return err
	}

	ctx.Fmt.Println(url)

	return nil
}
