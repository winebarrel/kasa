package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
)

type RmCmd struct {
	Path      string `arg:"" help:"Source post name/category/tag."`
	Force     bool   `short:"f" default:"false" help:"Skip confirmation of files to delete."`
	Page      int    `short:"p" default:"1" help:"Page number."`
	Recursive bool   `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *RmCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.ListOrTagSearch(cmd.Path, cmd.Page, cmd.Recursive)

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
		if cmd.Force {
			ctx.Fmt.Printf("rm '%s'\n", v.FullNameWithoutTags())
		}

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
