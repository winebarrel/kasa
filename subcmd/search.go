package subcmd

import (
	"sort"

	"github.com/winebarrel/kasa"
)

type SearchCmd struct {
	Query string `arg:"" help:"Search query."`
	Page  int    `short:"p" default:"1" help:"Page number."`
}

func (cmd *SearchCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.Search(cmd.Query, cmd.Page)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	for _, v := range posts {
		ctx.Fmt.Println(v.ListString())
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
