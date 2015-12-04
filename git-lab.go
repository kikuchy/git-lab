package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "git lab merge-request"
	app.Usage = "tekito-de"
	app.Commands = Commands

	app.Run(os.Args)
}
