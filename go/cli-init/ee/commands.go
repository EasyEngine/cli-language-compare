package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSite,
	commandStack,
	commandDebug,
	commandClean,
	commandSecure,
	commandUodate,
	commandInfo,
}

var commandSite = cli.Command{
	Name:  "site",
	Usage: "",
	Description: `
`,
	Action: doSite,
}

var commandStack = cli.Command{
	Name:  "stack",
	Usage: "",
	Description: `
`,
	Action: doStack,
}

var commandDebug = cli.Command{
	Name:  "debug",
	Usage: "",
	Description: `
`,
	Action: doDebug,
}

var commandClean = cli.Command{
	Name:  "clean",
	Usage: "",
	Description: `
`,
	Action: doClean,
}

var commandSecure = cli.Command{
	Name:  "secure",
	Usage: "",
	Description: `
`,
	Action: doSecure,
}

var commandUodate = cli.Command{
	Name:  "uodate",
	Usage: "",
	Description: `
`,
	Action: doUodate,
}

var commandInfo = cli.Command{
	Name:  "info",
	Usage: "",
	Description: `
`,
	Action: doInfo,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doSite(c *cli.Context) {
}

func doStack(c *cli.Context) {
}

func doDebug(c *cli.Context) {
}

func doClean(c *cli.Context) {
}

func doSecure(c *cli.Context) {
}

func doUodate(c *cli.Context) {
}

func doInfo(c *cli.Context) {
}
