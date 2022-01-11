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
	Search    bool   `short:"s" help:"Search posts. see https://docs.esa.io/posts/104"`
	Force     bool   `short:"f" help:"Skip confirmation of files to move."`
	WithCat   int    `short:"n" help:"Move with category."`
	Page      int    `short:"p" default:"1" help:"Page number."`
	Recursive bool   `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *MvCmd) Run(ctx *kasa.Context) error {
	var posts []*model.Post
	var hasMore bool
	var err error

	if cmd.Search {
		posts, hasMore, err = ctx.Driver.Search(cmd.Source, cmd.Page)
	} else {
		posts, hasMore, err = ctx.Driver.ListOrTagSearch(cmd.Source, cmd.Page, cmd.Recursive)
	}

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

	movePosts := make([]*model.MovePostBody, len(posts))

	for i, v := range posts {
		movePost := &model.MovePostBody{
			Name:     targetName,
			Category: postname.AppendCategoryN(targetCat, v.Category, cmd.WithCat),
		}

		movePosts[i] = movePost
	}

	if !cmd.Force {
		for i, v := range posts {
			movePost := movePosts[i]
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

	for i, v := range posts {
		movePost := movePosts[i]
		ctx.Fmt.Printf("mv '%s' '%s'\n", v.FullNameWithoutTags(), postname.Join(movePost.Category, movePost.Name))
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
