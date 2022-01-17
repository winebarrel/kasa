package subcmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/postname"
	"github.com/winebarrel/kasa/utils"
)

const (
	DefaultEditor = "vi"
)

type EditCmd struct {
	Path   string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
	Editor string `required:"" env:"EDITOR" help:"Editor to edit a post"`
	Notice bool   `negatable:"" default:"true" help:"Post with notify."`
}

func (cmd *EditCmd) Run(ctx *kasa.Context) error {
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

	newPost := &model.NewPostBody{}

	if post == nil {
		if num > 0 {
			return errors.New("post not found")
		} else {
			cat, name := postname.Split(cmd.Path)

			if name == "" {
				return errors.New("post name is empty")
			}

			post = &model.Post{}
			newPost.Name = name
			newPost.Category = cat
		}
	}

	tempDir, err := ioutil.TempDir("", "kasa")

	if err != nil {
		return err
	}

	srcBodyMd := strings.ReplaceAll(post.BodyMd, "\r\n", "\n")
	tempPost, err := createTempPost(tempDir, post.Number, srcBodyMd)

	if err != nil {
		return err
	}

	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = DefaultEditor
	}

	err = editTempPost(editor, tempPost)

	if err != nil {
		return err
	}

	rawBodyMd, err := ioutil.ReadFile(tempPost)
	bodyMd := string(rawBodyMd)

	if err != nil {
		return err
	}

	if bodyMd == srcBodyMd {
		return nil
	}

	newPost.BodyMd = bodyMd
	url, err := ctx.Driver.Post(newPost, post.Number, cmd.Notice)

	if err != nil {
		return err
	}

	ctx.Fmt.Println(url)

	return nil
}

func createTempPost(dir string, postNum int, postBody string) (string, error) {
	file := fmt.Sprintf("%s/%d.md", dir, postNum)
	bodyMd := strings.ReplaceAll(postBody, "\r\n", "\n")
	err := ioutil.WriteFile(file, []byte(bodyMd), 0600)

	if err != nil {
		return "", err
	}

	return file, nil
}

func editTempPost(editor string, file string) error {
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
