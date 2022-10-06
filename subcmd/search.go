package subcmd

import (
	"github.com/kanmu/kasa"
)

type SearchCmd struct {
	Query string `arg:"" help:"Search query. see https://docs.esa.io/posts/104"`
	Json  bool   `help:"Output as JSON"`
	Page  int    `short:"p" default:"1" help:"Page number."`
}

func (cmd *SearchCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.Search(cmd.Query, cmd.Page)

	if err != nil {
		return err
	}

	for _, v := range posts {
		if cmd.Json {
			out, err := v.Json()

			if err != nil {
				return err
			}

			ctx.Fmt.Println(string(out))
		} else {
			ctx.Fmt.Println(v.ListString())
		}
	}

	if hasMore && !cmd.Json {
		ctx.Fmt.Printf("(has more pages. current page is %d, try '-p %d')\n", cmd.Page, cmd.Page+1)
	}

	return nil
}
