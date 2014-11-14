package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ee"
	app.Version = Version
	app.Usage = ""
	app.Author = "Rahul Bansal"
	app.Email = "rahul.bansal@rtcamp.com"
	app.Commands = Commands

	app.Run(os.Args)
}
