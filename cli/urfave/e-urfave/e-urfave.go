package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

/// Full API Example
/*
"github.com/urfave/cli/v2"
*/

func main() {
	app := &cli.App{
		Name:  "e-urfave",
		Usage: "eldiiar create module for dart/flutter projects",
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"cr"},
				Usage:   "Create meta-app, feature, service, contants for dart/flutter projects",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "app",
						Aliases: []string{"a"},
						Usage:   "Name of the material app",
						Action: func(ctx *cli.Context, s string) error {
							CreateMetaApp("meta_app")
							cmd := exec.Command("dart", "format", ".")
							cmd.Run()
							return nil
						},
					},
					&cli.StringFlag{
						Name:    "feature",
						Aliases: []string{"f"},
						Usage:   "Name of the feature",
						Action: func(ctx *cli.Context, s string) error {
							CreateFeature(s)
							cmd := exec.Command("dart", "format", ".")
							cmd.Run()
							return nil
						},
					},
					&cli.StringFlag{
						Name:    "constant",
						Aliases: []string{"v"},
						Usage:   "Name of the constant class",
						Action: func(ctx *cli.Context, s string) error {
							CreateConstant(s)
							cmd := exec.Command("dart", "format", ".")
							cmd.Run()
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
