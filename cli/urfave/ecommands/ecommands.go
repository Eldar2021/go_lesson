package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/// subcommands

func main() {
	app := cli.NewApp()
	app.Name = "e-commands"
	app.Usage = "eldi-commands"

	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "show all files",
			Action: func(ctx *cli.Context) error {
				fmt.Println("List files command")
				return nil
			},
		},
		{
			Name:  "create",
			Usage: "create a new file",
			Action: func(ctx *cli.Context) error {
				fmt.Println("Create a new file command")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
