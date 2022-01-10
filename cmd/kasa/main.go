package main

import (
	"github.com/alecthomas/kong"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/subcmd"
)

var version string

var cli struct {
	Version kong.VersionFlag
	Team    string           `required:"" env:"ESA_TEAM" help:"esa team"`
	Token   string           `required:"" env:"ESA_TOKEN" help:"esa access token"`
	Debug   bool             `help:"Debug flag."`
	Cat     subcmd.CatCmd    `cmd:"" help:"Print post."`
	Cp      subcmd.CpCmd     `cmd:"" help:"Copy posts."`
	Info    subcmd.InfoCmd   `cmd:"" help:"Show post info."`
	Ls      subcmd.LsCmd     `cmd:"" help:"List posts."`
	Mv      subcmd.MvCmd     `cmd:"" help:"Move posts."`
	Mvcat   subcmd.MvcatCmd  `cmd:"" help:"Move category."`
	Mvs     subcmd.MvsCmd    `cmd:"" help:"Move posts from search results."`
	Post    subcmd.PostCmd   `cmd:"" help:"New/Update post."`
	Rm      subcmd.RmCmd     `cmd:"" help:"Delete posts."`
	Rmi     subcmd.RmiCmd    `cmd:"" help:"Delete post by number."`
	Rms     subcmd.RmsCmd    `cmd:"" help:"Delete posts from search results."`
	Search  subcmd.SearchCmd `cmd:"" help:"Search posts."`
}

func main() {
	ctx := kong.Parse(&cli, kong.Vars{"version": version})

	err := ctx.Run(&kasa.Context{
		Driver: esa.NewDriver(cli.Team, cli.Token, cli.Debug),
		Fmt:    &kasa.PrinterImpl{},
	})

	ctx.FatalIfErrorf(err)
}
