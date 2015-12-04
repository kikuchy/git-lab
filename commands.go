package main

import (
	"github.com/codegangsta/cli"
	"log"
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
			settings, err := CollectGitLabSettings()
			if err != nil {
				log.Println(err)
				return
			}
			MergeRequestDelegate(settings, c)
		},
	},
}
