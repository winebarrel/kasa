package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
)

type RmsCmd struct {
	Phrase string `arg:"" help:"Search phrase."`
	Force  bool   `short:"f" default:"false" help:"Skip confirmation of files to delete."`
	Page   int    `short:"p" default:"1" help:"Page number."`
}

func (cmd *RmsCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.Search(cmd.Phrase, cmd.Page)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	if !cmd.Force {
		for _, v := range posts {
			ctx.Fmt.Printf("rm '%s'\n", v.FullNameWithoutTags())
		}

		if hasMore {
			ctx.Fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
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
		ctx.Fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
