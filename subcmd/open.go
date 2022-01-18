package subcmd

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/utils"
)

type OpenCmd struct {
	Path string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
}

func getPostUrl(ctx *kasa.Context, path string) (string, error) {
	num, err := utils.GetPostNum(ctx.Team, path)

	if err != nil {
		return "", err
	}

	var post *model.Post

	if num > 0 {
		post, err = ctx.Driver.GetFromPageNum(num)
	} else {
		post, err = ctx.Driver.Get(path)
	}

	if err != nil {
		return "", err
	}

	if post == nil {
		return "", errors.New("post not found")
	}

	return post.URL, nil
}

func (cmd *OpenCmd) Run(ctx *kasa.Context) error {
	var postUrl string
	var err error

	if strings.HasSuffix(cmd.Path, "/") {
		cat := strings.TrimSuffix(cmd.Path, "/")

		if !strings.HasPrefix(cat, "/") {
			cat = "/" + cat
		}

		cat = url.QueryEscape(cat)
		postUrl = fmt.Sprintf("https://%s.esa.io/#path=%s", ctx.Team, cat)
	} else {
		postUrl, err = getPostUrl(ctx, cmd.Path)
	}

	if err != nil {
		return err
	}

	return utils.OpenInBrowser(postUrl)
}
