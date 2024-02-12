package subcmd

import (
	"errors"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/utils"
)

type RmiCmd struct {
	Path  string `arg:"" help:"Post name or Post URL('https://<TEAM>.esa.io/posts/<NUM>' or '//<NUM>')."`
	Force bool   `short:"f" help:"Skip confirmation of files to delete."`
}

func (cmd *RmiCmd) Run(ctx *kasa.Context) error {
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

	if !cmd.Force {
		ctx.Fmt.Printf("rm '%s'\n", post.FullNameWithoutTags())
		approval := prompter.YN("Do you want to delete a post?", false)

		if !approval {
			ctx.Fmt.Println("Delete cancelled.")
			return nil
		}
	}

	ctx.Fmt.Printf("rm '%s'\n", post.FullNameWithoutTags())

	return ctx.Driver.Delete(post.Number)
}
