package subcmd

import (
	"fmt"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
)

type UnwipCmd struct {
	Path      string `arg:"" help:"Post name/Post category/Post tag."`
	Search    bool   `short:"s" help:"Search posts. see https://docs.esa.io/posts/104"`
	Force     bool   `short:"f" help:"Skip confirmation of files to move."`
	Notice    bool   `negatable:"" help:"Unwip with notify."`
	Page      int    `short:"p" default:"1" help:"Page number."`
	Recursive bool   `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *UnwipCmd) Run(ctx *kasa.Context) error {
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
	newPosts := make([]*model.WipPostBody, len(posts))
	hasWip := false

	for i, v := range posts {
		if v.Wip {
			hasWip = true
		}

		newPost := &model.WipPostBody{
			Wip:       false,
			UpdatedAt: v.UpdatedAt,
		}

		newPosts[i] = newPost
	}

	if !hasWip {
		ctx.Fmt.Println("WIP posts missing.")
		return nil
	}

	if !cmd.Force {
		for _, oldPost := range posts {
			if oldPost.Wip {
				ctx.Fmt.Printf("unwip '%s'\n", oldPost.FullNameWithoutTags())
			}
		}

		if hasMore {
			ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
		}

		approval := prompter.YN("Do you want to unwip posts?", false)

		if !approval {
			ctx.Fmt.Println("Unwip cancelled.")
			return nil
		}
	}

	for i, oldPost := range posts {
		if !oldPost.Wip {
			continue
		}

		newPost := newPosts[i]
		ctx.Fmt.Printf("unwip '%s'\n", oldPost.FullNameWithoutTags())
		err := ctx.Driver.Wip(newPost, oldPost.Number, cmd.Notice)

		if err != nil {
			return fmt.Errorf("failed to unwip '%s':%w", oldPost.FullNameWithoutTags(), err)
		}
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
