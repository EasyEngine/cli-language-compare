
package main

import (
  "os"
  "github.com/codegangsta/cli"
)
func main() {
app := cli.NewApp()
app.Name = "say"
app.Commands = []cli.Command{
   site_commands,
   stack_commands,
}

app.Run(os.Args)
}

var site_commands = cli.Command{
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
                  args := c.Args()
                  if args.Present() {
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
                    } else {
                        println("Please provide site name")
                    }
                },
            },
        },//Site
    }

 var stack_commands = cli.Command{
        Name:      "stack",
        Usage:     "stack install/purge/remove web/mail/admin/mailscanner/all",
        Subcommands: []cli.Command{
            {
                Name:        "install",
                Usage:       "Install web/mail/admin/mailscanner/all packages",
                Description: "Installs packages",
                Flags: []cli.Flag{
                    cli.BoolFlag{
                        Name:  "web",
                        Usage: "Remove web packages",
                    },
                    cli.BoolFlag{
                        Name:  "mail",
                        Usage: "Remove mail packages",
                    },
                    cli.BoolFlag{
                        Name:  "admin",
                        Usage: "Remove admin packages",
                    },
                   cli.BoolFlag{
                        Name:  "mailscanner",
                        Usage: "Remove mailscanner packages",
                    },
                    cli.BoolFlag{
                        Name:  "all",
                        Usage: "Remove all packages",
                    },
                },
                Action: func(c *cli.Context) {
                    if c.Bool("web") {
                        println("Installed web package")
                    }else if c.Bool("mail") {
                        println("Installed mail packages")
                    }else if c.Bool("admin") {
                        println("Installed admin packages")
                    }else if c.Bool("mailscanner") {
                        println("Installed mailscanner packages")
                    }else if c.Bool("all") {
                        println("Installed all packages")
                    }else {
                        println("Installed web package")
                    }
                },
            },{
                Name:        "purge",
                Usage:       "purge web/mail/admin/mailscanner/all packages",
                Description: "Purge packages",
                Flags: []cli.Flag{
                    cli.BoolFlag{
                        Name:  "web",
                        Usage: "Remove web packages",
                    },
                    cli.BoolFlag{
                        Name:  "mail",
                        Usage: "Remove mail packages",
                    },
                    cli.BoolFlag{
                        Name:  "admin",
                        Usage: "Remove admin packages",
                    },
                   cli.BoolFlag{
                        Name:  "mailscanner",
                        Usage: "Remove mailscanner packages",
                    },
                    cli.BoolFlag{
                        Name:  "all",
                        Usage: "Remove all packages",
                    },
                },
                Action: func(c *cli.Context) {
                    if c.Bool("web") {
                        println("Purged web package")
                    }else if c.Bool("mail") {
                        println("Purged mail packages")
                    }else if c.Bool("admin") {
                        println("Purged admin packages")
                    }else if c.Bool("mailscanner") {
                        println("Purged mailscanner packages")
                    }else if c.Bool("all") {
                        println("Purged all packages")
                    }else {
                        println("Purged web package")
                    }
                },//end of Action
            },{
                Name:        "remove",
                Usage:       "remove web/mail/admin/mailscanner/all packages",
                Description: "remove packages",
                Flags: []cli.Flag{
                    cli.BoolFlag{
                        Name:  "web",
                        Usage: "Remove web packages",
                    },
                    cli.BoolFlag{
                        Name:  "mail",
                        Usage: "Remove mail packages",
                    },
                    cli.BoolFlag{
                        Name:  "admin",
                        Usage: "Remove admin packages",
                    },
                   cli.BoolFlag{
                        Name:  "mailscanner",
                        Usage: "Remove mailscanner packages",
                    },
                    cli.BoolFlag{
                        Name:  "all",
                        Usage: "Remove all packages",
                    },
                },
                Action: func(c *cli.Context) {
                    if c.Bool("web") {
                        println("Removed web package")
                    }else if c.Bool("mail") {
                        println("Removed mail packages")
                    }else if c.Bool("admin") {
                        println("Removed admin packages")
                    }else if c.Bool("mailscanner") {
                        println("Removed mailscanner packages")
                    }else if c.Bool("all") {
                        println("Removed all packages")
                    }else {
                        println("Removed web package")
                    }
                },//end of Action
            },
        },
    }
