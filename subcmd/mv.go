package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/esa/model"
)

type MvCmd struct {
	Source string `arg:"" help:"Source post name/category/tag."`
	Target string `arg:"" help:"Target post/category."`
	Force  bool   `short:"f" default:"false" help:"Skip confirmation of files to move."`
	Page   int    `short:"p" default:"1" help:"Page number."`
}

func (cmd *MvCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.ListOrTagSearch(cmd.Source, cmd.Page, true)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	if err != nil {
		return err
	}

	target := esa.NewPath(cmd.Target)

	if len(posts) > 1 && !target.IsCategory() {
		return fmt.Errorf("target '%s' is not a category", cmd.Target)
	}

	movePost := &model.MovePostBody{
		Name:     target.Name,
		Category: target.Category,
	}

	if !cmd.Force {
		for _, v := range posts {
			fmt.Printf("mv '%s' '%s'\n", v.FullNameWithoutTags(), movePost.String())
		}

		if hasMore {
			fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
		}

		approval := prompter.YN("Do you want to move posts?", false)

		if !approval {
			fmt.Println("Move cancelled.")
			return nil
		}
	}

	for _, v := range posts {
		err = ctx.Driver.Move(movePost, v.Number)

		if err != nil {
			return fmt.Errorf("failed to move '%s':%w", v.FullNameWithoutTags(), err)
		}
	}

	if hasMore {
		fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
