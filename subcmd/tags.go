package subcmd

import (
	"github.com/winebarrel/kasa"
)

type TagsCmd struct {
	Page int `short:"p" default:"1" help:"Page number."`
}

func (cmd *TagsCmd) Run(ctx *kasa.Context) error {
	tags, hasMore, err := ctx.Driver.GetTags(cmd.Page)

	if err != nil {
		return err
	}

	for _, t := range tags.Tags {
		ctx.Fmt.Printf("%9d  %s\n", t.PostsCount, t.Name)
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
