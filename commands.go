package main

import (
	"./internal"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	cli.Command{
		Name:  "merge-request",
		Usage: "Show merge requests.",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "branch, b",
				Usage: "Show branch name for each request.",
			},
		},
		Action: func(c *cli.Context) {
			settings := internal.CollectGitLabSettings()
			internal.MergeRequestDelegate(settings, c)
		},
	},
}
