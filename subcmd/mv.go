package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/postname"
)

type MvCmd struct {
	Source    string `arg:"" help:"Source post name/category/tag."`
	Target    string `arg:"" help:"Target post/category."`
	Force     bool   `short:"f" default:"false" help:"Skip confirmation of files to move."`
	Page      int    `short:"p" default:"1" help:"Page number."`
	Recursive bool   `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *MvCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.ListOrTagSearch(cmd.Source, cmd.Page, cmd.Recursive)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	if err != nil {
		return err
	}

	targetCat, targetName := postname.Split(cmd.Target)

	if len(posts) > 1 && targetName != "" {
		return fmt.Errorf("target '%s' is not a category", cmd.Target)
	}

	movePost := &model.MovePostBody{
		Name:     targetName,
		Category: targetCat,
	}

	if !cmd.Force {
		for _, v := range posts {
			ctx.Fmt.Printf("mv '%s' '%s'\n", v.FullNameWithoutTags(), postname.Join(movePost.Category, movePost.Name))
		}

		if hasMore {
			ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
		}

		approval := prompter.YN("Do you want to move posts?", false)

		if !approval {
			ctx.Fmt.Println("Move cancelled.")
			return nil
		}
	}

	for _, v := range posts {
		if cmd.Force {
			ctx.Fmt.Printf("mv '%s' '%s'\n", v.FullNameWithoutTags(), postname.Join(movePost.Category, movePost.Name))
		}

		err = ctx.Driver.Move(movePost, v.Number)

		if err != nil {
			return fmt.Errorf("failed to move '%s':%w", v.FullNameWithoutTags(), err)
		}
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
