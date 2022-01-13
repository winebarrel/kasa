package subcmd

import (
	"errors"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
)

type OpenCmd struct {
	Path string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
}

func openInBrowser(u string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", u).Start()
	case "linux":
		return exec.Command("xdg-open", u).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", u).Start()
	default:
		return fmt.Errorf("open browser failed: unsupported platform: %s", runtime.GOOS)
	}
}

func (cmd *OpenCmd) Run(ctx *kasa.Context) error {
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

	if flag.Lookup("test.v") == nil {
		if err := openInBrowser(post.URL); err != nil {
			return err
		}
	}
	return nil
}

func (cmd *OpenCmd) getPostNum(team string) (int, error) {
	r := regexp.MustCompile(`(?:https://` + team + `\.esa\.io/posts/|//)(\d+)$`)
	m := r.FindStringSubmatch(cmd.Path)

	if len(m) != 2 {
		return 0, nil
	}

	return strconv.Atoi(m[1])
}
