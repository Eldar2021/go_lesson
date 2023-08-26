package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "eldi-flag"
	app.Usage = "URFAVE FLAG kullanımı"

	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "name", Usage: "Your name"},
		&cli.IntFlag{Name: "age", Usage: "Your age"},
	}

	app.Action = func(ctx *cli.Context) error {
		fmt.Println("Hi, ", ctx.String("name"))
		fmt.Println("You are ", ctx.Int("age"), "years old")
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
	}
}
