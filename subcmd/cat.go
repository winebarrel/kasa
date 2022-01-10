package subcmd

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
)

type CatCmd struct {
	Path string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
}

func (cmd *CatCmd) Run(ctx *kasa.Context) error {
	num, err := cmd.getPostNum(ctx.Team)

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

func (cmd *CatCmd) getPostNum(team string) (int, error) {
	r := regexp.MustCompile(`(?:https://` + team + `\.esa\.io/posts/|//)(\d+)$`)
	m := r.FindStringSubmatch(cmd.Path)

	if len(m) != 2 {
		return 0, nil
	}

	return strconv.Atoi(m[1])
}
