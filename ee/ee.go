
package main

import (
  "github.com/codegangsta/cli"
)
func main() {
app := cli.NewApp()
app.Name = "EasyEngine Does amazing things"
app.Commands = []cli.Command{
   site_commands,
   stack_commands,
}

app.RunAndExitOnError()
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
                        println("site", c.Args().First(), "created with html option" )
                        }
                        if c.Bool("php") {
                            println("site", c.Args().First(), "created with php option")
                        }
                        if c.Bool("mysql") {
                            println("site", c.Args().First(), "created with mysql option")
                        }
                        if c.Bool("wp") {
                            println("site", c.Args().First(), "created with wp option")
                        }
                        if c.Bool("wpsubdir") {
                            println("site", c.Args().First(), "created with wpsubdir option")
                        }
                        if c.Bool("wpsubdomain") {
                            println("site", c.Args().First(), "created with wpsubdomain option")
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
                        Usage: "Install web packages",
                    },
                    cli.BoolFlag{
                        Name:  "mail",
                        Usage: "Install mail packages",
                    },
                    cli.BoolFlag{
                        Name:  "admin",
                        Usage: "Install admin packages",
                    },
                   cli.BoolFlag{
                        Name:  "mailscanner",
                        Usage: "Install mailscanner packages",
                    },
                    cli.BoolFlag{
                        Name:  "all",
                        Usage: "Install all packages",
                    },
                },
                Action: func(c *cli.Context) {
                    println(c.Bool("web"))
                    if c.Bool("web") {
                        println("Installed stack : web")
                    }else if c.Bool("mail") {
                        println("Installed stack : mail")
                    }else if c.Bool("admin") {
                        println("Installed stack : admin")
                    }else if c.Bool("mailscanner") {
                        println("Installed stack : mailscanner")
                    }else if c.Bool("all") {
                        println("Installed stack : all")
                    }
                },
            },{
                Name:        "purge",
                Usage:       "purge web/mail/admin/mailscanner packages",
                Description: "Purge packages",
                Flags: []cli.Flag{
                    cli.BoolFlag{
                        Name:  "web",
                        Usage: "Purge web packages",
                    },
                    cli.BoolFlag{
                        Name:  "mail",
                        Usage: "Purge mail packages",
                    },
                    cli.BoolFlag{
                        Name:  "admin",
                        Usage: "Purge admin packages",
                    },
                   cli.BoolFlag{
                        Name:  "mailscanner",
                        Usage: "Purge mailscanner packages",
                    },
                    cli.BoolFlag{
                        Name:  "all",
                        Usage: "Purge all packages",
                    },
                    
                },
                Action: func(c *cli.Context) {
                    if c.Bool("web") {
                        println("Purged stack : web")
                    }else if c.Bool("mail") {
                        println("Purged stack : mail")
                    }else if c.Bool("admin") {
                        println("Purged stack : admin")
                    }else if c.Bool("mailscanner") {
                        println("Purged stack : mailscanner")
                    }else if c.Bool("all") {
                        println("Purged stack : all")
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
                        println("Removed stack : web")
                    }else if c.Bool("mail") {
                        println("Removed stack : mail")
                    }else if c.Bool("admin") {
                        println("Removed stack : admin")
                    }else if c.Bool("mailscanner") {
                        println("Removed stack : mailscanner")
                    }else if c.Bool("all") {
                        println("Removed stack : all")
                    }
                },//end of Action
            },
        },
    }
