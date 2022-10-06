package subcmd

import (
	"encoding/json"

	"github.com/kanmu/kasa"
)

type StatsCmd struct {
}

func (cmd *StatsCmd) Run(ctx *kasa.Context) error {
	stats, err := ctx.Driver.GetStats()

	if err != nil {
		return err
	}

	rawJson, err := json.MarshalIndent(stats, "", "  ")

	if err != nil {
		return err
	}

	ctx.Fmt.Println(string(rawJson))

	return nil
}
