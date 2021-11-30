package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
)

type RmxCmd struct {
	Path  string `arg:"" help:"Source post name/category/tag."`
	Force bool   `short:"f" default:"false" help:"Skip confirmation of files to delete."`
	Page  int    `short:"p" default:"1" help:"Page number."`
}

func (cmd *RmxCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.ListOrTagSearch(cmd.Path, 1, true)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	if !cmd.Force {
		for _, v := range posts {
			fmt.Printf("rm '%s'\n", v.FullNameWithoutTags())
		}

		if hasMore {
			fmt.Println("(has more pages. Try increasing `-p NUM`)")
		}

		approval := prompter.YN("Do you want to delete posts?", false)

		if !approval {
			fmt.Println("Delete cancelled.")
			return nil
		}
	}

	for _, v := range posts {
		err = ctx.Driver.Delete(v.Number)

		if err != nil {
			if err != nil {
				return fmt.Errorf("failed to delete '%s':%w", v.FullNameWithoutTags(), err)
			}
		}
	}

	if hasMore {
		fmt.Println("(has more page. Try increasing `-p NUM`)")
	}

	return nil
}
