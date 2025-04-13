package subcmd

import (
	"fmt"
	pathpkg "path"
	"sort"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/postname"
)

type CpCmd struct {
	Source    string `arg:"" help:"Source post name/category/tag."`
	Target    string `arg:"" help:"Target post/category."`
	Force     bool   `short:"f" default:"false" help:"Skip confirmation of files to move."`
	Notice    bool   `negatable:"" help:"Copy with notify."`
	Page      int    `short:"p" default:"1" help:"Page number."`
	Recursive bool   `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *CpCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.ListOrTagSearch(cmd.Source, cmd.Page, cmd.Recursive)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })
	targetCat, targetName := postname.Split(cmd.Target)

	if len(posts) > 1 && targetName != "" {
		return fmt.Errorf("target '%s' is not a category", cmd.Target)
	}

	newPosts := make([]*model.NewPostBody, len(posts))
	withCat := postname.MinCategoryDepth(posts) + 1

	for i, v := range posts {
		newPost := &model.NewPostBody{
			Name:     v.Name,
			BodyMd:   v.BodyMd,
			Tags:     v.Tags,
			Category: postname.AppendCategoryN(targetCat, v.Category, withCat),
			Wip:      esa.Bool(v.Wip),
			Message:  v.Message,
		}

		if targetName != "" {
			newPost.Name = targetName
		}

		newPosts[i] = newPost
	}

	if !cmd.Force {
		for i, oldPost := range posts {
			newPost := newPosts[i]
			ctx.Fmt.Printf("cp '%s' '%s'\n", oldPost.FullNameWithoutTags(), postname.Join(newPost.Category, newPost.Name))
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

	for i, oldPost := range posts {
		newPost := newPosts[i]
		ctx.Fmt.Printf("cp '%s' '%s'\n", oldPost.FullNameWithoutTags(), postname.Join(newPost.Category, newPost.Name))
		url, err := ctx.Driver.Post(newPost, 0, cmd.Notice)

		if err != nil {
			return fmt.Errorf("failed to cp '%s':%w", oldPost.FullNameWithoutTags(), err)
		}

		urlDir := pathpkg.Dir(url)
		ctx.Fmt.Printf("%-*s  %s\n", len(urlDir)+9, url, postname.Join(newPost.Category, newPost.Name))
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
