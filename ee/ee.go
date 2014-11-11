
package main

import (
  "os"
  "github.com/codegangsta/cli"
)
func main() {
app := cli.NewApp()
app.Name = "say"
app.Commands = []cli.Command{
    {
        Name:        "site",
        Usage:       "Site create/delete/update",
        Description: "Site create/delete/update",
        Subcommands: []cli.Command{
            {
                Name:        "create",
                Usage:       "provide website site name",
                Description: "creates site",
                Flags: []cli.Flag{
                    cli.BoolFlag{
                        Name:  "html",
                        Usage: "html site",
                    },
                    cli.BoolFlag{
                        Name:  "php",
                        Usage: "php site",
                    },
                    cli.BoolFlag{
                        Name:  "mysql",
                        Usage: "mysql site",
                    },
                    cli.BoolFlag{
                        Name:  "wp",
                        Usage: "wordpress site",
                    },
                    cli.BoolFlag{
                        Name:  "wpsubdir",
                        Usage: "wordpress multisite subdirectory",
                    },
                    cli.BoolFlag{
                        Name:  "wpsubdomain",
                        Usage: "wordpress multisite subdomain",
                    },
                },
                Action: func(c *cli.Context) {
                    if c.Bool("html") {
                        println("Hello, html site", c.Args().First())
                    }
                    if c.Bool("php") {
                        println("Hello, php site", c.Args().First())
                    }
                    if c.Bool("mysql") {
                        println("Hello, mysql site", c.Args().First())
                    }
                    if c.Bool("wp") {
                        println("Hello, wp site", c.Args().First())
                    }
                    if c.Bool("wpsubdir") {
                        println("Hello, wpsubdir site", c.Args().First())
                    }
                    if c.Bool("wpsubdomain") {
                        println("Hello, wpsubdomain site", c.Args().First())
                    }
                },
            }, {
                Name:      "stack",
                Usage:     "sends a greeting in spanish",
                Flags: []cli.Flag{
                    cli.StringFlag{
                        Name:  "surname",
                        Value: "Jones",
                        Usage: "Surname of the person to greet",
                    },
                },
                Action: func(c *cli.Context) {
                    println("Hola, ", c.String("surname"))
                },
            }, {
                Name:      "french",
                ShortName: "fr",
                Usage:     "sends a greeting in french",
                Flags: []cli.Flag{
                    cli.StringFlag{
                        Name:  "nickname",
                        Value: "Stevie",
                        Usage: "Nickname of the person to greet",
                    },
                },
                Action: func(c *cli.Context) {
                    println("Bonjour, ", c.String("nickname"))
                },
            },
        },
    }, {
        Name:  "bye",
        Usage: "says goodbye",
        Action: func(c *cli.Context) {
            println("bye")
        },
    },
}

app.Run(os.Args)
}
