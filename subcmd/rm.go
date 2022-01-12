package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
)

type RmCmd struct {
	Path      string `arg:"" help:"Source post name/category/tag."`
	Search    bool   `short:"s" help:"Search posts. see https://docs.esa.io/posts/104"`
	Force     bool   `short:"f" help:"Skip confirmation of files to delete."`
	Page      int    `short:"p" default:"1" help:"Page number."`
	Recursive bool   `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *RmCmd) Run(ctx *kasa.Context) error {
	var posts []*model.Post
	var hasMore bool
	var err error

	if cmd.Search {
		posts, hasMore, err = ctx.Driver.Search(cmd.Path, cmd.Page)
	} else {
		posts, hasMore, err = ctx.Driver.ListOrTagSearch(cmd.Path, cmd.Page, cmd.Recursive)
	}

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	if !cmd.Force {
		for _, v := range posts {
			ctx.Fmt.Printf("rm '%s'\n", v.FullNameWithoutTags())
		}

		if hasMore {
			ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
		}

		approval := prompter.YN("Do you want to delete posts?", false)

		if !approval {
			ctx.Fmt.Println("Delete cancelled.")
			return nil
		}
	}

	for _, v := range posts {
		ctx.Fmt.Printf("rm '%s'\n", v.FullNameWithoutTags())
		err = ctx.Driver.Delete(v.Number)

		if err != nil {
			if err != nil {
				return fmt.Errorf("failed to delete '%s':%w", v.FullNameWithoutTags(), err)
			}
		}
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
